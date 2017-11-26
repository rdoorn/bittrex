package bittrex

import (
	"fmt"
	"log"
)

// SellLimitResult returns the uuid of the buy
type SellLimitResult struct {
	UUID string `json:"uuid"`
}

// SellLimit buys a coin
// market = market (eg. btc-ltc)
// quantity = the ammount to buy
// rate = the rate to buy at
func (u *User) SellLimit(market string, quantity float64, rate float64) (result SellLimitResult, err error) {
	var response jsonResponse
	r, err := u.getURL("GET", fmt.Sprintf("/api/v1.1/market/selllimit?market=%s&quantity=%.8f&rate=%.8f", market, quantity, rate), nil, true)
	if err != nil {
		return
	}
	log.Printf("Bittrex: %s, q:%f SellLimit: %s", market, quantity, string(r))

	err = parseData(r, &response)
	if err != nil {
		return
	}
	if response.Success == false {
		err = fmt.Errorf(response.Message)
		return
	}
	err = parseData(response.Result, &result)
	return
}
