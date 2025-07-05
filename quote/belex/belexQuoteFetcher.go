package belex

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)



func FetchQuoteFor(ticker, figi string) (BelexQuote, error) {
	sourceURL := "https://www.belex.rs/trgovanje/hartija/dnevni/" + ticker
	response, err := http.Get(sourceURL)
	if err != nil {
		return BelexQuote{}, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return BelexQuote{}, err
	}

	html := string(content)
	quoteHTMLTag := "<tr><td class='lb'>Cena</td><td>"
	index := strings.Index(html, quoteHTMLTag)
	htmlSliceContainingPrice := string(html[index+32:index+40])

	price := ""
	for _, character := range htmlSliceContainingPrice {
		if string(character) == "<" {
			break
		}
		if string(character) == "." {
			continue
		}
		price = price + string(character)
	}

	parsedPrice, err := strconv.Atoi(price)
	if err != nil {
		fmt.Println("Error parsing the price for ", ticker, err)
		return BelexQuote{}, err
	}

	return BelexQuote{ticker, float64(parsedPrice)}, nil
}