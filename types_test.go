package tiingo

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/gocarina/gocsv"
)

type kind int

const (
	csv_ kind = iota
	json_
)

func cmpFloat(f1, f2 float64) bool {
	return math.Abs(f1-f2) < .00001
}

func testUnmarshal[T any](path string, unmarshalType kind, correct T, cmpFunc func(a, b T) bool) error {
	rawBytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	var data T
	if unmarshalType == json_ {
		if err = json.Unmarshal(rawBytes, &data); err != nil {
			return fmt.Errorf("failed to unmarshall json bytes: %w", err)
		}
	} else if unmarshalType == csv_ {
		if err = gocsv.UnmarshalBytes(rawBytes, &data); err != nil {
			return fmt.Errorf("failed to unmarshall csv bytes: %w", err)
		}
	} else {
		return fmt.Errorf("unmarshalType not recognized: %v", unmarshalType)
	}

	if !cmpFunc(correct, data) {
		return fmt.Errorf("result not correct. expcected=%v got=%v", correct, data)
	}

	return nil
}

var correctEodPrice = []EodPrice{
	{
		Date:        time.Date(2012, 12, 12, 0, 0, 0, 0, time.UTC),
		Open:        547.77,
		High:        548,
		Low:         536.2701,
		Close:       539,
		Volume:      17398000,
		AdjOpen:     16.6501054665,
		AdjHigh:     16.6570965837,
		AdjLow:      16.3005526472,
		AdjClose:    16.3835311289,
		AdjVolume:   487144487,
		DivCash:     0,
		SplitFactor: 1,
	},
	{
		Date:        time.Date(2012, 12, 12, 0, 0, 0, 0, time.UTC),
		Open:        547.77,
		High:        548,
		Low:         536.2701,
		Close:       539,
		Volume:      17398000,
		AdjOpen:     16.6501054665,
		AdjHigh:     16.6570965837,
		AdjLow:      16.3005526472,
		AdjClose:    16.3835311289,
		AdjVolume:   487144487,
		DivCash:     0,
		SplitFactor: 1,
	},
	{},
}

func eodEqualFunc(a, b []EodPrice) bool {
	return slices.EqualFunc(a, b, eodEqual)
}

func eodEqual(a EodPrice, b EodPrice) bool {
	if !a.Date.Equal(b.Date) {
		return false
	}
	if !cmpFloat(a.Open, b.Open) {
		return false
	}
	if !cmpFloat(a.High, b.High) {
		return false
	}
	if !cmpFloat(a.Low, b.Low) {
		return false
	}
	if !cmpFloat(a.Close, b.Close) {
		return false
	}
	if a.Volume != b.Volume {
		return false
	}
	if !cmpFloat(a.AdjOpen, b.AdjOpen) {
		return false
	}
	if !cmpFloat(a.AdjHigh, b.AdjHigh) {
		return false
	}
	if !cmpFloat(a.AdjLow, b.AdjLow) {
		return false
	}
	if !cmpFloat(a.AdjClose, b.AdjClose) {
		return false
	}
	if a.AdjVolume != b.AdjVolume {
		return false
	}
	if !cmpFloat(a.DivCash, b.DivCash) {
		return false
	}
	if !cmpFloat(a.SplitFactor, b.SplitFactor) {
		return false
	}
	return true
}

func TestEodEodPrice_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/eod.csv"
	if err := testUnmarshal(path, csv_, correctEodPrice, eodEqualFunc); err != nil {
		t.Fatal(err)
	}
}

func TestEodPrice_UnmarshalJSON(t *testing.T) {
	path := "./test_data/eod.json"
	if err := testUnmarshal(path, json_, correctEodPrice, eodEqualFunc); err != nil {
		t.Fatal(err)
	}
}

