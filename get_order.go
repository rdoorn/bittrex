package bittrex

import "log"

type OrderG struct {
	AccountId                  string
	OrderUuid                  string `json:"OrderUuid"`
	Exchange                   string `json:"Exchange"`
	Type                       string
	Quantity                   float64 `json:"Quantity"`
	QuantityRemaining          float64 `json:"QuantityRemaining"`
	Limit                      float64 `json:"Limit"`
	Reserved                   float64
	ReserveRemaining           float64
	CommissionReserved         float64
	CommissionReserveRemaining float64
	CommissionPaid             float64
	Price                      float64 `json:"Price"`
	PricePerUnit               float64 `json:"PricePerUnit"`
	Opened                     string
	Closed                     string
	IsOpen                     bool
	Sentinel                   string
	CancelInitiated            bool
	ImmediateOrCancel          bool
	IsConditional              bool
	Condition                  string
	ConditionTarget            string
}

// GetOrder returns the status of an order
// uuid = orderid
func (u *User) GetOrder(orderID string) (result OrderG, err error) {
	var response jsonResponse
	r, err := u.getURL("GET", "/api/v1.1/account/getorder?uuid="+orderID, nil, true)
	if err != nil {
		return
	}
	log.Printf("Bittrex: GetOrder: %s", string(r))
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	return
}
