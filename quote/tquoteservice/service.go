package tquoteservice

import (
	"context"
	"errors"
	"fmt"

	"opensource.tbank.ru/invest/invest-go/investgo"
	investapi "opensource.tbank.ru/invest/invest-go/proto"
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

func GetBondPriceAndYield(client *investgo.Client, tickers []string) ([]BondQuote, []error) {
	parsedQuotes := []BondQuote{}
	errorList := []error{}

	mdService := client.NewMarketDataServiceClient()
	if mdService == nil {
		errorList = append(errorList, errors.New("failed to initialize Tinkoff's market data service"))
		return parsedQuotes, errorList
	}

	valuesTypes := []investapi.MarketValueType{}
	valuesTypes = append(valuesTypes, investapi.MarketValueType_INSTRUMENT_VALUE_LAST_PRICE)
	valuesTypes = append(valuesTypes, investapi.MarketValueType_INSTRUMENT_VALUE_YIELD)

	quotes, err := mdService.GetMarketValues(tickers, valuesTypes)
	if err != nil {
		errorList = append(errorList, err)
		return parsedQuotes, errorList
	}
	if quotes == nil {
		errorList = append(errorList, errors.New("The provided bond quote response is nil"))
		return parsedQuotes, errorList
	}

	for _, instrument := range quotes.Instruments {
		if instrument == nil {
			errorList = append(errorList, errors.New("Instrument is nil: "))
			continue
		}

		priceInPercentage, ytm := 0.0, 0.0
		priceInPercentageFound, ytmFound := false, false

		for _, mv := range instrument.Values {
			if mv == nil || mv.Type == nil || mv.Value == nil {
				errorList = append(errorList, errors.New("The market value or its type is nil"))
			}
			if *mv.Type == investapi.MarketValueType_INSTRUMENT_VALUE_LAST_PRICE {
				priceInPercentageFound = true
				priceInPercentage = mv.Value.ToFloat()
			}
			if *mv.Type == investapi.MarketValueType_INSTRUMENT_VALUE_YIELD {
				ytmFound = true
				ytm = mv.Value.ToFloat()
			}
			if !priceInPercentageFound || !ytmFound {
				errorList = append(errorList, errors.New("Either the quote or YTM was not found in the response"))
				continue
			}
			bondQuote := BondQuote{
				Ticker:            instrument.Ticker,
				QuoteAsPercentage: priceInPercentage,
				YTM:               ytm,
			}
			parsedQuotes = append(parsedQuotes, bondQuote)
		}
	}

	return parsedQuotes, errorList
}
