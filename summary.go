package missed

type Summary struct {
	BlockNum   int               `json:"block_num"`
	Timestamp  int64             `json:"timestamp"`
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
