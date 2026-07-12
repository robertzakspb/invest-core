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
func (b TBondQuote) GetQuoteAsPercentage() float64 {
	return b.quoteAsPercentage
}

func (b TBondQuote) GetYtm() float64 {
	return b.ytm
}

func (b TBondQuote) GetTicker() string {
	return b.ticker
}

func (b TBondQuote) GetFigi() string {
	return ""
}

func (b TBondQuote) GetTimestamp() time.Time {
	return b.timestamp
}
