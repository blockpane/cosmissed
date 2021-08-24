package missed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/savaki/geoip2"
	"net"
	"sync"
	"time"
)

var MMCache = &GeoCache{
	Nodes: make(map[string]*GeoNode),
}

type NoGeoKeyError struct{}
func (n NoGeoKeyError) Error() string {
	return "maxmind user or key not set"
}

type GeoNode struct {
	City string `json:"city"`
	Country string `json:"country"`
	Provider string `json:"provider"`
	LatLong point `json:"lat_long"`
}

type GeoCache struct {
	mux sync.Mutex
	Nodes map[string]*GeoNode

	id string
	key string
}

func (g *GeoCache) SetAuth(u, k string) bool {
	if u == "" || k == "" {
		return false
	}
	g.key = k
	g.id = u
	return true
}

func (g *GeoCache) hasAuth() bool {
	if g.key == "" || g.id == "" {
		return false
	}
	return true
}

func (g *GeoCache) Get(s string) *GeoNode {
	if !g.hasAuth() {
		return nil
	}
	g.mux.Lock()
	defer g.mux.Unlock()
	return g.Nodes[s]
}

func (g *GeoCache) set(ip string, node *GeoNode) bool {
	if !g.hasAuth() {
		return false
	}
	if node == nil || ip == "" {
		return false
	}
	g.mux.Lock()
	g.Nodes[ip] = node
	g.mux.Unlock()
	return true
}

func (g *GeoCache) Fetch(ip net.IP) (*GeoNode, error) {
	if !g.hasAuth() {
		return nil, NoGeoKeyError{}
	}
	if ip == nil || ip.String() == "" {
		return nil, errors.New("bad ip")
	}
	n := g.Get(ip.String())
	if n != nil {
		return n, nil
	}
	api := geoip2.New(g.id, g.key)
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	resp, err := api.City(ctx, ip.String())
	if err != nil {
		return nil, err
	}
	gn := &GeoNode{
		City:     resp.City.Names["en"],
		Country:  resp.Country.IsoCode,
		Provider: resp.Traits.Isp,
		LatLong:  point{float32(resp.Location.Latitude), float32(resp.Location.Longitude)},
	}
	switch "" {
	case gn.Country:
		gn.Country = "Unknown"
		fallthrough
	case gn.City:
		gn.City = "Unknown"
		fallthrough
	case gn.Provider:
		gn.Provider = "Unknown"
	}
	// annoying.
	if gn.Provider == "Digital Ocean" {
		gn.Provider = "DigitalOcean"
	}
	if len(gn.Provider) > 32 {
		gn.Provider = gn.Provider[:29]+"..."
	}
	switch float32(0) {
	case gn.LatLong[0], gn.LatLong[1]:
		return nil, errors.New("could not resolve location")
	default:
		if !g.set(ip.String(), gn) {
			return nil, errors.New("could not store geo node")
		}
	}
	return gn, nil
}



func (g *GeoCache) getLatLong(ipAddr string) (float32, float32, error) {
	ip := net.ParseIP(ipAddr)
	if ip.String() != ipAddr || IsPrivate(ip) {
		return 0, 0, fmt.Errorf("ip %s is invalid for geo lookup", ipAddr)
	}
	gn, e := g.Fetch(net.ParseIP(ipAddr))
	if e != nil {
		return 0,0,e
	}
	if gn == nil {
		return 0,0, errors.New("could not get lat long from cache")
	}
	return gn.LatLong[0], gn.LatLong[1], nil
}

func (g *GeoCache) getLocation(ipAddr string) (cityName, country, provider string, latLong point, err error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil || IsPrivate(ip) {
		err = fmt.Errorf("%s is not valid for geo lookup")
		return
	}
	var gn *GeoNode
	gn, err = g.Fetch(net.ParseIP(ipAddr))
	if err != nil {
		return
	}
	if gn == nil {
		err = errors.New("empty result for getLocation")
		return
	}
	return gn.City, gn.Country, gn.Provider, gn.LatLong, err
}

type PeerSet struct {
	Host        string `json:"host"`
	Coordinates point  `json:"coordinates"`
	Peers       []Peer `json:"peers"`
}

type Peer struct {
	Host        string `json:"host"`
	RpcPort     int    `json:"rpc_port"`
	Coordinates point  `json:"coordinates"`
	Outbound    bool   `json:"incoming"`
}

type PeerMap []PeerSet

func (pm PeerMap) ToLinesJson() (int, []byte, error) {
	allLines := make([]line3d, 0)
	uniq := make(map[string]bool)
	for _, peer := range pm {
		for _, h := range peer.Peers {
			uniq[h.Host] = true
		}
		lines, err := peer.toLines3d()
		if err != nil {
			l.Println(peer.Host, err)
			continue
		}
		allLines = append(allLines, lines...)
	}
	j, e := json.Marshal(allLines)
	return len(uniq), j, e
}

type point [2]float32
type line3d [2]point

func (ps PeerSet) toLines3d() ([]line3d, error) {
	if ps.Peers == nil || len(ps.Peers) == 0 {
		return nil, errors.New("no peers")
	}
	if ps.Coordinates == [2]float32{0, 0} {
		return nil, errors.New("host coordinates are [0,0] skipping")
	}
	result := make([]line3d, 0)
	for _, p := range ps.Peers {
		if p.Coordinates == [2]float32{0, 0} || IsPrivate(net.ParseIP(p.Host)) {
			// don't map 0,0 or private IPs
			continue
		}
		// switch order for inbound peers so line effect is not outbound.
		if !p.Outbound {
			result = append(result, [2]point{p.Coordinates, ps.Coordinates})
			continue
		}
		result = append(result, [2]point{ps.Coordinates, p.Coordinates})
	}
	return result, nil
}
