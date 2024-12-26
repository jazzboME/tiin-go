package tiingo

import (
	"context"
	"testing"
	"time"
)

var (
	startDate = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate   = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
)

func TestClient_EodPrice(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		ticker      string
		queryParams *EodPriceParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validArgs",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					startDate:    startDate,
					endDate:      endDate,
					resampleFreq: Weekly,
					sort:         DateDesc,
					respFormat:   CSV,
				},
			},
			wantErr: false,
		},
		{
			name: "invalidArgs",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					startDate:    startDate,
					endDate:      endDate,
					resampleFreq: Weekly,
					sort:         DateDesc,
					respFormat:   "BAD FORMAT",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[[]EodPrice]("EodPrice()", tt.wantErr, func() ([]EodPrice, error) {
				return getClient().EodPrice(tt.args.ctx, tt.args.ticker, tt.args.queryParams)
			}); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestEodPriceUrl(t *testing.T) {
	type args struct {
		ticker      string
		queryParams *EodPriceParams
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nilParams",
			args: args{
				ticker:      "AAPL",
				queryParams: nil,
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "zeroValParams",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					startDate:    time.Time{},
					endDate:      time.Time{},
					resampleFreq: "",
					sort:         "",
					respFormat:   "",
					columns:      nil,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "startDate",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					startDate: startDate,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?startDate=2022-01-01",
		},
		{
			name: "endDate",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					endDate: endDate,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?endDate=2024-01-01",
		},
		{
			name: "dailyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					resampleFreq: Daily,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=daily",
		},
		{
			name: "weeklyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					resampleFreq: Weekly,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=weekly",
		},
		{
			name: "monthlyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					resampleFreq: Monthly,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=monthly",
		},
		{
			name: "annuallyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					resampleFreq: Annually,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=annually",
		},
		{
			name: "dateAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: DateAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=date",
		},
		{
			name: "dateDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: DateDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-date",
		},
		{
			name: "openAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: OpenAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=open",
		},
		{
			name: "openDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: OpenDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-open",
		},
		{
			name: "highAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: HighAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=high",
		},
		{
			name: "highDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: HighDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-high",
		},
		{
			name: "lowAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: LowAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=low",
		},
		{
			name: "lowDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: LowDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-low",
		},
		{
			name: "closeAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: CloseAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=close",
		},
		{
			name: "closeDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: CloseDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-close",
		},
		{
			name: "volumeAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: VolumeAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=volume",
		},
		{
			name: "volumeDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: VolumeDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-volume",
		},
		{
			name: "adjOpenAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjOpenAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjOpen",
		},
		{
			name: "adjOpenDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjOpenDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjOpen",
		},
		{
			name: "adjHighAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjHighAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjHigh",
		},
		{
			name: "adjHighDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjHighDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjHigh",
		},
		{
			name: "adjLowAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjLowAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjLow",
		},
		{
			name: "adjLowDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjLowDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjLow",
		},
		{
			name: "adjCloseAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjCloseAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjClose",
		},
		{
			name: "adjCloseDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjCloseDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjClose",
		},
		{
			name: "adjVolumeAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjVolumeAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjVolume",
		},
		{
			name: "adjVolumeDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: AdjVolumeDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjVolume",
		},
		{
			name: "divCashAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: DivCashAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=divCash",
		},
		{
			name: "divCashDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: DivCashDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-divCash",
		},
		{
			name: "splitFactorAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: SplitFactorAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=splitFactor",
		},
		{
			name: "splitFactorDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					sort: SplitFactorDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-splitFactor",
		},
		{
			name: "csvRespFormat",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					respFormat: CSV,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?format=csv",
		},
		{
			name: "jsonRespFormat",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					respFormat: JSON,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?format=json",
		},
		{
			name: "nilColumns",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					columns: nil,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "emptyColumns",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					columns: []string{},
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "oneColumn",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					columns: []string{"open"},
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices" +
				"?columns=open",
		},
		{
			name: "manyColumns",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					columns: []string{"open", "high", "low", "close"},
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices" +
				"?columns=open,high,low,close",
		},
		{
			name: "allParams",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					startDate:  startDate,
					endDate:    endDate,
					sort:       DateDesc,
					respFormat: CSV,
					columns:    []string{"date", "open", "high", "low", "close", "volume"},
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?startDate=2022-01-01" +
				"&endDate=2024-01-01&sort=-date&format=csv&columns=date,open,high,low,close,volume",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EodPriceUrl(tt.args.ticker, tt.args.queryParams); got != tt.want {
				t.Errorf("EodPriceUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEodMetadataUrl(t *testing.T) {
	type args struct {
		ticker string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				ticker: "AAPL",
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EodMetadataUrl(tt.args.ticker); got != tt.want {
				t.Errorf("EodMetadataUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_EodMetadata(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		ticker string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validArgs",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
			},
			wantErr: false,
		},
		{
			name: "invalidArgs",
			args: args{
				ctx:    ctx,
				ticker: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := liveTest[EodMetadata]("EodMetadata()", tt.wantErr, func() (EodMetadata, error) {
				return getClient().EodMetadata(tt.args.ctx, tt.args.ticker)
			}); err != nil {
				t.Error(err)
			}
		})
	}
}
