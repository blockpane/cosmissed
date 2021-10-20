package missed

import "time"

type minSignatures struct {
	Result struct {
		Block struct {
			Header struct {
				ProposerAddress string `json:"proposer_address"`
				Time            string `json:"time"`
			} `json:"header"`
			LastCommit struct {
				Signatures []struct {
					ValidatorAddress string `json:"validator_address"`
				} `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
	} `json:"result"`
}

func (m minSignatures) parse() (proposer string, utime int64, signers []string) {
	if m.Result.Block.Header.ProposerAddress == "" || m.Result.Block.LastCommit.Signatures == nil {
		return "", 0, nil
	}
	signers = make([]string, 0)
	for i := range m.Result.Block.LastCommit.Signatures {
		if m.Result.Block.LastCommit.Signatures[i].ValidatorAddress != "" {
			signers = append(signers, m.Result.Block.LastCommit.Signatures[i].ValidatorAddress)
		}
	}
	t, err := time.Parse(`2006-01-02T15:04:05.999999999Z`, m.Result.Block.Header.Time)
	if err == nil {
		utime = t.UTC().UnixNano() / 1_000_000 // ms
	}
	return m.Result.Block.Header.ProposerAddress, utime, signers
}
