package tiingo

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// IexTopOfBookParams represents the query parameters for the [IEX].2.5.2 Top-of-Book
// & Last Price Endpoints
type IexTopOfBookParams struct {
	Tickers    []string
	RespFormat Format
}

// IexTopOfBook returns the last price item for the specified tickers from
// the [IEX].2.5.2 Top-of-Book & Last Price Endpoints.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) IexTopOfBook(ctx context.Context, queryParams *IexTopOfBookParams) ([]IexTopOfBook, error) {
	// Fetch the data
	rawBytes, err := c.IexTopOfBookRaw(ctx, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	var format string
	if queryParams != nil {
		format = queryParams.RespFormat
	}
	return Parse[[]IexTopOfBook](rawBytes, format)
}

// IexTopOfBookRaw functions the same as IexTopOfBook, except the raw response
// bytes are returned instead of the parsed type.
func (c *Client) IexTopOfBookRaw(ctx context.Context, queryParams *IexTopOfBookParams) ([]byte, error) {
	// Build url
	url := IexTopOfBookUrl(queryParams)

	// Fetch the data
	return c.get(ctx, url)
}

// IexTopOfBookUrl returns a built url for the given tickers from the
// [IEX].2.5.2 Top-of-Book & Last Price Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func IexTopOfBookUrl(queryParams *IexTopOfBookParams) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/iex")

	// No query params to add
	if queryParams == nil {
		return url.String()
	}

	// Build query string
	if len(queryParams.Tickers) > 0 {
		url.WriteString("/")
		url.WriteString(strings.Join(queryParams.Tickers, ","))
	}
	if queryParams.RespFormat != "" {
		url.WriteString("?format=")
		url.WriteString(queryParams.RespFormat)
	}

	return url.String()
}

type IexHistoryParams struct {
	StartDate    time.Time
	EndDate      time.Time
	ResampleFreq IexFreq
	AfterHours   bool
	ForceFill    bool
	RespFormat   Format
}

// IexHistory returns the intraday price response data for a given ticker with the
// provided params from the [IEX].2.5.3 Historical Intraday Prices Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) IexHistory(ctx context.Context, ticker string, queryParams *IexHistoryParams) ([]IexPrice, error) {
	// Fetch the data
	rawBytes, err := c.IexHistoryRaw(ctx, ticker, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	var format string
	if queryParams != nil {
		format = queryParams.RespFormat
	}
	return Parse[[]IexPrice](rawBytes, format)
}

// IexHistoryRaw functions the same as IexHistory, except the raw response bytes
// are returned instead of the parsed type.
func (c *Client) IexHistoryRaw(ctx context.Context, ticker string, queryParams *IexHistoryParams) ([]byte, error) {
	// Build url
	url := IexHistoryUrl(ticker, queryParams)

	// Fetch the data
	return c.get(ctx, url)
}

// IexHistoryUrl returns a built url for the given ticker with the provided params
// from the [IEX].2.5.3 Historical Intraday Prices Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func IexHistoryUrl(ticker string, queryParams *IexHistoryParams) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/iex/")
	url.WriteString(ticker)
	url.WriteString("/prices")

	// No query params to add
	if queryParams == nil {
		return url.String()
	}

	// Build query string
	first := true
	if !queryParams.StartDate.IsZero() {
		url.WriteString("?startDate=")
		url.WriteString(queryParams.StartDate.Format("2006-01-02"))
		first = false
	}
	if !queryParams.EndDate.IsZero() {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("endDate=")
		url.WriteString(queryParams.EndDate.Format("2006-01-02"))
	}
	if queryParams.ResampleFreq != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("resampleFreq=")
		url.WriteString(queryParams.ResampleFreq)
	}
	if queryParams.AfterHours {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("afterHours=true")
	}
	if queryParams.ForceFill {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("forceFill=true")
	}
	if queryParams.RespFormat != "" {
		if first {
			url.WriteString("?")
		} else {
			url.WriteString("&")
		}
		url.WriteString("format=")
		url.WriteString(queryParams.RespFormat)
	}

	return url.String()
}
