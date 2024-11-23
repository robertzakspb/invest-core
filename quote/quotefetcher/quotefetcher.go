package quotefetcher

import (
	"fmt"

	"github.com/compoundinvest/invest-core/quote/belex"
	"github.com/compoundinvest/invest-core/quote/entity"
	"github.com/compoundinvest/invest-core/quote/moexapi"
	"github.com/compoundinvest/invest-core/quote/yahooapi"
)

type TickerWithMarket struct {
	Ticker string
	MIC    string
}

// Universal quote fetcher
func FetchQuotesFor(tickers []TickerWithMarket) []entity.SimpleQuote {
	quotes := []entity.SimpleQuote{}

	moexTickers := []string{} //For MOEX it makes more sense to do a bulk download of quotes
	for _, security := range tickers {
		switch security.MIC {
		case "MOEX":
			moexTickers = append(moexTickers, security.Ticker)
		case "BELEX":
			belexQuote, err := belex.FetchQuoteFor(security.Ticker)
			if err != nil {
				fmt.Println(err)
				continue
			}
			quotes = append(quotes, belexQuote)
		case "XNAS":
			quotes = append(quotes, yahooapi.FetchQuotes([]string{security.Ticker})...)
		default:
			fmt.Println("Unsupported MIC")
		}
	}

	moexQuotes := moexapi.FetchQuotes(moexTickers)
	quotes = append(quotes, moexQuotes...)

	return quotes
}