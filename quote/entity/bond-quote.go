package entity

import "time"

type BondQuote interface {
	GetQuoteAsPercentage() float64
	GetYtm() float64
	GetTicker() string
	GetFigi() string
	GetTimestamp() time.Time
}
