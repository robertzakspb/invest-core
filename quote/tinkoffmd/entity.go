package tinkoffapi

import "time"

type TinkoffQuote struct {
	quote     float64
	figi      string
	isin      string
	timestamp time.Time
}

// MARK: Implementation of the SimpleQuote interface
func (quote TinkoffQuote) Quote() float64 {
	return quote.quote
}

func (quote TinkoffQuote) ISIN() string {
	return quote.isin
}

// TODO: To fix
func (quote TinkoffQuote) Ticker() string {
	return ""
}

func (quote TinkoffQuote) Timestamp() time.Time {
	return quote.timestamp
}
