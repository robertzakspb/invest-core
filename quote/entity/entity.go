package entity

type SimpleQuote interface {
	Quote() float64
	Ticker() string
	ISIN() string
}

// Converts slices of structs implementing SimpleQuote to a slice of SimpleQuote instances
func ConvertToSimpleQuote[T SimpleQuote](quotes []T) []SimpleQuote {
	quotesAsInterface := []SimpleQuote{}
	for _, quote := range quotes {
		quotesAsInterface = append(quotesAsInterface, quote)
	}
	return quotesAsInterface
}
