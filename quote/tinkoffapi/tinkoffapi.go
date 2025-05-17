package tinkoffapi

import (
	"context"
	"fmt"
	"time"

	tinkoff "github.com/russianinvestments/invest-api-go-sdk/investgo"
	"github.com/russianinvestments/invest-api-go-sdk/proto"
)

/*
type SimpleQuote interface {
	Quote() float64
	Ticker() string
}

type GetHistoricCandlesRequest struct {
	Instrument string
	Interval   pb.CandleInterval
	From       time.Time
	To         time.Time
	File       bool
	FileName   string
	Source     pb.GetCandlesRequest_CandleSource
}

*/

type TinkoffQuote struct {
	quote float64
	ticker string
}

func (quote TinkoffQuote) Quote() float64 {
	return quote.quote
}

func (quote TinkoffQuote) Ticker() string {
	return quote.ticker
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
		Interval: investapi.CandleInterval_CANDLE_INTERVAL_DAY,
		From: time.Now(), //Ignored by the API
		To: time.Now(), //Ignored by the API
		File: false,
		FileName: "",
		Source: investapi.GetCandlesRequest_CANDLE_SOURCE_UNSPECIFIED,
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
		fmt.Println(time.Unix(quote.Time.GetSeconds(), 0), ":", quote.Close.ToFloat())
	}

	return parsedQuotes, nil
}