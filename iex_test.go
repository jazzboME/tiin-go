package tiingo

import (
	"context"
	"testing"
)

var commonIexTopOfBookTests = []struct {
	name   string
	params *IexTopOfBookParams
	url    string
}{
	{
		name:   "nilParams",
		params: nil,
		url:    "https://api.tiingo.com/iex",
	},
	{
		name:   "zeroParams",
		params: &IexTopOfBookParams{},
		url:    "https://api.tiingo.com/iex",
	},
	{
		name: "oneTicker",
		params: &IexTopOfBookParams{
			Tickers: []string{"AAPL"},
		},
		url: "https://api.tiingo.com/iex/AAPL",
	},
	{
		name: "manyTickers",
		params: &IexTopOfBookParams{
			Tickers: []string{"AAPL", "MSFT", "GOOG"},
		},
		url: "https://api.tiingo.com/iex/AAPL,MSFT,GOOG",
	},
	{
		name: "respFormatCsv",
		params: &IexTopOfBookParams{
			RespFormat: CSV,
		},
		url: "https://api.tiingo.com/iex?format=csv",
	},
	{
		name: "respFormatJson",
		params: &IexTopOfBookParams{
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/iex?format=json",
	},
	{
		name: "allQueryParams",
		params: &IexTopOfBookParams{
			Tickers:    []string{"AAPL", "MSFT", "GOOG"},
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/iex/AAPL,MSFT,GOOG?format=json",
	},
}

func TestIexTopOfBookUrl(t *testing.T) {
	type args struct {
		queryParams *IexTopOfBookParams
	}
	type test struct {
		name string
		args args
		want string
	}
	var tests []test

	// Add common tests
	for _, tt := range commonIexTopOfBookTests {
		tests = append(tests, struct {
			name string
			args args
			want string
		}{
			name: tt.name,
			args: args{
				queryParams: tt.params,
			},
			want: tt.url,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IexTopOfBookUrl(tt.args.queryParams); got != tt.want {
				t.Errorf("IexTopOfBookUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_IexTopOfBook(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		queryParams *IexTopOfBookParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonIexTopOfBookTests {
		tests = append(tests, struct {
			name    string
			args    args
			wantErr bool
		}{
			name: tt.name,
			args: args{
				ctx:         ctx,
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
				queryParams: &IexTopOfBookParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidTickers",
			args: args{
				ctx: ctx,
				queryParams: &IexTopOfBookParams{
					Tickers: []string{"BAD TICKER// "},
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest("IexTopOfBook()", tt.wantErr, func() ([]IexTopOfBook, error) {
				return getClient().IexTopOfBook(tt.args.ctx, tt.args.queryParams)
			}); err != nil {
				t.Error(err)
			}
		})
	}
}

var commonIexHistoryTests = []struct {
	name   string
	ticker string
	params *IexHistoryParams
	url    string
}{
	{
		name:   "nilParams",
		ticker: "AAPL",
		params: nil,
		url:    "https://api.tiingo.com/iex/AAPL/prices",
	},
	{
		name:   "zeroParams",
		ticker: "AAPL",
		params: &IexHistoryParams{},
		url:    "https://api.tiingo.com/iex/AAPL/prices",
	},
	{
		name:   "startDate",
		ticker: "AAPL",
		params: &IexHistoryParams{
			StartDate: startDate,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?startDate=2022-01-01",
	},
	{
		name:   "endDate",
		ticker: "AAPL",
		params: &IexHistoryParams{
			EndDate: endDate,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?endDate=2024-01-01",
	},
	{
		name:   "resampleFreqOneMin",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ResampleFreq: OneMin,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=1min",
	},
	{
		name:   "resampleFreqFiveMin",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ResampleFreq: FiveMin,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=5min",
	},
	{
		name:   "resampleFreqFifteenMin",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ResampleFreq: FifteenMin,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=15min",
	},
	{
		name:   "resampleFreqThirtyMin",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ResampleFreq: ThirtyMin,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=30min",
	},
	{
		name:   "resampleFreqOneHour",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ResampleFreq: OneHour,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=1hour",
	},
	{
		name:   "resampleFreqTwoHour",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ResampleFreq: TwoHour,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=2hour",
	},
	{
		name:   "resampleFreqFourHour",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ResampleFreq: FourHour,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?resampleFreq=4hour",
	},
	{
		name:   "afterHours",
		ticker: "AAPL",
		params: &IexHistoryParams{
			AfterHours: true,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?afterHours=true",
	},
	{
		name:   "forceFill",
		ticker: "AAPL",
		params: &IexHistoryParams{
			ForceFill: true,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?forceFill=true",
	},
	{
		name:   "respFormatJson",
		ticker: "AAPL",
		params: &IexHistoryParams{
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?format=json",
	},
	{
		name:   "respFormatCsv",
		ticker: "AAPL",
		params: &IexHistoryParams{
			RespFormat: CSV,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?format=csv",
	},
	{
		name:   "allParams",
		ticker: "AAPL",
		params: &IexHistoryParams{
			StartDate:    startDate,
			EndDate:      endDate,
			ResampleFreq: FourHour,
			AfterHours:   true,
			ForceFill:    true,
			RespFormat:   CSV,
		},
		url: "https://api.tiingo.com/iex/AAPL/prices?startDate=2022-01-01&endDate=2024-01-01" +
			"&resampleFreq=4hour&afterHours=true&forceFill=true&format=csv",
	},
}

func TestIexHistoryUrl(t *testing.T) {
	type args struct {
		ticker      string
		queryParams *IexHistoryParams
	}
	type test struct {
		name string
		args args
		want string
	}
	var tests []test

	// Add common tests
	for _, tt := range commonIexHistoryTests {
		tests = append(tests, struct {
			name string
			args args
			want string
		}{
			name: tt.name,
			args: args{
				ticker:      tt.ticker,
				queryParams: tt.params,
			},
			want: tt.url,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IexHistoryUrl(tt.args.ticker, tt.args.queryParams); got != tt.want {
				t.Errorf("IexHistoryUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_IexHistory(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		ticker      string
		queryParams *IexHistoryParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonIexHistoryTests {
		tests = append(tests, struct {
			name    string
			args    args
			wantErr bool
		}{
			name: tt.name,
			args: args{
				ctx:         ctx,
				ticker:      tt.ticker,
				queryParams: tt.params,
			},
			wantErr: false,
		})
	}

	// Add invalid argument tests
	tests = append(tests, []test{
		{
			name: "invalidResampleFreq",
			args: args{
				ctx: ctx,
				queryParams: &IexHistoryParams{
					ResampleFreq: "BAD FREQUENCY",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidRespFormat",
			args: args{
				ctx: ctx,
				queryParams: &IexHistoryParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest("IexHistory()", tt.wantErr, func() ([]IexPrice, error) {
				return getClient().IexHistory(tt.args.ctx, tt.args.ticker, tt.args.queryParams)
			}); err != nil {
				t.Error(err)
			}
		})
	}
}
