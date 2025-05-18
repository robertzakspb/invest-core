package tinkoffapi

import (
	"context"
	"fmt"
	"time"

	tinkoff "github.com/russianinvestments/invest-api-go-sdk/investgo"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
)

type TinkoffQuote struct {
	Quote     float64
	Figi      string
	Timestamp time.Time
}

func FetchHistoricalQuotesFor(figi string, config tinkoff.Config) ([]TinkoffQuote, error) {
	client, err := tinkoff.NewClient(context.TODO(), config, nil)
	if err != nil {
		return []TinkoffQuote{}, err
	}

	mdService := client.NewMarketDataServiceClient()
	if mdService == nil {
		return []TinkoffQuote{}, fmt.Errorf("failed to initialize Tinkoff's market data service")
	}

	candleRequest := tinkoff.GetHistoricCandlesRequest{
		Instrument: figi,
		Interval:   investapi.CandleInterval_CANDLE_INTERVAL_DAY,
		From:       time.Now(), //Ignored by the API
		To:         time.Now(), //Ignored by the API
		File:       false,
		FileName:   "",
		Source:     investapi.GetCandlesRequest_CANDLE_SOURCE_INCLUDE_WEEKEND,
	}

	quotes, err := mdService.GetAllHistoricCandles(&candleRequest)
	if err != nil {
		return []TinkoffQuote{}, err
	}

	parsedQuotes := []TinkoffQuote{}
	for _, quote := range quotes {
		if quote == nil {
			continue
		}
		parsedQuote := TinkoffQuote{
			Quote:     quote.Close.ToFloat(),
			Figi:      figi,
			Timestamp: time.Unix(quote.Time.GetSeconds(), 0),
		}
		parsedQuotes = append(parsedQuotes, parsedQuote)
	}

	return parsedQuotes, nil
}
