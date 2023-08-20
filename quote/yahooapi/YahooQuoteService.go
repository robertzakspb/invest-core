package yahooapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	quote "github.com/compoundinvest/invest-core/quote/entity"
)

type SimpleQuote = quote.SimpleQuote

const user_agent_key = "User-Agent"
const user_agent_value = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/536.36 (KHTML, like Gecko) Chrome/58.0.2029.110 Safari/437.36"

func FetchQuotes(tickers []string) []SimpleQuote {
	quotes := []quote.SimpleQuote{}

	yahooQuotes := fetchQuotes(tickers)
	quotes = append(quotes, quote.ConvertToSimpleQuote(yahooQuotes)...)

	return quotes
}

func fetchQuotes(tickers []string) []YahooQuote {
	cookie := fetchAuthCookie()
	crumb := fetchCrumb(cookie)

	apiBase := "https://query2.finance.yahoo.com/v7/finance/quote?"
	queryParams := url.Values{
		"crumb":   {crumb},
		"symbols": {strings.Join((tickers), ",")},
	}
	quotesURL := apiBase + queryParams.Encode()

	request, err := http.NewRequest("GET", quotesURL, nil)
	request.AddCookie(cookie)
	request.Header.Set(user_agent_key, user_agent_value)
	if err != nil {
		fmt.Println("Unable to create a GET request using the following URL: ", quotesURL)
	}

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Unable to configure a cookie jar")
	}
	yahooCookie := &http.Cookie{
		Name:  cookie.Name,
		Value: cookie.Value,
	}
	request.AddCookie(yahooCookie)

	client := http.Client{Jar: cookieJar}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Unable to fetch quotes from ", quotesURL, ". ", err)
		return nil
	}
	defer response.Body.Close()

	var quotesDTO YahooQuotesDTO
	json.NewDecoder(response.Body).Decode(&quotesDTO)

	return quotesDTO.QuoteResponse.Result
}

// Workaround to make sure we can fetch quotes via Yahoo API
func fetchAuthCookie() *http.Cookie {
	cookieURL := "https://fc.yahoo.com/"
	response, err := http.Get(cookieURL)
	if err != nil {
		fmt.Println("Unable to fetch the auth cookie from ", cookieURL, ". ", err)
	}

	cookie := response.Cookies()[0]
	return cookie
}

func fetchCrumb(cookie *http.Cookie) string {
	crumbURL := "https://query2.finance.yahoo.com/v1/test/getcrumb"
	request, err := http.NewRequest("GET", crumbURL, nil)
	request.AddCookie(cookie)
	request.Header.Set(user_agent_key, user_agent_value)
	if err != nil {
		fmt.Println("Unable to create a GET request using the following URL: ", crumbURL)
	}

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Unable to configure a cookie jar")
	}
	yahooCookie := &http.Cookie{
		Name:  cookie.Name,
		Value: cookie.Value,
	}
	request.AddCookie(yahooCookie)

	client := http.Client{Jar: cookieJar}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Unable to fetch the crumb from ", crumbURL, ". ", err)
		return ""
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to read the crumb from the response: ", response)
	}
	responseString := string(responseData)

	return responseString
}
