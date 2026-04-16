package bondquote

import (
	"context"
	"fmt"

	"opensource.tbank.ru/invest/invest-go/investgo"
)

func FetchQuotesForFigis(figis []string, config investgo.Config) ([]TQuote, error) {
	client, err := investgo.NewClient(context.TODO(), config, nil)
	if err != nil {
		return []TQuote{}, err
	}

	mdService := client.NewMarketDataServiceClient()
	if mdService == nil {
		return []TQuote{}, fmt.Errorf("failed to initialize Tinkoff's market data service")
	}

	tinkoffQuotes, err := mdService.GetLastPrices(figis)
	if err != nil {
		return []TQuote{}, err
	}

	quotes := []TQuote{}
	for _, tQuote := range tinkoffQuotes.LastPrices {
		quote := TQuote{
			quote:     tQuote.GetPrice().ToFloat(), //TODO: Test if the service actually provides a quote as percentage
			figi:      tQuote.GetFigi(),
			timestamp: tQuote.GetTime().AsTime(),
			ticker:    tQuote.GetTicker(),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}
