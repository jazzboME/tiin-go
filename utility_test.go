package tiingo

import (
	"context"
	"testing"
)

var commonSearchTests = []struct {
	name   string
	query  string
	params *SearchParams
	url    string
}{
	{
		name:   "nilParams",
		query:  "AAPL",
		params: nil,
		url:    "https://api.tiingo.com/tiingo/utilities/search/AAPL",
	},
	{
		name:   "zeroParams",
		query:  "AAPL",
		params: &SearchParams{},
		url:    "https://api.tiingo.com/tiingo/utilities/search/AAPL",
	},
	{
		name:  "exactMatch",
		query: "AAPL",
		params: &SearchParams{
			ExactMatch: true,
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL?exactTickerMatch=true",
	},
	{
		name:  "includeDelisted",
		query: "AAPL",
		params: &SearchParams{
			IncludeDelisted: true,
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL?includeDelisted=true",
	},
	{
		name:  "limit",
		query: "AAPL",
		params: &SearchParams{
			Limit: 10,
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL?limit=10",
	},
	{
		name:  "respFormatCsv",
		query: "AAPL",
		params: &SearchParams{
			RespFormat: CSV,
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL?format=csv",
	},
	{
		name:  "respFormatJson",
		query: "AAPL",
		params: &SearchParams{
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL?format=json",
	},
	{
		name:  "nilColumns",
		query: "AAPL",
		params: &SearchParams{
			Columns: nil,
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL",
	},
	{
		name:  "emptyColumns",
		query: "AAPL",
		params: &SearchParams{
			Columns: []string{},
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL",
	},
	{
		name:  "oneColumn",
		query: "AAPL",
		params: &SearchParams{
			Columns: []string{"assetType"},
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL?columns=assetType",
	},
	{
		name:  "manyColumns",
		query: "AAPL",
		params: &SearchParams{
			Columns: []string{"ticker", "name", "assetType", "isActive",
				"permaTicker", "openFIGIComposite", "countryCode"},
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL" +
			"?columns=ticker,name,assetType,isActive,permaTicker,openFIGIComposite,countryCode",
	},
	{
		name:  "allQueryParams",
		query: "AAPL",
		params: &SearchParams{
			ExactMatch:      true,
			IncludeDelisted: true,
			Limit:           10,
			RespFormat:      JSON,
			Columns: []string{"ticker", "name", "assetType", "isActive",
				"permaTicker", "openFIGIComposite", "countryCode"},
		},
		url: "https://api.tiingo.com/tiingo/utilities/search/AAPL?exactTickerMatch=true&includeDelisted=true" +
			"&limit=10&format=json&columns=ticker,name,assetType,isActive,permaTicker,openFIGIComposite,countryCode",
	},
}

func TestSearchUrl(t *testing.T) {
	type args struct {
		query       string
		queryParams *SearchParams
	}
	type test struct {
		name string
		args args
		want string
	}
	var tests []test

	// Add common tests
	for _, tt := range commonSearchTests {
		tests = append(tests, struct {
			name string
			args args
			want string
		}{
			name: tt.name,
			args: args{
				query:       tt.query,
				queryParams: tt.params,
			},
			want: tt.url,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchUrl(tt.args.query, tt.args.queryParams); got != tt.want {
				t.Errorf("SearchUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Search(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		query       string
		queryParams *SearchParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonSearchTests {
		tests = append(tests, struct {
			name    string
			args    args
			wantErr bool
		}{
			name: tt.name,
			args: args{
				ctx:         ctx,
				query:       tt.query,
				queryParams: tt.params,
			},
			wantErr: false,
		})
	}

	// Add invalid argument tests
	tests = append(tests, []test{
		{
			name: "invalidRespFormat",
			args: args{
				ctx: ctx,
				queryParams: &SearchParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidColumns",
			args: args{
				ctx: ctx,
				queryParams: &SearchParams{
					Columns: []string{"BAD COLUMN// "},
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[[]SearchResult]("Search()", tt.wantErr, func() ([]SearchResult, error) {
				return getClient().Search(tt.args.ctx, tt.args.query, tt.args.queryParams)
			}); err != nil {
				t.Error(err)
			}
		})
	}
}
