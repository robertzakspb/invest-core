package quotefetcher

import (
	"github.com/compoundinvest/invest-core/quote/belex"
	"github.com/compoundinvest/invest-core/quote/entity"
	"github.com/compoundinvest/invest-core/quote/moexapi"
	"github.com/compoundinvest/invest-core/quote/yahooapi"
)

// Universal quote fetcher
func FetchQuotesFor(securities []entity.Security) []entity.SimpleQuote {
	quotes := []entity.SimpleQuote{}

	moexSecurities := []entity.Security{} //For MOEX it makes more sense to do a bulk download of quotes
	for _, security := range securities {
		switch security.MIC {
		case "MISX":
			moexSecurities = append(moexSecurities, security)
		case "XBEL":
			belexQuote, err := belex.FetchQuoteFor(security.Ticker, security.Figi)
			if err != nil {
				continue
			}
			quotes = append(quotes, belexQuote)
		case "XNAS":
			quotes = append(quotes, yahooapi.FetchQuotes([]entity.Security{security})...)
		default:
		}
	}
 
	moexQuotes := moexapi.FetchQuotes(moexSecurities)
	quotes = append(quotes, moexQuotes...)

	return quotes
}