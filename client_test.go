package missed

//FIXME: test unix socket?
//func TestFetchSummary(t *testing.T) {
//	var (
//		tendermintApi = "http://127.0.0.1:26657"
//		cosmosApi     = "http://127.0.0.1:1317"
//	)
//	if os.Getenv("COSMOS_API") != "" {
//		cosmosApi = os.Getenv("COSMOS_API")
//	}
//	if os.Getenv("TENDERMINT_API") != "" {
//		tendermintApi = os.Getenv("TENDERMINT_API")
//	}
//	height, _, err := CurrentHeight(tendermintApi)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	s, err := FetchSummary(height-1)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Printf("%+v\n", s)
//}
