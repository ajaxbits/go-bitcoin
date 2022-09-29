package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CryptowatchResponse struct {
    Result CryptowatchResult
    Allowance CryptowatchAllowance
}

type CryptowatchResult struct {
    14400 list
    180 list
    21600 list
    259200 list
    300 list
    3600 list
    43200 list
    60 list
    604800 list
    604800_Monday list
    7200 list
    86400 list
    900 list
}

type CryptowatchAllowance struct {
    Cost float32
    Remaining float32
    Upgrade string
}

func main() {
	time1 := 1664374620
	url1 := fmt.Sprintf("https://api.cryptowat.ch/markets/kraken/btcusd/ohlc?before=%d&after=%d", time1, time1)
	time2 := 1664374560
	url2 := fmt.Sprintf("https://api.cryptowat.ch/markets/kraken/btcusd/ohlc?before=%d&after=%d", time2, time2)

	resp, err := http.Get(url1)
	if err != nil {
		log.Fatalln(err)
	}

	resp2, err2 := http.Get(url2)
	if err != nil {
		log.Fatalln(err2)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)

	//We Read the response body on the line below.
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb2 := string(body2)
	fmt.Println(sb2)
}
