package bittrex

// Market contains basic market information
type Market struct {
	MarketCurrency     string  `json:"MarketCurrency"`
	BaseCurrency       string  `json:"BaseCurrency"`
	MarketCurrencyLong string  `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
	MinTradeSize       float64 `json:"MinTradeSize"`
	MarketName         string  `json:"MarketName"`
	IsActive           bool    `json:"IsActive"`
}

// GetMarkets returns all markets
func GetMarkets() (result []Market, err error) {
	var response jsonResponse
	r, err := getURL("GET", "/api/v1.1/public/getmarkets", nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}
