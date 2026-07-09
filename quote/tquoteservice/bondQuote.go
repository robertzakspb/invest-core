package tquoteservice

type TBondQuote struct {
	ticker            string
	quoteAsPercentage float64
	ytm               float64
}

func NewBondQuote(ticker string, quoteAsPercentage, ytm float64) TBondQuote {
	return TBondQuote{
		ticker:            ticker,
		quoteAsPercentage: quoteAsPercentage,
		ytm:               ytm,
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
