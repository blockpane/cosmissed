package missed

import (
	"fmt"
	"testing"
)

func TestParseValidatorsResp(t *testing.T) {
	v, e := ParseValidatorsResp([]byte(testValidators))
	if e != nil {
		t.Error(e)
	}
	if v == nil {
		t.Error("did not parse validators")
	}
	fmt.Printf("%+v\n", v)
}

const testValidators = `{
  "hist": {
    "header": {
      "version": {
        "block": "11",
        "app": "1"
      },
      "chain_id": "osmosis-1",
      "height": "471386",
      "time": "2021-07-23T18:43:54.270485119Z",
      "last_block_id": {
        "hash": "RDUYMMubJjgT+mAf3zWGdLisiYS2H7VUp3GKdK6bjW0=",
        "part_set_header": {
          "total": 1,
          "hash": "nttHpLCNB/d0JinUKFfYbrN7HGdIrD76Wr3bLMNNYU4="
        }
      },
      "last_commit_hash": "8QJvwYpzlBUF8aaO9V4f6O4FaNfRr/OPCsIqzLXyOeM=",
      "data_hash": "VH3YDk7oQA/y5+KRkevI+dJOmpxggyOztBmamlMxeF4=",
      "validators_hash": "aEvCNnKB0caB5wimaxDeVAkiD601qpTOcDZwrKUYWkk=",
      "next_validators_hash": "g6CNamiPBvzzf/jZF2kuaPMbdhYD2WwB4v6YdnkHX08=",
      "consensus_hash": "SwO0ATe7CtU37Is8WxxgnQMnXUfcMCjwZ42J93U+bgM=",
      "app_hash": "nfKMx7cVHjP3xya3oUH5DzQ2hjHXTE4c7pMwhZ5ZH/M=",
      "last_results_hash": "ZPXn5Ps5YD2QB3DNas5ZwO0DkC0c5VzXA7vKmxCTI/c=",
      "evidence_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
      "proposer_address": "FqFplRqHgkfb4lj93HFjj2YG0VY="
    },
    "valset": [
      {
        "operator_address": "osmovaloper1cyw4vw20el8e7ez8080md0r8psg25n0cq98a9n",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "b77zCh/VsRgVvfGXuW4dB+Dhg4PrMWWBC5G2K/qFgiU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "3059558121410",
        "delegator_shares": "3059558121410.000000000000000000",
        "description": {
          "moniker": "Sentinel",
          "identity": "045D374A62F15B56",
          "website": "https://sentinel.co",
          "security_contact": "ironman0x7b2@protonmail.com",
          "details": "A Blockchain-based decentralized Bandwidth sharing marketplace built with Cosmos SDK/Tendermint. Winner of phase-1A and phase-2 of Game of Zones. Secure and reliable. We do a 100% refund for the downtime slash. Maintained by @ironman0x7b2"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-06-20T15:08:58.607774082Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper16q8xd335y38xk2ul67mjg27vdnrcnklt4wx6kt",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "m/n7S1ZO+kUtI7+xgCGnklo5iXgMdI9Qvif2VcvHYjg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "1592861499643",
        "delegator_shares": "1592861499643.000000000000000000",
        "description": {
          "moniker": "StakeLab",
          "identity": "F12B081334CBE0C6",
          "website": "https://www.stakelab.fr",
          "security_contact": "securite@stakelab.fr",
          "details": "Le laboratoire de délégation est le premier validateur francophone prêt à vous accompagner dans vos investissements décentralisés"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.010000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T22:28:50.805274787Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1n3mhyp9fvcmuu8l0q8qvjy07x0rql8q4d3kvts",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "vNya6ReeIsgTR3h7cNq0uDwIX9Zcn4Sx/ZAFQW9oS7U="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "1558850455211",
        "delegator_shares": "1558850455211.000000000000000000",
        "description": {
          "moniker": "0base.vc (5% fees soon)",
          "identity": "67A577430DBBCEE0",
          "website": "https://0base.vc",
          "security_contact": "0@0base.vc",
          "details": "0base.vc is a validator who doesn't trust any blockchain. We validate it by ourselves."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.000000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-20T02:20:32.280248150Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper196ax4vc0lwpxndu9dyhvca7jhxp70rmcmmarz7",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "/k7X8YEOE3HtRC91Gq9rWh3Kddk2ls5gyeUThZjq4Do="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "1487931359208",
        "delegator_shares": "1487931359208.000000000000000000",
        "description": {
          "moniker": "SG-1",
          "identity": "48608633F99D1B60",
          "website": "https://sg-1.online",
          "security_contact": "",
          "details": "SG-1 - Your validator on Osmosis. We refund downtime slashing to 100%."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-06-19T16:51:56.512017341Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1clpqr4nrk4khgkxj78fcwwh6dl3uw4ep88n0y4",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "6Nz09YGHzwWxjczG0IhK4Iv0qY2IcX0P/5KitvRXTUc="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "1237128785372",
        "delegator_shares": "1237128785372.000000000000000000",
        "description": {
          "moniker": "Cosmostation",
          "identity": "AE4C403A6E7AA1AC",
          "website": "https://www.cosmostation.io",
          "security_contact": "admin@stamper.network",
          "details": "Cosmostation validator node. Delegate your tokens and Start Earning Staking Rewards"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.200000000000000000"
          },
          "update_time": "2021-06-23T08:39:31.241687333Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1hjct6q7npsspsg3dgvzk3sdf89spmlpf6t4agt",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "dfTEd6+krWYzqsBcpqdxySq+iwh7SGcwnBO8XaW2qKY="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "1063319310558",
        "delegator_shares": "1063319310558.000000000000000000",
        "description": {
          "moniker": "Figment",
          "identity": "E5F274B870BDA01D",
          "website": "https://figment.io",
          "security_contact": "",
          "details": "Figment"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-30T18:53:36.018255903Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1thsw3n94lzxy0knhss9n554zqp4dnfzx78j7sq",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "OPsZnvTy3S90rA8kv2FQEKmFigTn5hHdAop0qNZgNTc="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "913050014583",
        "delegator_shares": "913050014583.000000000000000000",
        "description": {
          "moniker": "wosmongton",
          "identity": "7C88A757E65A5445",
          "website": "",
          "security_contact": "",
          "details": "George Wosmongton, The Founding Father supporting your right to fair farms"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1t8qckan2yrygq7kl9apwhzfalwzgc2429p8f0s",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "galx4JN7FbjF2scxKj0u7h1pwNQ/NyPu96rhIermC6s="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "897098314417",
        "delegator_shares": "897098314417.000000000000000000",
        "description": {
          "moniker": "Imperator.co",
          "identity": "0878BA6BE556C132",
          "website": "https://imperator.co",
          "security_contact": "",
          "details": "Osmosis launch contributor, check all the charts at osmosis.imperator.co"
        },
        "unbonding_height": "190641",
        "unbonding_time": "2021-07-16T21:40:59.997949285Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-20T10:09:57.405465999Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper123nfq6m8f88m4g3sky570unsnk4zng4u6mkmvq",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "oboi39aVwyap22GNEnO/xV39Kvv+/njYlKzqXDHQ/u4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "739526404329",
        "delegator_shares": "739526404329.000000000000000000",
        "description": {
          "moniker": "B-Harvest",
          "identity": "8957C5091FBF4192",
          "website": "https://bharvest.io",
          "security_contact": "",
          "details": "B-Harvest provides validation services for multiple dPoS networks and is actively engaging in decentralized governance."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-21T07:24:18.998110219Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper15urq2dtp9qce4fyc85m6upwm9xul3049wh9czc",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "3g5KC6fJ2YYRoN583mKdstLi5egwt2CkyR0yiWIRlQw="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "625189320281",
        "delegator_shares": "625189320281.000000000000000000",
        "description": {
          "moniker": "Chorus One",
          "identity": "00B79D689B7DC1CE",
          "website": "https://chorus.one",
          "security_contact": "security@chorus.one",
          "details": "Secure Osmosis and shape its future by delegating to Chorus One, a highly secure and stable validator. By delegating, you agree to the terms of service at: https://chorus.one/tos"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.075000000000000000",
            "max_rate": "0.300000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1lzhlnpahvznwfv4jmay2tgaha5kmz5qxwmj9we",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "T1wp5ELzvqVOqBm6eGYiQQWe0TaC2LV2BThHaVxh/rk="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "540719211209",
        "delegator_shares": "540719211209.000000000000000000",
        "description": {
          "moniker": "Citadel.one",
          "identity": "EBB03EB4BB4CFCA7",
          "website": "https://citadel.one",
          "security_contact": "",
          "details": "Citadel.one is a multi-asset non-custodial staking platform that lets anyone become a part of decentralized infrastructure and earn passive income. Stake with our nodes or any other validator across multiple networks in a few clicks"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.030000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1x20lytyf6zkcrv5edpkfkn8sz578qg5s833swz",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "Yd2RYHyC0I3G34ZMld5rs94e4S5L1TnWaBDaDIhhpSg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "505113032863",
        "delegator_shares": "505113032863.000000000000000000",
        "description": {
          "moniker": "Cephalopod Equipment Corp",
          "identity": "6408AA029ADBE364",
          "website": "https://cephalopod.equipment",
          "security_contact": "squad@cephalopod.equipment",
          "details": "Cephalopod Equipment - infrastructure for decentralized intelligence"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.081100000000000000",
            "max_rate": "0.420000000000000000",
            "max_change_rate": "0.011800000000000000"
          },
          "update_time": "2021-06-18T22:35:32.759789953Z"
        },
        "min_self_delegation": "100000"
      },
      {
        "operator_address": "osmovaloper1ddle9tczl87gsvmeva3c48nenyng4n56yscals",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "uqhpfK4aZ8MaTC8ZxD4xhgcgXDCnko4ed6zEdz7jDes="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "434673109930",
        "delegator_shares": "434673109930.000000000000000000",
        "description": {
          "moniker": "Witval",
          "identity": "51468B615127273A",
          "website": "https://vitwit.com",
          "security_contact": "",
          "details": "Witval is the validator arm from Vitwit. Vitwit is into software consulting and services business since 2015. We are working closely with Cosmos ecosystem since 2018. We are also building tools for the ecosystem, Aneka explorer is one of them."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.150000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper10ymws40tepmjcu3a2wuy266ddna4ktas0zuzm4",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "95HMDW5006ssh50NO61qSXj94ZqRXGe1Ta78qy7u6Z0="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "321115011028",
        "delegator_shares": "321115011028.000000000000000000",
        "description": {
          "moniker": "Chandra Station ",
          "identity": "0BC47B3228CBF46C",
          "website": "https://chandrastation.com ",
          "security_contact": "",
          "details": "100% Uptime|100% Transparency|100% Slashing Protection"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.030000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-19T05:12:21.264217152Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1grgelyng2v6v3t8z87wu3sxgt9m5s03x7uy20c",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "/Jn6UX0QVWVOquEg5RWUxMkUR8Q9dRyc7gkmv92DBJQ="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "317517126726",
        "delegator_shares": "317517126726.000000000000000000",
        "description": {
          "moniker": "iqlusion",
          "identity": "DCB176E79AE7D51F",
          "website": "https://iqlusion.io",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "445979",
        "unbonding_time": "2021-08-04T21:02:53.491663911Z",
        "commission": {
          "commission_rates": {
            "rate": "0.070000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-26T20:53:06.510058058Z"
        },
        "min_self_delegation": "10000000000"
      },
      {
        "operator_address": "osmovaloper1eh5mwu044gd5ntkkc2xgfg8247mgc56f4dlt9h",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "jjcSFNFQ8q+t/xHr1uYdI2/c1sPwAD17uKg+AyOWTe4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "303476613470",
        "delegator_shares": "303476613470.000000000000000000",
        "description": {
          "moniker": "BouBouNode",
          "identity": "",
          "website": "https://boubounode.com/",
          "security_contact": "",
          "details": "AI-based Validator. #1 AI Validator on Game of Stakes. Fairly priced. Don't trust (humans), verify. Made with BouBou love."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.061000000000000000",
            "max_rate": "0.250000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-21T06:21:29.125321078Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper14kn0kk33szpwus9nh8n87fjel8djx0y0fhtak5",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "8YRkE4YdrBPXVK9hgiLYKbrgnNUYJsGGfc8Kd/tOWuo="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "283238246304",
        "delegator_shares": "283238246304.000000000000000000",
        "description": {
          "moniker": "Forbole",
          "identity": "2861F5EE06627224",
          "website": "https://forbole.com",
          "security_contact": "",
          "details": "Co-building the Interchain"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "1.000000000000000000"
          },
          "update_time": "2021-06-28T07:47:56.552899157Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1z89utvygweg5l56fsk8ak7t6hh88fd0axx2fya",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "wB25StLxbzmD0uTiFiH6xySZd0H13kyanNUvvlUpa34="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "266052461719",
        "delegator_shares": "266052461719.000000000000000000",
        "description": {
          "moniker": "Inotel",
          "identity": "975D494265B1AC25",
          "website": "https://inotel.ro",
          "security_contact": "",
          "details": "We do staking for a living"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.300000000000000000",
            "max_change_rate": "0.300000000000000000"
          },
          "update_time": "2021-06-19T07:32:09.964458968Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1r2u5q6t6w0wssrk6l66n3t2q3dw2uqny4gj2e3",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "mmYQm2nAnrUKK5KNy31FCV9lBMl9/KgRla4vBsh/JXA="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "245160128576",
        "delegator_shares": "245160128576.000000000000000000",
        "description": {
          "moniker": "pylonvalidator",
          "identity": "5B5AB9D8FBBCEDC60979483D4F669CFF",
          "website": "https://pylonvalidator.com",
          "security_contact": "",
          "details": "\"Laws against cryptography reach only so far as a nation's border and the arm of its violence\" -Eric Hughes"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.500000000000000000"
          },
          "update_time": "2021-06-22T21:32:39.682144803Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1ehkfl7palwrh6w2hhr2yfrgrq8jetguct4ddyl",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "d3qNk2Nc5DzapRRPNJ//djb6UJ8cRlJRtYxe48QlW+g="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "238436500000",
        "delegator_shares": "238436500000.000000000000000000",
        "description": {
          "moniker": "KalpaTech",
          "identity": "B4AD06F0EB355573",
          "website": "https://kalpatech.co",
          "security_contact": "",
          "details": "KalpaTech Staking Services"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-25T08:50:49.445338413Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper15czt5nhlnvayqq37xun9s9yus0d6y26d5jws45",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "nq6+oDBXYSKLIR2vksOeGM9iVT3+RtIyWAWEUqcOoO4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "212244669317",
        "delegator_shares": "212244669317.000000000000000000",
        "description": {
          "moniker": "binary_holdings",
          "identity": "3EB2AEED134D7138",
          "website": "https://www.binary.holdings/",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.500000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "800"
      },
      {
        "operator_address": "osmovaloper102ruvpv2srmunfffxavttxnhezln6fncrdjd27",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "nCD7wpBgaHm5CTMvsWro4W9ODmvzvLzphYoIAmSW5R8="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "207380052094",
        "delegator_shares": "207380052094.000000000000000000",
        "description": {
          "moniker": "ztake.org",
          "identity": "09A303A2C724C59",
          "website": "https://ztake.org",
          "security_contact": "",
          "details": "Support reliable independent validator"
        },
        "unbonding_height": "419411",
        "unbonding_time": "2021-08-02T20:40:15.806435244Z",
        "commission": {
          "commission_rates": {
            "rate": "0.070000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "1.000000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper12zwq8pcmmgwsl95rueqsf65avfg5zcj047ucw6",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "00VGuZFAeenn/Zf8rYqNowyUnTzPo3zcu/LKptGInHU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "150600792453",
        "delegator_shares": "150600792453.000000000000000000",
        "description": {
          "moniker": "OmniFlix Network",
          "identity": "535BF8D68742ACED",
          "website": "https://omniflix.nework",
          "security_contact": "",
          "details": "OmniFlix is a p2p network for creators, curators and their sovereign communities to mint, manage and monetize assets. Developed Cosmic Compass, winner of the Best Custom Zone category in Cosmos (GOZ) and run nodes on networks that share our vision."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.500000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-20T00:30:45.929127720Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1v5y0tg0jllvxf5c3afml8s3awue0ymjusax93d",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "SIinCf4GKmuqG1u01EZCy+PNY/f+8sQjPwynG2akkg8="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "150529261132",
        "delegator_shares": "150529261132.000000000000000000",
        "description": {
          "moniker": "Zero Knowledge Validator (ZKV)",
          "identity": "3E38E52A12F94561",
          "website": "https://zkvalidator.com",
          "security_contact": "security@chorus.one",
          "details": "Zero Knowledge Validator: Stake \u0026 Support ZKP Research \u0026 Privacy Tech"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.070000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-07-23T11:23:37.128801628Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1u6jr0pztvsjpvx77rfzmtw49xwzu9kas05lk04",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "ZH2TWvk7bd8MLjDHLItvOJUv+ElgfItJO4V4XnXedVw="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "135848617970",
        "delegator_shares": "135848617970.000000000000000000",
        "description": {
          "moniker": "Cros-nest",
          "identity": "5F1D6AC7EA588676",
          "website": "https://cros-nest.com",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-20T14:14:54.836138716Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1uxn3xw4xzyvu3xka0fr7krsfzyp8af3eyyurzh",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "I9sfSptsfnwEOIY8fRT3GezAiOs7H8in/mB0lkR/rE4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "133395083415",
        "delegator_shares": "133395083415.000000000000000000",
        "description": {
          "moniker": "Skystar Capital",
          "identity": "",
          "website": "",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-01T11:16:28.736569908Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1l7hln0l79erqaw6jdfdwx0hkfmj3dp27gr77ja",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "PA+RvSpVTcAWBD/+nxXltHLaOdaUzbrpNGfjMak/usU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "128930289823",
        "delegator_shares": "128930289823.000000000000000000",
        "description": {
          "moniker": "Smart Stake - osmosis.smartstake.io",
          "identity": "DD06F013A474ACA3",
          "website": "www.smartstake.io",
          "security_contact": "t.me/SmartStake",
          "details": "Performance analytics dashboard for community @ osmosis.smartstake.io. Transparent \u0026 professional staking validator with automated monitoring tools. Dedicated support @ t.me/SmartStake"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-19T15:02:53.692107491Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1kgddca7qj96z0qcxr2c45z73cfl0c75pf37k8l",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "U9J4uq5i6zKKHkReUKsMetJSvyCNrCaqLICyIoqXUEw="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "110202527864",
        "delegator_shares": "110202527864.000000000000000000",
        "description": {
          "moniker": "ChainLayer",
          "identity": "AD3CDBC91802F94A",
          "website": "https://www.chainlayer.io",
          "security_contact": "",
          "details": "Secure and reliable validator. TG: https://t.me/chainlayer"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-19T18:37:19.464632333Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1083svrca4t350mphfv9x45wq9asrs60c6rv0j5",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "d0hOXvDMHeFnCEnWbSb3qr/q/IT+SFSgzPn1f686xic="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "106624632783",
        "delegator_shares": "106624632783.000000000000000000",
        "description": {
          "moniker": "Notional",
          "identity": "3804A3D13B6CB379",
          "website": "https://ipfs.io/ipfs/QmYChijKhdkUCUmBBfA4LsyuB4y5quDjfjhwgw9uu17pFC",
          "security_contact": "",
          "details": "Open Source, Organic edge validation and relaying"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.090000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-18T08:25:10.960997367Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper13x77yexvf6qexfjg9czp6jhpv7vpjdwwpuclc8",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "eAeAUKnPvGxb6hEQ4VxzVO+6ZWvyqn1OjtzBjqzBKxQ="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "99506247970",
        "delegator_shares": "99506247970.000000000000000000",
        "description": {
          "moniker": "blockscape",
          "identity": "C46C8329BB5F48D8",
          "website": "https://blockscape.network/",
          "security_contact": "",
          "details": "By delegating, you confirm that you are aware of the risk of slashing and that M-Way Solutions GmbH is not liable for any potential damages to your investment."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-19T18:06:37.043063648Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1pjmngrwcsatsuyy8m3qrunaun67sr9x74vvvdk",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "s5PS5N/ichjOxyDowRowgjqZChE3Q/qEweTIK3Sn/X4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "96927207300",
        "delegator_shares": "96927207300.000000000000000000",
        "description": {
          "moniker": "Cypher Core",
          "identity": "5CCA4F526B9F85DA",
          "website": "https://cyphercore.io/",
          "security_contact": "",
          "details": "We are devoted to bring financial freedom to everyone"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1e8238v24qccht9mqc2w0r4luq462yxttfpaeam",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "zzZWZApRCaCX1YIzZQAxcHR5JX+r/zWJv/nT7mek4t0="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "85460454955",
        "delegator_shares": "85460454955.000000000000000000",
        "description": {
          "moniker": "POSTHUMAN ∞ DVS",
          "identity": "8A9FC930E1A980D6",
          "website": "https://github.com/Distributed-Validators-Synctems",
          "security_contact": "",
          "details": "full-time enthusiast"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-27T20:24:54.804088297Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1ceurjtrgxf0hfd4r0hez6fenevazenfhym7s4q",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "PjAFZo2+kvndcVlGzAIUQMYotuom4hUvaUOC9CSsR2k="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "80141890316",
        "delegator_shares": "80141890316.000000000000000000",
        "description": {
          "moniker": "Mandragora",
          "identity": "DFEAAB98E8D0975B",
          "website": "",
          "security_contact": "",
          "details": "Stake 'n Chill Out!"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-23T00:43:58.472086875Z"
        },
        "min_self_delegation": "1000000"
      },
      {
        "operator_address": "osmovaloper1000ya26q2cmh399q4c5aaacd9lmmdqp9cwpvl4",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "jSNHqQh5u31X5TqJ1dKVdFgeFFihqjP+D1Utvo5lezg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "78878336199",
        "delegator_shares": "78878336199.000000000000000000",
        "description": {
          "moniker": "Staking Fund",
          "identity": "805F39B20E881861",
          "website": "https://staking.fund",
          "security_contact": "",
          "details": "We've been actively engaging in the validating role for numerous novel Proof-of-Stake protocols since early 2018 and proving our commitment to secure decentralized blockchain networks with high availability and zero slashing."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.000000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.200000000000000000"
          },
          "update_time": "2021-06-24T02:37:22.882596422Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1pgl5usqpelz3a4c04g6t3yyvlrc9yseylmtcnt",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "cf+XPmAgtXNK/B4GuzvEpJdhwaFFo1fn69KPhaPXQOU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "78859634935",
        "delegator_shares": "78859634935.000000000000000000",
        "description": {
          "moniker": "Secure Secrets",
          "identity": "C5C28A947096C28A",
          "website": "https://www.securesecrets.org",
          "security_contact": "",
          "details": "Focusing Privacy and Transparency"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.010000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-19T19:18:05.304123449Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1d4pn8f3q5dn0g4uwr2egh5mgzrnf7fhh53tuu4",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "8Bt4TLbDFtJ6lfxPsSfln39zaaCEccnudGwZocEV7sA="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "77672173029",
        "delegator_shares": "77672173029.000000000000000000",
        "description": {
          "moniker": "Validating Chaos",
          "identity": "1ECD13F96C55C0CD",
          "website": "",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-02T06:08:18.114918491Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1zfcmwh56kmz4wqqg2t8pxrm228dx2c6hhzhtx7",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "4kc9OXQQiGyNRTXlrCxltrzgWLUpV504J4RhQ5YLpMU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "77608000000",
        "delegator_shares": "77608000000.000000000000000000",
        "description": {
          "moniker": "SkyNet | Validators",
          "identity": "1510797E867F484E",
          "website": "https://skynet.paullovette.com",
          "security_contact": "skynet@paullovette.com",
          "details": "SkyNet | Validators embraces Blockchain and Defi technology to make a better world.  We secure and support the Osmosis blockchain and community."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-06-24T05:49:28.697598978Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1juczud9nep06t0khghvm643hf9usw45r3jxhxn",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "YzA8yyHSvymqCplrh0r50T1DRhlNscmxK38Q4T0U4oI="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "71574557319",
        "delegator_shares": "71574557319.000000000000000000",
        "description": {
          "moniker": "ITA Stakers",
          "identity": "06E7A073C20F48EA",
          "website": "https://itastakers.com",
          "security_contact": "",
          "details": "Italian Stakers Community"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-28T09:19:34.386216512Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1vzcxcv9v4cymknqyj8a7qqp45apt7kyrvuwhxr",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "NSdK+xoO3oHloLV25r/yFpRNWfpfjH7KZuY+kY5MeBc="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "71473255322",
        "delegator_shares": "71473255322.000000000000000000",
        "description": {
          "moniker": "BasBlock",
          "identity": "E0A6A3980E464A66",
          "website": "",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.500000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-21T12:40:56.977026094Z"
        },
        "min_self_delegation": "1000000"
      },
      {
        "operator_address": "osmovaloper1gy0nyn2hscxxayj2pdyu8axmfvv75nnvhc079s",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "LOUaDTS0YZS70HjBwr9ARerXMXbXxFPsSxQ47uNmAvA="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "70639694598",
        "delegator_shares": "70639694598.000000000000000000",
        "description": {
          "moniker": "Provalidator",
          "identity": "3A7D5C9B0B88BEA1",
          "website": "https://provalidator.com",
          "security_contact": "",
          "details": "Supporting Blockchain Infrastructure"
        },
        "unbonding_height": "205711",
        "unbonding_time": "2021-07-18T00:16:50.480233876Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-20T02:27:53.416869133Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1c9ye54e3pzwm3e0zpdlel6pnavrj9qqv5qgdz9",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "euJubdwp0WNjX9xPVRYcMaBPd0NCP9mKKMcUVtPhqQk="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "67978324815",
        "delegator_shares": "67978324815.000000000000000000",
        "description": {
          "moniker": "StakeWithUs",
          "identity": "609F83752053AD57",
          "website": "https://StakeWith.Us",
          "security_contact": "",
          "details": "Secured Staking Made Easy. Put Your Crypto to Work - Hassle Free. Disclaimer: Delegators should understand that delegation comes with slashing risk. By delegating to StakeWithUs Pte Ltd, you acknowledge that StakeWithUs Pte Ltd is not liable for any losses on your investment."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-22T02:52:36.491248354Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1s33zct2zhhaf60x4a90cpe9yquw99jj0x8t08z",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "4qEagTlC7uJC7w84zvir7tnlKz30AL0DMJ3q0vmozSg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "65464750901",
        "delegator_shares": "65464750901.000000000000000000",
        "description": {
          "moniker": "dimi 🦙",
          "identity": "94FEC9A766EF8D04",
          "website": "https://dimi.sh",
          "security_contact": "dimiandre@gmail.com",
          "details": "The real Lama Validator"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-19T18:05:38.832547240Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper103s29j7dyc3utljce802m38p2kk89taj0r2hxa",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "HG3qcg4yqCzlBG5NqYtZaRKtVp5WKcxnETQZr0GTJmY="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "62609547867",
        "delegator_shares": "62609547867.000000000000000000",
        "description": {
          "moniker": "commercio.network",
          "identity": "ADBDB0178E4441BE",
          "website": "https://commercio.network",
          "security_contact": "",
          "details": "The Documents Blockchain"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.080000000000000000",
            "max_rate": "0.300000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-07-22T15:31:30.292863166Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1wke3ev9ja6rxsngld75r3vppcpet94xxwcmk4v",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "FxemoA8I2BR+58ooPCP5Dge1SU5npBWSPb9zccCtgMs="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "58666020222",
        "delegator_shares": "58666020222.000000000000000000",
        "description": {
          "moniker": "Cat Boss",
          "identity": "059BCF656623D0BE",
          "website": "",
          "security_contact": "",
          "details": "Trust me, I am a Cat"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "1.000000000000000000"
          },
          "update_time": "2021-06-23T14:07:03.463570601Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1j0vaeh27t4rll7zhmarwcuq8xtrmvqhu6m87mz",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "EvMa6sPpEJQz8bqalGMESOvDKq2YYXPo9+ib3qq+mz8="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "56918620684",
        "delegator_shares": "56918620684.000000000000000000",
        "description": {
          "moniker": "Chainflow",
          "identity": "103DCE407C9F1D13",
          "website": "https://chainflow.io/staking",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-26T11:46:52.142738782Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9ua8h4lf",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "DCJh91Oqh9P76gYC71pFSzd0cxnvqRDyZ6/PLFZfgLE="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "54954607944",
        "delegator_shares": "54954607944.000000000000000000",
        "description": {
          "moniker": "stakefish",
          "identity": "90B597A673FC950E",
          "website": "https://stake.fish",
          "security_contact": "",
          "details": "We are the leading staking service provider for blockchain projects. Join our community to help secure networks and earn rewards. We know staking."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1y0us8xvsvfvqkk9c6nt5cfyu5au5tww24nrlnx",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "LCCt4CBSrkDvSDHWMB7vepYZ4swoq1A0rOpmRM4+Y6c="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "50265551929",
        "delegator_shares": "50265551929.000000000000000000",
        "description": {
          "moniker": "Swiss Staking",
          "identity": "165F85FC0194320D",
          "website": "https://swiss-staking.ch",
          "security_contact": "",
          "details": "Experienced validator based in Switzerland. We offer a highly secure and stable staking infrastructure."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.250000000000000000",
            "max_change_rate": "0.025000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1et77usu8q2hargvyusl4qzryev8x8t9weceqyk",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "XgIu1lX318LSJylroCLzRjAliJoOI+nye7r3vG8S7Fc="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "48564151005",
        "delegator_shares": "48564151005.000000000000000000",
        "description": {
          "moniker": "Stargaze",
          "identity": "9203983F91296B66",
          "website": "https://stargaze.fi",
          "security_contact": "",
          "details": "Let the liquidity flow"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.500000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1pfh243e50apq0zut00vyhd3sqek0jthc8wvxws",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "f5DzEhtQbnmXE/WZQsX+I8RljPdEU0u0ncVGtniFyEM="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "47123538272",
        "delegator_shares": "47123538272.000000000000000000",
        "description": {
          "moniker": "Bi23 Labs",
          "identity": "EB3470949B3E89E2",
          "website": "https://bi23.com",
          "security_contact": "",
          "details": "Bi23 Labs is a trusted POS infrastructure provider and validator to comfortably stake your coins and earn rewards with Celo,Cosmos,IRISnet,Terra,KAVA,Solana"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-07-13T03:53:47.197385372Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1feh2keupglep6mvxf5c96eulh3puujjryj2h8v",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "dOk2pUzVZHq8JDdyNLOyHH9k+ry/Xvxl7d/vZivYE0I="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "44588064976",
        "delegator_shares": "44588064976.000000000000000000",
        "description": {
          "moniker": "Simply Staking",
          "identity": "F74595D6D5D568A2",
          "website": "https://simply-vc.com.mt/",
          "security_contact": "",
          "details": "Simply Staking runs highly reliable and secure infrastructure in our own datacentre in Malta, built with the aim of supporting the growth of the blockchain ecosystem."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-19T17:25:02.323439708Z"
        },
        "min_self_delegation": "1000000"
      },
      {
        "operator_address": "osmovaloper1jfqaxtg8g9ad80crl0rg2reg7mrje8qp258xsk",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "P7NEbDFMi9oR0jh8OkC96bE13goVUUdxxAct3yqcD8U="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "31455000000",
        "delegator_shares": "31455000000.000000000000000000",
        "description": {
          "moniker": "jetlife",
          "identity": "325A1EA86C1E3B12",
          "website": "https://jetlifestaking.org/",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-02T21:39:44.140137483Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper12rzd5qr2wmpseypvkjl0spusts0eruw2g35lkn",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "p7vsAmhAKlKj7MOYEx5Ht0N0wGua0On35WzW7q3eFh4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "28709771984",
        "delegator_shares": "28709771984.000000000000000000",
        "description": {
          "moniker": "Stakecito",
          "identity": "D16E26E5C8154E17",
          "website": "",
          "security_contact": "",
          "details": "Expect Chaos"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-30T07:49:26.632148547Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper17mggn4znyeyg25wd7498qxl7r2jhgue8td054x",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "OdVpTfLCvPyBLeE6jNxesgy3Hg1IiA+165lSusZDgLs="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "25210297444",
        "delegator_shares": "25210297444.000000000000000000",
        "description": {
          "moniker": "01node",
          "identity": "7BDD4C2E94392626",
          "website": "https://01node.com",
          "security_contact": "",
          "details": "01node Professional Staking Services for Cosmos, Iris, Terra, Solana, Kava, Polkadot, Skale"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.600000000000000000",
            "max_change_rate": "0.500000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1n24z8k7w8lxej8xe7uvpn3qc654n43c05unug6",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "Jr9+ftRFLf+Ixli7oIhaxV7E/3S4rEvbRSZi2ySNDIw="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "22483233730",
        "delegator_shares": "22483233730.000000000000000000",
        "description": {
          "moniker": "SolidStake",
          "identity": "A15273DFFD11E62E",
          "website": "https://solidstake.io",
          "security_contact": "",
          "details": "Securing the Decentralised Future"
        },
        "unbonding_height": "354012",
        "unbonding_time": "2021-07-28T23:24:04.857230155Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-19T08:09:54.284632701Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper133va8vgy0ygfjmvzy64pgay5w6lg4239nagwnk",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "ynIQ9OkSgNvGXqMDhGt+20i2Olzr5PQ99AgbBQAG1XU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "19702054073",
        "delegator_shares": "19702054073.000000000000000000",
        "description": {
          "moniker": "天照☀",
          "identity": "3912AE47B45446D7",
          "website": "",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-23T02:37:14.445650330Z"
        },
        "min_self_delegation": "1000000"
      },
      {
        "operator_address": "osmovaloper19j2hd230c3hw6ds843yu8akc0xgvdvyu4arf82",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "b2DVNBLrIFXkQ19WXmbzlUKh+sQrBaSY1s8hSPkFFMY="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "18585046987",
        "delegator_shares": "18585046987.000000000000000000",
        "description": {
          "moniker": "syncnode",
          "identity": "F422F328C14AFBFA",
          "website": "wallet.syncnode.ro",
          "security_contact": "",
          "details": "email: g@syncnode.ro || Telegram Channel: https://t.me/syncnodeValidator || Blog: https://medium.com/syncnode-validator"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.900000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-06-21T17:38:48.663932227Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1de86y3vdphqe904xyh6eh0fftw0283tnezqmug",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "gYhsVodd2ucRL6wxdVblaVypaUdkEEGxdS3jWfdQ5iU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "17097765408",
        "delegator_shares": "17097765408.000000000000000000",
        "description": {
          "moniker": "Alex (Bambarello) Validator",
          "identity": "A713F5C07C453731",
          "website": "https://keybase.io/bambarello",
          "security_contact": "https://keybase.io/bambarello",
          "details": "Experienced and Reliable Proof-of-Stake Validator, Dedicated and Diversed Hosting. Mainnet Validator for the Graph, Oasis, Solana, Certik, Bitsong, Sentinel, Regen, Centrifuge."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.000000000000000000",
            "max_rate": "0.300000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-19T21:20:45.348567054Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1jxv0u20scum4trha72c7ltfgfqef6nscqx0u46",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "RldYpYQCj26AeZq6tH6uC0ySp9vMYrwIuJowyWiVaTg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "16301000000",
        "delegator_shares": "16301000000.000000000000000000",
        "description": {
          "moniker": "ping",
          "identity": "6783E9F948541962",
          "website": "https://look.ping.pub",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "405937",
        "unbonding_time": "2021-08-01T20:19:39.377814269Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-30T23:27:37.942507163Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1m73mgwn3cm2e8x9a9axa0kw8nqz8a492vg4hp4",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "PwTrMSmKul2l8xfN3RmAryyVv2+PvwG9XpZxf4Vc4uM="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "15183439106",
        "delegator_shares": "15183439106.000000000000000000",
        "description": {
          "moniker": "#decentralizehk",
          "identity": "436039F82528A43A",
          "website": "https://decentralizehk.org/",
          "security_contact": "",
          "details": "Building Consensus among Hong Kong People"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.051000000000000000",
            "max_rate": "0.721000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-07-04T16:29:29.175409254Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1p8u2dda9ulfleg8wyqr7akvkywxfrhhvwgpv7q",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "uoEWo2QBO00/kiUxaMoxD1dM9IfAqW81FFS++O8NXtI="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "15101794022",
        "delegator_shares": "15101794022.000000000000000000",
        "description": {
          "moniker": "WhisperNode",
          "identity": "9C7571030BEF5157",
          "website": "https://www.whispernode.com",
          "security_contact": "",
          "details": "We love Osmossis"
        },
        "unbonding_height": "205815",
        "unbonding_time": "2021-07-18T00:27:44.029366743Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-25T03:38:27.487358620Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper13tk45jkxgf7w0nxquup3suwaz2tx483xe832ge",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "Z3A0fCkaKR09vvDT7x18I2IczxvVpezksPpsYyK1n2Y="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "14611969232",
        "delegator_shares": "14611969232.000000000000000000",
        "description": {
          "moniker": "bro_n_bro",
          "identity": "A57DAB9B09C7215D",
          "website": "https://osmosis.zone",
          "security_contact": "",
          "details": "We love Freedom"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.300000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-21T10:54:11.727926937Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1whdjuaspv2y3v02pjywncs7xmzpeudkfg9ukjw",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "veN3KDVhCBgMuyY0Q+/I6to7F72yHcgABmGhtJHKYwg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "13302570000",
        "delegator_shares": "13302570000.000000000000000000",
        "description": {
          "moniker": "mp20",
          "identity": "D0D9D1C2AEB79C5B",
          "website": "https://mp20.net/",
          "security_contact": "",
          "details": "staking@mp20.net"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1dpwecca283m0c20dmsfzztp4mma8h9ajxwvx52",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "B2C0sAwsppI3SMu+sUWlAa7VhvhWQ+zF/unbkYY3OoU="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "12274014996",
        "delegator_shares": "12274014996.000000000000000000",
        "description": {
          "moniker": "samurai",
          "identity": "",
          "website": "",
          "security_contact": "",
          "details": "Liberty Peace Free Markets !tyranny"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.010000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-05T18:30:09.592775919Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper13gvqalusnmapjgp3e6gnk9q832qv3g3ug7lcuh",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "HMvX3nk0Lt+a0ddiVKs3X8WW2cBffbCQaLjsKtfPTNc="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "12126000000",
        "delegator_shares": "12126000000.000000000000000000",
        "description": {
          "moniker": "SRU-OSMO",
          "identity": "",
          "website": "http://stake-r-us.com",
          "security_contact": "",
          "details": "Please visit stake-r-us.com"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.060000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1rhjvcesshc833dqwdlwskwcjz08gcnxq73dnqc",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "bDnUx51OClMaBjtynE3RLmCb0mlPlkaWLVZDBmrC1Uc="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "11801720649",
        "delegator_shares": "11801720649.000000000000000000",
        "description": {
          "moniker": "artifact",
          "identity": "70C162B0473634FD",
          "website": "",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-30T03:35:14.223278939Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1w4x44ek799hvg97x0mfwu6gg5dww2r8fdpqql4",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "GPlrdJ5z/EbXbUKc3xIcQn+LOU1UaF4kyk7CfvephJg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "11606753550",
        "delegator_shares": "11606753550.000000000000000000",
        "description": {
          "moniker": "AUDIT.one",
          "identity": "5736C325251A8046",
          "website": "https://audit.one",
          "security_contact": "",
          "details": "Validators of today, Auditors of tomorrow"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.070000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-22T10:24:28.957551817Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper14amduhjazqhwtkhm6kutdcy4ux5zazf5k803tq",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "Quejs9DfnCsgGElTXxULslAU06c4bzyrxDSTtDZ6VO8="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "11086180698",
        "delegator_shares": "11086180698.000000000000000000",
        "description": {
          "moniker": "Little_Cryptoman",
          "identity": "C89513E385AC6C88",
          "website": "",
          "security_contact": "",
          "details": "Little Cryptomen love BCNA and OSMO! Tokens self-bonded from delegator address linked to my Ledger."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.020000000000000000"
          },
          "update_time": "2021-06-20T14:25:54.974149195Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1ltch7uamdats4lphq6r6gl8ftyzhwzgvlnd38v",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "xOXWstC3+9gx+CR/VX7YHehECZZQuW6XpsGqzf76Ap4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "10936000000",
        "delegator_shares": "10936000000.000000000000000000",
        "description": {
          "moniker": "Delicious",
          "identity": "7AAAA066B64C3034",
          "website": "",
          "security_contact": "deliciousmail@protonmail.com",
          "details": "Yummy yummy stake."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-08T05:52:32.377110841Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1d7vvatc4rqwqf99z90rqf6dlc484p02q9pc0em",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "ASkOhXx65phXuzuAN6nCk/OHr2ArDDVyUKWS3Toim0E="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "10740164991",
        "delegator_shares": "10740164991.000000000000000000",
        "description": {
          "moniker": "c29r3",
          "identity": "9BDCB96F2AB4EAA9",
          "website": "http://github.com/c29r3",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.000000000000000000",
            "max_rate": "0.120000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-20T13:35:24.821512420Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1z0sh4s80u99l6y9d3vfy582p8jejeeu6tcucs2",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "XYCvhdcOlm/5g1gHSO+7/hAuCQ0kZPU7XKvc1C5e1U4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "9908000000",
        "delegator_shares": "9908000000.000000000000000000",
        "description": {
          "moniker": "[ block pane ]",
          "identity": "D75509198CE782A6",
          "website": "https://blockpane.com",
          "security_contact": "",
          "details": "🔥 Fast, datacenter-hosted, bare-metal validators 🔥"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-11T06:14:07.247832684Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1j7euyj85fv2jugejrktj540emh9353ltnz0lvc",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "omxGXH85dkmGesQ5GwQy0YKZ6qvm+/2e8fxoahlKorQ="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "9346000008",
        "delegator_shares": "9346000008.000000000000000000",
        "description": {
          "moniker": "Komikuri",
          "identity": "",
          "website": "",
          "security_contact": "",
          "details": "Komikuri have slashing guarantee and warranty"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-21T18:39:13.984592059Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper16gm3cvhluf9xfurkx9qgxq7pldvd479l0j6zms",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "mM9Nn9Xh0kDOybSEEVXypyN6N+dxd3wiRe9BS5BXU5M="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "8592231587",
        "delegator_shares": "8592231587.000000000000000000",
        "description": {
          "moniker": "SpacePotato",
          "identity": "B41FCF161C4B971B",
          "website": "https://spacepotayto.carrd.co/",
          "security_contact": "",
          "details": "We have a secure and reliable validator setup that adheres to best practices. Uptime and rewards generation for you is our top priority!"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1wtv0kp6ydt03edd8kyr5arr4f3yc52vpxy9qu6",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "B+g+ptuCXkZr20htGv/RiDQEnqXi28Ujwo2DFUceNys="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "8000000000",
        "delegator_shares": "8000000000.000000000000000000",
        "description": {
          "moniker": "kytzu",
          "identity": "909A480D5643CCC5",
          "website": "https://kytzu.com",
          "security_contact": "",
          "details": "Kytzu Validator"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1de7qx00pz2j6gn9k88ntxxylelkazfk3llxw6r",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "VO80I/0OSlZ1gjdxXAO53/e/bz9jBweFLVjaEcCH2qQ="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "7422012000",
        "delegator_shares": "7422012000.000000000000000000",
        "description": {
          "moniker": "Cosmic Validator",
          "identity": "FF4B91B50B71CEDA",
          "website": "https://cosmicvalidator.com",
          "security_contact": "",
          "details": "A reliable, passionate and service oriented validator. We are long term trusted community members; and have received delegation from both Tendermint and the Interchain Foundation (ICF) as a reward for our continuous support and effort."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.500000000000000000",
            "max_change_rate": "0.200000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1mwg6lt5eumhpnyx6gzwxaesj7lymyh8erj3xqq",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "yiLcmbSoLlYPaCgeMiob5EzL+bVpUB9sDt2Me7kEL64="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "7365828041",
        "delegator_shares": "7365828041.000000000000000000",
        "description": {
          "moniker": "hydrogen18",
          "identity": "",
          "website": "",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.020000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-23T23:56:17.139035857Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1rmdzkje2eh34fgerdq75lwgzy44u43erqdnnh3",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "hMROAIF7q2V/OSDikgrt1q+bgabGl3LH2XfG3KQllng="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "7144000000",
        "delegator_shares": "7144000000.000000000000000000",
        "description": {
          "moniker": "Staky.io",
          "identity": "C7D6DBE2CB576363",
          "website": "staky.io",
          "security_contact": "hello@staky.io",
          "details": "Staky.io is a user-centric staking as a service platform that will get you the best staking experience on the market. Governance, Telegram bots, Analytics \u0026 Rewards tracking.. Everything is here!"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-06-23T13:41:35.544464195Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1nskmzc8cqf8m9vwrxvnggtw4vw9cwx54jtfmuj",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "LWqk0T+Vbgheja80lL56hOq/yURJXKOpZ9ANdQLWa+c="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "6818082876",
        "delegator_shares": "6818082876.000000000000000000",
        "description": {
          "moniker": "rhinostake",
          "identity": "59C635D1CD02FEEC",
          "website": "",
          "security_contact": "support@rhinostake.com",
          "details": "Providing secure validation infrastructure for dPoS networks"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-07-15T02:30:05.691966235Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper17h2x3j7u44qkrq0sk8ul0r2qr440rwgjp38f93",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "C88upbte468p6VnUYxFStVugaV+Fzi6C3YckK4lwpgI="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "6161650809",
        "delegator_shares": "6161650809.000000000000000000",
        "description": {
          "moniker": "FreshOSMO.com",
          "identity": "63575EE3F0F9FAFC",
          "website": "https://FreshOSMO.com",
          "security_contact": "",
          "details": "FreshOSMO.com runs on bare metal in a SSAE16 SOC2 certified Tier 3 datacenter with geographically distributed private sentry nodes, YubiHSM2 hardware protected keys, with 24/7 monitoring, alerting, and analytics."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-07-08T06:29:55.867217302Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper14k3fu4yqlu3ddwe342mr2q0dz465vxyprvj3cw",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "tVGqxPAvgXlqaAK62+XDcH96VsXz+QGhwrN0jmNG9AE="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "5770005768",
        "delegator_shares": "5770005768.000000000000000000",
        "description": {
          "moniker": "BwareLabs",
          "identity": "E83A08BEEE7A70BD",
          "website": "bwarelabs.com",
          "security_contact": "flavian@bwarelabs.com",
          "details": "Guaranteed availability and up-time backed by a professional blockchain infrastructure team."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-07-21T10:28:46.137283721Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1pxphtfhqnx9ny27d53z4052e3r76e7qq495ehm",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "bU1m3aDQwxIpqFYubjZIhSvYCia1MQSfTaaXJBg38nw="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "5496426979",
        "delegator_shares": "5496426979.000000000000000000",
        "description": {
          "moniker": " BlockNgine.io",
          "identity": "5DCE5E12052FB516",
          "website": "https://blockngine.io",
          "security_contact": "",
          "details": "Low-fee | Secure | Reliable validator backed by enterprise-grade infrastructure distributed across multiple regions and hardware providers."
        },
        "unbonding_height": "271826",
        "unbonding_time": "2021-07-22T21:12:43.032800436Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-21T19:22:37.614436876Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper10rsjwrcj7maz4f0l8dh3smjd022n9ya8ycuyg7",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "rqsz3Pa5/AbpCuQEJUl/mh6oz2B/mbknJOxdVTeAkUM="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "5379061121",
        "delegator_shares": "5379061121.000000000000000000",
        "description": {
          "moniker": "AC Validator 🚀",
          "identity": "1C5486895FD0090C",
          "website": "https://www.acvalidator.net",
          "security_contact": "",
          "details": "We Skyrocket your Stake! 🚀 Highly Available Enterprise Infrastructure with multiple nodes around the world"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.500000000000000000",
            "max_change_rate": "0.500000000000000000"
          },
          "update_time": "2021-06-19T12:32:11.008990109Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1rcp29q3hpd246n6qak7jluqep4v006cd8q9sme",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "2cuVqUIby057xZ8ndvosCk3fuovzBoJGLMyB74JsZNg="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "5093057021",
        "delegator_shares": "5093057021.000000000000000000",
        "description": {
          "moniker": "in3s.com",
          "identity": "0CE19EE3E4BA48E5",
          "website": "https://in3s.com",
          "security_contact": "",
          "details": "Cosmos Hub \u0026 Starname (IOV) validator since block 1 of both chains."
        },
        "unbonding_height": "356061",
        "unbonding_time": "2021-07-29T03:02:20.691889571Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "1.000000000000000000"
          },
          "update_time": "2021-06-19T19:15:57.424091654Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1qskagvpanuyjuyc2cey7pme7glly02vmxzaj66",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "8EdCUsdmmvchHJUriByV+Fhizyo3npr9UpNIP1J/Dzw="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4815828145",
        "delegator_shares": "4815828145.000000000000000000",
        "description": {
          "moniker": "[SG] X Staking",
          "identity": "E42BBDAE71EC2A82",
          "website": "https://www.xstaking.sg",
          "security_contact": "security@xstaking.sg",
          "details": "Cross Chain validator from Singapore. Never been jailed before, and any slashing will be fully reimbursed."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.050000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-21T11:32:16.936262006Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1ualhu3fjgg77g485gmyswkq3w0dp7gysdcdgw2",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "A/NlPxYDzI3cgNZkl8VSYHXBe90BCv1mgqd5X88eYY0="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4608000000",
        "delegator_shares": "4608000000.000000000000000000",
        "description": {
          "moniker": "stake.systems",
          "identity": "7F82E4F0CAA26298",
          "website": "https://stake.systems",
          "security_contact": "",
          "details": "building infrastructure to support awesome projects running in the blockchain landscape"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper16vcwhpzd6hfsm5n0fxz9880lvkgjnpwyxqncyt",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "4XPXd9B7Nv6SJWvvGuThMPKb3mhFiAy8/lNvqlIXiBo="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4454496677",
        "delegator_shares": "4454496677.000000000000000000",
        "description": {
          "moniker": "EZStaking.io",
          "identity": "1534523421A364DB",
          "website": "https://ezstaking.io",
          "security_contact": "contact@ezstaking.io",
          "details": "EZStaking.io Node Validator"
        },
        "unbonding_height": "363977",
        "unbonding_time": "2021-07-29T17:06:18.865210570Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.030000000000000000"
          },
          "update_time": "2021-06-23T16:35:41.821193558Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper10valuxrrt5guceq5eu8tcf566vkegylksajxdk",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "sQ5oiPr9Bdhd59g3JY56vKmZZuYvDDvfTkm+mpaUB5U="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4390000000",
        "delegator_shares": "4390000000.000000000000000000",
        "description": {
          "moniker": "UbikCapital",
          "identity": "8265DEAF50B61DF7",
          "website": "https://ubik.capital",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "190710",
        "unbonding_time": "2021-07-16T21:48:23.938235810Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-07-23T16:47:36.323294098Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper16j5hsdrcaa6950ks0rf944rgmncukl74cs7yw6",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "gJvkPLiHo2+12JJGprhwVsABASXzXO/dSWtnk0ywNB4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4366829417",
        "delegator_shares": "4366829417.000000000000000000",
        "description": {
          "moniker": "SmartNodes",
          "identity": "D372724899D1EDC8",
          "website": "https://smartnodes.co.uk",
          "security_contact": "",
          "details": "Staking Service Provider"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-19T20:35:27.395152783Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1273zsxhxd5dlgcr2zjf5x25275hjcp3udgkz9z",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "jG3YOPzLxEOmfC2xN8qxqtF+A1IGLvn19AhTvhSmor8="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4249642336",
        "delegator_shares": "4249642336.000000000000000000",
        "description": {
          "moniker": "Coverlet",
          "identity": "11A2797A6DD3873D",
          "website": "https://coverlet.io/",
          "security_contact": "",
          "details": "Coverlet staking. Go you covered!"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-21T21:19:44.110346232Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper14suzzfw7rkz7uke8w9lhttnvewku7sd06djndg",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "K0P8fWZcajey9vaJfXPA0UYJezBseS8+CUpWPMqGx1Q="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4217050000",
        "delegator_shares": "4217050000.000000000000000000",
        "description": {
          "moniker": "StakeThat",
          "identity": "281A045D9C119F34",
          "website": "stakethat.io",
          "security_contact": "",
          "details": "Secure, Reliable and Experienced PoS node operator"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-07-15T16:41:21.951090733Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1crg24mgm8l3qe7ej9l8t7tsttp40aaas4xmkhx",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "xP04sdg7xwZqPpOYaoTIjOhCgQjXK5adgghQLpkP6MM="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "4158301726",
        "delegator_shares": "4158301726.000000000000000000",
        "description": {
          "moniker": "ushakov",
          "identity": "2E3A8285E6B547B2",
          "website": "https://stake2.me",
          "security_contact": "",
          "details": "Individual staking service"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.050000000000000000"
          },
          "update_time": "2021-06-21T19:07:37.254586108Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1h2c47vd943scjlfum6yc5frvu2l279lwjep5d6",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "cNRUrOt4I5rJcsAVhUOiBvOLYRSFx4dso22bcUEhtnI="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "3806356988",
        "delegator_shares": "3806356988.000000000000000000",
        "description": {
          "moniker": "CryptoCrew Validators ✅",
          "identity": "9AE70F9E3EDA8956",
          "website": "https://cryptocrew.cc",
          "security_contact": "",
          "details": "CryptoCrew Validator Service for Osmosis Network (osmosis-1). Reliable \u0026 secured by Sentry-Node architecture. Based in Europe. t.me/cryptocrew_osmo_validator"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.080000000000000000"
          },
          "update_time": "2021-07-15T14:51:38.517579143Z"
        },
        "min_self_delegation": "3000000"
      },
      {
        "operator_address": "osmovaloper1vkfmegxrsveefn2wmudh7apxuzu2n77654ad62",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "FVTcxQeQQs4CexEmSLVkMFrxpYSKv7to7rOdlw+7ej0="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "3786598641",
        "delegator_shares": "3786598641.000000000000000000",
        "description": {
          "moniker": "blitz",
          "identity": "",
          "website": "https://osmosis.zone",
          "security_contact": "",
          "details": "Cranking Osmois"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.060000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-07-04T19:59:10.658756630Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper19yvvl0cz068qu4sg7v2pxln0cqm8nt7756fa6m",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "TyshR1RRrhqT8QwVsCw9DAPaLKymxWKqo2w833D/Dc4="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "3550300000",
        "delegator_shares": "3550300000.000000000000000000",
        "description": {
          "moniker": "Noderunners",
          "identity": "812E82D12FEA3493",
          "website": "http://noderunners.biz",
          "security_contact": "info@noderunners.biz",
          "details": "Noderunners is a professional validator in POS networks. We have a huge node running experience, reliable soft and hardware. Our commissions are always low, our support to delegators is always full. Stake and get rewards with us!"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-21T23:18:04.122810483Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1gf5lfrstxcv8764x35360tmf62d0gewzzsw3ze",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "g1acgtYuvxV442kbmjsUfJZg9nF4ReoSXKK8MfAYqIo="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "3383700000",
        "delegator_shares": "3383700000.000000000000000000",
        "description": {
          "moniker": "Staker Space",
          "identity": "59850BC3A3C5F039",
          "website": "https://staker.space",
          "security_contact": "hello@staker.space",
          "details": "secure independent validator ❤️  osmosis"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-19T19:20:12.172919569Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1t29pjdugzyetxaehf62a8x7hhq4u0v4un9mxg5",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "qM8LJ7CG8X6EsagOJcsitlnTfyZRob1V8KuMoBuM7Nk="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "3365905339",
        "delegator_shares": "3365905339.000000000000000000",
        "description": {
          "moniker": "Staketab",
          "identity": "D55266E648F3F70B",
          "website": "https://staketab.com",
          "security_contact": "",
          "details": "Staking Provider. Secure and non-custodial staking."
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.100000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-20T11:58:00.867279552Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper16jn3383fn4v4vuuvgclr3q7rumeglw8kdq6e48",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "5wqRcwcA0u1Cjor+ywAL9AC9QjD3vDP5pM9pUZhVK24="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "3184400000",
        "delegator_shares": "3184400000.000000000000000000",
        "description": {
          "moniker": "SOLAR Validator",
          "identity": "6257A55EA42BA680",
          "website": "https://validator.solar",
          "security_contact": "security@validator.solar",
          "details": "Reliable \u0026 secure validator from Estonia. Built on top of government-grade infrastructure by SOLAR Labs."
        },
        "unbonding_height": "287788",
        "unbonding_time": "2021-07-24T01:29:32.586603215Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "1.000000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-22T10:39:28.001242920Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1hz22q65qyjqwml4s2ny86u7fpdxgz5euxwppgt",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "FPaM0nlLXyQnN1EEHTxnlEFdld2pAKuYut6KrjYnme0="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "2884000000",
        "delegator_shares": "2884000000.000000000000000000",
        "description": {
          "moniker": "みんみん.net",
          "identity": "693B9EB0FB22B623",
          "website": "",
          "security_contact": "",
          "details": "Based in Japan.We strive for stable node management.---------日本を拠点にしています。安定したノード運営を心がけています。"
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.100000000000000000"
          },
          "update_time": "2021-06-24T00:47:21.608115663Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper14g9r3rwutmpc80e8kqnnq8uvdcptdusc0nm3gc",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "jHyba/XE8/qBdd+0xiGDPvTHoywcFAY8nMQ0CHg4Xtk="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "2785000000",
        "delegator_shares": "2785000000.000000000000000000",
        "description": {
          "moniker": "Tuzem",
          "identity": "4006E2C214ACAA86",
          "website": "",
          "security_contact": "",
          "details": "In Node We Trust"
        },
        "unbonding_height": "366925",
        "unbonding_time": "2021-07-29T22:23:31.931089226Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-21T21:45:20.519142191Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1cyepvt5kayjzsa76ft98ud8mrpvh3acxv83pwv",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "m2mRSBFHx6k1UBe1vLk1lpdzNBAke61NBg0J6ggDXG8="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "2500000000",
        "delegator_shares": "2500000000.000000000000000000",
        "description": {
          "moniker": "wave",
          "identity": "E5378C73E7776BA8",
          "website": "https://bity.hns.to/",
          "security_contact": "",
          "details": ""
        },
        "unbonding_height": "0",
        "unbonding_time": "1970-01-01T00:00:00Z",
        "commission": {
          "commission_rates": {
            "rate": "0.100000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-18T17:00:00Z"
        },
        "min_self_delegation": "1"
      },
      {
        "operator_address": "osmovaloper1z5xyynz9ewuf044uaweswldut34z34z3cwpt7y",
        "consensus_pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "gWt6q+82LR2nMWHo7k2q5LGLsXzuZGajZQVFyqkzygY="
        },
        "jailed": false,
        "status": "BOND_STATUS_BONDED",
        "tokens": "1873048132",
        "delegator_shares": "1873048132.000000000000000000",
        "description": {
          "moniker": "Stakely.io",
          "identity": "55A5F88B4ED52D3E",
          "website": "https://stakely.io",
          "security_contact": "admin@stakely.io",
          "details": ""
        },
        "unbonding_height": "420782",
        "unbonding_time": "2021-08-02T23:09:53.979417592Z",
        "commission": {
          "commission_rates": {
            "rate": "0.050000000000000000",
            "max_rate": "0.200000000000000000",
            "max_change_rate": "0.010000000000000000"
          },
          "update_time": "2021-06-22T07:30:12.913246822Z"
        },
        "min_self_delegation": "1"
      }
    ]
  }
}`
