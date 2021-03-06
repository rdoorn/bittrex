package bittrex

import "fmt"

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
func (u *User) GetMarketsSummary() (result []MarketSummary, err error) {
	var response jsonResponse
	r, err := u.getURL("GET", "/api/v1.1/public/getmarketsummaries", nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}

// GetMarketSummary returns summary of 1 market
func (u *User) GetMarketSummary(market string) (result MarketSummary, err error) {
	var response jsonResponse
	var results []MarketSummary
	r, err := u.getURL("GET", "/api/v1.1/public/getmarketsummary?market="+market, nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	if response.Success == false {
		err = fmt.Errorf("bittrex returned: %s (market requested:%s)", response.Message, market)
		return
	}
	err = parseData(response.Result, &results)
	result = results[0]
	return
}
