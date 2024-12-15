package tiingo

import (
	"testing"
	"time"
)

func TestStmtDefsUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		tickers    []string
		respFormat Format
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{},
			want: "https://api.tiingo.com/tiingo/fundamentals/definitions",
		},
		{
			name: "tickersOne",
			args: args{
				tickers: []string{"AAPL"},
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/definitions?tickers=AAPL",
		},
		{
			name: "tickersMany",
			args: args{
				tickers: []string{"AAPL", "MSFT", "GOOG", "TSLA"},
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/definitions?tickers=AAPL,MSFT,GOOG,TSLA",
		},
		{
			name: "respFormat",
			args: args{
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/definitions?format=csv",
		},
		{
			name: "combined",
			args: args{
				tickers:    []string{"AAPL", "MSFT", "GOOG", "TSLA"},
				respFormat: JSON,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/definitions?tickers=AAPL,MSFT,GOOG,TSLA&format=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StmtDefsUrl(tt.args.tickers, tt.args.respFormat); got != tt.want {
				t.Errorf("StmtDefsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStmtValsUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		ticker     string
		asReported bool
		startDate  time.Time
		endDate    time.Time
		sort       Sort
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
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements",
		},
		{
			name: "asReported",
			args: args{
				ticker:     "AAPL",
				asReported: true,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?asReported=true",
		},
		{
			name: "startDate",
			args: args{
				ticker:    "AAPL",
				startDate: startDate,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?startDate=2024-01-01",
		},
		{
			name: "endDate",
			args: args{
				ticker:  "AAPL",
				endDate: endDate,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?endDate=2024-01-02",
		},
		{
			name: "sort",
			args: args{
				ticker: "AAPL",
				sort:   DateAsc,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?sort=date",
		},
		{
			name: "respFormat",
			args: args{
				ticker:     "AAPL",
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?format=csv",
		},
		{
			name: "combinedTwo",
			args: args{
				ticker:  "AAPL",
				endDate: endDate,
				sort:    DateDesc,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?endDate=2024-01-02&sort=-date",
		},
		{
			name: "combinedAll",
			args: args{
				ticker:     "AAPL",
				asReported: true,
				startDate:  startDate,
				endDate:    endDate,
				sort:       DateDesc,
				respFormat: JSON,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?asReported=true&startDate=2024-01-01&endDate=2024-01-02&sort=-date&format=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StmtDataUrl(
				tt.args.ticker,
				tt.args.asReported,
				tt.args.startDate,
				tt.args.endDate,
				tt.args.sort,
				tt.args.respFormat,
			); got != tt.want {
				t.Errorf("StmtDataUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDailyFundamentalUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		ticker     string
		startDate  time.Time
		endDate    time.Time
		sort       Sort
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
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily",
		},
		{
			name: "startDate",
			args: args{
				ticker:    "AAPL",
				startDate: startDate,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?startDate=2024-01-01",
		},
		{
			name: "endDate",
			args: args{
				ticker:  "AAPL",
				endDate: endDate,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?endDate=2024-01-02",
		},
		{
			name: "sort",
			args: args{
				ticker: "AAPL",
				sort:   MktCapAsc,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=marketCap",
		},
		{
			name: "respFormat",
			args: args{
				ticker:     "AAPL",
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?format=csv",
		},
		{
			name: "combinedTwo",
			args: args{
				ticker:  "AAPL",
				endDate: endDate,
				sort:    PBRatioDesc,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?endDate=2024-01-02&sort=-pbRatio",
		},
		{
			name: "combinedAll",
			args: args{
				ticker:     "AAPL",
				startDate:  startDate,
				endDate:    endDate,
				sort:       DateAsc,
				respFormat: JSON,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?startDate=2024-01-01&endDate=2024-01-02&sort=date&format=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DailyFundamentalUrl(
				tt.args.ticker,
				tt.args.startDate,
				tt.args.endDate,
				tt.args.sort,
				tt.args.respFormat,
			); got != tt.want {
				t.Errorf("DailyFundamentalUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFundamentalMetadataUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		tickers    []string
		respFormat Format
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{},
			want: "https://api.tiingo.com/tiingo/fundamentals/meta",
		},
		{
			name: "TickersOne",
			args: args{
				tickers: []string{"AAPL"},
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/meta?tickers=AAPL",
		},
		{
			name: "TickersMany",
			args: args{
				tickers: []string{"AAPL", "MSFT", "GOOG", "TSLA"},
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/meta?tickers=AAPL,MSFT,GOOG,TSLA",
		},
		{
			name: "respFormat",
			args: args{
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/meta?format=csv",
		},
		{
			name: "combined",
			args: args{
				tickers:    []string{"AAPL", "MSFT", "GOOG", "TSLA"},
				respFormat: JSON,
			},
			want: "https://api.tiingo.com/tiingo/fundamentals/meta?tickers=AAPL,MSFT,GOOG,TSLA&format=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FundamentalMetadataUrl(tt.args.tickers, tt.args.respFormat); got != tt.want {
				t.Errorf("FundamentalMetadataUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
