package main

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	missed "github.com/blockpane/cosmissed"
	"github.com/gorilla/websocket"
	"github.com/textileio/go-threads/broadcast"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	lagBlocks      = 2
	pollDiscovered = 5
)

type savedState struct {
	CachedResult []byte
	CachedTop    []byte
	CachedChart  []byte
	CachedParams []byte
	BlockError   []byte
	Results      []*missed.Summary
	Discovered   *missed.Discovered
	PeerMap      missed.PeerMap
	Successful   int
	GeoCache     *missed.GeoCache
}

func main() {
	l := log.New(os.Stderr, "cosmissed | ", log.Lshortfile|log.LstdFlags)

	var (
		current, successful, track, listen                                          int
		cosmosApi, tendermintApi, prefix, networkName, socket, cacheFile, xRpcHosts, user, apiKey string
		ready, stdout                                                               bool
	)

	flag.StringVar(&cosmosApi, "c", "http://127.0.0.1:1317", "cosmos http API endpoint")
	flag.StringVar(&tendermintApi, "t", "http://127.0.0.1:26657", "tendermint http API endpoint")
	flag.StringVar(&prefix, "p", "cosmos", "address prefix, ex- cosmos = cosmosvaloper, cosmosvalcons ...")
	flag.StringVar(&socket, "socket", "", "filename for unix socket to listen on, if set will disable TCP listener")
	flag.StringVar(&cacheFile, "cache", "cosmissed.dat", "filename for caching previous blocks")
	flag.StringVar(&xRpcHosts, "extra-rpc", "", "extra tendermint RPC endpoints to poll for peer info, comma seperated list of URLs")
	flag.StringVar(&user, "user", "", "Required: Username for GeoIP2 Precision Web Service")
	flag.StringVar(&apiKey, "key", "", "Required: Key for GeoIP2 Precision Web Service")
	flag.IntVar(&listen, "l", 8080, "webserver port to listen on")
	flag.IntVar(&track, "n", 3000, "most recent blocks to track")
	flag.BoolVar(&stdout, "v", false, "log new records to stdout (error logs already on stderr)")

	flag.Parse()

	if user == "" || apiKey == "" {
		l.Fatal("the '-user' and '-key' options are required")
	}


	switch {
	case strings.HasPrefix(cosmosApi, "unix://"):
		l.Println("Using socket:", strings.Replace(cosmosApi, `unix://`, "", 1))
		missed.CClient = &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", strings.Replace(cosmosApi, `unix://`, "", 1))
				},
			},
		}
		missed.CUrl = `http://unix`
	default:
		missed.CClient = http.DefaultClient
		missed.CUrl = cosmosApi
	}
	switch {
	case strings.HasPrefix(tendermintApi, `unix://`):
		l.Println("Using socket:", strings.Replace(tendermintApi, `unix://`, "", 1))
		missed.TClient = &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", strings.Replace(tendermintApi, `unix://`, "", 1))
				},
			},
		}
		missed.TUrl = `http://unix`
	default:
		missed.TClient = http.DefaultClient
		missed.TUrl = tendermintApi
	}

	cachedResult := []byte("{}")
	cachedTop := []byte("{}")
	cachedChart := []byte("{}")
	cachedParams, _ := json.Marshal(missed.Params{})
	cachedMap := []byte("[]")
	cachedNetStats := []byte("{}")
	blockError := []byte(`{"missing":{"fetch error":""}}`)
	pm := missed.PeerMap(make([]missed.PeerSet, 0))
	discovered := missed.NewDiscovered()
	var results []*missed.Summary

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//closeDb := make(chan interface{})

	go func() {
		sig := <-sigs
		//if missed.GeoDb != nil {
		//	// prevent race using channel to close
		//	close(closeDb)
		//}
		if socket != "" {
			os.Remove(socket)
		}
		l.Println("received", sig, "attempting to save state")
		f, e := os.OpenFile(cacheFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if e != nil {
			l.Fatal(e)
		}
		defer f.Close()
		out := gob.NewEncoder(f)
		ss := &savedState{
			CachedResult: cachedResult,
			CachedTop:    cachedTop,
			CachedChart:  cachedChart,
			CachedParams: cachedParams,
			BlockError:   blockError,
			Results:      results,
			PeerMap:      pm,
			Discovered:   discovered,
			Successful:   successful,
			GeoCache:     missed.MMCache,
		}
		e = out.Encode(ss)
		if e != nil {
			l.Fatal(e)
		}
		l.Fatal("exiting")
	}()

	var bcastMissed, bcastTop, bcastChart, bcastMap, bcastNetstats broadcast.Broadcaster
	defer func() {
		bcastMissed.Discard()
		bcastTop.Discard()
		bcastTop.Discard()
		bcastMap.Discard()
		bcastNetstats.Discard()
	}()

	// Additional tendermint RPC endpoints to poll for net_info
	xtraHosts := make([]string, 0)
	if xRpcHosts != "" {
		split := strings.Split(xRpcHosts, ",")
		for i := range split {
			if _, e := url.Parse(strings.Trim(split[i], " ")); e != nil {
				l.Fatalf(`invalid -extra-rpc value: %s, is not a URL.`, split[i])
			}
			xtraHosts = append(xtraHosts, split[i])
		}
	}

	top := func() {
		t, err := missed.TopMissed(results, track, prefix)
		if err != nil {
			l.Println(err)
			return
		}
		j, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			l.Println(err)
			return
		}
		if string(cachedTop) != string(j) {
			cachedTop = j
			err = bcastTop.Send(j)
			if err != nil {
				l.Println(err)
			}
		}
	}

	push := func(sum *missed.Summary) {
		if results[len(results)-1] != nil && results[len(results)-1].Timestamp != 0 {
			sum.DeltaSec = float64(sum.Timestamp-results[len(results)-1].Timestamp) / 1_000.0
		}
		results = append(results[1:], sum)
		if ready {
			cachedChart, _ = missed.SummariesToChart(results)
			cachedResult, _ = json.Marshal(results)
			j, _ := json.Marshal(sum)
			e := bcastMissed.Send(j)
			if e != nil {
				_ = l.Output(2, e.Error())
			}
			if stdout {
				fmt.Println(string(j))
			}

			e = bcastChart.Send(missed.SummaryToUpdate(sum))
			if e != nil {
				_ = l.Output(2, e.Error())
			}

			cachedParams, e = json.Marshal(missed.Params{
				Depth: track,
				Power: sum.VotePower,
				Chain: networkName,
			})
			if e != nil {
				_ = l.Output(2, e.Error())
			}
			top()
		}
	}

	newBlock := func() (new bool) {
		var h int
		var e error
		h, networkName, e = missed.CurrentHeight()
		if e != nil {
			_ = l.Output(2, e.Error())
			return false
		}
		if h-lagBlocks <= current {
			return false
		}
		current = h - lagBlocks
		return true
	}

	logmod := 100
	refresh := func() {
		for i := successful + 1; i < current; i++ {
			catchingUp := false
			if i < current-1 {
				catchingUp = true
			}
			summary, e := missed.FetchSummary(i, catchingUp)
			if e != nil {
				_ = l.Output(2, e.Error())
				return
			}
			if summary.BlockNum%logmod == 0 {
				l.Println("block", summary.BlockNum)
			}

			successful = i
			push(summary)
		}
	}

	if !newBlock() {
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		l.Fatalln("cannot get current block, giving up")
	}
	if !func() bool {
		f, e := os.Open(cacheFile)
		if e != nil {
			l.Println("could not load existing cache file, starting with clean state:", e.Error())
			return false
		}
		defer f.Close()
		enc := gob.NewDecoder(f)
		state := &savedState{}
		e = enc.Decode(state)
		if e != nil {
			l.Println("could not decode cache file, removing old file and starting with clean state:", e.Error())
			defer os.Remove(cacheFile)
			return false
		}
		if state.BlockError == nil || state.Results == nil || state.CachedTop == nil || state.CachedParams == nil || state.CachedChart == nil || state.CachedResult == nil {
			l.Println("cache file had invalid data, starting with clean state")
			defer os.Remove(cacheFile)
			return false
		}
		cachedResult = state.CachedResult
		cachedTop = state.CachedTop
		cachedChart = state.CachedChart
		cachedParams = state.CachedParams
		blockError = state.BlockError
		results = state.Results
		pm = state.PeerMap
		discovered = state.Discovered
		successful = state.Successful
		missed.MMCache = state.GeoCache
		if !missed.MMCache.SetAuth(user, apiKey) {
			l.Fatal("could not set maxmind credentials")
		}
		return true
	}() {
		if !missed.MMCache.SetAuth(user, apiKey) {
			l.Fatal("could not set maxmind credentials")
		}
		results = make([]*missed.Summary, track)
		successful = current - track - lagBlocks - 1
		l.Printf("fetching last %d blocks, please wait\n", track)
		refresh()
		top()
	}

	go func() {
		j := make([]byte, 0)
		netStats := missed.NetworkStats{
			CityLabels:    make([]string, 0),
			CityCounts:    make([]int, 0),
			CountryLabels: make([]string, 0),
			CountryCounts: make([]int, 0),
		}
		cachedNetStats, _ = json.Marshal(netStats)

		// if loading cache didn't work, these will be nil
		if discovered == nil {
			discovered = missed.NewDiscovered()
		}
		if pm == nil {
			pm = make([]missed.PeerSet, 0)
		}

		var busy bool
		newDiscovered := missed.PeerMap(make([]missed.PeerSet, 0))
		findPeers := func() {
			if busy {
				l.Println("skipping update, one is already in progress")
				return
			}
			busy = true
			defer func() { busy = false }()
			var e error

			l.Println("updating remote peers from reserved nodes")
			newPeers := missed.FetchPeers(xtraHosts)
			newPeers = append(newPeers, missed.FetchPeers(nil)...)
			// only poll discovered peers every 5 minutes to be polite:
			l.Println("updating remote peers from discovered nodes")
			pollNeighbors := make(map[string]bool)
			for i := range pm {
				for ii := range pm[i].Peers {
					ip := net.ParseIP(pm[i].Peers[ii].Host)
					if !missed.IsPrivate(ip) && pm[i].Peers[ii].RpcPort != 0 {
						if err := discovered.Add(ip, pm[i].Peers[ii].RpcPort); err != nil {
							l.Printf(`could not add discovered peer %s: %s`, pm[i].Peers[ii].Host, err)
							continue
						}
					}
					if !discovered.Skip(pm[i].Peers[ii].Host) {
						pollNeighbors[`http://`+pm[i].Peers[ii].Host+`:`+strconv.Itoa(pm[i].Peers[ii].RpcPort)] = true
					}
				}
			}
			rpcs := make([]string, 0)
			for k := range pollNeighbors {
				rpcs = append(rpcs, k)
			}
			tmpDisco := missed.FetchPeers(rpcs)
			if len(tmpDisco) >= len(newDiscovered) {
				newDiscovered = tmpDisco
			}

			l.Println("done updating discovered peers")

			log.Println("newPeers", len(newPeers), "newDiscovered", len(newDiscovered))
			pm = append(newPeers, newDiscovered...)
			l.Printf("tracking %d hosts", len(discovered.Nodes))
			var count int
			if count, j, e = pm.ToLinesJson(); e != nil {
				l.Println(`error converting lines3d to json:`, e)
			} else {
				if j != nil {
					cachedMap = j
					e = bcastMap.Send(cachedMap)
					if e != nil {
						l.Println(e)
					}
				}
			}
			netStats = missed.NetworkSummary(discovered, pm)
			netStats.PeersDiscovered = count
			cachedNetStats, e = json.Marshal(netStats)
			if e == nil {
				_ = bcastNetstats.Send(cachedNetStats)
			}
			l.Println("done updating reserved peers")
		}
		findPeers()
		tick := time.NewTicker(pollDiscovered * time.Minute)
		for {
			select {
			case <-tick.C:
				//findPeers()
			//case <-closeDb:
			//	_ = missed.GeoDb.Close()
			}
		}
	}()

	ready = true
	logmod = 10
	l.Println("cache populated, starting server.")
	go func() {
		tick := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-tick.C:
				if newBlock() {
					refresh()
				}
			}
		}
	}()

	setJsonHeader := func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Server", "cosmissed")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	broadcaster := func(writer http.ResponseWriter, request *http.Request, b *broadcast.Broadcaster) {
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			l.Print("upgrade:", err)
			return
		}
		defer c.Close()
		sub := b.Listen()
		defer sub.Discard()
		for message := range sub.Channel() {
			if e := c.WriteMessage(websocket.TextMessage, message.([]byte)); e != nil {
				l.Println(request.RemoteAddr, e)
				return
			}
		}
	}

	// Something very strange going on with http.FS ... if using switch below does not send mime types?
	http.Handle("/js/", http.FileServer(http.FS(missed.StaticContent)))
	http.Handle("/img/", http.FileServer(http.FS(missed.StaticContent)))
	http.Handle("/css/", http.FileServer(http.FS(missed.StaticContent)))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/missed/ws":
			broadcaster(writer, request, &bcastMissed)
		case "/top/ws":
			broadcaster(writer, request, &bcastTop)
		case "/chart/ws":
			broadcaster(writer, request, &bcastChart)
		case "/map/ws":
			broadcaster(writer, request, &bcastMap)
		case "/net/ws":
			broadcaster(writer, request, &bcastNetstats)
		case "/net":
			setJsonHeader(writer)
			_, _ = writer.Write(cachedNetStats)
		case "/chart":
			setJsonHeader(writer)
			_, _ = writer.Write(cachedChart)
		case "/missed":
			setJsonHeader(writer)
			_, _ = writer.Write(cachedResult)
		case "/top":
			setJsonHeader(writer)
			_, _ = writer.Write(cachedTop)
		case "/params":
			setJsonHeader(writer)
			_, _ = writer.Write(cachedParams)
		case "/map":
			setJsonHeader(writer)
			_, _ = writer.Write(cachedMap)
		case "/block":
			params := request.URL.Query()
			if params["num"] == nil || len(params["num"]) != 1 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			block, err := strconv.ParseUint(params["num"][0], 10, 32)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			var s *missed.Summary
			s, err = missed.FetchSummary(int(block), true)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write(blockError)
				return
			}
			var j []byte
			j, err = json.Marshal(s)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write(blockError)
				return
			}
			setJsonHeader(writer)
			_, _ = writer.Write(j)
		case "/network.html":
			writer.Header().Set("Server", "cosmissed")
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			// todo appropriate security headers.
			_, _ = writer.Write(missed.NetHtml)
		case "/missed.html":
			writer.Header().Set("Server", "cosmissed")
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			// todo appropriate security headers.
			_, _ = writer.Write(missed.MissedHtml)
		case "/", "/index.html":
			writer.Header().Set("Server", "cosmissed")
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			// todo appropriate security headers.
			_, _ = writer.Write(missed.IndexHtml)
		}
	})
	if socket != "" {
		server := http.Server{}
		unixListener, err := net.Listen("unix", socket)
		if err != nil {
			l.Fatal(err)
		}
		defer os.Remove(socket)
		l.Fatal(server.Serve(unixListener))
	}

	l.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listen), nil))
}
