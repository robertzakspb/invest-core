package yahooapi

import "time"

type YahooQuotesDTO struct {
	QuoteResponse QuoteResponseDTO `json:"quoteResponse"`
}

type QuoteResponseDTO struct {
	Result []YahooQuote `json:"result"`
}

type YahooQuote struct {
	Symbol             string  `json:"symbol"`
	DisplayName        string  `json:"displayName"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
	PriceToBook        float64 `json:"priceToBook"`
	ForwardPE          float64 `json:"forwardPE"`
	MarketCap          float64 `json:"marketCap"`
	TrailingPE         float64 `json:"trailingPE"`
	figi               string
}

func (quote YahooQuote) Quote() float64 {
	return quote.RegularMarketPrice
}

func (quote YahooQuote) Currency() string {
	return "USD"
}

func (quote YahooQuote) Figi() string {
	return quote.figi
}

func (quote YahooQuote) Timestamp() time.Time {
	return time.Now()
}
