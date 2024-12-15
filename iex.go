package tiingo

import (
	"context"
	"strings"
	"time"
)

// IexTopOfBook returns the last price item for the specified tickers from
// the [IEX].2.5.2 Top-of-Book & Last Price Endpoints.
//
// Any zero value arguments will be left off the query string.
func (c *Client) IexTopOfBook(ctx context.Context, tickers []string, respFormat Format) ([]byte, error) {
	// Build url
	url := IexTopOfBookUrl(tickers, respFormat)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultIexTopOfBook returns the last price item for the all tickers from
// the [IEX].2.5.2 Top-of-Book & Last Price Endpoints.
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultIexTopOfBook(ctx context.Context) ([]byte, error) {
	// Build url
	url := IexTopOfBookUrl(nil, "")

	// Fetch the data
	return c.get(ctx, url)
}

// IexTopOfBookUrl returns a built url for the given tickers from the
// [IEX].2.5.2 Top-of-Book & Last Price Endpoint.
//
// Any zero value arguments will be left off the query string.
func IexTopOfBookUrl(tickers []string, respFormat Format) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/iex")
	if len(tickers) > 0 {
		url.WriteString("/")
		url.WriteString(strings.Join(tickers, ","))
	}

	// Build query string
	if respFormat != "" {
		url.WriteString("?format=")
		url.WriteString(respFormat)
	}

	return url.String()
}

// IexHistory returns the intraday price response data for a given ticker with the
// provided params from the [IEX].2.5.3 Historical Intraday Prices Endpoint.
//
// Any zero value arguments will be left off the query string & whatever Tiingo's
// default for an empty query string will be returned.
func (c *Client) IexHistory(ctx context.Context, ticker string, startDate, endDate time.Time,
	resampleFreq IexFreq, afterHours, forceFill bool, respFormat Format) ([]byte, error) {
	// Build url
	url := IexHistoryUrl(ticker, startDate, endDate, resampleFreq, afterHours, forceFill, respFormat)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultIexHistory returns the intraday price response for the given ticker from the
// [IEX].2.5.3 Historical Intraday Prices Endpoint.
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultIexHistory(ctx context.Context, ticker string) ([]byte, error) {
	// Build url
	url := IexHistoryUrl(ticker, time.Time{}, time.Time{}, "", false, false, "")

	// Fetch the data
	return c.get(ctx, url)
}

// IexHistoryUrl returns a built url for the given ticker with the provided params
// from the [IEX].2.5.3 Historical Intraday Prices Endpoint.
//
// Any zero value arguments will be left off the query string.
func IexHistoryUrl(ticker string, startDate, endDate time.Time, resampleFreq IexFreq,
	afterHours, forceFill bool, respFormat Format) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/iex/")
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
	if afterHours {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("afterHours=true")
	}
	if forceFill {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("forceFill=true")
	}
	if respFormat != "" {
		if first {
			url.WriteString("?")
		} else {
			url.WriteString("&")
		}
		url.WriteString("format=")
		url.WriteString(respFormat)
	}

	return url.String()
}
