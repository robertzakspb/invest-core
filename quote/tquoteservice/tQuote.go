package bondquote

import "time"

type TQuote struct {
	QuoteAsPercentage float64
	Figi              string
	Timestamp         time.Time
	Ticker            string
}
