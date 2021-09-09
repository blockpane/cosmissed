package missed

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/btcsuite/btcutil/bech32"
	"strconv"
	"strings"
	"time"
)

var Prefix = "cosmos"

type ConsensusPubkey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}

type CommisionRates struct {
	Rate          float32 `json:"rate"`
	MaxRate       float32 `json:"max_rate"`
	MaxChangeRate float32 `json:"max_change_rate"`
}

type Commission struct {
	CommisionRates CommisionRates `json:"commision_rates"`
	UpdateTime     time.Time      `json:"update_time"`
}

type Validator struct {
	OperatorAddress   string          `json:"operator_address"`
	ConsensusPubkey   ConsensusPubkey `json:"consensus_pubkey"`
	Valcons           string          `json:"valcons,omitempty"`           // this is not part of the struct provided by Cosmos
	ValidatorAddress  string          `json:"validator-address,omitempty"` // this is not part of the struct provided by Cosmos
	Jailed            bool            `json:"jailed"`
	Status            string          `json:"status"`
	Tokens            uint64          `json:"tokens"`
	DelegatorShares   float64         `json:"delegator_shares"`
	Description       Description     `json:"description"`
	UnbondingHeight   uint64          `json:"unbonding_height"`
	UnbondingTime     time.Time       `json:"unbonding_time"`
	Commission        Commission      `json:"commission"`
	MinSelfDelegation float64         `json:"min_self_delegation"`
}

type Pagination struct {
	NextKey string `json:"next_key"`
	Total   uint32 `json:"total"`
}

type HistValidatorsResp struct {
	Hist struct {
		Validators []Validator `json:"valset"`
	} `json:"hist"`
	Validators []Validator `json:"validators"`
}

func ParseValidatorsResp(body []byte, history bool) ([]Validator, error) {
	var (
		err      error
		response = HistValidatorsResp{}
		raw      = make(map[string]interface{})
		valset   = "valset"
	)
	if !history {
		valset = "validators"
	}
	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}
	if raw["hist"] == nil && raw["validators"] == nil {
		return nil, errors.New("not a valid HistValidatorsResp or Validators structure")
	}
	var bi bool
	if history {
		if raw, bi = raw["hist"].(map[string]interface{}); !bi {
			return nil, errors.New("no hist in validators response")
		}
	}
	response.Hist.Validators = make([]Validator, 0)

	toNum := func(s, t string) interface{} {
		switch t {
		case "float64":
			f, _ := strconv.ParseFloat(s, 64)
			return f
		case "float32":
			f, _ := strconv.ParseFloat(s, 32)
			return float32(f)
		case "uint64":
			i, _ := strconv.ParseUint(s, 10, 64)
			return i
		case "uint32":
			i, _ := strconv.ParseUint(s, 10, 64)
			return uint32(i)
		default:
			return 0
		}
	}

	toStr := func(v interface{}) string {
		if s, ok := v.(string); ok {
			return s
		}
		return ""
	}

	if _, ok := raw[valset].([]interface{}); !ok {
		return nil, errors.New("no validators found")
	}

	for _, val := range raw[valset].([]interface{}) {
		validator := Validator{}
		if _, ok := val.(map[string]interface{}); !ok {
			continue
		}
		for k, v := range val.(map[string]interface{}) {
			switch toStr(k) {
			case "operator_address":
				validator.OperatorAddress = toStr(v)
			case "consensus_pubkey":
				validator.ConsensusPubkey = ConsensusPubkey{}
				if pk, ok := v.(map[string]interface{}); ok {
					for j, u := range pk {
						switch toStr(j) {
						case "@type":
							validator.ConsensusPubkey.Type = toStr(u)
						case "key":
							validator.ConsensusPubkey.Key = toStr(u)
							validator.Valcons, _ = PubToCons(Prefix+"valcons", validator.ConsensusPubkey.Key)
							validator.ValidatorAddress, _ = pubToHex(validator.ConsensusPubkey.Key)
						}
					}
				}
			case "jailed":
				switch v.(type) {
				case string:
					validator.Jailed, _ = strconv.ParseBool(toStr(v))
				case bool:
					validator.Jailed, _ = v.(bool)
				}
			case "status":
				validator.Status = toStr(v)
			case "tokens":
				s := toStr(v)
				if len(s) <= Precision {
					s = "0"
				} else {
					s = s[:len(s)-Precision]
				}
				validator.Tokens = toNum(s, "uint64").(uint64)
			case "delegator_shares":
				validator.DelegatorShares = toNum(toStr(v), "float64").(float64)
			case "description":
				validator.Description = Description{}
				if desc, ok := v.(map[string]interface{}); ok {
					for j, u := range desc {
						switch toStr(j) {
						case "moniker":
							validator.Description.Moniker = toStr(u)
						case "identity":
							validator.Description.Identity = toStr(u)
						case "website":
							validator.Description.Website = toStr(u)
						case "security_contact":
							validator.Description.SecurityContact = toStr(u)
						case "details":
							validator.Description.Details = toStr(u)
						}
					}
				}
			case "unbonding_height":
				validator.UnbondingHeight = toNum(toStr(v), "uint64").(uint64)
			case "unbonding_time":
				validator.UnbondingTime, _ = time.Parse("2006-01-02T15:04:05.999999999Z", toStr(v))
			case "commission":
				validator.Commission = Commission{}
				if com, ok := v.(map[string]interface{}); ok {
					for j, u := range com {
						switch toStr(j) {
						case "commission_rates":
							if comr, kk := u.(map[string]interface{}); kk {
								for j, u := range comr {
									switch toStr(j) {
									case "rate":
										validator.Commission.CommisionRates.Rate = toNum(toStr(u), "float32").(float32)
									case "max_rate":
										validator.Commission.CommisionRates.MaxRate = toNum(toStr(u), "float32").(float32)
									case "max_change_rate":
										validator.Commission.CommisionRates.MaxChangeRate = toNum(toStr(u), "float32").(float32)
									}
								}
							}
						case "update_time":
							validator.Commission.UpdateTime, _ = time.Parse("2006-01-02T15:04:05.999999999Z", toStr(u))
						}
					}
				}
			case "min_self_delegation":
				validator.MinSelfDelegation = toNum(toStr(v), "float64").(float64)
			}
		}
		response.Hist.Validators = append(response.Hist.Validators, validator)
	}

	return response.Hist.Validators, nil
}

func PubToCons(prefix, pub string) (string, error) {
	b := make([]byte, 32)
	_, err := base64.StdEncoding.Decode(b, []byte(pub))
	if err != nil {
		return "", err
	}
	sha := sha256.New()
	sha.Write(b)
	conv, err := bech32.ConvertBits(sha.Sum(nil)[:20], 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode(prefix, conv)
}

func pubToHex(pub string) (string, error) {
	b := make([]byte, 32)
	_, err := base64.StdEncoding.Decode(b, []byte(pub))
	if err != nil {
		return "", err
	}
	sha := sha256.New()
	sha.Write(b)
	return strings.ToUpper(hex.EncodeToString(sha.Sum(nil)[:20])), nil
}
