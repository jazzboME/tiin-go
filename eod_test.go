package tiingo

import (
	"context"
	"testing"
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
			name: "emptyParams",
			args: args{
				ctx:         ctx,
				ticker:      "AAPL",
				queryParams: nil,
			},
		},
		{
			name: "startDate",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					StartDate: startDate,
				},
			},
			wantErr: false,
		},
		{
			name: "endDate",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					EndDate: endDate,
				},
			},
			wantErr: false,
		},
		{
			name: "dailyResampleFreq",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Daily,
				},
			},
			wantErr: false,
		},
		{
			name: "weeklyResampleFreq",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Weekly,
				},
			},
			wantErr: false,
		},
		{
			name: "monthlyResampleFreq",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Monthly,
				},
			},
			wantErr: false,
		},
		{
			name: "annualResampleFreq",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Annually,
				},
			},
			wantErr: false,
		},
		{
			name: "dateAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DateAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "dateDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DateDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "openAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: OpenAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "openDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: OpenDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "highAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: HighAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "highDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: HighDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "lowAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: LowAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "lowDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: LowDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "closeAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: CloseAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "closeDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: CloseDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "volumeAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: VolumeAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "volumeDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: VolumeDesc,
				},
			},
			wantErr: false,
		},

		{
			name: "AdjOpenAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjOpenAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "AdjOpenDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjOpenDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjHighAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjHighAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjHighDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjHighDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjLowAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjLowAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjLowDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjLowDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjCloseAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjCloseAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjCloseDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjCloseDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjVolumeAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjVolumeAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "adjVolumeDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjVolumeDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "divCashAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DivCashAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "divCashDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DivCashDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "splitFactorAscSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: SplitFactorAsc,
				},
			},
			wantErr: false,
		},
		{
			name: "splitFactorDescSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: SplitFactorDesc,
				},
			},
			wantErr: false,
		},
		{
			name: "invalidTicker",
			args: args{
				ctx:         ctx,
				ticker:      "",
				queryParams: nil,
			},
			wantErr: true,
		},
		{
			name: "invalidResampleFreq",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: "INVALID FREQUENCY",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidSort",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: "INVALID SORT",
				},
			},
			wantErr: true,
		},
		{
			name: "invalidFormat",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					RespFormat: "INVALID FORMAT",
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
				ticker:      "AAPL",
				queryParams: &EodPriceParams{},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "startDate",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					StartDate: startDate,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?startDate=2022-01-01",
		},
		{
			name: "endDate",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					EndDate: endDate,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?endDate=2024-01-01",
		},
		{
			name: "dailyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Daily,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=daily",
		},
		{
			name: "weeklyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Weekly,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=weekly",
		},
		{
			name: "monthlyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Monthly,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=monthly",
		},
		{
			name: "annuallyFreq",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					ResampleFreq: Annually,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?resampleFreq=annually",
		},
		{
			name: "dateAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DateAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=date",
		},
		{
			name: "dateDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DateDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-date",
		},
		{
			name: "openAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: OpenAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=open",
		},
		{
			name: "openDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: OpenDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-open",
		},
		{
			name: "highAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: HighAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=high",
		},
		{
			name: "highDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: HighDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-high",
		},
		{
			name: "lowAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: LowAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=low",
		},
		{
			name: "lowDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: LowDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-low",
		},
		{
			name: "closeAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: CloseAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=close",
		},
		{
			name: "closeDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: CloseDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-close",
		},
		{
			name: "volumeAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: VolumeAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=volume",
		},
		{
			name: "volumeDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: VolumeDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-volume",
		},
		{
			name: "adjOpenAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjOpenAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjOpen",
		},
		{
			name: "adjOpenDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjOpenDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjOpen",
		},
		{
			name: "adjHighAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjHighAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjHigh",
		},
		{
			name: "adjHighDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjHighDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjHigh",
		},
		{
			name: "adjLowAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjLowAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjLow",
		},
		{
			name: "adjLowDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjLowDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjLow",
		},
		{
			name: "adjCloseAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjCloseAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjClose",
		},
		{
			name: "adjCloseDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjCloseDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjClose",
		},
		{
			name: "adjVolumeAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjVolumeAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=adjVolume",
		},
		{
			name: "adjVolumeDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: AdjVolumeDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-adjVolume",
		},
		{
			name: "divCashAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DivCashAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=divCash",
		},
		{
			name: "divCashDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: DivCashDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-divCash",
		},
		{
			name: "splitFactorAsc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: SplitFactorAsc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=splitFactor",
		},
		{
			name: "splitFactorDesc",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Sort: SplitFactorDesc,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?sort=-splitFactor",
		},
		{
			name: "csvRespFormat:",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					RespFormat: CSV,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?format=csv",
		},
		{
			name: "jsonRespFormat:",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					RespFormat: JSON,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices?format=json",
		},
		{
			name: "nilColumns",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Columns: nil,
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "emptyColumns",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Columns: []string{},
				},
			},
			want: "https://api.tiingo.com/tiingo/daily/AAPL/prices",
		},
		{
			name: "oneColumn",
			args: args{
				ticker: "AAPL",
				queryParams: &EodPriceParams{
					Columns: []string{"open"},
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
					Columns: []string{"open", "high", "low", "close"},
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
					StartDate:  startDate,
					EndDate:    endDate,
					Sort:       DateDesc,
					RespFormat: CSV,
					Columns:    []string{"date", "open", "high", "low", "close", "volume"},
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
			name: "validTicker",
			args: args{
				ctx:    ctx,
				ticker: "AAPL",
			},
			wantErr: false,
		},
		{
			name: "invalidTicker",
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
