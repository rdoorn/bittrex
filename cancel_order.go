package bittrex

// CancelOrder cancels an order
func CancelOrder(orderID string) (err error) {
	_, err = getURL("GET", "/api/v1.1/market/cancel?uuid="+orderID, nil, true)
	return
}
