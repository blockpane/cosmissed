package missed

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse_minSignatures(t *testing.T) {
	m := minSignatures{}
	err := json.Unmarshal([]byte(testSigs), &m)
	if err != nil {
		t.Error(err)
		return
	}
	proposer, ts, signers := m.parse()
	if proposer == "" {
		t.Error("did not parse proposer")
	}
	if signers == nil || len(signers) == 0 {
		t.Error("did not parse signers")
	}
	if ts < 1 {
		t.Error("bad timestamp")
	}
	fmt.Println(proposer, "\n", signers)

}

// http://192.168.255.252:26657/commit?height=471386

const testSigs = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "signed_header": {
      "header": {
        "version": {
          "block": "11",
          "app": "1"
        },
        "chain_id": "osmosis-1",
        "height": "471386",
        "time": "2021-07-23T18:43:54.270485119Z",
        "last_block_id": {
          "hash": "44351830CB9B263813FA601FDF358674B8AC8984B61FB554A7718A74AE9B8D6D",
          "parts": {
            "total": 1,
            "hash": "9EDB47A4B08D07F7742629D42857D86EB37B1C6748AC3EFA5ABDDB2CC34D614E"
          }
        },
        "last_commit_hash": "F1026FC18A73941505F1A68EF55E1FE8EE0568D7D1AFF38F0AC22ACCB5F239E3",
        "data_hash": "547DD80E4EE8400FF2E7E29191EBC8F9D24E9A9C608323B3B4199A9A5331785E",
        "validators_hash": "684BC2367281D1C681E708A66B10DE5409220FAD35AA94CE703670ACA5185A49",
        "next_validators_hash": "83A08D6A688F06FCF37FF8D917692E68F31B761603D96C01E2FE987679075F4F",
        "consensus_hash": "4B03B40137BB0AD537EC8B3C5B1C609D03275D47DC3028F0678D89F7753E6E03",
        "app_hash": "9DF28CC7B7151E33F7C726B7A141F90F34368631D74C4E1CEE9330859E591FF3",
        "last_results_hash": "64F5E7E4FB39603D900770CD6ACE59C0ED03902D1CE55CD703BBCA9B109323F7",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "16A169951A878247DBE258FDDC71638F6606D156"
      },
      "commit": {
        "height": "471386",
        "round": 0,
        "block_id": {
          "hash": "4275FDDA420C5BEA88DDA8F6112F534FF9E6A2CACC4CD84FC98BF699437FD308",
          "parts": {
            "total": 1,
            "hash": "C46640E29C30E060CF020384361DC31BBECA7AFD3C6A1655D8259071E2A92449"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "16A169951A878247DBE258FDDC71638F6606D156",
            "timestamp": "2021-07-23T18:44:00.676133333Z",
            "signature": "jT8Zgf52xzbZky2AzxYUzKGRRNZTo/uV+FVRGGeGDnkTz4+MD4W8PLqM1CABwJoSp93+nZ442Xg19/U2RdgGCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "9E7CAE009EFFF4D163F3FB8781A07B25C2B10B32",
            "timestamp": "2021-07-23T18:44:00.61338236Z",
            "signature": "aHHJVggmsD00lELRC/6i9St+R3ZXrlARevV10DJ6C689W8GWd6zHHrozqZwnb6xvDC0jtKLX5f+ld+Xn7e4hAA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "76F706AE73A8251652BC72CB801E4294E2135AFB",
            "timestamp": "2021-07-23T18:44:00.646044412Z",
            "signature": "Oj0fzB8ZpejSefT80wvTxCtFxl+Yk7eGJk7iG3aDHKq7L5YqZ7h0lQh4oBpWKuhobMYFHR+I4yxKZQ03NlV+DQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "6239A498C22DF3EC3FB0CA2F96D15535F6F3387A",
            "timestamp": "2021-07-23T18:44:00.718540733Z",
            "signature": "HwudCNyfZBW3KJPX8nkctqATjQkFJrNY4Ft0pk+Oef6VDxppVRj2DTcX2d5bNyGTaU7buLr1RerF5yPWaiVuBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "CB5A63B91E8F4EE8DB935942CBE25724636479E0",
            "timestamp": "2021-07-23T18:44:00.577099142Z",
            "signature": "kZ8AuDtKSlPdVAwSi4zgEP7JcbXbYkszuIMODGB8oOU47jwBaOuUjanUU0P1BHNEt5zMy7XyShqP/F91LdjLDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "9D0281786872D3BBE53C58FBECA118D86FA82177",
            "timestamp": "2021-07-23T18:44:00.628358549Z",
            "signature": "6Ou/Vb2dYxVTnxTis0os1C1WPMH+/QYbuveHzj4BswFtOGpIDVQrywQgKLPOC1pAEhwFe2vWy2Cyk+oBTcf+DA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "844290531EE59B40FEEFDE5259857368BF7119EC",
            "timestamp": "2021-07-23T18:44:00.665834187Z",
            "signature": "p09YV5Z2FvRIxzIlwRx5ChuSo3mXwWLnHkeWcr+qWsyqSGN/vZSqCRdjSJK8D7qJ0q1A1zYQB76OjW4emSRpCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "1B002B6EBEB8653C721301B1B56472B1B4DE7247",
            "timestamp": "2021-07-23T18:44:00.658899817Z",
            "signature": "Li09FldPz+JboHc6WXaoillhPY168ZKqcrGZcmCnP1jcgzW3DG3bBvwH/a6I7MwN+zoF5ZA+VfMZBgKgYyJTBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F194DD4A8AD83323C3E9C2A93DB25F049621C7B4",
            "timestamp": "2021-07-23T18:44:00.578030714Z",
            "signature": "P+W4rTLHiVMiFFhwwj72Exx9tznOvfsvC2WivyZ6NIHOLP8UIR/49fvqxCdszH7F0cI397+fgMCkHNO3nuCnDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E08FBA0FE999707D1496BAAB743EAB27784DC1C5",
            "timestamp": "2021-07-23T18:44:00.72697219Z",
            "signature": "QilDZ2TetAUNBMtjRPnsKjfR4puENKnYhGWcohUHwzkWNQdt/oLvuJk8VsrqjX0K/N/2ls2ZiyGTP48kAN0+Dg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "99063B919404B6950A79A6A31E370378FE07020D",
            "timestamp": "2021-07-23T18:44:00.632669841Z",
            "signature": "OO1uPp/bwK9mxidSz0pnhElh7C0e2v+21Vm7W1UYQ9LlmS6PIAty2pCFed44ZrmwOL/n1YwDGbudjgzadLQFBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A16E480524D636B2DA2AD18483327C2E10A5E8A0",
            "timestamp": "2021-07-23T18:44:00.797683447Z",
            "signature": "KgRZgAzeyMFHgBEHegUAGvjAA5G260ziDyW10T6/qkd5GMyfDakNezdi6XCzhKn2yEH/Jzw7R5Mzl0WnbEA5CA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F233E036248A36FC73C154FFA79261BCBDC4BB76",
            "timestamp": "2021-07-23T18:44:00.655429891Z",
            "signature": "gv6/fZ9v6DNg2hPWVVW5ob0kMOfH1EQDBbuwizN44y7RfGPRU9zBRsbGy9hSy6XK9td4Yj2A2UGJYXzCaf20BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "04C83AA20F7563BBCBCF6AA150EF6B0C81808DAA",
            "timestamp": "2021-07-23T18:44:00.807180488Z",
            "signature": "bl0iipYd6tixEA9OCMQpr2y4rQkQNHX4UWc2EmmAlKpzz7lH2sZhRJ3iaATsr8QYL4h6tas7x8wk0rk1zqx/Ag=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "20EFE186DA91A00AC7F042CD6CB6A1E882C583C7",
            "timestamp": "2021-07-23T18:43:55.270485119Z",
            "signature": "PbI6DL+QKjVDN+2Mz/qynqdc0R381aeszhI2jue90f8r+WooGnCNGTi5Y/StRlyLCHxRU7Nw5ricBKmoJ1vXCg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F3F55DA24BB47DA60B0FB71EC1A9C9274BCEEDB2",
            "timestamp": "2021-07-23T18:44:00.634804594Z",
            "signature": "k6CGdtv6leIVZY5KnuZmgPf/uXrXE5GhhhUo5EC3N6R3axgHFGQbJ79xuLHp1BOrCdXcO6ELrMSjY2Czbyd2CA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D8A6C54C54A236D4843BA566520BA03F60F09E35",
            "timestamp": "2021-07-23T18:44:00.622897345Z",
            "signature": "tyswLd4oV7VcbCOPfUIyuz8QdhTNPkKTi+e12QyIDJvXWGkTRU90psdw0ek456mw6iDANocmavDqu43dEVA8BA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "66B69666EBF776E7EBCBE197ABA466A712E27076",
            "timestamp": "2021-07-23T18:44:00.610207019Z",
            "signature": "r31aUdzXod5O293mxcdHUGUyllOCpN8Obur20oCEAQNyL0SQ5GBq1bydNiFJnYPk0qe74WLnNO+w0G8Q0dOHDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "138FD9AB7ABE0BAED14CA7D41D885B78052A4AA1",
            "timestamp": "2021-07-23T18:44:00.618574947Z",
            "signature": "qP3lDvIYYlS/NDecnIQ8COoWNZ2u7//TMgREpo8Xif2T+sidiP1pidXeInRDl8eK8DKXbC42WfWEOA4h850YCg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E06DADEB413829558F7C95339FFB61499C5A1BB9",
            "timestamp": "2021-07-23T18:44:00.746690966Z",
            "signature": "+RQUF+ohjfo/5QCTonlkfKgiUdn2Rf6PJ/Pcmrj/2lNy0d8SKwnmNGtRScNu3AvQpq0tf0cNPrKliFj71aImAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "2022FE8CC49E48630C76160E11A880459219D244",
            "timestamp": "2021-07-23T18:44:00.622301122Z",
            "signature": "wy8cInFIgVx3qq3P8cNA64ITgyP1LLCcX9z6+R+PxVJZqI4eP+7jAN+ELGm70F1GX734G52cjyHerWgiYcedCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "2F89D7D3D1E1478F88EF3AD8AAD76A88189F6124",
            "timestamp": "2021-07-23T18:44:00.603720382Z",
            "signature": "1b1DJXQTO7ZyxiiFW1OlP2ip0EDlaG8EnE3RYyNMTLEQADBRLvMXzxRKgv/TwwCz18BE1yKAIRcBKkevgwUXDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A06B5B682B425AD206A35CAF246FD70DD098E506",
            "timestamp": "2021-07-23T18:44:00.632772553Z",
            "signature": "/RRSfaFMRnKYjyH94+uK2NfZ9Vlxt+3TWfv+jOKgKAG8mrXAN01vf7AGXaJUugDUbqjbLdsF0ZgPIXnj6bnmBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E191E654D06B9F721568BB945B2EB51DDC1C8FDC",
            "timestamp": "2021-07-23T18:44:00.687421506Z",
            "signature": "etzy0YycPPA937LVC0/WqjrAsxNTzkqHk5yAINb6ErVRrCatuaNln7Ee3D/gSh0VUzjkqI28n60GVLhzJddvDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "97AFE45395B74E784C88D45E5CCA2995019FAE08",
            "timestamp": "2021-07-23T18:44:00.596339956Z",
            "signature": "wRvs9e4No/8BYRRGdbYzv0dtqaWcdQR3Rv3OY6XfLXs94PGHZs8Hrq+Gj0gFFPTecC2NlQ2n/7D7A2bB/kmNCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "C02F531D9BBBA4907511EF2680421CE714A11E3B",
            "timestamp": "2021-07-23T18:44:00.829440644Z",
            "signature": "EWwCMed6pVNSkyuG3TV+iaiGSh3uybR08oTwGwnJUu/QH72CO1HuM2bd0hBTjUrBMJgb91YoogkMNrTOF8clDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95B002DE67707313D123D06492F1A7A58478E546",
            "timestamp": "2021-07-23T18:44:00.634692461Z",
            "signature": "JQgAvOKjzB3hl36rjdBT/GXJTtwuL1qk//+9osVR9FmigP4F458TpNkhfbUc/01MmmE1MUPkC89+tH/KhrCjBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E5CBA199E045E7036711D814E57E2B66C3CC0391",
            "timestamp": "2021-07-23T18:44:00.605920109Z",
            "signature": "BimChSiSGhuPjI5QV1QC0BofwSFpGAtL3NFmMMXPp0oUTzKMhjZst3+5nmy+NM6KVA2NY/PFV/WsajDOR92SBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D9EC9739CCCF051A05861ACB8A2218A9A4756390",
            "timestamp": "2021-07-23T18:44:00.703227007Z",
            "signature": "lQ9UEREUtXineVnr58PDV/c53xFdU2q8JGu2FwqkrL4/Bj75eu8/BUxbXcp5rlKRPAKlnBheEB/1mcrgUIwkCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "9496535A8F2945BDB60572015D2D6F721AB6FED9",
            "timestamp": "2021-07-23T18:44:00.58352196Z",
            "signature": "2+bY7D7d2zXA4L6BnpbW6nthI1cCeglBzC/jdXBsafVXeWWNHhq207lKRy4IsT87Axl/fNr70orfJnobNeITCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E242DB2CB929D6F44A1A2FE485CC7D3F620FFAEB",
            "timestamp": "2021-07-23T18:44:00.675337984Z",
            "signature": "4SsKGQH/XpZn8U6XEREMp0TPOAO79RPXs1h10YBZaOhg3lzZnLyP5aamXHgPrI2i70r4j1VEt60gi2pc6GYEAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7EDB006522610C58283E30644A14F27BCC0D32ED",
            "timestamp": "2021-07-23T18:44:00.604494157Z",
            "signature": "32XeOUGV/RiUOcpQruonTk4lmd3iSX2yfEV/wDKHaZgjY34UnDJ3R1vKM3o4eWpyKa3IoYlatd7HqZtBBXf2DA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "9F5BC08868DF50484F24520CB87D75D43F6DC23B",
            "timestamp": "2021-07-23T18:44:00.61463332Z",
            "signature": "3a7koppqbu69CxZuHcDhmcJKDIbUrw2GC/iDbAukrvYutzLTbo0D+ngRDlO+jGiUrLQldatWarR+CAbGWYqUDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "0960EF3FD58FE7DBB8F20FC98269D3B840451603",
            "timestamp": "2021-07-23T18:44:00.628256993Z",
            "signature": "Oi8Rvrcq9UPvcYL46nTHudEoqJ2VhPXmPWCT/yDnWQD94YA9HZfNw0k9peJCJa6RemXjhHbH3B3P6ObI9Mf5BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "20658BF40ED48ED01A2D087C7FF7874F21A56333",
            "timestamp": "2021-07-23T18:44:00.561179425Z",
            "signature": "jQtmX6VTcFc3f+2YY1xCOiM7gAb7LvaVGhtjNaO7otjWx6gbKlnlWihkHe+mgzR6nkhLiQaWvNn3xcfT6bwKAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "EF4F7A6EAC883B6E491E5466E6BEC764C1FB99C3",
            "timestamp": "2021-07-23T18:44:00.663851683Z",
            "signature": "EV8kSr2WvaNh6bfval3T1E9bI5BqgPfRFfJvSxsJKue7vLcqBgHgWTjImUICDMgWB8X4xUC4Jk4D7mw4RtQdCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "800FF47897D19DCDF7E20CF308C993E304E6EA93",
            "timestamp": "2021-07-23T18:44:00.641366715Z",
            "signature": "9Ig0ZAVl9wsDS/TYtdcIMWfu69Wtwyq/gYgggUmyO2/YmVqnfd3Dprap2mMIlbnrtt+ENGkScGtBiHIzLWVnBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7C5AA87E5203C66EA35C64262F576EDD29BAD980",
            "timestamp": "2021-07-23T18:44:00.613430966Z",
            "signature": "QMFe9oBWEm6xQGtItthPNSGbJiq8RFbQuGVJooEf4/ktLE4HdRyXNGPA4eefGcFNZLOYuK6u83XonAibVxWuAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "DDFF1B21F85EB0A3300360E3A3380CE32DFE5484",
            "timestamp": "2021-07-23T18:44:00.588952741Z",
            "signature": "bvols5SpVSYqTjMfdM9V0uh/ST059bKstYeJ0erCmEAMCNG16AOopxS4E2PCLs8FZoXOj5/M8XqwPt6QlGRMBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "06F45C36FCB957E55D923A6D4E905C2D715115AD",
            "timestamp": "2021-07-23T18:44:00.792374424Z",
            "signature": "Yx532bfZ0dDT4LgK/dwwWsly5QSwN7LBNbP7LbQZ3raqGtv5YDp8weqpo6ZidyDCgErrHV5hn2nfvsIpEhcDBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3E88E7C54F64642A98B2E1DDD5BDBA48794F06C7",
            "timestamp": "2021-07-23T18:44:00.742306795Z",
            "signature": "zwKMVYH0rKXHOVuT527xtzNpGM9F9f39PFpWb/1sa8J6BbU8DMMyv0AdUJotGPOE5Hu/tp9iJKU950NkBM87Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "2F4D6730476407195AF3C1BF438B61CB6D785B95",
            "timestamp": "2021-07-23T18:44:00.587081086Z",
            "signature": "VFsCReUnJuG7hQAO+syszJQGXDsb96MSqgNi3Qm3JAqA37tWqBLoDTDbygBxp/vl2/ebcstGfzBqVN7TMfjrCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "C02ACBA7653AC3782750B53D03A672E191F00361",
            "timestamp": "2021-07-23T18:44:00.621665078Z",
            "signature": "4qo4ZYHq0YBdTzEkgZmllbNGI6kN8NfHQqffXXSd6+7as5c+kN6iM6LCF8DFTUOcM4LYkupWkvfX9tlJhaLaCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "736BBB1AB1BE4467B5E2138E2217D7385E38CBC5",
            "timestamp": "2021-07-23T18:44:00.795855908Z",
            "signature": "kINz75Wg10XhL2x2FqLpOYiQSEmhvEn4HxxUJ6BikQxo4voFO9Nu0aK9uJEER2CmhjRUKkZ+A0/dC5fC4hsADQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "99938495407C09B343562AAEC3AB551A5C246232",
            "timestamp": "2021-07-23T18:44:00.62125067Z",
            "signature": "01DnoWgv9pciA93O9OyyGIRaNWo9DkXb4Mc59PhjMPhRGyrsDWbaFJSzDqzGHnQODJaPaJ8pu9eADgliKjR2Ag=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "71DF8D9879C20563A4E2ABEDA95CD1FC57DBF6AA",
            "timestamp": "2021-07-23T18:44:00.702335315Z",
            "signature": "wOOsqccEo8Mo1KFw5SrqiJoiQhVqh8wOhng3PTZXhjkiBOlHGtvWsJNSxlU3wFp4GD28f5xIL3pw3yvrSzcTBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "5E809E91EAB69D385784D191140E9C8CF6DD1037",
            "timestamp": "2021-07-23T18:44:00.608423014Z",
            "signature": "d68UVqGZWMzCfg1UMYsW6FJok8SPYVBu5tQaY5wGNNE49NsidQyPnDBmJej2DQaGeQ1YRj2d/YKTItaoCkg9Bw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "901FD122CC512EF13DE8E1A3D7953BFDDC0786D6",
            "timestamp": "2021-07-23T18:44:00.76205Z",
            "signature": "rTn6omLyRia/Z7iKLZKYw/j/LwbKVdgc8wbV/m2DYy2BzFkV/9uIMQD9PKZqsb+L4mw9HdQDdRWor6Ab3WT9Dg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "0C55C18D9C6689B8CD6F775DD7EC46331FF662E2",
            "timestamp": "2021-07-23T18:44:00.706617364Z",
            "signature": "7AGdpzY/bc56d4R8W7F2J88oPdb0N1ytvz1ktw2T/tGTfNorzSWgoMUEBGoZ1U4MwTnrmIDzbz/59OjuMhT5CQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "DA4AF19A378C09B54C26C3467CB0ADF889292954",
            "timestamp": "2021-07-23T18:44:00.630955964Z",
            "signature": "tW/koF14BxFfaRKLKQ4dQT/pL1Av0UXJ8kkMqSMRDnannRnUaxSTHmdNyi8upPMsxGGtLBRL13ta0Hyc03/pDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3B6C59535788FE666D7EFCC50DC0586D509B3089",
            "timestamp": "2021-07-23T18:44:00.701713442Z",
            "signature": "RzXB5H0Keks84SCDzxvnMQLYTUFKg8XP88Ub3IU0sESpozY+3jX3OjzRT9aIXV7W6e3Fl2RSRYWsHuomfXm6Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D24B7A32413338C2AA26FC0016D91FBE73BB5EAE",
            "timestamp": "2021-07-23T18:44:00.665557Z",
            "signature": "COEfncYMUcqfmCOFSJWy3lvfTrPS8jqc0gzJcnVGrwjdRKDvim9b64Q0yuxdD4ZVe0wb8kJA939r80GQ3sBpBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "03C016AB7EC32D9F8D77AFDB191FBF53EA08D917",
            "timestamp": "2021-07-23T18:44:00.677048356Z",
            "signature": "hMje5PrPZVMCAaRplFdDcGQfGVBkc2EsJvhyTmUd98IOzxGvxdn87JTHybrKT5PHl8qRzFIxf+TB/zuMr8W8Dw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "AB7751B2E741D4608B2C07AC93347279B31E81EC",
            "timestamp": "2021-07-23T18:44:00.583040221Z",
            "signature": "Gr8dyAbTXBCMqZbNJwuDjo9S2zdXSB/AVkZsMACAc0287OK8o0s3Bwp0VGCvcl2hlZFnpBTr7yxMu2SXsBwvDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "2BC2A0C3ABAF936778030C004585B4750A862C1D",
            "timestamp": "2021-07-23T18:44:00.662745727Z",
            "signature": "hYZ1lwhQFiVjb6WFgQ6gYbLP/BQUAMYql9dfy26GFfOoMlGyWIp6leHP5yNZQD4pINEjPJk+WQWOmuZic7WQCg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "22BA59AC2918AFA4C1B56D3E6F86083E470CD8CB",
            "timestamp": "2021-07-23T18:44:00.660418784Z",
            "signature": "ZeZfK3KaP+zsnN9mdyBe4xYODG+at7d2w05sk3d3QQcOS79Qz3IVpjLi18uAjmSum3qKiGQQ6kZUqyk4EfSPBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A15694A887FA6688B65C28C78369E86C7BEF9B92",
            "timestamp": "2021-07-23T18:44:00.597068017Z",
            "signature": "g9C8v2tkf67xhwnKnIBBoLOjvETfuxLYKFKIptOSaMsUdWvWvkM0ouzRtHIgCfAsZlEgh3W7ni7HSIZ74eY2AQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "966FD89B1DB51535F2D898CF9B0F14DA374EFB96",
            "timestamp": "2021-07-23T18:44:00.609717018Z",
            "signature": "R1iGPKvznyZVdxEqkVMwWqj2XObACAgvhWFztunjDYMNmhBqYub8KPrQOWjCrwwlg7v2L7sH4l4UPhP62EXsDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "9CBF2EFFD5570B3A9A41346244757CDA3E18D401",
            "timestamp": "2021-07-23T18:44:00.635337771Z",
            "signature": "ef6cuMe5T8o6u1CxcT/zHPiKbG97vYmbHBT2nLjI8oVTo2z5RSFMN3FMj8oxFN+68cFRzaYHkEPg4aA4LnjrCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F65237239119095D18AC790B727CA4E8F78A1728",
            "timestamp": "2021-07-23T18:44:00.632695972Z",
            "signature": "+At1T2Pt+BmtD7p+zERHZS4Y6umSKgm0+VvF6BxEHPNaw8M5vVnVWOC39tp86UgImxLWnE76FJmpbcn3bCYsAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "63481F6DCAAF733D2FC953A335C2200EE190862C",
            "timestamp": "2021-07-23T18:44:00.692058916Z",
            "signature": "lCsNq/obssV459zifZIzWQi5xvtdiDntk6FKtbgrtHn3y3onCgATvgALiHVMnk2CIkDneyb8/Q1smiSA+La+Cw=="
          },
          {
            "block_id_flag": 1,
            "validator_address": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "signature": null
          },
          {
            "block_id_flag": 2,
            "validator_address": "7D5B402E18AF250EFA95CD9402D2D821DFACB876",
            "timestamp": "2021-07-23T18:44:00.60066645Z",
            "signature": "/yQTY+6hQQWYD4wiciXjDpgF0zrR75wY5Jys4ue1OVksZZ2qFVx+yybv6rmnUvQuwxW8th5dFZWlsUiTEDtGBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3D49630F3AA68CBCC53301E0CE857E7686BB3C80",
            "timestamp": "2021-07-23T18:44:00.699206483Z",
            "signature": "4Z4e6B7DE0bObHu0r6k93p1xU/BXG2gQplVMtgHgCn2tcZJKyozs98IDnVm6xJckoJhJjfPyvsSv4T1jd3VlBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E80D1F5519A5B3C9D290D3EA314FA05564535C1A",
            "timestamp": "2021-07-23T18:44:00.748703013Z",
            "signature": "uazcN3xA1ADEpBLEXCy09b7m7/zuL5EX2WJ00j5DE/q5EEeRzC787Lsbv8EKngmG/Ve4QfrTe4r5/xnFb5CRBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "273F72EE55987AFA771B27D370FA131F608B83AC",
            "timestamp": "2021-07-23T18:44:00.644707278Z",
            "signature": "hOtlxM/LtRWV4MQU0jR3iH9LbSDO4kM5tjoSZLDxSAvGLSZl4P1zD7THVpkkP4Mmn9D/v4539ogejfVnER+BCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7D53D76F2DB86BE30A9B26CADEA69078531AB9BB",
            "timestamp": "2021-07-23T18:44:00.628642169Z",
            "signature": "Q9aMl9znjhpqBDZOEi/qb2kPeiCvv2RgVjR822yMZdFdq6KyVWso8fy+imMJYh7xCXbyb9MgjcB/XbCGPxc6BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "B346EF6CDDD5BFC7FE71832FE97D297E98D5C4A4",
            "timestamp": "2021-07-23T18:44:00.581490769Z",
            "signature": "ljjc5tZRxGJOqT/eMRVaOWudxH84+LTov4SBO+MCrKQkeiXKe7CdBOTiYLMx1wncZ7gDg/2WGKzMXWFR25gLAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "C935258AA0A0D5C97B037238BA9FE83C23B25D65",
            "timestamp": "2021-07-23T18:44:00.621985765Z",
            "signature": "GDhJdJM+IyCf4VJ7RYz+6DEbpf4O5xjXIUomArbyLuBPfFBstWBnJkdeJt9tJWMdQwqDHocTSyqPCp0g6siUDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "000A5959634B4296E4DE536481DE00A8A0EB9A58",
            "timestamp": "2021-07-23T18:44:00.675854202Z",
            "signature": "FwBMvH59GwLdZsOX6U8hqrs8kHTTbsMlWh/se9gmX3vvp/w12TVzd34l/fIK6/ftM0WHABccyQ/kVqssvrSoAA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E20004515311B205618FAD504FB529A3DEEE2E71",
            "timestamp": "2021-07-23T18:44:00.582963023Z",
            "signature": "SjhecbU/JbOKzFBBMRc5eq1h6MkHkf68nhhNeeSixBIcOW81OXxYEQFaIjtaJS5szNXXHJwvh70Q90nXm/hLDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A572FC790EDB3653F0A82DCC92C865975CA4925D",
            "timestamp": "2021-07-23T18:44:00.762541428Z",
            "signature": "8kTVdAxVhtRKfsvhnkPa4UTjT1l1hmv0RRp3svuxzpfa73QQjt9XU2ucWDiRWFj2aoV5XySJdeb7jTdWXdoWDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "6A0DBE5D0B92E571465AF52E2B77665838C2E51C",
            "timestamp": "2021-07-23T18:44:00.613492744Z",
            "signature": "Um9l8q6YqKlrs58CCIrTRi75w0QD8sAke/TI+3cXUp/5haYrLDTLyzwgRF12u5zOFt6Fey+5VCYYHipsY5MBBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D48911753EDB5AD78DFD921F5D92637713B38368",
            "timestamp": "2021-07-23T18:44:00.673474778Z",
            "signature": "XcFYbO4e836tQ+MAeIEkz77TBSH+o+66dq5C/0HxbG2RH5649dNYKwEyPo65h4KpRogwE4uTDonYXn34xarFAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4AC2B026A4B3992292610EC8588644BF3D5B34FF",
            "timestamp": "2021-07-23T18:44:00.68755285Z",
            "signature": "jV5h3TArqhXv5v2axxmhp4wdhe61Udx2Cj3Q5X5b+eypA80S5pPPM7Ch6a4V5btcmLY8H+NBEWfrb9eYWvucBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "2C0D20F60F1EDA345AD50E7419F29EDEB17ADBD5",
            "timestamp": "2021-07-23T18:44:00.596849709Z",
            "signature": "owVplEqRM26dn7yOBAkc8t4oV5XiDYH90HzVymrlcicGG2MPC5Y8X4GYPrXJB5GHfFVCmHkzpgW7BjC+m6H6BA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4EEBC2AAAEF1C058AB7355982594D5AC1A31E454",
            "timestamp": "2021-07-23T18:44:00.580980615Z",
            "signature": "njFRz3MnL4DRpU84l74y5TEQkX73dhmXhYdV0SWtkLe3eUde+FkJuUc/aPm+Gfh0fzlb+WMGnsqTzEpVA6W2CQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "41B543E91479A95CD5CA9F109C26DFAC149126FA",
            "timestamp": "2021-07-23T18:44:00.626585489Z",
            "signature": "6hAE3dCoFvMJHjOXIFPBBPtPb8XzqnUPI904hG/Y7aPVZdljDV1WWPp9oeZ/lD7k66EYchMdEdogPwra9pcBAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "41A79767A7E4914CAB7E4246D9D8FEB5E96DEE97",
            "timestamp": "2021-07-23T18:44:00.738168977Z",
            "signature": "KVumRbyN4sSAMFlHl1YtuNL1LninkQT91agmecqdi7OcM6rlZkNp0iAijepGLy1wrU8gRUqkKDoi93jUi7odDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "68A393C7ED496871150C0A7CAD0CAC09B8E458FB",
            "timestamp": "2021-07-23T18:44:00.638011989Z",
            "signature": "sSq2NHAskeGp+qMNCE3P2STiwMhVX6rilsdneCxVw2834YlaE0aiTmZV0Ji6sz9uSJzudJyfjZ+QzUVeaAlJDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4146FD7A1AB8B861B7018978BCD13D2D1FA63EBE",
            "timestamp": "2021-07-23T18:44:00.612062901Z",
            "signature": "1u+ENKWGaO1Gr+7x4YHVEVJ9f3kuIQScZ58B2kvPBzMLqAyxERWgDStflgh96BjZhVMkxXgdK3sehaAilF27DA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F6783D8FB30E283081C16398293F482DCA0E912D",
            "timestamp": "2021-07-23T18:44:00.651283626Z",
            "signature": "JVkJN6e/HOJCBxESispVqyWjVLCPKqGBePHmY15jz4VCRm0UwjhoKVBev+aIuyqrzZwzV6fKbWO4ojni6M7yDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95A04D023E6AEABA5BB79062753EFA1D4B90CF45",
            "timestamp": "2021-07-23T18:44:00.644536481Z",
            "signature": "yujSD+pxs81+XsrJnCJ2G4sVahAHYhxNR3fxJWIMeEufibE3nXPPhgx9h1MScPGBvOQBLvM2N7VDua8q7NQDBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "972A684F364CCE31469B37A9D439985115EB5A40",
            "timestamp": "2021-07-23T18:44:00.574950652Z",
            "signature": "YYvb2H4TPcn4y7sfgiU00sirKZaaL3V6tWiprwlPn+K6CRCUpbjnRtrqQgGT/pj64KS6vONn0xnCPyIxS+AJDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3645027DEE53FE9A6723D8FEDE2B903EFD792AA5",
            "timestamp": "2021-07-23T18:44:00.737892014Z",
            "signature": "98+44w+cJNJ8ju/trBtT/TFM48H3UKhyVrsz/M0rDVWyiep2k4o9F+n58cPdlbvLtBHk5t38li4sSXVIKFC+CQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "30599A9C8D779F13C0A7B5A2359D5FCFCE1E4175",
            "timestamp": "2021-07-23T18:44:00.67773029Z",
            "signature": "F/MESCVprK8gOlJHPjobtLj4UwYcGrzcqlWWb8GO9OrXQkZD+jLk4tIOIaZtBj5XRuxS/28E4gTAuJK0vGAYCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3FF719F1664BEE93D482B480677C03A47EC0B643",
            "timestamp": "2021-07-23T18:44:00.616575283Z",
            "signature": "3QypKtzwa1zvENOW5o+2zj9paIA6ElyFDp8OH92t6AUWxeAc37AxCOrr54TbW+rxtHkOz/d56nUEK0HWlf/SBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "191E896A11C0A77A96A99ABEE986A2A40355C044",
            "timestamp": "2021-07-23T18:44:00.626500221Z",
            "signature": "as5gyqRLN2uSbLPPB1yd1p2ICYouqcWmBkA3Eq/5T5PmLtTSmO0lcpxe3pyIe7NtrFvKwzU/7u75WrHqSl9VCg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "46E5338EF19A939D3D3B0B0B78A1C665F0FA19E8",
            "timestamp": "2021-07-23T18:44:00.658752478Z",
            "signature": "Z6tpdnmzAOdvncTowJWpKoFCWBgJtqlZiVJ9+lAC0yVsrJjnO5lurL6Nx20MNu3HX/9ZVrr1NiboefQI127FDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "8014BA212ED388597510D064258F5E30AA30D591",
            "timestamp": "2021-07-23T18:44:00.645789037Z",
            "signature": "7cOygfTjDGhJOsjKb9EQXuoFX16RO8KI6eNUtdLt1pp0N7ta3iqKgOjE+ucWSCLlzA6Sc2BM53+uNAMxzDslCQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "72B1489EFB57A680577A838A5BAAEBE162A7C802",
            "timestamp": "2021-07-23T18:44:00.592295144Z",
            "signature": "M4aEtkIPVCBFPTfN4+ecvApnFVsmYKgtkcP0J6rhJnmReZ1iJ0AekehAsHpG7nztV4UibVtcKAn8sB5WwVG7Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4E2F0E49E1A479B2A213A841E5E8A1F3BC76B3F7",
            "timestamp": "2021-07-23T18:44:00.672919491Z",
            "signature": "Fwop+V6JAxA5CcBTUFo3W/o4FJucfIgWbaovNH9KogluoEgbfwQ5+IqFqbqHemt3eiikzhsXifW6d3XIRQkMBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "92FBEE90B1646A57F842722C658B6776B5BD575A",
            "timestamp": "2021-07-23T18:44:00.633656896Z",
            "signature": "vuRHNh32UFkj1Xsp4FQpMTMC5d6lpboWg20Ctf1R02YcbLV+PNIpAYILpVcsYO8cEvD7DVgF5gIfA5Dyvs09Dg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "0526389B63CAA7A69AD84376CDCFFD0C8FCC6ED4",
            "timestamp": "2021-07-23T18:44:00.663416329Z",
            "signature": "0kRHo6foIkeHq9VdthgDsFqn8HlnNv65PqQeOVuPwGTGV7v8OSrdnCt/2NAE9alNIYs/22pwCKF278gUyg/rBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "35884D02480A1EE78855E404AC365EB7F3438FB1",
            "timestamp": "2021-07-23T18:44:00.720999439Z",
            "signature": "A63A33diGFgCTXddpaxYsPkqHnFE8SSZfrrEbF9Fm/VbBaafANTH1IHiC70JDxVftceGI2ioqKCgdJilz3T9BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "0B97B2BD62680B733C9FB7A4A309BBD40F3E770F",
            "timestamp": "2021-07-23T18:44:00.586457753Z",
            "signature": "vvV9GOPa7sNr92ZhlkQ19DwK1XeTfut/kRRnna5eIwANk8LWR5RQSnoe76Y/xTuBGAOrDWHr+HTO/XSHm/mkBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "45F34C471F998918F4242785C14B3BC68C5733FA",
            "timestamp": "2021-07-23T18:44:00.823866743Z",
            "signature": "jC6FujiewNduUPd0+oyhbpXW4+0bjRxLvahdzORxxKghZkzenfhPzonBEHQmu6kINZPsAd+ESo3Hyg7Ct8kjCQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "59E84313576C62727D6623BA76AE20012A6BBA41",
            "timestamp": "2021-07-23T18:44:00.64436817Z",
            "signature": "Ph4vCAoYo9SVPQfq8g13feKQqaw7NTLNvWh3Y2603W7PqRCQBalbPk09vcCVTf/t9bPp3Jm0m9Kjmm6g3HjoDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "026D6F178191E370C313F3098C54526443A817DA",
            "timestamp": "2021-07-23T18:44:00.581549132Z",
            "signature": "tSVsjd1mRgyo6LQf2cXOPB8LOhTd8erSH0mqJ7DGQi4/IWNTjPdpkrbbuiE4P9YiSF8PFu5BX0QvHM+j3317CQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "C9E615289D1D92E50292C3B0BD8358D9B2E40290",
            "timestamp": "2021-07-23T18:44:00.63401002Z",
            "signature": "kkVCd08GOCTu3rvQv4nEuuV6kydiFV8RZ8YRj8QPNpsaH7md8K0oP7hepMl47Vbm3b2FZBiQ0bVsMuk9dilYAQ=="
          }
        ]
      }
    },
    "canonical": true
  }
}`
