package bondquote

import "time"

type BondQuote interface {
	QuoteAsPercentage() float64
	Figi() string
	Timestamp() time.Time
}

type TinkoffBondQuote struct {
	quoteAsPercentage float64
	figi              string
	timestamp         time.Time
	ticker            string
}

func (quote TinkoffBondQuote) QuoteAsPercentage() float64 {
	return quote.quoteAsPercentage
}

func (quote TinkoffBondQuote) Figi() string {
	return quote.figi
}

func (quote TinkoffBondQuote) Timestamp() time.Time {
	return quote.timestamp
}
