package missed

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"
	"io"
	"net/http"
	"strconv"
	"time"
)

const trackUnconfirmed = 900 // length of ringbuffer for tracking mempool. TODO: rethink this approach



func mkTxPoint(scaling int) []*txCountPoint {
	now := time.Now().UTC()
	u := make([]*txCountPoint, trackUnconfirmed/scaling)
	for i := len(u) - 1; i > 0; i-- {
		u[i] = &txCountPoint{
			Time:    now.Add(time.Duration(-i*scaling) * time.Second),
			Pending: 0,
		}
	}
	return u
}

var UnconfirmedCache = []byte(`[]`)
var unconfirmed = mkTxPoint(1)
var confirmed = mkTxPoint(1)


func WatchUnconfirmed(ctx context.Context, updates chan []byte, client *http.Client, baseUrl, origApi string, started chan interface{}) {
	<- started
	failed := make(chan interface{})
	points := make(chan *txCountPoint)
	go streamMemPool(ctx, client, baseUrl, points, failed)

	for {
		subClient, _ := rpchttp.New(origApi, "/websocket")
		err := subClient.Start()
		if err != nil {
			l.Fatal(origApi, err)
		}
		query := "tm.event = 'NewBlock'"
		freshBlocks, err := subClient.Subscribe(ctx, "test-client", query)
		if err != nil {
			l.Fatal(origApi, err)
		}

		func() {
			defer subClient.Stop()
			for {
				select {
				case <-ctx.Done():
					l.Println("WatchUnconfirmed() exiting: parent exited")
					return
				case <-failed:
					l.Println("WatchUnconfirmed() exiting: streamMemPool failed")
					return
				case b := <-freshBlocks:
					updates <- []byte(fmt.Sprintf(`["%s",%d,"Confirmed Tx"]`,
						b.Data.(types.EventDataNewBlock).Block.Time.Format(time.RFC3339),
						len(b.Data.(types.EventDataNewBlock).Block.Txs),
					))
					confirmed[len(confirmed)-1] = &txCountPoint{
						Time:    b.Data.(types.EventDataNewBlock).Block.Time,
						Pending: len(b.Data.(types.EventDataNewBlock).Block.Txs),
					}
				case p := <-points:
					updates <- []byte(fmt.Sprintf(`["%s",%d,"Pending Tx"]`, p.Time.Format(time.RFC3339), p.Pending))
					confirmed = append(confirmed[1:], &txCountPoint{
						Time:    p.Time,
						Pending: 0,
					})
					unconfirmed = append(unconfirmed[1:], p)
					if time.Now().Second()%5 == 0 {
						buf := bytes.NewBufferString(`[[`)
						prefix := ""
						for i := range unconfirmed {
							buf.WriteString(fmt.Sprintf(`%s["%s",%d,"Pending Tx"]`, prefix, unconfirmed[i].Time.UTC().Format(time.RFC3339), unconfirmed[i].Pending))
							prefix = ","
						}
						buf.WriteString(`],[`)
						prefix = ""
						for i := range confirmed {
							buf.WriteString(fmt.Sprintf(`%s["%s",%d,"Confirmed Tx"]`, prefix, confirmed[i].Time.UTC().Format(time.RFC3339), confirmed[i].Pending))
							prefix = ","
						}
						buf.WriteString(`]]`)
						UnconfirmedCache = buf.Bytes()
					}
				}
			}
		}()
		time.Sleep(10*time.Second)
	}
}

type txCountPoint struct {
	Time    time.Time `json:"time"`
	Pending int       `json:"pending"`
}

// unconfirmResp is a stripped down version of RPC response only holding total count
type unconfirmResp struct {
	Result struct {
		Total string `json:"total"`
	} `json:"result"`
}

func (u *unconfirmResp) total() int {
	if u == nil {
		return 0
	}
	i, _ := strconv.Atoi(u.Result.Total)
	return i
}

// streamMemPool sends the result from 'num_unconfirmed_txs' to a channel.
// this is going to work a lot better with a unix socket since tcp gets expensive when polling.
func streamMemPool(ctx context.Context, client *http.Client, baseUrl string, points chan *txCountPoint, failed chan interface{}) {
	defer close(failed)
	tick := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			go func() {
				var (
					body []byte
					err  error
					resp *http.Response
					t    = time.Now().UTC()
				)
				c, cancel := context.WithTimeout(context.Background(), time.Second-1)
				defer cancel()
				req, _ := http.NewRequestWithContext(c, "GET", baseUrl+`/num_unconfirmed_txs`, nil)
				resp, err = client.Do(req)
				if err != nil {
					l.Println("refresh mempool length:", err)
					points <- &txCountPoint{
						Time:    t,
						Pending: 0,
					}
					return
				}
				body, err = io.ReadAll(resp.Body)
				if err != nil {
					l.Println(err)
					points <- &txCountPoint{
						Time:    t,
						Pending: 0,
					}
					return
				}
				defer resp.Body.Close()
				ucr := &unconfirmResp{}
				err = json.Unmarshal(body, ucr)
				if err != nil {
					l.Println(err)
					points <- &txCountPoint{
						Time:    t,
						Pending: 0,
					}
					return
				}
				points <- &txCountPoint{
					Time:    t,
					Pending: ucr.total(),
				}
			}()
		}
	}
}
