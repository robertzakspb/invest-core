package yahooapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	quote "github.com/compoundinvest/invest-core/quote/entity"
)

type SimpleQuote = quote.SimpleQuote

func fetchQuotes(tickers []string) []YahooQuote {
	quotesURL := "https://query1.finance.yahoo.com/v6/finance/quote?symbols=" + strings.Join((tickers), ",")

	response, err := http.Get(quotesURL)
	if err != nil {
		fmt.Println("Unable to fetch quotes from ", quotesURL, ". ", err)
		return nil
	}
	defer response.Body.Close()

	var quotesDTO YahooQuotesDTO
	json.NewDecoder(response.Body).Decode(&quotesDTO)

	return quotesDTO.QuoteResponse.Result
}

func FetchQuotes(tickers []string) []SimpleQuote {
	quotes := []quote.SimpleQuote{}

	yahooQuotes := fetchQuotes(tickers)
	quotes = append(quotes, quote.ConvertToSimpleQuote(yahooQuotes)...)

	return quotes
}
