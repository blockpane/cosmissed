package main

import (
	"encoding/json"
	"flag"
	"fmt"
	missed "github.com/frameloss/playground/cosmos-missed"
	"github.com/gorilla/websocket"
	"github.com/textileio/go-threads/broadcast"
	"log"
	"net/http"
	"os"
	"time"
)

const lagBlocks = 3

func main() {
	l := log.New(os.Stderr, "cosmiss | ", log.Lshortfile|log.LstdFlags)

	var (
		current, successful, track, listen int
		cosmosApi, tendermintApi, prefix   string
		ready, stdout bool
	)

	flag.StringVar(&cosmosApi, "c", "http://127.0.0.1:1317", "cosmos http API endpoint")
	flag.StringVar(&tendermintApi, "t", "http://127.0.0.1:26657", "tendermint http API endpoint")
	flag.StringVar(&prefix, "p", "cosmos", "address prefix, ex- cosmos = cosmosvaloper, cosmosvalcons ...")
	flag.IntVar(&listen, "l", 8080, "webserver port to listen on")
	flag.IntVar(&track, "n", 3000, "most recent blocks to track")
	flag.BoolVar(&stdout, "v", false, "log new records to stdout (error logs on stderr")

	flag.Parse()

	results := make([]*missed.Summary, track)
	cachedResult := []byte("not ready")
	var bcast broadcast.Broadcaster
	defer bcast.Discard()

	push := func(sum *missed.Summary) {
		results = append(results[1:], sum)
		if ready {
			cachedResult, _ = json.Marshal(results)
			j, _ := json.Marshal(sum)
			e := bcast.Send(j)
			if stdout {
				fmt.Println(string(j))
			}
			if e != nil {
				_ = l.Output(2, e.Error())
			}
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
			if summary.BlockNum % logmod == 0 {
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


	http.HandleFunc("/missed", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Server", "cosmissd")
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("X-You-Kids", "get off my lawn")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(cachedResult)
	})

	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	http.HandleFunc("/missed/ws", func(writer http.ResponseWriter, request *http.Request) {
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			l.Print("upgrade:", err)
			return
		}
		defer c.Close()
		sub := bcast.Listen()
		defer sub.Discard()
		for b := range sub.Channel() {
			if e := c.WriteMessage(websocket.TextMessage, b.([]byte)); e != nil {
				l.Println(request.RemoteAddr, e)
				return
			}
		}
	})

	l.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listen), nil))

}
