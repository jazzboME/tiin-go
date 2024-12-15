package tiingo

import (
	"testing"
	"time"
)

func TestIexTopOfBookUrl(t *testing.T) {
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
			want: "https://api.tiingo.com/iex",
		},
		{
			name: "oneSymbol",
			args: args{
				tickers: []string{"AAPL"},
			},
			want: "https://api.tiingo.com/iex/AAPL",
		},
		{
			name: "manySymbols",
			args: args{
				tickers: []string{"AAPL", "MSFT", "GOOG", "TSLA"},
			},
			want: "https://api.tiingo.com/iex/AAPL,MSFT,GOOG,TSLA",
		},
		{
			name: "respFormat",
			args: args{
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/iex?format=csv",
		},
		{
			name: "combined",
			args: args{
				tickers:    []string{"AAPL", "MSFT", "GOOG", "TSLA"},
				respFormat: JSON,
			},
			want: "https://api.tiingo.com/iex/AAPL,MSFT,GOOG,TSLA?format=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IexTopOfBookUrl(tt.args.tickers, tt.args.respFormat); got != tt.want {
				t.Errorf("IexTopOfBookUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIexHistoryUrl(t *testing.T) {
	t.Parallel()

	type args struct {
		ticker       string
		startDate    time.Time
		endDate      time.Time
		resampleFreq IexFreq
		afterHours   bool
		forceFill    bool
		respFormat   Format
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
			want: "https://api.tiingo.com/iex/AAPL/prices",
		},
		{
			name: "startDate",
			args: args{
				ticker:    "AAPL",
				startDate: startDate,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?startDate=2024-01-01",
		},
		{
			name: "endDate",
			args: args{
				ticker:  "AAPL",
				endDate: endDate,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?endDate=2024-01-02",
		},
		{
			name: "frequency",
			args: args{
				ticker:       "AAPL",
				resampleFreq: OneMin,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=1min",
		},
		{
			name: "afterHours",
			args: args{
				ticker:     "AAPL",
				afterHours: true,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?afterHours=true",
		},
		{
			name: "forceFill",
			args: args{
				ticker:    "AAPL",
				forceFill: true,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?forceFill=true",
		},
		{
			name: "respFormat",
			args: args{
				ticker:     "AAPL",
				respFormat: CSV,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?format=csv",
		},
		{
			name: "combinedTwo",
			args: args{
				ticker:       "AAPL",
				endDate:      endDate,
				resampleFreq: OneHour,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?endDate=2024-01-02&resampleFreq=1hour",
		},
		{
			name: "combinedAll",
			args: args{
				ticker:       "AAPL",
				startDate:    startDate,
				endDate:      endDate,
				resampleFreq: OneHour,
				afterHours:   true,
				forceFill:    true,
				respFormat:   JSON,
			},
			want: "https://api.tiingo.com/iex/AAPL/prices?startDate=2024-01-01&endDate=2024-01-02&resampleFreq=1hour&afterHours=true&forceFill=true&format=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IexHistoryUrl(
				tt.args.ticker,
				tt.args.startDate,
				tt.args.endDate,
				tt.args.resampleFreq,
				tt.args.afterHours,
				tt.args.forceFill,
				tt.args.respFormat,
			); got != tt.want {
				t.Errorf("IexHistoryUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
