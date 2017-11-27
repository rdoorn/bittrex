package bittrex

import "fmt"

// Balance is the balance of your account for each currency
type Balance struct {
	Currency      string  `json:"Currency"`
	Balance       float64 `json:"Balance"`
	Available     float64 `json:"Available"`
	Pending       float64 `json:"Pending"`
	CryptoAddress string  `json:"CryptoAddress"`
	Requested     bool    `json:"Requested"`
	UUID          string  `json:"Uuid"`
}

// GetBalances returns summary of your belances
func (u *User) GetBalances() (result []Balance, err error) {
	var response jsonResponse
	r, err := u.getURL("GET", "/api/v1.1/account/getbalances", nil, true)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}

// GetBalance returns balance of your currency
// currency = currency (e.g. btc)
func (u *User) GetBalance(currency string) (result Balance, err error) {
	var response jsonResponse
	r, err := u.getURL("GET", "/api/v1.1/account/getbalance?currency="+currency, nil, true)
	fmt.Printf("Got data 1: %+v\n", string(r))
	err = parseData(r, &response)
	if err != nil {
		return
	}
	if response.Success == false {
		err = fmt.Errorf("bittrex returned: %s (currency requested:%s)", response.Message, currency)
		return
	}
	fmt.Printf("Got data result: %+v\n", string(response.Result))
	err = parseData(response.Result, &result)
	return
}
