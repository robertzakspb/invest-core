package main

import (
	"github.com/compoundinvest/invest-core/quote/tinkoffapi"
	tinkoff "github.com/russianinvestments/invest-api-go-sdk/investgo"
)

func main() {
	config, err := tinkoff.LoadConfig("tinkoffAPIconfig.yaml")
	if err != nil {
		panic("Unable to initialize the config for Tinkoff API")
	}
	tinkoffapi.FetchHistoricalQuotesFor("BBG004730RP0", config)
}
