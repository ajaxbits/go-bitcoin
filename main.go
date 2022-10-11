package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	time1 := 1664374620
	url1 := fmt.Sprintf("https://api.cryptowat.ch/markets/kraken/btcusd/ohlc?before=%d&after=%d", time1, time1)

	resp, err := http.Get(url1)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	wow := string(body)
	fmt.Println(wow)
}
