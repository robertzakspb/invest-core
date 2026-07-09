package tquoteservice

import "time"

type TBondQuote struct {
	ticker            string
	quoteAsPercentage float64
	ytm               float64
	timestamp         time.Time
}

func NewBondQuote(ticker string, quoteAsPercentage, ytm float64, timestamp time.Time) TBondQuote {
	return TBondQuote{
		ticker:            ticker,
		quoteAsPercentage: quoteAsPercentage,
		ytm:               ytm,
		timestamp:         timestamp,
	}
}

// Implementation of the BondQuote interface
func (b TBondQuote) QuoteAsPercentage() float64 {
	return b.quoteAsPercentage
}

func (b TBondQuote) Ytm() float64 {
	return b.ytm
}

func (b TBondQuote) Ticker() string {
	return b.ticker
}

func (b TBondQuote) Figi() string {
	return ""
}

func (b TBondQuote) Timestamp() time.Time {
	return b.timestamp
}
