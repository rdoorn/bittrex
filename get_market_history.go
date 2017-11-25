package bittrex

import (
	"log"
	"strings"
	"time"
)

// History is the history of buy and sell orders
type History struct {
	ID           int64     `json:"Id"`
	TmpTimeStamp Time      `json:"TimeStamp"`
	TimeStamp    time.Time `json:"TimeStampConv"`
	Quantity     float64   `json:"Quantity"`
	Rate         float64   `json:"Price"`
	Value        float64   `json:"Total"`
	FillType     string    `json:"FillType"`
	OrderType    string    `json:"OrderType"`
}

type Time struct {
	time.Time
}

// UnmarshalJSON returns time.Now() no matter what!
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	// you can now parse b as thoroughly as you want

	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	// 2017-08-23T14:12:49.2
	t.Time, err = time.Parse("2006-01-02T15:04:05", s)

	//*t = Time{time.Now()}
	return err
}

// GetHistory returns order history
// market = market (eg. btc-ltc)
func GetHistory(market string) (result []History, err error) {
	var response jsonResponse
	r, err := getURL("GET", "/api/v1.1/public/getmarkethistory?market="+market, nil, false)
	err = parseData(r, &response)
	if err != nil {
		return
	}
	err = parseData(response.Result, &result)
	if err != nil {
		log.Printf("Error parsing HISTORY: %s", err)
	}
	tmp := result
	for id := range tmp {
		result[id].TimeStamp = result[id].TmpTimeStamp.Time
	}
	return
}
