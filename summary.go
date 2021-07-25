package missed

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Summary struct {
	BlockNum   int               `json:"block_num"`
	Timestamp  int64             `json:"timestamp"`
	DeltaSec   float64           `json:"delta_sec"`
	Missed     int               `json:"missed"`
	Validators map[string]string `json:"missing"`
	Proposer   string            `json:"proposer"`
}

func summarize(blocknum int, ts int64, proposer string, signers []string, addrs map[string]bool, valcons map[string]string, validators []Validator) *Summary {
	s := Summary{
		BlockNum:   blocknum,
		Timestamp:  ts,
		Missed:     len(addrs) - len(signers),
		Validators: make(map[string]string),
	}
	names := make(map[string]string)
	for _, v := range validators {
		names[v.ValidatorAddress] = v.Description.Moniker
	}
	s.Proposer = names[proposer]
	for i := range signers {
		addrs[signers[i]] = true
	}
	for k, v := range addrs {
		if !v {
			s.Validators[names[k]] = valcons[k]
		}
	}
	return &s
}

// Top holds info about who missed rounds, and is used for building a polar area chart, where weight increases with votes
type Top struct {
	Moniker   string  `json:"moniker"`
	Missed    int     `json:"missed"`
	MissedPct float32 `json:"missed_pct"`
	Votes     uint64  `json:"votes"`
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
			t.Votes = tokens / 1_000_000
		}
		t.Moniker = weights.Validators[index[k]].Description.Moniker
		top = append(top, t)
	}
	return top, nil
}
