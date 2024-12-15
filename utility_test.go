package tiingo

import "testing"

func TestSearchUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		query            string
		exactTickerMatch bool
		includeDelisted  bool
		limit            int
		respFormat       Format
		columns          []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				query: "AAPL",
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL",
		},
		{
			name: "exactTickerMatch",
			args: args{
				query:            "AAPL",
				exactTickerMatch: true,
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?exactTickerMatch=true",
		},
		{
			name: "includeDelisted",
			args: args{
				query:           "AAPL",
				includeDelisted: true,
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?includeDelisted=true",
		},
		{
			name: "limit",
			args: args{
				query: "AAPL",
				limit: 19,
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?limit=19",
		},
		{
			name: "respFormat",
			args: args{
				query:      "AAPL",
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?format=csv",
		},
		{
			name: "columnsOne",
			args: args{
				query:   "AAPL",
				columns: []string{"ticker"},
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?columns=ticker",
		},
		{
			name: "columnsMany",
			args: args{
				query:   "AAPL",
				columns: []string{"ticker", "name", "isActive"},
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?columns=ticker,name,isActive",
		},
		{
			name: "combinedTwo",
			args: args{
				query:      "AAPL",
				limit:      19,
				respFormat: JSON,
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?limit=19&format=json",
		},
		{
			name: "combinedAll",
			args: args{
				query:            "AAPL",
				exactTickerMatch: true,
				includeDelisted:  true,
				limit:            19,
				respFormat:       CSV,
				columns:          []string{"ticker", "name", "isActive"},
			},
			want: "https://api.tiingo.com/tiingo/utilities/search/AAPL?exactTickerMatch=true&includeDelisted=true&limit=19&format=csv&columns=ticker,name,isActive",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchUrl(
				tt.args.query,
				tt.args.exactTickerMatch,
				tt.args.includeDelisted,
				tt.args.limit,
				tt.args.respFormat,
				tt.args.columns,
			); got != tt.want {
				t.Errorf("SearchUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
