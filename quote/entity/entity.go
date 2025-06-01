package entity

import "time"

type SimpleQuote interface {
	Quote() float64
	Ticker() string
	ISIN() string
	Timestamp() time.Time
}

// Converts slices of structs implementing SimpleQuote to a slice of SimpleQuote instances
func ConvertToSimpleQuote[T SimpleQuote](quotes []T) []SimpleQuote {
	quotesAsInterface := []SimpleQuote{}
	for _, quote := range quotes {
		quotesAsInterface = append(quotesAsInterface, quote)
	}
	return quotesAsInterface
}
