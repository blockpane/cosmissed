package missed

import (
	"fmt"
	"net"
	"strings"
)

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

func (nif netInfoResp) getListenerIp() (string, error) {
	list := strings.Split(nif.Result.Listeners[0], `//`)
	if len(list) != 2 {
		return "", fmt.Errorf(`could not parse %+v into hostname`, nif.Result.Listeners)
	}
	host := strings.Split(list[1], `:`)[0]
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
	if isPrivate(ipAddr) {
		return "", fmt.Errorf("host %s is a bad address, skipping", host)
	}
	return host, nil
}

var privateBlocks = [...]*net.IPNet{
	parseCIDR("10.0.0.0/8"),     // RFC 1918 IPv4 private network address
	//parseCIDR("100.64.0.0/10"),  // RFC 6598 IPv4 shared address space ... TODO: no longer private?
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

func isPrivate(ip net.IP) bool {
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
