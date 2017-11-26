package bittrex

// Currency contains basic currency information
type Currency struct {
	Currency         string  `json:"Currency"`
	CurrencyLong     string  `json:"CurrencyLong"`
	MinConfirmations int     `json:"MinConfirmations"`
	TxFee            float64 `json:"TxFee"`
	IsActive         bool    `json:"IsActive"`
	CoinType         string  `json:"CoinType"`
	BaseAddress      string  `json:"BaseAddress"`
}

// GetCurrencies returns all currencies
func (u *User) GetCurrencies() (result []Currency, err error) {
	var response jsonResponse
	r, err := u.getURL("GET", "/api/v1.1/public/getcurrencies", nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}
