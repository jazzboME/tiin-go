package tiingo

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/gocarina/gocsv"
)

// DefaultSymbolList returns the full list of SymbolRespItem's from the
// [End-of-Day].2.1.3.supported_tickers.zip Endpoint
//
// Note: This returns the raw zipped csv bytes from the endpoint, you must unzip
// before unmarshalling the csv bytes (or pass the raw bytes into FilteredSymbolList
// or ParseSymbolListCSV).
func (c *Client) DefaultSymbolList(ctx context.Context) ([]byte, error) {
	// Base URL
	url := "https://apimedia.tiingo.com/docs/tiingo/daily/supported_tickers.zip"

	return c.get(ctx, url)
}

// FilteredSymbolList returns a filtered list of SymbolRespItem's from the
// [End-of-Day].2.1.3.supported_tickers.zip Endpoint
func (c *Client) FilteredSymbolList(ctx context.Context, f SymbolFilterFunc) ([]SymbolItem, error) {
	// Get the raw data
	rawBytes, err := c.DefaultSymbolList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticker list: %w", err)
	}

	// Parse the zipped csv
	symbolItems, err := ParseSymbolListCSV(rawBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ticker list bytes: %w", err)
	}

	// Filter the tickers
	var valid []SymbolItem
	for _, symbolItem := range symbolItems {
		if f(symbolItem) {
			valid = append(valid, symbolItem)
		}
	}

	return valid, nil
}

// ParseSymbolListCSV accepts the raw zipped csv bytes returned from the
// End-of-Day].2.1.3.supported_tickers.zip Endpoint and parses them into a list
// of the individual SymbolItem
func ParseSymbolListCSV(rawData []byte) ([]SymbolItem, error) {
	// Parse the raw bytes
	zipReader, err := zip.NewReader(bytes.NewReader(rawData), int64(len(rawData)))
	if err != nil {
		return nil, fmt.Errorf("could not read bytes from zip: %w", err)
	}
	if len(zipReader.File) < 1 {
		return nil, errors.New("nothing in zip")
	}

	file, err := zipReader.File[0].Open()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	var data []SymbolItem
	if err = gocsv.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal csv bytes: %w", err)
	}

	return data, nil
}
