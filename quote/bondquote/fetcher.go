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

	quotes, err := mdService.GetLastPrices(figis)
	if err != nil {
		return []TinkoffBondQuote{}, err
	}
	for _, quote := range quotes.LastPrices {
		fmt.Println(quote)
	}

	return []TinkoffBondQuote{}, nil //FIXME
}
