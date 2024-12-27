package tiingo

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type SearchParams struct {
	ExactMatch      bool
	IncludeDelisted bool
	Limit           int
	RespFormat      Format
	Columns         []string
}

// Search returns the search result response for a given query from the
// [Utility].4.1.2 Search Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func (c *Client) Search(ctx context.Context, query string, queryParams *SearchParams) ([]SearchResult, error) {
	// Fetch the data
	rawBytes, err := c.SearchRaw(ctx, query, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %w", err)
	}

	// Parse
	var format string
	if queryParams != nil {
		format = queryParams.RespFormat
	}
	return Parse[[]SearchResult](rawBytes, format)
}

// SearchRaw functions the same as Search, except the raw response bytes are
// returned instead of the parsed type.
func (c *Client) SearchRaw(ctx context.Context, query string, queryParams *SearchParams) ([]byte, error) {
	// Build url
	url := SearchUrl(query, queryParams)

	// Fetch the data
	return c.get(ctx, url)
}

// SearchUrl returns a built url for the given query from the
// [Utility].4.1.2 Search Endpoint.
//
// If queryParams is non-nil, any non-zero struct values will be applied to the
// url. Zero value items will be left out and Tiingo defaults will be used. A
// nil queryParams results in all Tiingo defaults.
func SearchUrl(query string, queryParams *SearchParams) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/utilities/search/")
	url.WriteString(query)

	// No query params to add
	if queryParams == nil {
		return url.String()
	}

	// Build query string
	first := true
	if queryParams.ExactMatch {
		url.WriteString("?exactTickerMatch=true")
		first = false
	}
	if queryParams.IncludeDelisted {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("includeDelisted=true")
	}
	if queryParams.Limit > 0 {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("limit=")
		url.WriteString(strconv.Itoa(queryParams.Limit))
	}
	if queryParams.RespFormat != "" {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("format=")
		url.WriteString(queryParams.RespFormat)
	}
	if len(queryParams.Columns) > 0 {
		if first {
			url.WriteString("?")
		} else {
			url.WriteString("&")
		}
		url.WriteString("columns=")
		url.WriteString(strings.Join(queryParams.Columns, ","))
	}

	return url.String()
}
