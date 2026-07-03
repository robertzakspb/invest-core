package tquoteservice

import "time"

type BondQuote struct {
	QuoteAsPercentage float64
	YTM               float64
}

type TQuote struct {
	quote     float64
	figi      string
	timestamp time.Time
	ticker    string
}

func New(quote float64, figi, ticker string, timestamp time.Time) TQuote {
	return TQuote{
		quote:     quote,
		figi:      figi,
		timestamp: timestamp,
		ticker:    ticker,
	}
}

func (t *TQuote) Quote() float64 {
	return t.quote
}
func (t *TQuote) Figi() string {
	return t.figi
}
func (t *TQuote) Currency() string {
	return ""
}
func (t *TQuote) Timestamp() time.Time {
	return t.timestamp
}
