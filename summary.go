package missed

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

var Precision int

type Params struct {
	Depth int    `json:"depth"`
	Power uint64 `json:"power"`
	Chain string `json:"chain"`
}

type Summary struct {
	BlockNum          int               `json:"block_num"`
	Timestamp         int64             `json:"timestamp"`
	DeltaSec          float64           `json:"delta_sec,omitempty"`
	Missed            int               `json:"missed"`
	MissingValidators map[string]string `json:"missing"`
	PresentValidators map[string]string `json:"-"`
	Proposer          string            `json:"proposer"`
	VotePower         uint64            `json:"vote_power"`
	VoteMissing       uint64            `json:"vote_missing"`
	JailedUnbonding   map[string]string `json:"jailed_unbonding"`
}

func summarize(blocknum int, ts int64, proposer string, signers []string, addrs map[string]bool, valcons map[string]string, validators, jailed []Validator, includeJailed bool) *Summary {
	s := Summary{
		BlockNum:          blocknum,
		Timestamp:         ts,
		Missed:            len(addrs) - len(signers),
		MissingValidators: make(map[string]string),
		PresentValidators: make(map[string]string),
		JailedUnbonding:   make(map[string]string),
	}
	names := make(map[string]string)
	powers := make(map[string]uint64)
	for _, v := range validators {
		names[v.ValidatorAddress] = bm.Sanitize(v.Description.Moniker)
		powers[v.ValidatorAddress] = v.Tokens
		s.VotePower += v.Tokens
	}
	if includeJailed {
		for _, v := range jailed {
			if v.Jailed == false {
				continue
			}
			s.JailedUnbonding[bm.Sanitize(v.Description.Moniker)] = "" // FIXME: get valoper for lookups.
		}
	}
	s.Proposer = names[proposer]
	for i := range signers {
		addrs[signers[i]] = true
	}
	for k, v := range addrs {
		if !v {
			s.MissingValidators[strings.TrimSpace(bm.Sanitize(names[k]))] = valcons[k]
			s.VoteMissing += powers[k]
			delete(addrs, k)
		}
	}
	for k := range addrs {
		s.PresentValidators[strings.TrimSpace(bm.Sanitize(names[k]))] = valcons[k]
	}
	return &s
}

// Top holds info about who missed rounds, and is used for building a polar area chart, where weight increases with votes
type Top struct {
	Moniker   string  `json:"moniker"`
	Missed    int     `json:"missed"`
	MissedPct float32 `json:"missed_pct"`
	Votes     int64   `json:"votes"`
}

type minValidators struct {
	Validators []struct {
		ConsensusPubkey struct {
			Key string `json:"key"`
		} `json:"consensus_pubkey"`
		Tokens      string `json:"tokens"`
		Description struct {
			Moniker string `json:"moniker"`
		} `json:"description"`
	} `json:"validators"`
}

func TopMissed(summaries []*Summary, blocks int, prefix string) ([]*Top, error) {
	// get current vote power:
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", CUrl+"/cosmos/staking/v1beta1/validators?pagination.limit=1000", nil)
	if err != nil {
		return nil, err
	}
	resp, err := CClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	weights := minValidators{}
	err = json.Unmarshal(body, &weights)
	if err != nil {
		return nil, err
	}
	// tally up misses
	counts := make(map[string]int)
	for _, s := range summaries {
		if s != nil {
			for _, v := range s.MissingValidators {
				counts[v] += 1
			}
			for _, v := range s.PresentValidators {
				counts[v] += 0
			}
		}
	}
	index := make(map[string]int)
	for i := range weights.Validators {
		valcons, e := PubToCons(prefix+"valcons", weights.Validators[i].ConsensusPubkey.Key)
		if e != nil {
			continue
		}
		index[valcons] = i
	}
	top := make([]*Top, 0)
	for k, v := range counts {
		t := &Top{
			Missed:    v,
			MissedPct: 100.0 * float32(v) / float32(blocks),
		}
		if len(weights.Validators[index[k]].Tokens) > Precision{
			tokens, e := strconv.ParseUint(weights.Validators[index[k]].Tokens[:len(weights.Validators[index[k]].Tokens)-Precision], 10, 64)
			if e == nil {
				t.Votes = int64(tokens) / -1_000_000
			} else {
				log.Println("error calculating votes", e)
			}
		}

		t.Moniker = bm.Sanitize(weights.Validators[index[k]].Description.Moniker)
		top = append(top, t)
	}
	for i := range top {
		top[i].Moniker = strings.TrimRight(top[i].Moniker, "🟢")
		if top[i].Missed == 0 {
			top[i].Moniker = top[i].Moniker + " 🟢"
		}
	}
	sort.Slice(top, func(i, j int) bool {
		switch false {
		case top[i].Missed == top[j].Missed:
			return top[i].Missed > top[j].Missed
		case top[i].Votes == top[j].Votes:
			return top[i].Votes > top[j].Votes
		default:
			return top[i].Moniker[0] > top[j].Moniker[0]
		}
	})
	return top, nil
}

// BlockChart is just another way of presenting the same data that is easier to use with some charting libraries, where
// each series is presented as an array
type BlockChart struct {
	Blocks  []int     `json:"blocks"`          // block number
	Time    []string  `json:"time"`            // time as a string
	Missed  []int     `json:"missed"`          // number of missing validators
	MissPct []float64 `json:"missing_percent"` // vote power missing (shows actual risk of losing consensus.)
	Took    []float64 `json:"took"`            // time since last block in seconds
}

func SummariesToChart(s []*Summary) ([]byte, error) {
	bc := BlockChart{
		Blocks:  make([]int, len(s)),
		Time:    make([]string, len(s)),
		Missed:  make([]int, len(s)),
		Took:    make([]float64, len(s)),
		MissPct: make([]float64, len(s)),
	}
	for i := range s {
		bc.Blocks[i] = s[i].BlockNum
		bc.Time[i] = time.Unix(s[i].Timestamp/1000, 0).UTC().Format(time.Stamp)
		bc.Missed[i] = s[i].Missed
		bc.Took[i] = s[i].DeltaSec
		if s[i].VotePower > 0 {
			bc.MissPct[i] = (float64(s[i].VoteMissing) / float64(s[i].VotePower)) * 100.0
		}
	}
	return json.MarshalIndent(bc, "", "  ")
}

type ChartUpdate struct {
	Block   int     `json:"block"`
	Time    string  `json:"time"`
	Missed  int     `json:"missed"`
	MissPct float64 `json:"missing_percent"`
	Took    float64 `json:"took"`
}

func SummaryToUpdate(s *Summary) []byte {
	u := ChartUpdate{
		Block:  s.BlockNum,
		Time:   time.Unix(s.Timestamp/1000, 0).UTC().Format(time.Stamp),
		Missed: s.Missed,
		Took:   s.DeltaSec,
	}
	if s.VotePower > 0 {
		u.MissPct = (float64(s.VoteMissing) / float64(s.VotePower)) * 100.0
	}
	j, _ := json.Marshal(u)
	return j
}
