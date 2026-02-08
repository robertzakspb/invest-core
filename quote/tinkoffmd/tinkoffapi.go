package tinkoffapi

import (
	"context"
	"fmt"
	"time"

	tinkoff "github.com/russianinvestments/invest-api-go-sdk/investgo"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
)

func FetchHistoricalQuotesFor(figi string, config tinkoff.Config, startDate time.Time, endDate time.Time) ([]TinkoffQuote, error) {
	if startDate.After(time.Now()) || endDate.After(time.Now()) || startDate.After(endDate) {
		return []TinkoffQuote{}, fmt.Errorf("attempting to fetch historical candles for invalid dates: start date: %d, end date: %d", startDate.Unix(), endDate.Unix())
	}

	if startDate.Unix() == 0 && endDate.Unix() == 0 {
		return FetchAllHistoricalQuotesFor(figi, config)
	}

	client, err := tinkoff.NewClient(context.TODO(), config, nil)
	if err != nil {
		return []TinkoffQuote{}, err
	}

	mdService := client.NewMarketDataServiceClient()
	if mdService == nil {
		return []TinkoffQuote{}, fmt.Errorf("failed to initialize Tinkoff API's market data service")
	}

	candles, err := mdService.GetCandles(
		figi,
		investapi.CandleInterval_CANDLE_INTERVAL_DAY,
		startDate,
		endDate,
		investapi.GetCandlesRequest_CANDLE_SOURCE_INCLUDE_WEEKEND
	)
	if err != nil {
		return []TinkoffQuote{}, err
	}

	parsedQuotes := mapTinkoffCandlesToQuotes(figi, candles.Candles)

	return parsedQuotes, nil
}

func FetchAllHistoricalQuotesFor(figi string, config tinkoff.Config) ([]TinkoffQuote, error) {
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

	candles, err := mdService.GetAllHistoricCandles(&candleRequest)
	if err != nil {
		return []TinkoffQuote{}, err
	}

	parsedQuotes := mapTinkoffCandlesToQuotes(figi, candles)

	return parsedQuotes, nil
}

func mapTinkoffCandlesToQuotes(figi string, candles []*investapi.HistoricCandle) []TinkoffQuote {
	parsedQuotes := []TinkoffQuote{}
	for _, quote := range candles {
		if quote == nil {
			continue
		}
		parsedQuote := TinkoffQuote{
			quote:     quote.Close.ToFloat(),
			figi:      figi,
			timestamp: time.Unix(quote.Time.GetSeconds(), 0),
		}
		parsedQuotes = append(parsedQuotes, parsedQuote)
	}

	return parsedQuotes
}
