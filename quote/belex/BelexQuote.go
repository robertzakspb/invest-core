package belex

import "time"

type BelexQuote struct {
	figi         string
	ticker       string
	currentQuote float64
}

// MARK: Implementation of the SimpleQuote interface
func (quote BelexQuote) Quote() float64 {
	return quote.currentQuote
}

func (quote BelexQuote) Timestamp() time.Time {
	return time.Now()
}

func (quote BelexQuote) Figi() string {
	return quote.figi
}

func (quote BelexQuote) Currency() string {
	return "RSD"
}
