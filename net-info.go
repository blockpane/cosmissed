package missed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const evictDuration = 14 * 24 * time.Hour

// MinNeighbor is a stripped down response from the API with only a moniker and IP
type MinNeighbor struct {
	RemoteIp string `json:"remote_ip"`
	NodeInfo struct {
		IsOutbound bool   `json:"is_outbound"`
		Moniker    string `json:"moniker"`
		Other      struct {
			// tendermint RPC API we will try to connect and get more peers (future feature)
			// if not listening on 127.0.0.1 or unix://
			RpcAddress string `json:"rpc_address"`
		} `json:"other"`
	} `json:"node_info"`
}

type netInfoResp struct {
	Result struct {
		Listeners []string      `json:"listeners"`
		Peers     []MinNeighbor `json:"peers"`
	} `json:"result"`
}

func (nif netInfoResp) getListenerIp(node string) (string, error) {
	h := nif.Result.Listeners[0]
	list := make([]string, 0)
	var host string

	switch {
	case h == `Listener(@)`:
		host = node
	case strings.Contains(h, `//`):
		list = strings.Split(h, `//`)
		if len(list) != 2 {
			return "", fmt.Errorf(`could not parse %+v into hostname`, nif.Result.Listeners)
		}
		host = strings.Split(list[1], `:`)[0]
	case strings.HasPrefix(h, `Listener(@`):
		host = strings.Split(strings.TrimPrefix(h, `Listener(@`), `:`)[0]
	}

	ipAddr := net.ParseIP(host)
	if ipAddr.String() != host {
		// got a DNS name
		ips, err := net.LookupIP(host)
		if err != nil || len(ips) == 0 {
			return "", fmt.Errorf("could not get ip for %s: %v", host, err)
		}
		host = ips[0].String()
		ipAddr = ips[0]
	}
	if IsPrivate(ipAddr) {
		return "", fmt.Errorf("host %s is a bad address, skipping", host)
	}
	return host, nil
}

var privateBlocks = [...]*net.IPNet{
	parseCIDR("0.0.0.0/32"),     // RFC 1918 IPv4 private network address
	parseCIDR("10.0.0.0/8"),     // RFC 1918 IPv4 private network address
	parseCIDR("100.64.0.0/10"),  // RFC 6598 IPv4 shared address space
	parseCIDR("127.0.0.0/8"),    // RFC 1122 IPv4 loopback address
	parseCIDR("169.254.0.0/16"), // RFC 3927 IPv4 link local address
	parseCIDR("172.16.0.0/12"),  // RFC 1918 IPv4 private network address
	parseCIDR("192.0.0.0/24"),   // RFC 6890 IPv4 IANA address
	parseCIDR("192.0.2.0/24"),   // RFC 5737 IPv4 documentation address
	parseCIDR("192.168.0.0/16"), // RFC 1918 IPv4 private network address
	parseCIDR("::1/128"),        // RFC 1884 IPv6 loopback address
	parseCIDR("fe80::/10"),      // RFC 4291 IPv6 link local addresses
	parseCIDR("fc00::/7"),       // RFC 4193 IPv6 unique local addresses
	parseCIDR("fec0::/10"),      // RFC 1884 IPv6 site-local addresses
	parseCIDR("2001:db8::/32"),  // RFC 3849 IPv6 documentation address
}

func parseCIDR(s string) *net.IPNet {
	_, block, err := net.ParseCIDR(s)
	if err != nil {
		// invariant OK to panic
		panic(fmt.Sprintf("Bad CIDR %s: %s", s, err))
	}
	return block
}

func IsPrivate(ip net.IP) bool {
	if ip == nil {
		return true // presumes a true result gets rejected
	}
	for _, priv := range privateBlocks {
		if priv.Contains(ip) {
			return true
		}
	}
	return false
}

type DiscoveredNode struct {
	Ip         string
	Port       int
	Skip       bool
	ValidUntil time.Time
}

type Discovered struct {
	mux sync.RWMutex

	Nodes map[string]*DiscoveredNode
}

func NewDiscovered() *Discovered {
	return &Discovered{
		Nodes: make(map[string]*DiscoveredNode),
	}
}

func (d *Discovered) Skip(s string) bool {
	d.mux.RLock()
	defer d.mux.RUnlock()
	if d.Nodes[s] == nil || d.Nodes[s].Skip == true {
		return true
	}
	return false
}

