package missed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	blockPath         = `/block?height=`
	validatorsetsPath = `/validatorsets/`
	historicalPath    = `/cosmos/staking/v1beta1/historical_info/`
	unbondingPath     = `/cosmos/staking/v1beta1/validators?status=BOND_STATUS_UNBONDING`
	timeout           = 2 * time.Second
)

type statusResp struct {
	Result struct {
		NodeInfo struct {
			Network string `json:"network"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHeight string `json:"latest_block_height"`
			CatchingUp        bool   `json:"catching_up"`
		} `json:"sync_info"`
	} `json:"result"`
}

func CurrentHeight() (curHeight int, networkName string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", TUrl+"/status", nil)
	if err != nil {
		return 0, "", err
	}
	resp, err := TClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	sr := &statusResp{}
	err = json.Unmarshal(body, sr)
	if err != nil {
		return 0, "", err
	}
	if sr.Result.SyncInfo.CatchingUp {
		return 0, "", errors.New("node is catching up")
	}
	if NetworkId == "" {
		NetworkId = sr.Result.NodeInfo.Network
	}
	curHeight, err = strconv.Atoi(sr.Result.SyncInfo.LatestBlockHeight)
	networkName = sr.Result.NodeInfo.Network
	return
}

func fetch(height int, client *http.Client, baseUrl, path string, page int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var req *http.Request
	var err error
	switch height {
	case 0:
		req, err = http.NewRequestWithContext(ctx, "GET", baseUrl+path, nil)
	default:
		if page > 0 {
			req, err = http.NewRequestWithContext(ctx, "GET", baseUrl+path+strconv.Itoa(height)+`?page=`+strconv.Itoa(page), nil)
		} else {
			req, err = http.NewRequestWithContext(ctx, "GET", baseUrl+path+strconv.Itoa(height), nil)
		}
	}
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func FetchSummary(height int, catchingUp bool) (*Summary, error) {
	m := minSignatures{}
	b, err := fetch(height, TClient, TUrl, blockPath, 0)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	proposer, ts, signers := m.parse()

	v := minValidatorSet{}
	b, err = fetch(height, CClient, CUrl, validatorsetsPath, 0)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}
	// can't get more than 100 from the validatorsets endpoint! This is a dirty hack.
	num, _ := strconv.Atoi(v.Result.Total)
	if num > 100 {
		v2 := minValidatorSet{}
		b, err = fetch(height, CClient, CUrl, validatorsetsPath, 2)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, &v2)
		if err != nil {
			return nil, err
		}
		for i := range v2.Result.Validators {
			v.Result.Validators = append(v.Result.Validators, v2.Result.Validators[i])
		}
	}
	addrs, cons := v.parse()

	b, err = fetch(0, CClient, CUrl, unbondingPath, 0)
	if err != nil {
		return nil, err
	}
	jailed, err := ParseValidatorsResp(b, false)
	if err != nil {
		return nil, err
	}

	b, err = fetch(height, CClient, CUrl, historicalPath, 0)
	if err != nil {
		return nil, err
	}
	vals, err := ParseValidatorsResp(b, true)
	if err != nil {
		return nil, err
	}
	return summarize(height, ts, proposer, signers, addrs, cons, vals, jailed, !catchingUp), nil
}

func FetchPeers(xtra []string) (peers PeerMap) {
	peers = make([]PeerSet, 0)
	if xtra == nil {
		_, pm, err := GetNeighbors("")
		if err != nil {
			l.Println(err)
		}
		peers = append(peers, pm)
	}
	for _, s := range xtra {
		_, neighbor, e := GetNeighbors(s)
		if e != nil {
			l.Println(e)
			continue
		}
		peers = append(peers, neighbor)
	}
	return
}

var cachedPoints = make(map[string]point)
var cacheMux = sync.Mutex{}

// GetNeighbors calls the RCP endpoint asking for neighbors.
func GetNeighbors(node string) (source string, peers PeerSet, e error) {
	empty := PeerSet{}
	//if GeoDb == nil {
	//	return "", empty, errors.New("no geoip database is loaded, skipping peer discovery")
	//}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var client = TClient
	var url = TUrl
	if node != "" {
		client = http.DefaultClient
		url = node
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url+`/net_info`, nil)
	if err != nil {
		return "", empty, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", empty, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", empty, fmt.Errorf("GetNeighbors %s: %s", node, err.Error())
	}
	defer resp.Body.Close()
	ni := &netInfoResp{}
	err = json.Unmarshal(body, ni)
	if err != nil {
		return "", empty, err
	}
	listenerIp, err := ni.getListenerIp(strings.Split(strings.TrimPrefix(node, `http://`), `:`)[0])
	if err != nil {
		return "", empty, err
	}

	var lat, long float32
	cacheMux.Lock()
	LongLat := cachedPoints[listenerIp]
	if LongLat[0] == 0 {
		long, lat, e = MMCache.getLatLong(listenerIp)
		if e != nil {
			return "", empty, e
		}
		LongLat = point{lat, long}
		cachedPoints[listenerIp] = LongLat
	}
	cacheMux.Unlock()

	result := PeerSet{
		Host:        listenerIp,
		Coordinates: LongLat,
		Peers:       make([]Peer, 0),
	}
	for _, p := range ni.Result.Peers {
		cacheMux.Lock()
		// for now skip ipv6
		if strings.Contains(p.RemoteIp, `[`) {
			cacheMux.Unlock()
			continue
		}
		ll := cachedPoints[p.RemoteIp]
		cacheMux.Unlock()
		if ll[0] == 0 {
			long, lat, e = MMCache.getLatLong(p.RemoteIp)
			if e != nil {
				continue
			}
			LongLat = point{lat, long}
			cacheMux.Lock()
			cachedPoints[p.RemoteIp] = LongLat
			cacheMux.Unlock()
		}
		port := 0
		if p.NodeInfo.Other.RpcAddress == `@` {
			port = 26657
		} else if !strings.Contains(p.NodeInfo.Other.RpcAddress, `127.0.0.1`) && !strings.Contains(p.NodeInfo.Other.RpcAddress, `unix://`) {
			pSplit := strings.Split(p.NodeInfo.Other.RpcAddress, `:`)
			if port, err = strconv.Atoi(pSplit[len(pSplit)-1]); err != nil {
				l.Printf(`could not parse port from %s: %s`, p.NodeInfo.Other.RpcAddress, err.Error())
				continue
			}
		}
		result.Peers = append(result.Peers, Peer{
			Host:        p.RemoteIp,
			RpcPort:     port,
			Coordinates: ll,
			Outbound:    p.NodeInfo.IsOutbound,
		})
	}
	return listenerIp, result, nil
}
