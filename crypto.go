package tiingo

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

// CrytoPriceParams represents the query parameters for the [Real-time & Historical Prices].2.3.2
// Crypto Endpoint
type CryptoPriceParams struct {
	Exchanges	 []string
	StartDate    time.Time
	EndDate      time.Time
	ResampleFreq IexFreq
}

// EodPrice returns the daily price response data for a given ticker with the
// provided params from the [End-of-Day].2.1.2 End-of-Day Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) CryptoPrice(ctx context.Context, ticker []string,
	queryParams *CryptoPriceParams) ([]CryptoResult, error) {
	// Fetch the data
	rawBytes, err := c.CryptoPriceRaw(ctx, ticker, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	return Parse[[]CryptoResult](rawBytes, "json")
}

// EodPriceRaw functions the same as EodPrice, except the raw response bytes are
// returned instead of the parsed type.
func (c *Client) CryptoPriceRaw(ctx context.Context, ticker []string,
	queryParams *CryptoPriceParams) ([]byte, error) {
	// Build URL
	url := CryptoPriceUrl(ticker, queryParams)

	// Fetch the data
	return c.get(ctx, url)
}

// EodPriceUrl returns a built url for the given ticker with the provided params
// from the [End-of-Day].2.1.2 End-of-Day Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func CryptoPriceUrl(ticker []string, queryParams *CryptoPriceParams) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/crypto/prices")

	// No query params to add
	if queryParams == nil || len(ticker) == 0 {
		return url.String()
	}

	// Build query string
	url.WriteString("?tickers=")
	url.WriteString(strings.Join(ticker, ","))

	if len(queryParams.Exchanges) > 0 {
		url.WriteString("&exchanges=")
		url.WriteString(strings.Join(queryParams.Exchanges, ","))
	}
	if !queryParams.StartDate.IsZero() {
		url.WriteString("&startDate=")
		url.WriteString(queryParams.StartDate.Format("2006-01-02"))
	}
	if !queryParams.EndDate.IsZero() {
		url.WriteString("&endDate=")
		url.WriteString(queryParams.EndDate.Format("2006-01-02"))
	}
	if queryParams.ResampleFreq != "" {
		url.WriteString("&resampleFreq=")
		url.WriteString(queryParams.ResampleFreq)
	}
	url.WriteString("&includeRawExchangeData=true")
	
	log.Printf(url.String())
	return url.String()
}

// EodMetadata returns the eod metadata response data for a given ticker
// from the [End-of-Day].2.1.3 Meta Endpoint
func (c *Client)  CryptoMetadata(ctx context.Context, ticker string) (EodMetadata, error) {
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
func (c *Client) CryptoMetadataRaw(ctx context.Context, ticker []string) ([]byte, error) {
	// Build URL
	url := CryptoMetadataUrl(ticker)

	// Fetch the data
	return c.get(ctx, url)
}

// EodMetadataUrl returns a built url for the given ticker from the
// [End-of-Day].2.1.3 Meta Endpoint.
func CryptoMetadataUrl(ticker []string) string {
	return fmt.Sprintf("https://api.tiingo.com/tiingo/daily/%s", ticker)
}
