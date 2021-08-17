package missed

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	commitPath        = `/commit?height=`
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
	curHeight, err = strconv.Atoi(sr.Result.SyncInfo.LatestBlockHeight)
	networkName = sr.Result.NodeInfo.Network
	return
}

func fetch(height int, client *http.Client, baseUrl, path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var req *http.Request
	var err error
	switch height {
	case 0:
		req, err = http.NewRequestWithContext(ctx, "GET", baseUrl+path, nil)
	default:
		req, err = http.NewRequestWithContext(ctx, "GET", baseUrl+path+strconv.Itoa(height), nil)
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
	b, err := fetch(height, TClient, TUrl, commitPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	proposer, ts, signers := m.parse()

	v := minValidatorSet{}
	b, err = fetch(height, CClient, CUrl, validatorsetsPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}
	addrs, cons := v.parse()

	b, err = fetch(0, CClient, CUrl, unbondingPath)
	if err != nil {
		return nil, err
	}
	jailed, err := ParseValidatorsResp(b, false)
	if err != nil {
		return nil, err
	}

	b, err = fetch(height, CClient, CUrl, historicalPath)
	if err != nil {
		return nil, err
	}
	vals, err := ParseValidatorsResp(b, true)
	if err != nil {
		return nil, err
	}
	return summarize(height, ts, proposer, signers, addrs, cons, vals, jailed, !catchingUp), nil
}

// TODO: build the peermap and return json to send to clients

func FetchPeers() (j []byte, err error){
	// FIXME: mock data
	// randomly update to test websocket
	//return mkPex(), nil
	_, pm, err := getNeighbors(nil)
	if err != nil {
		return nil, err
	}
	return pm.ToLinesJson()
}

var cachedPoints = make(map[string]point)

// getNeighbors calls the RCP endpoint asking for neighbors.
// TODO: the peers map isn't used here anymore ... consider changing it to a slice of id@host:port
//       that can be used in config.toml
func getNeighbors(nodes []string) (source string, peers PeerMap, e error) {
	if GeoDb == nil {
		return "", nil, errors.New("no geoip database is loaded, skipping peer discovery")
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// TODO: for now the nodes variable isn't used, but will be adding the ability to poll more nodes soon,
	// including trying our discovered peer's RPC. Need to investigate how difficult it would be to use
	// native pex instead of API.
	req, err := http.NewRequestWithContext(ctx, "GET", TUrl+`/net_info`, nil)
	if err != nil {
		return "", nil, err
	}
	resp, err := TClient.Do(req)
	if err != nil {
		return "", nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	ni := &netInfoResp{}
	err = json.Unmarshal(body, ni)
	if err != nil {
		return "", nil, err
	}
	listenerIp, err := ni.getListenerIp()
	if err != nil {
		return "", nil, err
	}

	var lat, long float32
	LongLat := cachedPoints[listenerIp]
	if LongLat[0] == 0 {
		long, lat, e = getLatLong(listenerIp)
		if e != nil {
			return "", nil, e
		}
		LongLat = point{lat, long}
		cachedPoints[listenerIp] = LongLat
	}

	result := PeerSet{
		Host:        listenerIp,
		Coordinates: LongLat,
		Peers:       make([]Peer, 0),
	}
	for _, p := range ni.Result.Peers {
		ll := cachedPoints[p.RemoteIp]
		if ll[0] == 0 {
			long, lat, e = getLatLong(p.RemoteIp)
			if e != nil {
				l.Println(e)
				continue
			}
			LongLat = point{lat, long}
			cachedPoints[p.RemoteIp] = LongLat
		}
		result.Peers = append(result.Peers, Peer{
			Host:        p.RemoteIp,
			Coordinates: ll,
			Outbound:    p.NodeInfo.IsOutbound,
		})
	}

	return listenerIp, PeerMap{result}, nil
}
