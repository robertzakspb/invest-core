package historicalquotes

type HistoricalQuotesDTO []struct {
	Charsetinfo struct {
		Name string `json:"name"`
	} `json:"charsetinfo,omitempty"`
	History []any `json:"history,omitempty"`
}

type HistoricalQuoteDTO struct {
	Close float64 `json:"CLOSE"`
}
