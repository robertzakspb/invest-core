package main

import "github.com/compoundinvest/invest-core/quote/moexapi/historicalquotes"

func main() {
	historicalquotes.FetchHistoricalQuotesFor("SBER")
}
