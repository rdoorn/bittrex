package bittrex

// OrderBook is the book of buy and sell orders
type OrderBook struct {
	Buy  []Orderb `json:"buy"`
	Sell []Orderb `json:"sell"`
}

// Orderb is bookentry
type Orderb struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
}

// GetOrderBook returns order book
// market = market (eg. btc-ltc)
// booktype = what book to get (eg. buy, sell or both)
func (u *User) GetOrderBook(market, booktype string) (result OrderBook, err error) {
	var response jsonResponse
	r, err := u.getURL("GET", "/api/v1.1/public/getorderbook?market="+market+"&type="+booktype, nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}
