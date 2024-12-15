package tiingo

import (
	"context"
	"strconv"
	"strings"
)

// Search returns the search result response for a given query from the
// [Utility].4.1.2 Search Endpoint.
//
// Any zero value arguments will be left off the query string & whatever Tiingo's
// default for an empty query string will be returned.
func (c *Client) Search(ctx context.Context, query string, exactMatch, includeDelisted bool,
	limit int, respFormat Format, columns []string) ([]byte, error) {
	// Build url
	url := SearchUrl(query, exactMatch, includeDelisted, limit, respFormat, columns)

	// Fetch the data
	return c.get(ctx, url)
}

// DefaultSearch search result response for a given query from the
// [Utility].4.1.2 Search Endpoint.
//
// Only the required params are added to the url & query string, everything else
// will be the Tiingo defaults.
func (c *Client) DefaultSearch(ctx context.Context, query string) ([]byte, error) {
	// Build url
	url := SearchUrl(query, false, false, 0, "", nil)

	// Fetch the data
	return c.get(ctx, url)
}

// SearchUrl returns a built url for the given query from the
// [Utility].4.1.2 Search Endpoint.
//
// Any zero value arguments will be left off the query string.
func SearchUrl(query string, exactTickerMatch, includeDelisted bool, limit int,
	respFormat Format, columns []string) string {
	var url strings.Builder

	// Build base endpoint url
	url.WriteString("https://api.tiingo.com/tiingo/utilities/search/")
	url.WriteString(query)

	// Build query string
	first := true
	if exactTickerMatch {
		url.WriteString("?exactTickerMatch=true")
		first = false
	}
	if includeDelisted {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("includeDelisted=true")
	}
	if limit > 0 {
		if first {
			url.WriteString("?")
			first = false
		} else {
			url.WriteString("&")
		}
		url.WriteString("limit=")
		url.WriteString(strconv.Itoa(limit))
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
