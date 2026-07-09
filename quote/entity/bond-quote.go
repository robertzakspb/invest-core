package entity

type BondQuote interface {
	QuoteAsPercentage() float64
	Ytm() float64
	Ticker() string
	Figi() string
}
