package tiingo

import (
	"context"
	"testing"
)

var commonStmtDefsTests = []struct {
	name   string
	params *StmtDefsParams
	url    string
}{
	{
		name:   "nilParams",
		params: nil,
		url:    "https://api.tiingo.com/tiingo/fundamentals/definitions",
	},
	{
		name:   "zeroParams",
		params: &StmtDefsParams{},
		url:    "https://api.tiingo.com/tiingo/fundamentals/definitions",
	},
	{
		name: "oneTicker",
		params: &StmtDefsParams{
			Tickers: []string{"AAPL"},
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/definitions?tickers=AAPL",
	},
	{
		name: "manyTickers",
		params: &StmtDefsParams{
			Tickers: []string{"AAPL", "MSFT", "GOOG"},
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/definitions?tickers=AAPL,MSFT,GOOG",
	},
	{
		name: "respFormatCsv",
		params: &StmtDefsParams{
			RespFormat: CSV,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/definitions?format=csv",
	},
	{
		name: "respFormatJson",
		params: &StmtDefsParams{
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/definitions?format=json",
	},
	{
		name: "allQueryParams",
		params: &StmtDefsParams{
			Tickers:    []string{"AAPL", "MSFT", "GOOG"},
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/definitions?tickers=AAPL,MSFT,GOOG&format=json",
	},
}

func TestStmtDefsUrl(t *testing.T) {
	type args struct {
		queryParams *StmtDefsParams
	}
	type test struct {
		name string
		args args
		want string
	}
	var tests []test

	// Add common tests
	for _, tt := range commonStmtDefsTests {
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
			if got := StmtDefsUrl(tt.args.queryParams); got != tt.want {
				t.Errorf("StmtDefsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StmtDefs(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		queryParams *StmtDefsParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonStmtDefsTests {
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
				queryParams: &StmtDefsParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidTickers",
			args: args{
				ctx: ctx,
				queryParams: &StmtDefsParams{
					Tickers: []string{"BAD TICKER// "},
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[[]StmtDef]("StmtDefs()", tt.wantErr, func() ([]StmtDef, error) {
				return getClient().StmtDefs(tt.args.ctx, tt.args.queryParams)
			}); err != nil {
				t.Error(err)
			}
		})
	}
}
