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

var commonStmtDataTests = []struct {
	name   string
	ticker string
	params *StmtDataParams
	url    string
}{
	{
		name:   "nilParams",
		ticker: "AAPL",
		params: nil,
		url:    "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements",
	},
	{
		name:   "zeroParams",
		ticker: "AAPL",
		params: &StmtDataParams{},
		url:    "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements",
	},
	{
		name:   "asReported",
		ticker: "AAPL",
		params: &StmtDataParams{
			AsReported: true,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?asReported=true",
	},
	{
		name:   "startDate",
		ticker: "AAPL",
		params: &StmtDataParams{
			StartDate: startDate,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?startDate=2022-01-01",
	},
	{
		name:   "endDate",
		ticker: "AAPL",
		params: &StmtDataParams{
			EndDate: endDate,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?endDate=2024-01-01",
	},
	{
		name:   "sortDateAsc",
		ticker: "AAPL",
		params: &StmtDataParams{
			Sort: DateAsc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?sort=date",
	},
	{
		name:   "sortDateDesc",
		ticker: "AAPL",
		params: &StmtDataParams{
			Sort: DateDesc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?sort=-date",
	},
	{
		name:   "respFormatJson",
		ticker: "AAPL",
		params: &StmtDataParams{
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?format=json",
	},
	{
		name:   "respFormatCsv",
		ticker: "AAPL",
		params: &StmtDataParams{
			RespFormat: CSV,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?format=csv",
	},
	{
		name:   "allParams",
		ticker: "AAPL",
		params: &StmtDataParams{
			AsReported: true,
			StartDate:  startDate,
			EndDate:    endDate,
			Sort:       DateDesc,
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/statements?asReported=true" +
			"&startDate=2022-01-01&endDate=2024-01-01&sort=-date&format=json",
	},
}

func TestStmtDataUrl(t *testing.T) {
	type args struct {
		ticker      string
		queryParams *StmtDataParams
	}
	type test struct {
		name string
		args args
		want string
	}
	var tests []test

	// Add common tests
	for _, tt := range commonStmtDataTests {
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
			if got := StmtDataUrl(tt.args.ticker, tt.args.queryParams); got != tt.want {
				t.Errorf("StmtDataUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StmtDataFlat(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		ticker      string
		queryParams *StmtDataParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonStmtDataTests {
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
			name: "invalidSort",
			args: args{
				ctx: ctx,
				queryParams: &StmtDataParams{
					Sort: "BAD SORT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidRespFormat",
			args: args{
				ctx: ctx,
				queryParams: &StmtDataParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[[]StmtDataFlat]("StmtDataFlat()", tt.wantErr,
				func() ([]StmtDataFlat, error) {
					return getClient().StmtDataFlat(tt.args.ctx, tt.args.ticker, tt.args.queryParams)
				},
			); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestClient_StmtDataNested(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		ticker      string
		queryParams *StmtDataParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonStmtDataTests {
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
			name: "invalidSort",
			args: args{
				ctx: ctx,
				queryParams: &StmtDataParams{
					Sort: "BAD SORT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidRespFormat",
			args: args{
				ctx: ctx,
				queryParams: &StmtDataParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[[]StmtDataNested]("StmtDataNested()", tt.wantErr,
				func() ([]StmtDataNested, error) {
					return getClient().StmtDataNested(tt.args.ctx, tt.args.ticker, tt.args.queryParams)
				},
			); err != nil {
				t.Error(err)
			}
		})
	}
}

var commonDailyFundamentalTests = []struct {
	name   string
	ticker string
	params *DailyFundamentalParams
	url    string
}{
	{
		name:   "nilParams",
		ticker: "AAPL",
		params: nil,
		url:    "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily",
	},
	{
		name:   "zeroParams",
		ticker: "AAPL",
		params: &DailyFundamentalParams{},
		url:    "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily",
	},
	{
		name:   "startDate",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			StartDate: startDate,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?startDate=2022-01-01",
	},
	{
		name:   "endDate",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			EndDate: endDate,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?endDate=2024-01-01",
	},
	{
		name:   "sortDateAsc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: DateAsc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=date",
	},
	{
		name:   "sortDateDesc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: DateDesc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=-date",
	},
	{
		name:   "sortMktCapAsc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: MktCapAsc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=marketCap",
	},
	{
		name:   "sortMktCapDesc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: MktCapDesc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=-marketCap",
	},
	{
		name:   "sortEntValAsc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: EntValAsc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=enterpriseVal",
	},
	{
		name:   "sortEntValDesc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: EntValDesc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=-enterpriseVal",
	},
	{
		name:   "sortPERatioAsc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: PERatioAsc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=peRatio",
	},
	{
		name:   "sortPERatioDesc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: PERatioDesc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=-peRatio",
	},
	{
		name:   "sortPBRatioAsc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: PBRatioAsc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=pbRatio",
	},
	{
		name:   "sortPBRatioDesc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: PBRatioDesc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=-pbRatio",
	},
	{
		name:   "sortTrailPEGAsc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: TrailPEGAsc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=trailingPEG1Y",
	},
	{
		name:   "sortTrailPEGDesc",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			Sort: TrailPEGDesc,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?sort=-trailingPEG1Y",
	},
	{
		name:   "respFormatJson",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?format=json",
	},
	{
		name:   "respFormatCsv",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			RespFormat: CSV,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?format=csv",
	},
	{
		name:   "allParams",
		ticker: "AAPL",
		params: &DailyFundamentalParams{
			StartDate:  startDate,
			EndDate:    endDate,
			Sort:       DateDesc,
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/AAPL/daily?startDate=2022-01-01" +
			"&endDate=2024-01-01&sort=-date&format=json",
	},
}

func TestDailyFundamentalUrl(t *testing.T) {
	type args struct {
		ticker      string
		queryParams *DailyFundamentalParams
	}
	type test struct {
		name string
		args args
		want string
	}
	var tests []test

	// Add common tests
	for _, tt := range commonDailyFundamentalTests {
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
			if got := DailyFundamentalUrl(tt.args.ticker, tt.args.queryParams); got != tt.want {
				t.Errorf("DailyFundamentalUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DailyFundamental(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		ticker      string
		queryParams *DailyFundamentalParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonDailyFundamentalTests {
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
			name: "invalidSort",
			args: args{
				ctx: ctx,
				queryParams: &DailyFundamentalParams{
					Sort: "BAD SORT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidRespFormat",
			args: args{
				ctx: ctx,
				queryParams: &DailyFundamentalParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[[]DailyFundamental]("DailyFundamental()", tt.wantErr,
				func() ([]DailyFundamental, error) {
					return getClient().DailyFundamental(tt.args.ctx, tt.args.ticker, tt.args.queryParams)
				},
			); err != nil {
				t.Error(err)
			}
		})
	}
}

var commonFundamentalMetadataTests = []struct {
	name   string
	params *FundamentalMetadataParams
	url    string
}{
	{
		name:   "nilParams",
		params: nil,
		url:    "https://api.tiingo.com/tiingo/fundamentals/meta",
	},
	{
		name:   "zeroParams",
		params: &FundamentalMetadataParams{},
		url:    "https://api.tiingo.com/tiingo/fundamentals/meta",
	},
	{
		name: "oneTicker",
		params: &FundamentalMetadataParams{
			Tickers: []string{"AAPL"},
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/meta?tickers=AAPL",
	},
	{
		name: "manyTicker",
		params: &FundamentalMetadataParams{
			Tickers: []string{"AAPL", "MSFT", "GOOG"},
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/meta?tickers=AAPL,MSFT,GOOG",
	},
	{
		name: "respFormatJson",
		params: &FundamentalMetadataParams{
			RespFormat: JSON,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/meta?format=json",
	},
	{
		name: "respFormatCsv",
		params: &FundamentalMetadataParams{
			RespFormat: CSV,
		},
		url: "https://api.tiingo.com/tiingo/fundamentals/meta?format=csv",
	},
}

func TestFundamentalMetadataUrl(t *testing.T) {
	type args struct {
		ticker      string
		queryParams *FundamentalMetadataParams
	}
	type test struct {
		name string
		args args
		want string
	}
	var tests []test

	// Add common tests
	for _, tt := range commonFundamentalMetadataTests {
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
			if got := FundamentalMetadataUrl(tt.args.queryParams); got != tt.want {
				t.Errorf("FundamentalMetadataUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FundamentalMetadata(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		ticker      string
		queryParams *FundamentalMetadataParams
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	// Add common tests
	for _, tt := range commonFundamentalMetadataTests {
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
				queryParams: &FundamentalMetadataParams{
					RespFormat: "BAD FORMAT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidTickers",
			args: args{
				ctx: ctx,
				queryParams: &FundamentalMetadataParams{
					Tickers: []string{"BAD TICKER// -/"},
				},
			},
			wantErr: true,
		},
	}...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[[]FundamentalMetadata]("FundamentalMetadata()", tt.wantErr,
				func() ([]FundamentalMetadata, error) {
					return getClient().FundamentalMetadata(tt.args.ctx, tt.args.queryParams)
				},
			); err != nil {
				t.Error(err)
			}
		})
	}
}
