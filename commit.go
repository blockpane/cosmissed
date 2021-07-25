package missed

import "time"

type minSignatures struct {
	Result struct {
		SignedHeader struct {
			Header struct {
				ProposerAddress string `json:"proposer_address"`
				Time            string `json:"time"`
			} `json:"header"`
			Commit struct {
				Signatures []struct {
					ValidatorAddress string `json:"validator_address"`
				} `json:"signatures"`
			} `json:"commit"`
		} `json:"signed_header"`
	} `json:"result"`
}

func (m minSignatures) parse() (proposer string, utime int64, signers []string) {
	if m.Result.SignedHeader.Header.ProposerAddress == "" || m.Result.SignedHeader.Commit.Signatures == nil {
		return "", 0, nil
	}
	signers = make([]string, 0)
	for i := range m.Result.SignedHeader.Commit.Signatures {
		if m.Result.SignedHeader.Commit.Signatures[i].ValidatorAddress != "" {
			signers = append(signers, m.Result.SignedHeader.Commit.Signatures[i].ValidatorAddress)
		}
	}
	t, err := time.Parse(`2006-01-02T15:04:05.999999999Z`, m.Result.SignedHeader.Header.Time)
	if err == nil {
		utime = t.UTC().UnixNano() / 1_000_000 // ms
	}
	return m.Result.SignedHeader.Header.ProposerAddress, utime, signers
}
