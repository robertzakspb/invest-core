package historicalquotes

import (
	"time"
)

type QuoteHistory struct {
	Ticker   string
	Currency string
	Quotes   []HistoricalQuote
}

type HistoricalQuote struct {
	Quote float64
	Date  time.Time
}
