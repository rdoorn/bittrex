package bittrex

import (
	"encoding/json"
)

// ParseData parses client json and formats the data
func parseData(data []byte, feedback interface{}) (err error) {
	err = json.Unmarshal(json.RawMessage(data), &feedback)
	return
}

// Register a user to set variables
/*func Register(key, secret string) {
	apiKey = key
	apiSecret = secret
}*/
