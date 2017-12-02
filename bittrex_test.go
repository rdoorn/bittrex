package bittrex

import (
	"testing"
)

var testorder = "xxx-xxx-xxx-xxx"

func TestBuyAndSell(t *testing.T) {

	/*
		TestMode = false
		if err := LoadConfig("../../tiger.yaml"); err != nil {
			log.Fatalf("Error reading tiget.yaml: %s", err)
		}
		log.Printf("Config: %+v", Get())

		Register(Get().Bittrex.Key, Get().Bittrex.Secret)

		x, err := GetOrder(testorder)
		if err != nil {
			log.Fatalf("Could not get order: %s", err)
		}
		log.Printf("OrderDetails: %+v", x)

		/*
			order, err := SellLimit("BTC-NEO", 0.00001, 0.01)
			if err != nil {
				log.Fatalf("Could not buy: %s", err)
			}
			orderID := order.UUID
			log.Printf("Order UUID:%s", orderID)
			CancelOrder(orderID)
	*/

}