func (d *Discovered) Add(ip net.IP, port int) error {
	if NetworkId == "" {
		return errors.New("cannot add dynamic peer, not ready: undetermined network id")
	}
	ipAddr := ip.String()
	if ipAddr == "" {
		return fmt.Errorf("invalid IP %+v", ip)
	}
	d.mux.RLock()
	if d.Nodes[ipAddr] != nil {
		if !d.Nodes[ipAddr].Skip {
			d.Nodes[ipAddr].ValidUntil = time.Now().Add(evictDuration)
		}
		d.mux.RUnlock()
		return nil
	}
	d.mux.RUnlock()
	d.mux.Lock()
	defer d.mux.Unlock()
	d.Nodes[ipAddr] = &DiscoveredNode{
		Ip:         ipAddr,
		Port:       port,
		Skip:       true,
		ValidUntil: time.Now().Add(evictDuration),
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://"+ipAddr+":"+strconv.Itoa(port)+"/status", nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ir := &statusResp{}
	err = json.Unmarshal(body, ir)
	if err != nil {
		return err
	}
	d.Nodes[ipAddr] = &DiscoveredNode{
		Ip:         ipAddr,
		Port:       port,
		ValidUntil: time.Now().Add(evictDuration),
	}
	if NetworkId == ir.Result.NodeInfo.Network {
		d.Nodes[ipAddr].Skip = false
	} else {
		return fmt.Errorf("network id %s does not match %s", ir.Result.NodeInfo.Network, NetworkId)
	}
	return nil
}

func (d *Discovered) Trim() {
	d.mux.Lock()
	defer d.mux.Unlock()
	for k := range d.Nodes {
		if d.Nodes[k].ValidUntil.Before(time.Now()) {
			delete(d.Nodes, k)
		}
	}
}

type locationCounts struct {
	label string
	count int
}

type Sunburst struct {
	Name string `json:"name"`
	Value int `json:"value"`
	Children []Sunburst `json:"children,omitempty"`
}

type NetworkStats struct {
	PeersDiscovered int       `json:"peers_discovered"`
	RpcDiscovered   int       `json:"rpc_discovered"`
	CityLabels      []string  `json:"city_labels"`
	CityCounts      []int     `json:"city_counts"`
	CountryLabels   []string  `json:"country_labels"`
	CountryCounts   []int     `json:"country_counts"`
	LastUpdated     time.Time `json:"last_updated"`
	Sunburst        []Sunburst `json:"sunburst"`
	Providers       []Sunburst `json:"providers"`
}

type NodeLocation struct {
	Region     string `json:"region"`
	Country    string `json:"country"`
	Coordinate point  `json:"coordinate"`
	Isp        string `json:"isp"`
}

var nodeLocCache = make(map[string]*NodeLocation)
var ispList = make(map[string]map[string]map[string]int)

func NetworkSummary(d *Discovered, p PeerMap) NetworkStats {
	if d == nil {
		return NetworkStats{}
	}
	d.Trim()
	ns := NetworkStats{
		CityLabels:      make([]string, 0),
		CityCounts:      make([]int, 0),
		CountryLabels:   make([]string, 0),
		CountryCounts:   make([]int, 0),
		Sunburst:        make([]Sunburst, 0),
		Providers:       make([]Sunburst, 0),
		RpcDiscovered:   len(p),
	}
	allNodes := make(map[string]bool)
	for _, pSet := range p {
		for _, peer := range pSet.Peers {
			if IsPrivate(net.ParseIP(peer.Host)){
				continue
			}
			allNodes[peer.Host] = true
		}
	}
	cityFound := make(map[string]int)
	countryFound := make(map[string]int)
	increment := func(s string, w map[string]int) {
		w[s] += 1
	}
	for k := range allNodes {
		if nodeLocCache[k] != nil {
			increment(fmt.Sprintf("%s (%s)", nodeLocCache[k].Region, nodeLocCache[k].Country), cityFound)
			increment(nodeLocCache[k].Country, countryFound)
			continue
		}
		city, country, isp, latlong, err := MMCache.getLocation(k)
		if err != nil {
			continue
		}
		nodeLocCache[k] = &NodeLocation{
			Region:     city,
			Country:    country,
			Coordinate: latlong,
			Isp:        isp,
		}
		increment(fmt.Sprintf("%s (%s)", nodeLocCache[k].Region, nodeLocCache[k].Country), cityFound)
		increment(nodeLocCache[k].Country, countryFound)
		switch {
		case ispList[isp] == nil:
			ispList[isp] = make(map[string]map[string]int)
			fallthrough
		case ispList[isp][country] == nil:
			ispList[isp][country] = make(map[string]int)
		}
		ispList[isp][country][city] += 1
	}
	delete(ispList, "Unknown")
	for ispName, countryMap := range ispList {
		var ispCounter int
		iBurst := Sunburst{
			Name:     ispName,
			Children: make([]Sunburst, 0),
		}
		for countryName, cityMap := range countryMap {
			var countryCounter int
			countryBurst := Sunburst{
				Name:     countryName,
				Children: make([]Sunburst, 0),
			}
			for cityName, cityCount := range cityMap {
				countryBurst.Children = append(countryBurst.Children, Sunburst{
					Name:     cityName,
					Value:    cityCount,
				})
				ispCounter += cityCount
				countryCounter += cityCount
			}
			countryBurst.Value = countryCounter
			iBurst.Children = append(iBurst.Children, countryBurst)
		}
		iBurst.Value = ispCounter
		ns.Providers = append(ns.Providers, iBurst)
	}
	sort.Slice(ns.Providers, func(i, j int) bool {
		return ns.Providers[i].Value > ns.Providers[j].Value
	})
	cities := make([]locationCounts, 0)
	countries := make([]locationCounts, 0)
	for k, v := range cityFound {
		cities = append(cities, locationCounts{
			label: k,
			count: v,
		})
	}
	sort.Slice(cities, func(i, j int) bool {
		return cities[i].count > cities[j].count
	})
	for k, v := range countryFound {
		countries = append(countries, locationCounts{
			label: k,
			count: v,
		})
	}
	sort.Slice(countries, func(i, j int) bool {
		return countries[i].count > countries[j].count
	})
	for _, v := range cities {
		ns.CityLabels = append(ns.CityLabels, v.label)
		ns.CityCounts = append(ns.CityCounts, v.count)
	}
	for _, v := range countries {
		ns.CountryLabels = append(ns.CountryLabels, v.label)
		ns.CountryCounts = append(ns.CountryCounts, v.count)
	}
	for _, c := range countries {
		s := Sunburst{
			Name:     c.label,
			Value:    c.count,
			Children: make([]Sunburst, 0),
		}
		match := "("+s.Name+")"
		for _, ci := range cities {
			if strings.HasSuffix(ci.label, match) {
				s.Children = append(s.Children, Sunburst{
					Name:     ci.label,
					Value:    ci.count,
				})
			}
		}
		ns.Sunburst = append(ns.Sunburst, s)
	}
	ns.LastUpdated = time.Now().UTC()
	return ns
}
