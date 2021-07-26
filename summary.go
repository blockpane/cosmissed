package missed

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Params struct {
	Depth int `json:"depth"`
	Power uint64 `json:"power"`
}

type Summary struct {
	BlockNum    int               `json:"block_num"`
	Timestamp   int64             `json:"timestamp"`
	DeltaSec    float64           `json:"delta_sec"`
	Missed      int               `json:"missed"`
	Validators  map[string]string `json:"missing"`
	Proposer    string            `json:"proposer"`
	VotePower   uint64            `json:"vote_power"`
	VoteMissing uint64            `json:"vote_missing"`
}

func summarize(blocknum int, ts int64, proposer string, signers []string, addrs map[string]bool, valcons map[string]string, validators []Validator) *Summary {
	s := Summary{
		BlockNum:   blocknum,
		Timestamp:  ts,
		Missed:     len(addrs) - len(signers),
		Validators: make(map[string]string),
	}
	names := make(map[string]string)
	powers := make(map[string]uint64)
	for _, v := range validators {
		names[v.ValidatorAddress] = v.Description.Moniker
		powers[v.ValidatorAddress] = v.Tokens
		s.VotePower += v.Tokens
	}
	s.Proposer = names[proposer]
	for i := range signers {
		addrs[signers[i]] = true
	}
	for k, v := range addrs {
		if !v {
			s.Validators[names[k]] = valcons[k]
			s.VoteMissing += powers[k]
		}
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

func TopMissed(summaries []*Summary, blocks int, prefix, cosmosApi string) ([]*Top, error) {
	// get current vote power:
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", cosmosApi+"/cosmos/staking/v1beta1/validators?pagination.limit=1000", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
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
		if s != nil && s.Missed > 0 {
			for _, v := range s.Validators {
				counts[v] += 1
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
		if tokens, e := strconv.ParseUint(weights.Validators[index[k]].Tokens, 10, 64); e == nil {
			t.Votes = int64(tokens) / -1_000_000
		}
		t.Moniker = weights.Validators[index[k]].Description.Moniker
		top = append(top, t)
	}
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
