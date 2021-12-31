package missed

type minValidatorSet struct {
	Result struct {
		Validators []struct {
			Address string `json:"address"`
			PubKey  struct {
				Value string `json:"value"`
			} `json:"pub_key"`
		} `json:"validators"`
		Total string `json:"total"`
	} `json:"result"`
}

func (mv minValidatorSet) parse() (addresses map[string]bool, valcons map[string]string) {
	if mv.Result.Validators == nil || len(mv.Result.Validators) == 0 {
		return nil, nil
	}
	addresses = make(map[string]bool)
	valcons = make(map[string]string)

	for _, val := range mv.Result.Validators {
		addr, err := pubToHex(val.PubKey.Value)
		if err != nil {
			continue
		}
		addresses[addr] = false
		valcons[addr] = val.Address
	}
	return
}
