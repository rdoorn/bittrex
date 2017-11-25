package bittrex

// Ticker gives the last bid, ask and last
type Ticker struct {
	Bid  float64 `json:"Bid"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}

// GetTicker returns summary of 1 market
// market = market (eg. btc-ltc)
func GetTicker(market string) (result Ticker, err error) {
	var response jsonResponse
	r, err := getURL("GET", "/api/v1.1/public/getticker?market="+market, nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}
