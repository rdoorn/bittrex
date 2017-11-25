package bittrex

// MarketSummary contains the summary of a market
type MarketSummary struct {
	MarketName     string  `json:"MarketName"`
	High           float64 `json:"High"`
	Low            float64 `json:"Low"`
	Ask            float64 `json:"Ask"`
	Bid            float64 `json:"Bid"`
	OpenBuyOrders  int     `json:"OpenBuyOrders"`
	OpenSellOrders int     `json:"OpenSellOrders"`
	Volume         float64 `json:"Volume"`
	Last           float64 `json:"Last"`
	BaseVolume     float64 `json:"BaseVolume"`
	PrevDay        float64 `json:"PrevDay"`
	TimeStamp      string  `json:"TimeStamp"`
}

// GetMarketsSummary returns all market status
func GetMarketsSummary() (result []MarketSummary, err error) {
	var response jsonResponse
	r, err := getURL("GET", "/api/v1.1/public/getmarketsummaries", nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}

// GetMarketSummary returns summary of 1 market
func GetMarketSummary(market string) (result MarketSummary, err error) {
	var response jsonResponse
	r, err := getURL("GET", "/api/v1.1/public/getmarketsummary?market="+market, nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}
