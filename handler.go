package bittrex

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

//
// We are currently restricting orders to 500 open orders and 200,000 orders a day
//

func (u *User) getURL(method string, path string, data interface{}, auth bool) ([]byte, error) {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, bittrexURL+path, strings.NewReader(string(jsondata)))
	if err != nil {
		return nil, err
	}

	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	if auth == true {
		nonce := time.Now().UnixNano()
		q := req.URL.Query()
		if u.APIKey == "" || u.APISecret == "" {
			log.Fatalln("No api Key or Secret set!")
		}
		q.Set("apikey", u.APIKey)
		q.Set("nonce", fmt.Sprintf("%d", nonce))
		req.URL.RawQuery = q.Encode()
		mac := hmac.New(sha512.New, []byte(u.APISecret))
		_, err = mac.Write([]byte(req.URL.String()))
		if err != nil {
			return nil, err
		}
		sig := hex.EncodeToString(mac.Sum(nil))
		req.Header.Add("apisign", sig)
	}

	if testMode {
		p := strings.Split(req.URL.Path, "?")
		ps := strings.Split(p[0], "/")
		log.Printf("TEST MODE: Path: %s Request: %+v ", ps[len(ps)-1], req)
		switch ps[len(ps)-1] {
		case "selllimit":
			id, _ := uuid.NewV4()
			return []byte("{ \"success\" : true, \"message\" : \"\", \"result\" : { \"uuid\" : \"" + id.String() + "\" } }"), nil
		case "buylimit":
			id, _ := uuid.NewV4()
			return []byte("{ \"success\" : true, \"message\" : \"\", \"result\" : { \"uuid\" : \"" + id.String() + "\" } }"), nil
		case "cancel":
			return []byte("{ \"success\" : true, \"message\" : \"\", \"result\" : { } }"), nil
			//return []byte{}, nil
		case "getorder":
			return []byte("{ \"success\" : true, \"message\" : \"\", \"result\" : { \"IsOpen\" : true, \"Quantity\": 1, \"QuantityRemaining\": 0 } }"), nil
		}
	}
	// Custom dialer with timeouts
	dialer := &net.Dialer{
		//LocalAddr: &localTCPAddr,
		Timeout:   time.Duration(httpTimeout) * time.Second,
		KeepAlive: 10 * time.Second,
		Deadline:  time.Now().Add(httpTimeout * 10 * time.Second),
		DualStack: true,
	}

	tlsConfig := &tls.Config{}

	// Overwrite default transports with our own for checking the correct node
	tr := &http.Transport{
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
		ResponseHeaderTimeout: time.Duration(httpTimeout) * time.Second,
		TLSHandshakeTimeout:   time.Duration(httpTimeout) * time.Second,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, network, addr)
		},
	}

	//log.Printf("Request: %+v\n", req)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//log.Printf("Response: %+v\n", resp)

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	//log.Printf("Body: %+v\n", string(response))
	//fmt.Println(fmt.Sprintf("reponse %s", response), err)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err

}
