package missed

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

const trackUnconfirmed = 3600 // length of ringbuffer for tracking mempool. TODO: rethink this approach

var UnconfirmedCache = []byte(`[]`)
var unconfirmed = func() []*memPoolPoint {
	now := time.Now().UTC()
	u := make([]*memPoolPoint, trackUnconfirmed)
	for i := trackUnconfirmed - 1; i > 0; i-- {
		u[i] = &memPoolPoint{
			Time:    now.Add(time.Duration(-i) * time.Second),
			Pending: 0,
		}
	}
	return u
}()

func WatchUnconfirmed(ctx context.Context, updates chan []byte, client *http.Client, baseUrl string) {
	failed := make(chan interface{})
	points := make(chan *memPoolPoint)
	go streamMemPool(ctx, client, baseUrl, points, failed)
	for {
		select {
		case <-ctx.Done():
			l.Println("WatchUnconfirmed() exiting: parent exited")
			return
		case <-failed:
			l.Println("WatchUnconfirmed() exiting: streamMemPool failed")
			return
		case p := <-points:
			if j, _ := json.Marshal(p); j != nil {
				updates <-j
			}
			unconfirmed = append(unconfirmed[1:], p)
			if time.Now().Second()%5 == 0 {
				j, e := json.Marshal(unconfirmed)
				if e != nil {
					l.Println(e)
					continue
				}
				UnconfirmedCache = j
			}
		}
	}
}

type memPoolPoint struct {
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
func streamMemPool(ctx context.Context, client *http.Client, baseUrl string, points chan *memPoolPoint, failed chan interface{}) {
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
					points <- &memPoolPoint{
						Time:    t,
						Pending: 0,
					}
					return
				}
				body, err = io.ReadAll(resp.Body)
				if err != nil {
					l.Println(err)
					points <- &memPoolPoint{
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
					points <- &memPoolPoint{
						Time:    t,
						Pending: 0,
					}
					return
				}
				points <- &memPoolPoint{
					Time:    t,
					Pending: ucr.total(),
				}
			}()
		}
	}
}
