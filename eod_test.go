package tiingo

import (
	"testing"
	"time"
)

var (
	startDate = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate   = time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
)

func TestEodPriceURL(t *testing.T) {
	t.Parallel()

	type args struct {
		ticker       string
		startDate    time.Time
		endDate      time.Time
		resampleFreq EodFreq
		sort         Sort
		respFormat   Format
		columns      []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				ticker: "AAPL",
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "startDate",
			args: args{
				ticker:    "AAPL",
				startDate: startDate,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?startDate=2024-01-01",
		},
		{
			name: "endDate",
			args: args{
				ticker:  "AAPL",
				endDate: endDate,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?endDate=2024-01-02",
		},
		{
			name: "frequency",
			args: args{
				ticker:       "AAPL",
				resampleFreq: Annually,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=annually",
		},
		{
			name: "sort",
			args: args{
				ticker: "AAPL",
				sort:   DateAsc,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=date",
		},
		{
			name: "respFormat",
			args: args{
				ticker:     "AAPL",
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?format=csv",
		},
		{
			name: "columnsOne",
			args: args{
				ticker:  "AAPL",
				columns: []string{"adjClose"},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?columns=adjClose",
		},
		{
			name: "columnsMany",
			args: args{
				ticker:  "AAPL",
				columns: []string{"adjClose", "open", "volume"},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?columns=adjClose,open,volume",
		},
		{
			name: "combinedTwo",
			args: args{
				ticker:     "AAPL",
				respFormat: CSV,
				columns:    []string{"adjClose", "open", "volume"},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?format=csv&columns=adjClose,open,volume",
		},
		{
			name: "combinedAll",
			args: args{
				ticker:       "AAPL",
				startDate:    startDate,
				endDate:      endDate,
				resampleFreq: Daily,
				sort:         DateDesc,
				respFormat:   JSON,
				columns:      []string{"volume", "adjClose", "divCash"},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?startDate=2024-01-01&endDate=2024-01-02&resampleFreq=daily&sort=-date&format=json&columns=volume,adjClose,divCash",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EodPriceUrl(
				tt.args.ticker,
				tt.args.startDate,
				tt.args.endDate,
				tt.args.resampleFreq,
				tt.args.sort,
				tt.args.respFormat,
				tt.args.columns,
			); got != tt.want {
				t.Errorf("EodUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEodMetadataUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		ticker     string
		respFormat Format
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				ticker: "AAPL",
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL",
		},
		{
			name: "respFormat",
			args: args{
				ticker:     "AAPL",
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL?format=csv",
		},
		{
			name: "combined",
			args: args{
				ticker:     "AAPL",
				respFormat: JSON,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL?format=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EodMetadataUrl(tt.args.ticker, tt.args.respFormat); got != tt.want {
				t.Errorf("EodMetadataUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
