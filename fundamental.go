package tiingo

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type StmtDefsParams struct {
	Tickers    []string
	RespFormat Format
}

// StmtDefs returns the statement definition response data with the
// provided params from the [Fundamentals].2.6.2 Definitions Data Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) StmtDefs(ctx context.Context, queryParams *StmtDefsParams) ([]StmtDef, error) {
	// Fetch the data
	rawBytes, err := c.StmtDefsRaw(ctx, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	var format string
	if queryParams != nil {
		format = queryParams.RespFormat
	}
	return Parse[[]StmtDef](rawBytes, format)
}

// StmtDefsRaw functions the same as StmtDefs, except the raw response
// bytes are returned instead of the parsed type.
func (c *Client) StmtDefsRaw(ctx context.Context, queryParams *StmtDefsParams) ([]byte, error) {
	// Build url
	url := StmtDefsUrl(queryParams)

	// Fetch the data
	return c.get(ctx, url)
}

// StmtDefsUrl returns a built url for with the provided params
// from the [Fundamentals].2.6.2 Definitions Data Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func StmtDefsUrl(queryParams *StmtDefsParams) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/fundamentals/definitions")

	// No query params to add
	if queryParams == nil {
		return url.String()
	}

	// Build query string
	first := true
	if len(queryParams.Tickers) > 0 {
		url.WriteString("?tickers=")
		url.WriteString(strings.Join(queryParams.Tickers, ","))
		first = false
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

type StmtDataParams struct {
	AsReported bool
	StartDate  time.Time
	EndDate    time.Time
	Sort       Sort
	RespFormat Format
}

// StmtDataFlat returns the statement values response data for the given ticker
// with the provided params from the [Fundamentals].2.6.3 Statement Data Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) StmtDataFlat(ctx context.Context, ticker string,
	queryParams *StmtDataParams) ([]StmtDataFlat, error) {
	// Ensure the response format is csv (required for flat statement data)
	if queryParams == nil {
		queryParams = &StmtDataParams{RespFormat: CSV}
	} else {
		queryParams.RespFormat = CSV
	}

	// Fetch the data
	rawBytes, err := c.StmtDataRaw(ctx, ticker, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	var format string
	if queryParams != nil {
		format = queryParams.RespFormat
	}
	return Parse[[]StmtDataFlat](rawBytes, format)
}

// StmtDataNested returns the statement values response data for the given ticker
// with the provided params from the [Fundamentals].2.6.3 Statement Data Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) StmtDataNested(ctx context.Context, ticker string,
	queryParams *StmtDataParams) ([]StmtDataNested, error) {
	// Ensure the response format is json (required for nested statement data)
	if queryParams == nil {
		queryParams = &StmtDataParams{RespFormat: JSON}
	} else {
		queryParams.RespFormat = JSON
	}

	// Fetch the data
	rawBytes, err := c.StmtDataRaw(ctx, ticker, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	var format string
	if queryParams != nil {
		format = queryParams.RespFormat
	}
	return Parse[[]StmtDataNested](rawBytes, format)
}

// StmtDataRaw functions the same as StmtDataFlat and StmtDataNested, except the raw
// response bytes are returned instead of the parsed type.
func (c *Client) StmtDataRaw(ctx context.Context, ticker string,
	queryParams *StmtDataParams) ([]byte, error) {
	// Build url
	url := StmtDataUrl(ticker, queryParams)

	// Fetch the data
	return c.get(ctx, url)
}

// StmtDataUrl returns a built url for with the provided params
// from the [Fundamentals].2.6.3 Statement Data Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func StmtDataUrl(ticker string, queryParams *StmtDataParams) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/fundamentals/")
	url.WriteString(ticker)
	url.WriteString("/statements")

	// No query params to add
	if queryParams == nil {
		return url.String()
	}

	// Build query string
	first := true
	if queryParams.AsReported {
		url.WriteString("?asReported=true")
		first = false
	}
	if !queryParams.StartDate.IsZero() {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("startDate=")
		url.WriteString(queryParams.StartDate.Format("2006-01-02"))
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
	if queryParams.Sort != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("sort=")
		url.WriteString(queryParams.Sort)
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

// DailyFundamental returns the daily fundamental response data for the given ticker
// with the provided params from the [Fundamentals].2.6.4 daily Data Endpoint.
//
// Any zero value arguments will be left off the query string & whatever Tiingo's
// default for an empty query string will be returned.
func (c *Client) DailyFundamental(ctx context.Context, ticker string,
	startDate, endDate time.Time, sort Sort, respFormat Format) ([]byte, error) {
	// Build url
	url := DailyFundamentalUrl(ticker, startDate, endDate, sort, respFormat)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultDailyFundamental returns the statement values response data for the given ticker
// from the [Fundamentals].2.6.4 daily Data Endpoint.
//
// Only the required params are added to the url & query string, everything else will
// be the Tiingo defaults.
func (c *Client) DefaultDailyFundamental(ctx context.Context, ticker string) ([]byte, error) {
	// Build url
	url := DailyFundamentalUrl(ticker, time.Time{}, time.Time{}, "", "")

	// Fetch the data
	return c.get(ctx, url)
}

// DailyFundamentalUrl returns a built url for with the provided params
// from the [Fundamentals].2.6.4 daily Data Endpoint.
//
// Any zero value arguments will be left off the query string.
func DailyFundamentalUrl(ticker string, startDate, endDate time.Time, sort Sort, respFormat Format) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/fundamentals/")
	url.WriteString(ticker)
	url.WriteString("/daily")

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
		} else {
			url.WriteString("&")
		}
		url.WriteString("format=")
		url.WriteString(respFormat)
	}

	return url.String()
}

// FundamentalMetadata returns the daily fundamental metadata for the given ticker(s)
// with the provided params from the [Fundamentals].2.6.5 MetaData Endpoint.
//
// Any zero value arguments will be left off the query string & whatever Tiingo's
// default for an empty query string will be returned.
func (c *Client) FundamentalMetadata(ctx context.Context, tickers []string, respFormat Format) ([]byte, error) {
	// Build url
	url := FundamentalMetadataUrl(tickers, respFormat)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultFundamentalMetadata returns the statement values response data for the
// all tickers from the [Fundamentals].2.6.5 MetaData Endpoint.
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultFundamentalMetadata(ctx context.Context) ([]byte, error) {
	// Build url
	url := FundamentalMetadataUrl(nil, "")

	// Fetch the data
	return c.get(ctx, url)
}

// FundamentalMetadataUrl returns a built url for with the provided params
// from the [Fundamentals].2.6.5 MetaData Endpoint
//
// Any zero value arguments will be left off the query string.
func FundamentalMetadataUrl(tickers []string, respFormat Format) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/fundamentals/meta")

	// Build query string
	first := true
	if len(tickers) > 0 {
		url.WriteString("?tickers=")
		url.WriteString(strings.Join(tickers, ","))
		first = false
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
