package bittrex

// CancelOrder cancels an order
func (u *User) CancelOrder(orderID string) (err error) {
	_, err = u.getURL("GET", "/api/v1.1/market/cancel?uuid="+orderID, nil, true)
	return
}
