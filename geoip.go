package missed

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

type GeoLightErr struct{}

func (g GeoLightErr) Error() string {
	return "geolight database not loaded"
}

func getLatLong(ipAddr string) (float32, float32, error) {
	if GeoDb == nil {
		return 0, 0, GeoLightErr{}
	}
	ip := net.ParseIP(ipAddr)
	if ip.String() != ipAddr || IsPrivate(ip) {
		return 0, 0, fmt.Errorf("ip %s is invalid for geo lookup", ipAddr)
	}
	record, err := GeoDb.City(ip)
	if err != nil {
		return 0, 0, err
	}
	return float32(record.Location.Latitude), float32(record.Location.Longitude), err
}

func getLocation(ipAddr string) (cityName string, country string, latLong point, err error) {
	if GeoDb == nil {
		err = GeoLightErr{}
		return
	}
	ip := net.ParseIP(ipAddr)
	if ip == nil || IsPrivate(ip) {
		err = fmt.Errorf("%s is not valid for geo lookup")
		return
	}
	city := &geoip2.City{}
	city, err = GeoDb.City(ip)

	if err != nil {
		return
	}
	cityName = city.City.Names["en"]
	if cityName == "" {
		cityName = "Unknown"
		if isp, e := GeoDb.ISP(ip); e == nil && isp.ISP != "" {
			cityName = isp.ISP
		}
	}
	country = city.Country.Names["en"]
	if country == "" {
		country = "Unknown"
	}
	return cityName, country, point{float32(city.Location.Latitude), float32(city.Location.Longitude)}, err
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

func (pm PeerMap) ToLinesJson() ([]byte, error) {
	allLines := make([]line3d, 0)
	for _, peer := range pm {
		lines, err := peer.toLines3d()
		if err != nil {
			l.Println(err)
			continue
		}
		allLines = append(allLines, lines...)
	}
	return json.Marshal(allLines)
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
