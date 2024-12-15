package tiingo

import (
	"context"
	"strings"
	"time"
)

// EodPrice returns the daily price response data for a given ticker with the provided
// params from the [End-of-Day].2.1.2 End-of-Day Endpoint.
//
// Any zero value arguments will be left off the query string & whatever Tiingo's
// default for an empty query string will be returned.
func (c *Client) EodPrice(ctx context.Context, ticker string, startDate, endDate time.Time,
	resampleFreq EodFreq, sort Sort, respFormat Format, columns []string) ([]byte, error) {
	// Build URL
	url := EodPriceUrl(ticker, startDate, endDate, resampleFreq, sort, respFormat, columns)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultEodPrice returns the daily price response data for the given ticker from the
// [End-of-Day].2.1.2 End-of-Day Endpoint.
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultEodPrice(ctx context.Context, symbol string) ([]byte, error) {
	// Build URL
	url := EodPriceUrl(symbol, time.Time{}, time.Time{}, "", "", "", nil)

	// Fetch the data
	return c.get(ctx, url)
}

// EodPriceUrl returns a built url for the given ticker with the provided params
// from the [End-of-Day].2.1.2 End-of-Day Endpoint.
//
// Any zero value arguments will be left off the query string.
func EodPriceUrl(ticker string, startDate, endDate time.Time, resampleFreq EodFreq,
	sort Sort, respFormat Format, columns []string) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/daily/")
	url.WriteString(ticker)
	url.WriteString("/prices")

	// Build query string
	first := true
	if !startDate.IsZero() {
		url.WriteString("?startDate=")
		url.WriteString(startDate.Format("2006-01-02"))
		first = false
	}
	if !endDate.IsZero() {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("endDate=")
		url.WriteString(endDate.Format("2006-01-02"))
	}
	if resampleFreq != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("resampleFreq=")
		url.WriteString(resampleFreq)
	}
	if sort != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("sort=")
		url.WriteString(sort)
	}
	if respFormat != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("format=")
		url.WriteString(respFormat)
	}
	if len(columns) > 0 {
		if first {
			url.WriteString("?")
		} else {
			url.WriteString("&")
		}
		url.WriteString("columns=")
		url.WriteString(strings.Join(columns, ","))
	}

	return url.String()
}

// EodMetadata returns the eod metadata response data for a given ticker
// from the [End-of-Day].2.1.3 Meta Endpoint
func (c *Client) EodMetadata(ctx context.Context, ticker string, respFormat Format) ([]byte, error) {
	// Build URL
	url := EodMetadataUrl(ticker, respFormat)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultEodMetadata returns the eod metadata response data for a given ticker
// // from the [End-of-Day].2.1.3 Meta Endpoint
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultEodMetadata(ctx context.Context, symbol string) ([]byte, error) {
	// Build URL
	url := EodMetadataUrl(symbol, "")

	// Fetch the data
	return c.get(ctx, url)
}

// EodMetadataUrl returns a built url for the given ticker from the
// [End-of-Day].2.1.3 Meta Endpoint.
func EodMetadataUrl(ticker string, respFormat Format) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/daily/")
	url.WriteString(ticker)

	// Build query string
	if respFormat != "" {
		url.WriteString("?format=")
		url.WriteString(respFormat)
	}

	return url.String()
}
