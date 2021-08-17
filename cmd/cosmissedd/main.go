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

const lagBlocks = 3

type savedState struct {
	CachedResult []byte
	CachedTop    []byte
	CachedChart  []byte
	CachedParams []byte
	BlockError   []byte
	Results      []*missed.Summary
	Successful   int
}

func main() {
	l := log.New(os.Stderr, "cosmissed | ", log.Lshortfile|log.LstdFlags)

	var (
		current, successful, track, listen                                          int
		cosmosApi, tendermintApi, prefix, networkName, socket, cacheFile, xRpcHosts string
		ready, stdout                                                               bool
	)

	flag.StringVar(&cosmosApi, "c", "http://127.0.0.1:1317", "cosmos http API endpoint")
	flag.StringVar(&tendermintApi, "t", "http://127.0.0.1:26657", "tendermint http API endpoint")
	flag.StringVar(&prefix, "p", "cosmos", "address prefix, ex- cosmos = cosmosvaloper, cosmosvalcons ...")
	flag.StringVar(&socket, "socket", "", "filename for unix socket to listen on, if set will disable TCP listener")
	flag.StringVar(&cacheFile, "cache", "cosmissed.dat", "filename for caching previous blocks")
	flag.StringVar(&xRpcHosts, "extra-rpc", "", "extra tendermint RPC endpoints to poll for peer info, comma seperated list of URLs")
	flag.IntVar(&listen, "l", 8080, "webserver port to listen on")
	flag.IntVar(&track, "n", 3000, "most recent blocks to track")
	flag.BoolVar(&stdout, "v", false, "log new records to stdout (error logs already on stderr)")

	flag.Parse()

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

	cachedResult := []byte("not ready")
	cachedTop := []byte("not ready")
	cachedChart := []byte("not ready")
	cachedParams := []byte("not ready")
	cachedMap := []byte("[]")
	blockError := []byte(`{"missing":{"fetch error":""}}`)
	var results []*missed.Summary

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	closeDb := make(chan interface{})

	go func() {
		sig := <-sigs
		if missed.GeoDb != nil {
			// prevent race using channel to close
			close(closeDb)
		}
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
			Successful:   successful,
		}
		e = out.Encode(ss)
		if e != nil {
			l.Fatal(e)
		}
		l.Fatal("exiting")
	}()

	var bcastMissed, bcastTop, bcastChart, bcastMap broadcast.Broadcaster
	defer func() {
		bcastMissed.Discard()
		bcastTop.Discard()
		bcastTop.Discard()
		bcastMap.Discard()
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

	go func() {
		fp := func() {
			j, e := missed.FetchPeers(xtraHosts)
			if e != nil {
				l.Println(e)
				return
			}
			if j != nil {
				// TODO: ensure it really has changed, this may be a lot of data.
				cachedMap = j
				e = bcastMap.Send(cachedMap)
				if e != nil {
					l.Println(e)
				}
			}
		}
		fp()
		tick := time.NewTicker(time.Minute)
		for {
			select {
			case <-tick.C:
				fp()
			case <-closeDb:
				_ = missed.GeoDb.Close()
			}
		}
	}()

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
		successful = state.Successful
		return true
	}() {
		results = make([]*missed.Summary, track)
		successful = current - track - lagBlocks - 1
		l.Printf("fetching last %d blocks, please wait\n", track)
		refresh()
		top()
	}
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
