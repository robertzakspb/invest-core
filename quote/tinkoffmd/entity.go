package tinkoffapi

import "time"

type TinkoffQuote struct {
	quote     float64
	figi      string
	timestamp time.Time
}

// MARK: Implementation of the SimpleQuote interface
func (quote TinkoffQuote) Quote() float64 {
	return quote.quote
}

func (quote TinkoffQuote) Figi() string {
	return quote.figi
}

func (quote TinkoffQuote) Currency() string {
	return "RUB"
}

func (quote TinkoffQuote) Timestamp() time.Time {
	return quote.timestamp
}
