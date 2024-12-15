package tiingo

import (
	"context"
	"strings"
	"time"
)

// StmtDefs returns the statement definition response data with the
// provided params from the [Fundamentals].2.6.2 Definitions Data Endpoint.
//
// Any zero value arguments will be left off the query string & whatever Tiingo's
// default for an empty query string will be returned.
func (c *Client) StmtDefs(ctx context.Context, tickers []string, respFormat Format) ([]byte, error) {
	// Build url
	url := StmtDefsUrl(tickers, respFormat)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultStmtDefs returns the statement definition response data from the
// [Fundamentals].2.6.2 Definitions Data Endpoint.
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultStmtDefs(ctx context.Context) ([]byte, error) {
	// Build url
	url := StmtDefsUrl(nil, "")

	// Fetch the data
	return c.get(ctx, url)
}

// StmtDefsUrl returns a built url for with the provided params
// from the [Fundamentals].2.6.2 Definitions Data Endpoint.
//
// Any zero value arguments will be left off the query string.
func StmtDefsUrl(tickers []string, respFormat Format) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/fundamentals/definitions")

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

// StmtData returns the statement values response data for the given ticker
// with the provided params from the [Fundamentals].2.6.3 Statement Data Endpoint.
//
// Any zero value arguments will be left off the query string & whatever Tiingo's
// default for an empty query string will be returned.
func (c *Client) StmtData(ctx context.Context, ticker string, asReported bool,
	startDate, endDate time.Time, sort Sort, respFormat Format) ([]byte, error) {
	// Build url
	url := StmtDataUrl(ticker, asReported, startDate, endDate, sort, respFormat)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultStmtData returns the statement values response data for the given ticker
// from the [Fundamentals].2.6.3 Statement Data Endpoint.
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultStmtData(ctx context.Context, ticker string) ([]byte, error) {
	// Build url
	url := StmtDataUrl(ticker, false, time.Time{}, time.Time{}, "", "")

	// Fetch the data
	return c.get(ctx, url)
}

// StmtDataUrl returns a built url for with the provided params
// from the [Fundamentals].2.6.3 Statement Data Endpoint.
//
// Any zero value arguments will be left off the query string.
func StmtDataUrl(ticker string, asReported bool, startDate, endDate time.Time,
	sort Sort, respFormat Format) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/fundamentals/")
	url.WriteString(ticker)
	url.WriteString("/statements")

	// Build query string
	first := true
	if asReported {
		url.WriteString("?asReported=true")
		first = false
	}
	if !startDate.IsZero() {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("startDate=")
		url.WriteString(startDate.Format("2006-01-02"))
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
