package historicalquotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchHistoricalQuotesFor(ticker string) []HistoricalQuote {
	quotesURL := "http://iss.moex.com/iss/history/engines/stock/markets/shares/boards/TQBR/securities/" + ticker + ".json?iss.json=extended&from=2000-01-01"

	response, err := http.Get(quotesURL)
	if err != nil {
		fmt.Println("Unable to fetch historical quotes from ", quotesURL)
		return nil
	}
	defer response.Body.Close()

	var quotesDTO HistoricalQuotesDTO
	error := json.NewDecoder(response.Body).Decode(&quotesDTO)
	if error != nil {
		println(error)
	}

	return []HistoricalQuote{}
}
