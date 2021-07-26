package main

import (
	"encoding/json"
	"flag"
	"fmt"
	missed "github.com/blockpane/cosmissed"
	"github.com/gorilla/websocket"
	"github.com/textileio/go-threads/broadcast"
	"log"
	"net/http"
	"os"
	"time"
)

const lagBlocks = 3

func main() {
	l := log.New(os.Stderr, "cosmissed | ", log.Lshortfile|log.LstdFlags)

	var (
		current, successful, track, listen int
		cosmosApi, tendermintApi, prefix   string
		ready, stdout                      bool
	)

	flag.StringVar(&cosmosApi, "c", "http://127.0.0.1:1317", "cosmos http API endpoint")
	flag.StringVar(&tendermintApi, "t", "http://127.0.0.1:26657", "tendermint http API endpoint")
	flag.StringVar(&prefix, "p", "cosmos", "address prefix, ex- cosmos = cosmosvaloper, cosmosvalcons ...")
	flag.IntVar(&listen, "l", 8080, "webserver port to listen on")
	flag.IntVar(&track, "n", 3000, "most recent blocks to track")
	flag.BoolVar(&stdout, "v", false, "log new records to stdout (error logs already on stderr)")

	flag.Parse()

	cachedResult := []byte("not ready")
	cachedTop := []byte("not ready")
	cachedChart := []byte("not ready")
	cachedParams := []byte("not ready")

	results := make([]*missed.Summary, track)
	var bcastMissed, bcastTop, bcastChart broadcast.Broadcaster
	defer bcastMissed.Discard()
	defer bcastTop.Discard()

	top := func() {
		t, err := missed.TopMissed(results, track, prefix, cosmosApi)
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
			})
			if e != nil {
				_ = l.Output(2, e.Error())
			}
			top()
		}
	}

	newBlock := func() (new bool) {
		h, e := missed.CurrentHeight(tendermintApi)
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
			summary, e := missed.FetchSummary(cosmosApi, tendermintApi, i)
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
	successful = current - track - lagBlocks - 1
	l.Printf("fetching last %d blocks, please wait\n", track)
	refresh()
	top()
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
	http.HandleFunc("/missed/ws", func(writer http.ResponseWriter, request *http.Request) {
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			l.Print("upgrade:", err)
			return
		}
		defer c.Close()
		sub := bcastMissed.Listen()
		defer sub.Discard()
		for b := range sub.Channel() {
			if e := c.WriteMessage(websocket.TextMessage, b.([]byte)); e != nil {
				l.Println(request.RemoteAddr, e)
				return
			}
		}
	})

	http.HandleFunc("/top/ws", func(writer http.ResponseWriter, request *http.Request) {
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			l.Print("upgrade:", err)
			return
		}
		defer c.Close()
		sub := bcastTop.Listen()
		defer sub.Discard()
		for b := range sub.Channel() {
			if e := c.WriteMessage(websocket.TextMessage, b.([]byte)); e != nil {
				l.Println(request.RemoteAddr, e)
				return
			}
		}
	})

	http.HandleFunc("/chart/ws", func(writer http.ResponseWriter, request *http.Request) {
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			l.Print("upgrade:", err)
			return
		}
		defer c.Close()
		sub := bcastChart.Listen()
		defer sub.Discard()
		for b := range sub.Channel() {
			if e := c.WriteMessage(websocket.TextMessage, b.([]byte)); e != nil {
				l.Println(request.RemoteAddr, e)
				return
			}
		}
	})

	http.Handle("/js/", http.FileServer(http.FS(missed.Js)))
	http.Handle("/img/", http.FileServer(http.FS(missed.Js)))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
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
		case "/", "/index.html":
			writer.Header().Set("Server", "cosmissed")
			writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			// todo appropriate security headers.
			_, _ = writer.Write(missed.IndexHtml)
		}
	})

	l.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listen), nil))
}
