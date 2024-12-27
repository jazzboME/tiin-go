package tiingo

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// EodPriceParams represents the query parameters for the [End-of-Day].2.1.2
// End-of-Day Endpoint
type EodPriceParams struct {
	startDate    time.Time
	endDate      time.Time
	resampleFreq EodFreq
	sort         Sort
	respFormat   Format
	columns      []string
}

// EodPrice returns the daily price response data for a given ticker with the
// provided params from the [End-of-Day].2.1.2 End-of-Day Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) EodPrice(ctx context.Context, ticker string,
	queryParams *EodPriceParams) ([]EodPrice, error) {
	// Fetch the data
	rawBytes, err := c.EodPriceRaw(ctx, ticker, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	var format string
	if queryParams != nil {
		format = queryParams.respFormat
	}
	return Parse[[]EodPrice](rawBytes, format)
}

// EodPriceRaw functions the same as EodPrice, except the raw response bytes are
// returned instead of the parsed type.
func (c *Client) EodPriceRaw(ctx context.Context, ticker string,
	queryParams *EodPriceParams) ([]byte, error) {
	// Build URL
	url := EodPriceUrl(ticker, queryParams)

	// Fetch the data
	return c.get(ctx, url)
}

// EodPriceUrl returns a built url for the given ticker with the provided params
// from the [End-of-Day].2.1.2 End-of-Day Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func EodPriceUrl(ticker string, queryParams *EodPriceParams) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/daily/")
	url.WriteString(ticker)
	url.WriteString("/prices")

	// No query params to add
	if queryParams == nil {
		return url.String()
	}

	// Build query string
	first := true
	if !queryParams.startDate.IsZero() {
		url.WriteString("?startDate=")
		url.WriteString(queryParams.startDate.Format("2006-01-02"))
		first = false
	}
	if !queryParams.endDate.IsZero() {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("endDate=")
		url.WriteString(queryParams.endDate.Format("2006-01-02"))
	}
	if queryParams.resampleFreq != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("resampleFreq=")
		url.WriteString(queryParams.resampleFreq)
	}
	if queryParams.sort != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("sort=")
		url.WriteString(queryParams.sort)
	}
	if queryParams.respFormat != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("format=")
		url.WriteString(queryParams.respFormat)
	}
	if len(queryParams.columns) > 0 {
		if first {
			url.WriteString("?")
		} else {
			url.WriteString("&")
		}
		url.WriteString("columns=")
		url.WriteString(strings.Join(queryParams.columns, ","))
	}

	return url.String()
}

// EodMetadata returns the eod metadata response data for a given ticker
// from the [End-of-Day].2.1.3 Meta Endpoint
func (c *Client) EodMetadata(ctx context.Context, ticker string) (EodMetadata, error) {
	// Get the data
	rawBytes, err := c.EodMetadataRaw(ctx, ticker)
	if err != nil {
		return EodMetadata{}, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	return Parse[EodMetadata](rawBytes, JSON)
}

// EodMetadataRaw functions the same as EodMetadata, except the raw response
// bytes are returned instead of the parsed type.
func (c *Client) EodMetadataRaw(ctx context.Context, ticker string) ([]byte, error) {
	// Build URL
	url := EodMetadataUrl(ticker)

	// Fetch the data
	return c.get(ctx, url)
}

// EodMetadataUrl returns a built url for the given ticker from the
// [End-of-Day].2.1.3 Meta Endpoint.
func EodMetadataUrl(ticker string) string {
	return fmt.Sprintf("https://api.tiingo.com/tiingo/daily/%s", ticker)
}
