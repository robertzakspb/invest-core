package moexapi

import "time"

type MoexQuotesDTO struct {
	MarketData MarketData `json:"marketdata"`
}

type MarketData struct {
	Data [][]any `json:"data"`
}

type MoexQuote struct {
	ticker  string
	CurrentQuote float64
	MarketCap    float64
	figi         string
}

// MARK: Implementation of the SimpleQuote interface
func (quote MoexQuote) Quote() float64 {
	return quote.CurrentQuote
}

// TODO: To fix
func (quote MoexQuote) Figi() string {
	return quote.figi
}

func (quote MoexQuote) Currency() string {
	return "RUB"
}

func (quote MoexQuote) Timestamp() time.Time {
	return time.Now()
}
