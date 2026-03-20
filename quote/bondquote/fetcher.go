package bondquote

import (
	"context"
	"fmt"

	"opensource.tbank.ru/invest/invest-go/investgo"
)

func FetchQuotesForFigis(figis []string, config investgo.Config) ([]TinkoffBondQuote, error) {
	client, err := investgo.NewClient(context.TODO(), config, nil)
	if err != nil {
		return []TinkoffBondQuote{}, err
	}

	mdService := client.NewMarketDataServiceClient()
	if mdService == nil {
		return []TinkoffBondQuote{}, fmt.Errorf("failed to initialize Tinkoff's market data service")
	}

	tinkoffQuotes, err := mdService.GetLastPrices(figis)
	if err != nil {
		return []TinkoffBondQuote{}, err
	}

	quotes := []TinkoffBondQuote{}
	for _, tQuote := range tinkoffQuotes.LastPrices {
		quote := TinkoffBondQuote{
			quoteAsPercentage: tQuote.GetPrice().ToFloat(), //TODO: Test if the service actually provides a quote as percentage
			figi:              tQuote.GetFigi(),
			timestamp:         tQuote.GetTime().AsTime(),
			ticker:            tQuote.GetTicker(),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}