func TestEodMetadata_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/eod_metadata.csv"
	correctEodMetadata := []EodMetadata{
		{
			Ticker:       "AAPL",
			Name:         "Apple Inc",
			ExchangeCode: "NASDAQ",
			Description:  `b"Apple Inc. (Apple) designs, manufactures and markets mobile communication and media devices, personal computers, and portable digital music players, and a variety of related software, services, peripherals, networking solutions, and third-party digital content and applications. The Company's products and services include iPhone, iPad, Mac, iPod, Apple TV, a portfolio of consumer and professional software applications, the iOS and OS X operating systems, iCloud, and a variety of accessory, service and support offerings. The Company also delivers digital content and applications through the iTunes Store, App StoreSM, iBookstoreSM, and Mac App Store. The Company distributes its products worldwide through its retail stores, online stores, and direct sales force, as well as through third-party cellular network carriers, wholesalers, retailers, and value-added resellers. In February 2012, the Company acquired app-search engine Chomp."`,
			StartDate:    time.Date(1980, 12, 12, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC),
		},
		{},
	}
	equalFunc := func(a, b []EodMetadata) bool {
		return slices.Equal(a, b)
	}
	if err := testUnmarshal(path, csv_, correctEodMetadata, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestEodMetadata_UnmarshalJSON(t *testing.T) {
	path := "./test_data/eod_metadata.json"
	correctEodMetadata := EodMetadata{
		Ticker:       "AAPL",
		Name:         "Apple Inc",
		ExchangeCode: "NASDAQ",
		Description:  `Apple Inc. (Apple) designs, manufactures and markets mobile communication and media devices, personal computers, and portable digital music players, and a variety of related software, services, peripherals, networking solutions, and third-party digital content and applications. The Company's products and services include iPhone, iPad, Mac, iPod, Apple TV, a portfolio of consumer and professional software applications, the iOS and OS X operating systems, iCloud, and a variety of accessory, service and support offerings. The Company also delivers digital content and applications through the iTunes Store, App StoreSM, iBookstoreSM, and Mac App Store. The Company distributes its products worldwide through its retail stores, online stores, and direct sales force, as well as through third-party cellular network carriers, wholesalers, retailers, and value-added resellers. In February 2012, the Company acquired app-search engine Chomp.`,
		StartDate:    time.Date(1980, 12, 12, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC),
	}
	equalFunc := func(a, b EodMetadata) bool {
		if a.Ticker != b.Ticker {
			return false
		}
		if a.Name != b.Name {
			return false
		}
		if a.Description != b.Description {
			return false
		}
		if !a.StartDate.Equal(b.StartDate) {
			return false
		}
		if !a.EndDate.Equal(b.EndDate) {
			return false
		}
		return true
	}

	if err := testUnmarshal(path, json_, correctEodMetadata, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestSymbolItem_ParseZip(t *testing.T) {
	rawBytes, err := os.ReadFile("./test_data/symbol_list.zip")
	if err != nil {
		t.Fatalf("failed to open file: %s", err)
	}

	symbolList, err := ParseSymbolListCSV(rawBytes)
	if err != nil {
		t.Fatalf("failed to parse raw zipped csv bytes: %s", err)
	}

	if len(symbolList) == 0 {
		t.Fatalf("parsed symbol list is empty")
	}
}

func TestSymbolItem_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/supported_tickers.csv"
	correctSymbolItems := []SymbolItem{
		{
			Ticker:        "AAPL",
			Exchange:      "NASDAQ",
			AssetType:     "Stock",
			PriceCurrency: "USD",
			StartDate:     time.Date(1980, 12, 12, 0, 0, 0, 0, time.UTC),
			EndDate:       time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			Ticker:        "AAPL",
			Exchange:      "NASDAQ",
			AssetType:     "Stock",
			PriceCurrency: "USD",
			StartDate:     time.Date(1980, 12, 12, 0, 0, 0, 0, time.UTC),
			EndDate:       time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC),
		},
		{},
	}
	equalFunc := func(a, b []SymbolItem) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, csv_, correctSymbolItems, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestIexTopOfBook_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/iex_top_of_book.csv"
	correctTopOfBook := []IexTopOfBook{
		{
			Ticker:            "GOOG",
			Timestamp:         time.Date(2024, 12, 13, 19, 31, 32, 16361786, time.UTC),
			QuoteTimestamp:    time.Date(2024, 12, 13, 19, 31, 32, 16361786, time.UTC),
			LastSaleTimestamp: time.Date(2024, 12, 13, 19, 31, 24, 604004330, time.UTC),
			Last:              193.09,
			LastSize:          200,
			TngoLast:          193.085,
			PrevClose:         193.63,
			Open:              193.08,
			High:              194.335,
			Low:               191.665,
			Mid:               193.085,
			Volume:            339489,
			BidSize:           100,
			BidPrice:          193.07,
			AskSize:           102,
			AskPrice:          193.1,
		},
		{
			Ticker:            "AAPL",
			Timestamp:         time.Date(2024, 12, 13, 19, 31, 30, 747495598, time.UTC),
			QuoteTimestamp:    time.Date(2024, 12, 13, 19, 31, 30, 747495598, time.UTC),
			LastSaleTimestamp: time.Date(2024, 12, 13, 19, 31, 26, 199530570, time.UTC),
			Last:              247.96,
			LastSize:          4,
			TngoLast:          247.975,
			PrevClose:         247.96,
			Open:              249.2,
			High:              249.325,
			Low:               246.245,
			Mid:               247.975,
			Volume:            460522,
			BidSize:           200,
			BidPrice:          247.95,
			AskSize:           100,
			AskPrice:          248,
		},
		{
			Ticker:            "AAPL",
			Timestamp:         time.Date(2024, 12, 13, 21, 0, 0, 0, time.UTC),
			QuoteTimestamp:    time.Date(2024, 12, 13, 21, 0, 0, 0, time.UTC),
			LastSaleTimestamp: time.Date(2024, 12, 13, 21, 0, 0, 0, time.UTC),
			Last:              248.13,
			LastSize:          0,
			TngoLast:          248.13,
			PrevClose:         247.96,
			Open:              247.815,
			High:              249.2902,
			Low:               246.24,
			Mid:               0,
			Volume:            33155290,
			BidSize:           0,
			BidPrice:          0,
			AskSize:           0,
			AskPrice:          0,
		},
		{
			Ticker:            "ACST",
			Timestamp:         time.Date(2024, 10, 24, 16, 51, 0, 0, time.UTC),
			QuoteTimestamp:    time.Time{},
			LastSaleTimestamp: time.Time{},
			Last:              0,
			LastSize:          0,
			TngoLast:          3.3,
			PrevClose:         3.11,
			Open:              3.3,
			High:              3.3,
			Low:               3.3,
			Mid:               0,
			Volume:            42,
			BidSize:           0,
			BidPrice:          0,
			AskSize:           0,
			AskPrice:          0,
		},
		{},
		{
			Ticker:            "GPCR",
			Timestamp:         time.Date(2024, 12, 26, 21, 0, 0, 120115, time.UTC),
			QuoteTimestamp:    time.Date(2024, 12, 26, 21, 0, 0, 120115, time.UTC),
			LastSaleTimestamp: time.Date(2024, 12, 26, 20, 59, 55, 4092910, time.UTC),
			Last:              29.58,
			LastSize:          1,
			TngoLast:          29.58,
			PrevClose:         28.25,
			Open:              28.5,
			High:              30.41,
			Low:               28.225,
			Mid:               0,
			Volume:            39247,
			BidSize:           0,
			BidPrice:          0,
			AskSize:           0,
			AskPrice:          0,
		},
	}
	equalFunc := func(a, b []IexTopOfBook) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, csv_, correctTopOfBook, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestIexTopOfBook_UnmarshalJSON(t *testing.T) {
	path := "./test_data/iex_top_of_book.json"
	correctTopOfBook := []IexTopOfBook{
		{
			Ticker:            "GOOG",
			Timestamp:         time.Date(2024, 12, 13, 19, 31, 20, 167224203, time.UTC),
			QuoteTimestamp:    time.Date(2024, 12, 13, 19, 31, 20, 167224203, time.UTC),
			LastSaleTimestamp: time.Date(2024, 12, 13, 19, 31, 12, 366566237, time.UTC),
			Last:              193.07,
			LastSize:          96,
			TngoLast:          193.075,
			PrevClose:         193.63,
			Open:              193.08,
			High:              194.335,
			Low:               191.665,
			Mid:               192.68,
			Volume:            338830,
			BidSize:           205,
			BidPrice:          191.36,
			AskSize:           102,
			AskPrice:          194,
		},
		{
			Ticker:            "AAPL",
			Timestamp:         time.Date(2024, 12, 13, 19, 31, 19, 890886829, time.UTC),
			QuoteTimestamp:    time.Date(2024, 12, 13, 19, 31, 19, 890886829, time.UTC),
			LastSaleTimestamp: time.Date(2024, 12, 13, 19, 31, 0, 23335459, time.UTC),
			Last:              247.99,
			LastSize:          1,
			TngoLast:          247.98,
			PrevClose:         247.96,
			Open:              249.2,
			High:              249.325,
			Low:               246.245,
			Mid:               247.98,
			Volume:            460418,
			BidSize:           104,
			BidPrice:          247.96,
			AskSize:           100,
			AskPrice:          248,
		},
		{
			Ticker:            "ACST",
			Timestamp:         time.Date(2024, 10, 24, 16, 51, 0, 0, time.UTC),
			QuoteTimestamp:    time.Time{},
			LastSaleTimestamp: time.Time{},
			Last:              0,
			LastSize:          0,
			TngoLast:          3.3,
			PrevClose:         3.11,
			Open:              3.3,
			High:              3.3,
			Low:               3.3,
			Mid:               0,
			Volume:            42,
			BidSize:           0,
			BidPrice:          0,
			AskSize:           0,
			AskPrice:          0,
		},
		{},
	}
	equalFunc := func(a, b []IexTopOfBook) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, json_, correctTopOfBook, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestIexPrice_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/iex_price.csv"
	correct := []IexPrice{
		{
			Date:   time.Date(2024, 6, 18, 16, 25, 0, 0, time.UTC),
			Open:   213.14,
			High:   213.315,
			Low:    212.975,
			Close:  213.205,
			Volume: 4647,
		},
		{
			Date:   time.Date(2024, 6, 18, 16, 30, 0, 0, time.UTC),
			Open:   213.22,
			High:   213.59,
			Low:    213.205,
			Close:  213.505,
			Volume: 3178,
		},
		{},
	}
	equalFunc := func(a, b []IexPrice) bool {
		return slices.Equal(a, b)
	}
	if err := testUnmarshal(path, csv_, correct, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestIexPrice_UnmarshalJSON(t *testing.T) {
	path := "./test_data/iex_price.json"
	correct := []IexPrice{
		{
			Date:   time.Date(2024, 6, 18, 16, 25, 0, 0, time.UTC),
			Open:   213.14,
			High:   213.315,
			Low:    212.975,
			Close:  213.205,
			Volume: 4647,
		},
		{
			Date:   time.Date(2024, 6, 18, 16, 30, 0, 0, time.UTC),
			Open:   213.22,
			High:   213.59,
			Low:    213.205,
			Close:  213.505,
			Volume: 3178,
		},
		{},
	}
	equalFunc := func(a, b []IexPrice) bool {
		return slices.Equal(a, b)
	}
	if err := testUnmarshal(path, json_, correct, equalFunc); err != nil {
		t.Fatal(err)
	}
}

var correctStmtDef = []StmtDef{
	{
		DataCode:      "rps",
		Name:          "Revenue Per Share",
		Description:   "Revenue per share",
		StatementType: "overview",
		Units:         "$",
	},
	{
		DataCode:      "roa",
		Name:          "Return on Assets ROA",
		Description:   "Net Income/Total Assets",
		StatementType: "overview",
		Units:         "%",
	},
	{},
}

func TestStmtDef_UnmarshalCSV(t *testing.T) {
	path := "./test_data/statement_definitions.csv"
	equalFunc := func(a, b []StmtDef) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, csv_, correctStmtDef, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestStmtDef_UnmarshalJSON(t *testing.T) {
	path := "./test_data/statement_definitions.json"
	equalFunc := func(a, b []StmtDef) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, json_, correctStmtDef, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestStmtDataNested_UnmarshalJSON(t *testing.T) {
	path := "./test_data/statement_value_nested.json"
	equalFunc := func(a, b []StmtDataNested) bool {
		return slices.EqualFunc(a, b, func(c StmtDataNested, d StmtDataNested) bool {
			if !c.Date.Equal(d.Date) {
				return false
			}
			if c.Year != d.Year {
				return false
			}
			if c.Quarter != d.Quarter {
				return false
			}
			if !slices.Equal(c.StatementData.IncomeStatement, d.StatementData.IncomeStatement) {
				return false
			}
			if !slices.Equal(c.StatementData.CashFlow, d.StatementData.CashFlow) {
				return false
			}
			if !slices.Equal(c.StatementData.Overview, d.StatementData.Overview) {
				return false
			}
			if !slices.Equal(c.StatementData.BalanceSheet, d.StatementData.BalanceSheet) {
				return false
			}

			return true
		})
	}
	correctStmtVal := []StmtDataNested{
		{
			Date:    time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC),
			Year:    2024,
			Quarter: 4,
			StatementData: struct {
				BalanceSheet    []StmtDataField
				IncomeStatement []StmtDataField
				CashFlow        []StmtDataField
				Overview        []StmtDataField
			}{
				BalanceSheet: []StmtDataField{
					{
						DataCode: "acctPay",
						Value:    68960000000,
					},
					{
						DataCode: "deposits",
						Value:    0,
					},
				},
				IncomeStatement: []StmtDataField{
					{
						DataCode: "sga",
						Value:    6523000000,
					},
					{
						DataCode: "shareswaDil",
						Value:    15242855000,
					},
				},
				CashFlow: []StmtDataField{
					{
						DataCode: "depamor",
						Value:    2911000000,
					},
					{
						DataCode: "capex",
						Value:    -2908000000,
					},
				},
				Overview: []StmtDataField{
					{
						DataCode: "bookVal",
						Value:    56950000000,
					},
					{
						DataCode: "longTermDebtEquity",
						Value:    1.50570676031607,
					},
				},
			},
		},
		{
			Date:    time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC),
			Year:    2024,
			Quarter: 4,
			StatementData: struct {
				BalanceSheet    []StmtDataField
				IncomeStatement []StmtDataField
				CashFlow        []StmtDataField
				Overview        []StmtDataField
			}{
				BalanceSheet: []StmtDataField{
					{
						DataCode: "acctPay",
						Value:    68960000000,
					},
					{
						DataCode: "deposits",
						Value:    0,
					},
				},
				IncomeStatement: []StmtDataField{
					{
						DataCode: "sga",
						Value:    6523000000,
					},
					{
						DataCode: "shareswaDil",
						Value:    15242855000,
					},
				},
				CashFlow: []StmtDataField{
					{
						DataCode: "depamor",
						Value:    2911000000,
					},
					{
						DataCode: "capex",
						Value:    -2908000000,
					},
				},
				Overview: []StmtDataField{
					{
						DataCode: "bookVal",
						Value:    56950000000,
					},
					{
						DataCode: "longTermDebtEquity",
						Value:    1.50570676031607,
					},
				},
			},
		},
		{},
		{},
	}

	if err := testUnmarshal(path, json_, correctStmtVal, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestStmtDataFlat_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/statement_value_flat.csv"
	equalFunc := func(a, b []StmtDataFlat) bool {
		return slices.Equal(a, b)
	}
	correctStmtValFlat := []StmtDataFlat{
		{
			Date:          time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC),
			Year:          2024,
			Quarter:       4,
			StatementType: "balanceSheet",
			DataCode:      "acctPay",
			Value:         68960000000,
		},
		{
			Date:          time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC),
			Year:          2024,
			Quarter:       4,
			StatementType: "balanceSheet",
			DataCode:      "deposits",
			Value:         0,
		},
		{},
	}

	if err := testUnmarshal(path, csv_, correctStmtValFlat, equalFunc); err != nil {
		t.Fatal(err)
	}
}

var correctStmtValFlat = []DailyFundamental{
	{
		Date:                 time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC),
		MarketCap:            3770017810520,
		EnterpriseValue:      3811475810520,
		PriceToEarningsRatio: 40.2195294286,
		PriceToBookRatio:     66.198732406,
		TrailingYearPEG:      -1.2065858829,
	},
	{
		Date:                 time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC),
		MarketCap:            3770017810520,
		EnterpriseValue:      3811475810520,
		PriceToEarningsRatio: 40.2195294286,
		PriceToBookRatio:     66.198732406,
		TrailingYearPEG:      -1.2065858829,
	},
	{},
}

func dailyFundamentalCmp(a, b DailyFundamental) bool {
	if !a.Date.Equal(b.Date) {
		return false
	}
	if !cmpFloat(a.MarketCap, b.MarketCap) {
		return false
	}
	if !cmpFloat(a.EnterpriseValue, b.EnterpriseValue) {
		return false
	}
	if !cmpFloat(a.PriceToEarningsRatio, b.PriceToEarningsRatio) {
		return false
	}
	if !cmpFloat(a.PriceToBookRatio, b.PriceToBookRatio) {
		return false
	}
	if !cmpFloat(a.TrailingYearPEG, b.TrailingYearPEG) {
		return false
	}

	return true
}

func TestDailyFundamental_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/daily_fundamental.csv"
	equalFunc := func(a, b []DailyFundamental) bool {
		return slices.EqualFunc(a, b, dailyFundamentalCmp)
	}

	if err := testUnmarshal(path, csv_, correctStmtValFlat, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestDailyFundamental_UnmarshalJSON(t *testing.T) {
	path := "./test_data/daily_fundamental.json"
	equalFunc := func(a, b []DailyFundamental) bool {
		return slices.EqualFunc(a, b, dailyFundamentalCmp)
	}

	if err := testUnmarshal(path, json_, correctStmtValFlat, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestFundamentalMetadata_UnmarshalCSVWithFields(t *testing.T) {
	path := "./test_data/fundamental_metadata.csv"
	correctFundamentalMetadata := []FundamentalMetadata{
		{
			PermaTicker:             "US000000000038",
			Ticker:                  "aapl",
			Name:                    "Apple Inc",
			IsActive:                true,
			IsADR:                   false,
			Sector:                  "Technology",
			Industry:                "Consumer Electronics",
			SicCode:                 3571,
			SicSector:               "Manufacturing",
			SicIndustry:             "Electronic Computers",
			ReportingCurrency:       "usd",
			Location:                "California, USA",
			CompanyWebsite:          "http://www.apple.com",
			SecFilingWebsite:        "https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000320193",
			StatementLastUpdated:    time.Date(2024, 11, 02, 1, 1, 16, 780400000, time.UTC),
			DailyLastUpdated:        time.Date(2024, 12, 13, 14, 15, 9, 502872000, time.UTC),
			DataProviderPermaTicker: "199059",
		},
		{
			PermaTicker:             "US000000000038",
			Ticker:                  "aapl",
			Name:                    "Apple Inc",
			IsActive:                true,
			IsADR:                   false,
			Sector:                  "Technology",
			Industry:                "Consumer Electronics",
			SicCode:                 3571,
			SicSector:               "Manufacturing",
			SicIndustry:             "Electronic Computers",
			ReportingCurrency:       "usd",
			Location:                "California, USA",
			CompanyWebsite:          "http://www.apple.com",
			SecFilingWebsite:        "https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000320193",
			StatementLastUpdated:    time.Date(2024, 11, 02, 1, 1, 16, 780400000, time.UTC),
			DailyLastUpdated:        time.Date(2024, 12, 13, 14, 15, 9, 502872000, time.UTC),
			DataProviderPermaTicker: "199059",
		},
		{
			PermaTicker:             "US000000009823",
			Ticker:                  "aac",
			Name:                    "American Addiction Centers",
			IsActive:                false,
			IsADR:                   false,
			Sector:                  "",
			Industry:                "",
			SicCode:                 0,
			SicSector:               "",
			SicIndustry:             "",
			ReportingCurrency:       "",
			Location:                "",
			CompanyWebsite:          "",
			SecFilingWebsite:        "",
			StatementLastUpdated:    time.Date(2021, 4, 5, 22, 1, 56, 682496000, time.UTC),
			DailyLastUpdated:        time.Date(2021, 4, 19, 20, 41, 37, 359000000, time.UTC),
			DataProviderPermaTicker: "",
		},
		{
			PermaTicker:             "US000000073255",
			Ticker:                  "aav1",
			Name:                    "AVATEX CORP",
			IsActive:                false,
			IsADR:                   false,
			Sector:                  "Real Estate",
			Industry:                "Real Estate Services",
			SicCode:                 6500,
			SicSector:               "Finance Insurance And Real Estate",
			SicIndustry:             "Real Estate",
			ReportingCurrency:       "usd",
			Location:                "Texas, USA",
			CompanyWebsite:          "",
			SecFilingWebsite:        "https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000716644",
			StatementLastUpdated:    time.Date(2020, 8, 6, 22, 8, 27, 305372000, time.UTC),
			DailyLastUpdated:        time.Time{},
			DataProviderPermaTicker: "",
		},
		{},
	}
	equalFunc := func(a, b []FundamentalMetadata) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, csv_, correctFundamentalMetadata, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestFundamentalMetadata_UnmarshalJSON(t *testing.T) {
	path := "./test_data/fundamental_metadata.json"
	correctFundamentalMetadata := []FundamentalMetadata{
		{
			PermaTicker:             "US000000000038",
			Ticker:                  "aapl",
			Name:                    "Apple Inc",
			IsActive:                true,
			IsADR:                   false,
			Sector:                  "Technology",
			Industry:                "Consumer Electronics",
			SicCode:                 3571,
			SicSector:               "Manufacturing",
			SicIndustry:             "Electronic Computers",
			ReportingCurrency:       "usd",
			Location:                "California, USA",
			CompanyWebsite:          "http://www.apple.com",
			SecFilingWebsite:        "https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000320193",
			StatementLastUpdated:    time.Date(2024, 11, 02, 1, 1, 16, 780000000, time.UTC),
			DailyLastUpdated:        time.Date(2024, 12, 13, 14, 15, 9, 502000000, time.UTC),
			DataProviderPermaTicker: "199059",
		},
		{
			PermaTicker:             "US000000000038",
			Ticker:                  "aapl",
			Name:                    "Apple Inc",
			IsActive:                true,
			IsADR:                   false,
			Sector:                  "Technology",
			Industry:                "Consumer Electronics",
			SicCode:                 3571,
			SicSector:               "Manufacturing",
			SicIndustry:             "Electronic Computers",
			ReportingCurrency:       "usd",
			Location:                "California, USA",
			CompanyWebsite:          "http://www.apple.com",
			SecFilingWebsite:        "https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000320193",
			StatementLastUpdated:    time.Date(2024, 11, 02, 1, 1, 16, 780000000, time.UTC),
			DailyLastUpdated:        time.Date(2024, 12, 13, 14, 15, 9, 502000000, time.UTC),
			DataProviderPermaTicker: "199059",
		},
		{},
	}
	equalFunc := func(a, b []FundamentalMetadata) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, json_, correctFundamentalMetadata, equalFunc); err != nil {
		t.Fatal(err)
	}
}

var correctSearchResult = []SearchResult{
	{
		Name:              "Apple Inc",
		Ticker:            "AAPL",
		PermaTicker:       "US000000000038",
		OpenFIGIComposite: "BBG000B9XRY4",
		AssetType:         "Stock",
		IsActive:          true,
		CountryCode:       "US",
	},
	{
		Name:              "Pineapple Inc",
		Ticker:            "PNPL",
		PermaTicker:       "US000000047877",
		OpenFIGIComposite: "BBG000D05J52",
		AssetType:         "Stock",
		IsActive:          true,
		CountryCode:       "US",
	},
	{},
}

func TestSearchResult_UnmarshalCSV(t *testing.T) {
	path := "./test_data/search_result.csv"
	equalFunc := func(a, b []SearchResult) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, csv_, correctSearchResult, equalFunc); err != nil {
		t.Fatal(err)
	}
}

func TestSearchResult_UnmarshalJSON(t *testing.T) {
	path := "./test_data/search_result.json"
	equalFunc := func(a, b []SearchResult) bool {
		return slices.Equal(a, b)
	}

	if err := testUnmarshal(path, json_, correctSearchResult, equalFunc); err != nil {
		t.Fatal(err)
	}
}
