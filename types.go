package tiingo

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

const (
	JSON Format = "json"
	CSV  Format = "csv"

	Daily    EodFreq = "daily"
	Weekly   EodFreq = "weekly"
	Monthly  EodFreq = "monthly"
	Annually EodFreq = "annually"

	OneMin     IexFreq = "1min"
	FiveMin    IexFreq = "5min"
	FifteenMin IexFreq = "15min"
	ThirtyMin  IexFreq = "30min"
	OneHour    IexFreq = "1hour"
	TwoHour    IexFreq = "2hour"
	FourHour   IexFreq = "4hour"
	OneDay	   IexFreq = "1day"

	DateAsc         Sort = "date"
	DateDesc        Sort = "-date"
	OpenAsc         Sort = "open"
	OpenDesc        Sort = "-open"
	HighAsc         Sort = "high"
	HighDesc        Sort = "-high"
	LowAsc          Sort = "low"
	LowDesc         Sort = "-low"
	CloseAsc        Sort = "close"
	CloseDesc       Sort = "-close"
	VolumeAsc       Sort = "volume"
	VolumeDesc      Sort = "-volume"
	AdjOpenAsc      Sort = "adjOpen"
	AdjOpenDesc     Sort = "-adjOpen"
	AdjHighAsc      Sort = "adjHigh"
	AdjHighDesc     Sort = "-adjHigh"
	AdjLowAsc       Sort = "adjLow"
	AdjLowDesc      Sort = "-adjLow"
	AdjCloseAsc     Sort = "adjClose"
	AdjCloseDesc    Sort = "-adjClose"
	AdjVolumeAsc    Sort = "adjVolume"
	AdjVolumeDesc   Sort = "-adjVolume"
	DivCashAsc      Sort = "divCash"
	DivCashDesc     Sort = "-divCash"
	SplitFactorAsc  Sort = "splitFactor"
	SplitFactorDesc Sort = "-splitFactor"
	MktCapAsc       Sort = "marketCap"
	MktCapDesc      Sort = "-marketCap"
	EntValAsc       Sort = "enterpriseVal"
	EntValDesc      Sort = "-enterpriseVal"
	PERatioAsc      Sort = "peRatio"
	PERatioDesc     Sort = "-peRatio"
	PBRatioAsc      Sort = "pbRatio"
	PBRatioDesc     Sort = "-pbRatio"
	TrailPEGAsc     Sort = "trailingPEG1Y"
	TrailPEGDesc    Sort = "-trailingPEG1Y"
)

// Format is the respFormat the response is sent in.
// Possible predefined constants are:
//   - CSV
//   - JSON
type Format = string

// EodFreq is the eod price frequency requested.
// Possible predefined constants are:
//   - Daily
//   - Weekly
//   - Monthly
//   - Annually
type EodFreq = string

// IexFreq is the IexHistory price frequency requested.
// Possible predefined constants are:
//   - OneMin
//   - FiveMin
//   - FifteenMin
//   - ThirtyMin
//   - OneHour
//   - TwoHour
//   - FourHour
type IexFreq = string

// Sort defines the sorting order of the response data.
// Possible predefined constants are:
//   - DateAsc
//   - DateDesc
//   - OpenAsc
//   - OpenDesc
//   - HighAsc
//   - HighDesc
//   - LowAsc
//   - LowDesc
//   - CloseAsc
//   - CloseDesc
//   - VolumeAsc
//   - VolumeDesc
//   - AdjOpenAsc
//   - AdjOpenDesc
//   - AdjHighAsc
//   - AdjHighDesc
//   - AdjLowAsc
//   - AdjLowDesc
//   - AdjCloseAsc
//   - AdjCloseDesc
//   - AdjVolumeAsc
//   - AdjVolumeDesc
//   - DivCashAsc
//   - DivCashDesc
//   - SplitFactorAsc
//   - SplitFactorDesc
//   - MktCapAsc
//   - MktCapDesc
//   - EntValAsc
//   - EntValDesc
//   - PERatioAsc
//   - PERatioDesc
//   - PBRatioAsc
//   - PBRatioDesc
//   - TrailPEGAsc
//   - TrailPEGDesc
type Sort = string

// SymbolFilterFunc is a function that takes in a SymbolRespItem and returns a boolean
// for if that ticker item should be added to the overall ticker list. A value of nil
// will return every SymbolRespItem from the list.
//
// Example: Only include NYSE & NASDAQ stocks that have a startDate & endDate date
//
//	func(asset SymbolRespItem) bool {
//		if asset.Exchange != "NYSE" && asset.Exchange != "NASDAQ" {
//			return false
//		}
//		if asset.AssetType != "stock" {
//			return false
//		}
//		if asset.StartDate.IsZero() || asset.EndDate.IsZero() {
//			return false
//		}
//
//		return true
//	}
type SymbolFilterFunc func(asset SymbolItem) bool

// EodPrice corresponds to the [End-of-Day].2.1.2 End-of-Day Endpoint
type EodPrice struct {
	Date        time.Time `json:"date,omitempty" csv:"date"`
	Open        float64   `json:"open,omitempty" csv:"open"`
	High        float64   `json:"high,omitempty" csv:"high"`
	Low         float64   `json:"low,omitempty" csv:"low"`
	Close       float64   `json:"close,omitempty" csv:"close"`
	Volume      int64     `json:"volume,omitempty" csv:"volume"`
	AdjOpen     float64   `json:"adjOpen,omitempty" csv:"adjOpen"`
	AdjHigh     float64   `json:"adjHigh,omitempty" csv:"adjHigh"`
	AdjLow      float64   `json:"adjLow,omitempty" csv:"adjLow"`
	AdjClose    float64   `json:"adjClose,omitempty" csv:"adjClose"`
	AdjVolume   int64     `json:"adjVolume,omitempty" csv:"adjVolume"`
	DivCash     float64   `json:"divCash,omitempty" csv:"divCash"`
	SplitFactor float64   `json:"splitFactor,omitempty" csv:"splitFactor"`
}

func (e *EodPrice) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "date":
		date, err := parseTime(time.DateOnly, value)
		if err != nil {
			return err
		}
		e.Date = date
	case "open":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.Open = f
	case "high":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.High = f
	case "low":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.Low = f
	case "close":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.Close = f
	case "volume":
		i, err := parseInt(value)
		if err != nil {
			return err
		}
		e.Volume = i
	case "adjOpen":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.AdjOpen = f
	case "adjHigh":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.AdjHigh = f
	case "adjLow":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.AdjLow = f
	case "adjClose":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.AdjClose = f
	case "adjVolume":
		i, err := parseInt(value)
		if err != nil {
			return err
		}
		e.AdjVolume = i
	case "divCash":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.DivCash = f
	case "splitFactor":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		e.SplitFactor = f
	}

	return nil
}

// EodMetadata corresponds to the [End-of-Day].2.1.3 Meta Endpoint
type EodMetadata struct {
	Ticker       string    `json:"ticker,omitempty" csv:"ticker,omitempty"`
	Name         string    `json:"name,omitempty" csv:"name,omitempty"`
	ExchangeCode string    `json:"exchangeCode,omitempty" csv:"exchangeCode,omitempty"`
	Description  string    `json:"description,omitempty" csv:"description,omitempty"`
	StartDate    time.Time `json:"startDate,omitempty" csv:"startDate,omitempty"`
	EndDate      time.Time `json:"endDate,omitempty" csv:"endDate,omitempty"`
}

func (e *EodMetadata) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "ticker":
		e.Ticker = value
	case "name":
		e.Name = value
	case "exchangeCode":
		e.ExchangeCode = value
	case "description":
		e.Description = value
	case "startDate":
		date, err := parseTime(time.DateOnly, value)
		if err != nil {
			return err
		}
		e.StartDate = date
	case "endDate":
		date, err := parseTime(time.DateOnly, value)
		if err != nil {
			return err
		}
		e.EndDate = date
	}

	return nil
}

func (e *EodMetadata) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		Ticker       string `json:"ticker"`
		Name         string `json:"name"`
		ExchangeCode string `json:"exchangeCode"`
		Description  string `json:"description"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
	}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	start, err := parseTime(time.DateOnly, tmp.StartDate)
	if err != nil {
		return fmt.Errorf("could not parse startDate: %w", err)
	}
	end, err := parseTime(time.DateOnly, tmp.EndDate)
	if err != nil {
		return fmt.Errorf("could not parse endDate: %w", err)
	}

	e.Ticker = tmp.Ticker
	e.Name = tmp.Name
	e.ExchangeCode = tmp.ExchangeCode
	e.Description = tmp.Description
	e.StartDate = start
	e.EndDate = end

	return nil
}

// SymbolItem corresponds to [End-of-Day].2.1.3.supported_tickers.zip
type SymbolItem struct {
	Ticker        string    `json:"ticker,omitempty" csv:"ticker,omitempty"`
	Exchange      string    `json:"exchange,omitempty" csv:"exchange,omitempty"`
	AssetType     string    `json:"assetType,omitempty" csv:"assetType,omitempty"`
	PriceCurrency string    `json:"priceCurrency,omitempty" csv:"priceCurrency,omitempty"`
	StartDate     time.Time `json:"startDate,omitempty" csv:"startDate,omitempty"`
	EndDate       time.Time `json:"EndDate,omitempty" csv:"EndDate,omitempty"`
}

func (s *SymbolItem) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "ticker":
		s.Ticker = value
	case "exchange":
		s.Exchange = value
	case "assetType":
		s.AssetType = value
	case "priceCurrency":
		s.PriceCurrency = value
	case "startDate":
		date, err := parseTime(time.DateOnly, value)
		if err != nil {
			return err
		}
		s.StartDate = date
	case "endDate":
		date, err := parseTime(time.DateOnly, value)
		if err != nil {
			return err
		}
		s.EndDate = date
	}

	return nil
}

// IexTopOfBook corresponds to the [IEX].2.5.2 Top-of-Book and Last Price Endpoints
type IexTopOfBook struct {
	Ticker            string    `json:"ticker,omitempty" csv:"ticker,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty" csv:"timestamp,omitempty"`
	QuoteTimestamp    time.Time `json:"quoteTimestamp,omitempty" csv:"quoteTimestamp,omitempty"`
	LastSaleTimestamp time.Time `json:"lastSaleTimestamp,omitempty" csv:"lastSaleTimestamp,omitempty"`
	Last              float64   `json:"last,omitempty" csv:"last,omitempty"`
	LastSize          int32     `json:"lastSize,omitempty" csv:"lastSize,omitempty"`
	TngoLast          float64   `json:"tngoLast,omitempty" csv:"tngoLast,omitempty"`
	PrevClose         float64   `json:"prevClose,omitempty" csv:"prevClose,omitempty"`
	Open              float64   `json:"open,omitempty" csv:"open,omitempty"`
	High              float64   `json:"high,omitempty" csv:"high,omitempty"`
	Low               float64   `json:"low,omitempty" csv:"low,omitempty"`
	Mid               float64   `json:"mid,omitempty" csv:"mid,omitempty"`
	Volume            int64     `json:"volume,omitempty" csv:"volume,omitempty"`
	BidSize           float64   `json:"bidSize,omitempty" csv:"bidSize,omitempty"`
	BidPrice          float64   `json:"bidPrice,omitempty" csv:"bidPrice,omitempty"`
	AskSize           float64   `json:"askSize,omitempty" csv:"askSize,omitempty"`
	AskPrice          float64   `json:"askPrice,omitempty" csv:"askPrice,omitempty"`
}

func (i *IexTopOfBook) UnmarshalCSVWithFields(key, value string) error {
	timeLayouts := []string{time.RFC3339, "2006-01-02T15:04:05"}

	switch key {
	case "ticker":
		i.Ticker = value
	case "timestamp":
		date, err := parseTimeMultiLayout(timeLayouts, value)
		if err != nil {
			return err
		}
		i.Timestamp = date.UTC()
	case "quoteTimestamp":
		date, err := parseTimeMultiLayout(timeLayouts, value)
		if err != nil {
			return err
		}
		i.QuoteTimestamp = date.UTC()
	case "lastSaleTimestamp":
		date, err := parseTimeMultiLayout(timeLayouts, value)
		if err != nil {
			return err
		}
		i.LastSaleTimestamp = date.UTC()
	case "last":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Last = f
	case "lastSize":
		// This can sometimes come in as a float formatted as x.0000, so we
		// need to trim the "." and trailing zeros
		if strings.Contains(value, ".") {
			value = strings.Split(strings.TrimRight(value, "0"), ".")[0]
		}
		int_, err := parseInt(value)
		if err != nil {
			return err
		}
		i.LastSize = int32(int_)
	case "tngoLast":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.TngoLast = f
	case "prevClose":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.PrevClose = f
	case "open":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Open = f
	case "high":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.High = f
	case "low":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Low = f
	case "mid":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Mid = f
	case "volume":
		int_, err := parseInt(value)
		if err != nil {
			return err
		}
		i.Volume = int_
	case "bidSize":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.BidSize = f
	case "bidPrice":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.BidPrice = f
	case "askSize":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.AskSize = f
	case "askPrice":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.AskPrice = f
	}

	return nil
}

func (i *IexTopOfBook) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		Ticker            string  `json:"ticker,omitempty"`
		Timestamp         string  `json:"timestamp,omitempty"`
		QuoteTimestamp    string  `json:"quoteTimestamp,omitempty"`
		LastSaleTimestamp string  `json:"lastSaleTimestamp,omitempty"`
		Last              float64 `json:"last,omitempty"`
		LastSize          int32   `json:"lastSize,omitempty"`
		TngoLast          float64 `json:"tngoLast,omitempty"`
		PrevClose         float64 `json:"prevClose,omitempty"`
		Open              float64 `json:"open,omitempty"`
		High              float64 `json:"high,omitempty"`
		Low               float64 `json:"low,omitempty"`
		Mid               float64 `json:"mid,omitempty"`
		Volume            int64   `json:"volume,omitempty"`
		BidSize           float64 `json:"bidSize,omitempty"`
		BidPrice          float64 `json:"bidPrice,omitempty"`
		AskSize           float64 `json:"askSize,omitempty"`
		AskPrice          float64 `json:"askPrice,omitempty"`
	}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	layouts := []string{time.RFC3339Nano, "2006-01-02T15:04:05"}
	timestamp, err := parseTimeMultiLayout(layouts, tmp.Timestamp)
	if err != nil {
		return err
	}
	quoteTimestamp, err := parseTimeMultiLayout(layouts, tmp.QuoteTimestamp)
	if err != nil {
		return err
	}
	lastSaleTimestamp, err := parseTimeMultiLayout(layouts, tmp.LastSaleTimestamp)
	if err != nil {
		return err
	}

	i.Ticker = tmp.Ticker
	i.Timestamp = timestamp.UTC()
	i.QuoteTimestamp = quoteTimestamp.UTC()
	i.LastSaleTimestamp = lastSaleTimestamp.UTC()
	i.Last = tmp.Last
	i.LastSize = tmp.LastSize
	i.TngoLast = tmp.TngoLast
	i.PrevClose = tmp.PrevClose
	i.Open = tmp.Open
	i.High = tmp.High
	i.Low = tmp.Low
	i.Mid = tmp.Mid
	i.Volume = tmp.Volume
	i.BidSize = tmp.BidSize
	i.BidPrice = tmp.BidPrice
	i.AskSize = tmp.AskSize
	i.AskPrice = tmp.AskPrice

	return nil
}

// IexPrice corresponds to the [IEX].2.5.3 Historical Intraday Prices Endpoint
type IexPrice struct {
	Date   time.Time `json:"date,omitempty" csv:"date,omitempty"`
	Open   float64   `json:"open,omitempty" csv:"open,omitempty"`
	High   float64   `json:"high,omitempty" csv:"high,omitempty"`
	Low    float64   `json:"low,omitempty" csv:"low,omitempty"`
	Close  float64   `json:"close,omitempty" csv:"close,omitempty"`
	Volume int64     `json:"volume,omitempty" csv:"volume,omitempty"`
}

func (i *IexPrice) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "date":
		date, err := parseTime("2006-01-02 15:04:05-07:00", value)
		if err != nil {
			return err
		}
		i.Date = date.UTC()
	case "open":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Open = f
	case "high":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.High = f
	case "low":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Low = f
	case "close":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Close = f
	case "volume":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		i.Volume = int64(f)
	}

	return nil
}

// StmtDef corresponds to the [Fundamentals].2.6.2 Definitions Data Endpoint
type StmtDef struct {
	DataCode      string `json:"dataCode,omitempty" csv:"dataCode,omitempty"`
	Name          string `json:"name,omitempty" csv:"name,omitempty"`
	Description   string `json:"description,omitempty" csv:"description,omitempty"`
	StatementType string `json:"statementType,omitempty" csv:"statementType,omitempty"`
	Units         string `json:"units,omitempty" csv:"units,omitempty"`
}

// StmtDataNested corresponds to the [Fundamentals].2.6.3 Statement Data Endpoint
// when a json respFormat is requested
type StmtDataNested struct {
	Date          time.Time `json:"date,omitempty"`
	Year          int32     `json:"year,omitempty"`
	Quarter       int32     `json:"quarter,omitempty"`
	StatementData struct {
		BalanceSheet    []StmtDataField
		IncomeStatement []StmtDataField
		CashFlow        []StmtDataField
		Overview        []StmtDataField
	}
}

func (s *StmtDataNested) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		Date          string `json:"date,omitempty"`
		Year          int32  `json:"year,omitempty"`
		Quarter       int32  `json:"quarter,omitempty"`
		StatementData struct {
			BalanceSheet    []StmtDataField
			IncomeStatement []StmtDataField
			CashFlow        []StmtDataField
			Overview        []StmtDataField
		}
	}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	date, err := parseTime(time.DateOnly, tmp.Date)
	if err != nil {
		return err
	}

	s.Date = date
	s.Year = tmp.Year
	s.Quarter = tmp.Quarter
	s.StatementData.BalanceSheet = tmp.StatementData.BalanceSheet
	s.StatementData.IncomeStatement = tmp.StatementData.IncomeStatement
	s.StatementData.CashFlow = tmp.StatementData.CashFlow
	s.StatementData.Overview = tmp.StatementData.Overview

	return nil
}

// StmtDataField corresponds to the [Fundamentals].2.6.3 Statement Data Endpoint
// statement data fields
type StmtDataField struct {
	DataCode string  `json:"dataCode,omitempty"`
	Value    float64 `json:"value,omitempty"`
}

// StmtDataFlat corresponds to the [Fundamentals].2.6.3 Statement Data Endpoint
// when a csv respFormat is requested
type StmtDataFlat struct {
	Date          time.Time `json:"date" csv:"date"`
	Year          int       `json:"year" csv:"year"`
	Quarter       int       `json:"quarter" csv:"quarter"`
	StatementType string    `json:"statementType" csv:"statementType"`
	DataCode      string    `json:"dataCode" csv:"dataCode"`
	Value         float64   `json:"value" csv:"value"`
}

func (s *StmtDataFlat) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "date":
		date, err := parseTime(time.DateOnly, value)
		if err != nil {
			return err
		}
		s.Date = date
	case "year":
		i, err := parseInt(value)
		if err != nil {
			return err
		}
		s.Year = int(i)
	case "quarter":
		i, err := parseInt(value)
		if err != nil {
			return err
		}
		s.Quarter = int(i)
	case "statementType":
		s.StatementType = value
	case "dataCode":
		s.DataCode = value
	case "value":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		s.Value = f
	}

	return nil
}

// DailyFundamental corresponds to the [Fundamentals].2.6.4 Daily Data Endpoint
type DailyFundamental struct {
	Date                 time.Time `json:"date,omitempty" csv:"date,omitempty"`
	MarketCap            float64   `json:"marketCap,omitempty" csv:"marketCap,omitempty"`
	EnterpriseValue      float64   `json:"enterpriseVal,omitempty" csv:"enterpriseVal,omitempty"`
	PriceToEarningsRatio float64   `json:"peRatio,omitempty" csv:"peRatio,omitempty"`
	PriceToBookRatio     float64   `json:"pbRatio,omitempty" csv:"pbRatio,omitempty"`
	TrailingYearPEG      float64   `json:"trailingPEG1Y,omitempty" csv:"trailingPEG1Y,omitempty"`
}

func (d *DailyFundamental) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "date":
		date, err := parseTime(time.DateOnly, value)
		if err != nil {
			return err
		}
		d.Date = date
	case "marketCap":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		d.MarketCap = f
	case "enterpriseVal":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		d.EnterpriseValue = f
	case "peRatio":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		d.PriceToEarningsRatio = f
	case "pbRatio":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		d.PriceToBookRatio = f
	case "trailingPEG1Y":
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		d.TrailingYearPEG = f
	}

	return nil
}

// FundamentalMetadata corresponds to the [Fundamentals].2.6.5 Meta Data Endpoint
type FundamentalMetadata struct {
	PermaTicker             string    `json:"permaTicker,omitempty" csv:"permaTicker,omitempty"`
	Ticker                  string    `json:"ticker,omitempty" csv:"ticker,omitempty"`
	Name                    string    `json:"name,omitempty" csv:"name,omitempty"`
	IsActive                bool      `json:"isActive,omitempty" csv:"isActive,omitempty"`
	IsADR                   bool      `json:"isADR,omitempty" csv:"isADR,omitempty"`
	Sector                  string    `json:"sector,omitempty" csv:"sector,omitempty"`
	Industry                string    `json:"industry,omitempty" csv:"industry,omitempty"`
	SicCode                 int32     `json:"sicCode,omitempty" csv:"sicCode,omitempty"`
	SicSector               string    `json:"sicSector,omitempty" csv:"sicSector,omitempty"`
	SicIndustry             string    `json:"sicIndustry,omitempty" csv:"sicIndustry,omitempty"`
	ReportingCurrency       string    `json:"reportingCurrency,omitempty" csv:"reportingCurrency,omitempty"`
	Location                string    `json:"location,omitempty" csv:"location,omitempty"`
	CompanyWebsite          string    `json:"companyWebsite,omitempty" csv:"companyWebsite,omitempty"`
	SecFilingWebsite        string    `json:"secFilingWebsite,omitempty" csv:"secFilingWebsite,omitempty"`
	StatementLastUpdated    time.Time `json:"statementLastUpdated,omitempty" csv:"statementLastUpdated,omitempty"`
	DailyLastUpdated        time.Time `json:"dailyLastUpdated,omitempty" csv:"dailyLastUpdated,omitempty"`
	DataProviderPermaTicker string    `json:"dataProviderPermaTicker,omitempty" csv:"dataProviderPermaTicker,omitempty"`
}

func (f *FundamentalMetadata) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "permaTicker":
		f.PermaTicker = value
	case "ticker":
		f.Ticker = value
	case "name":
		f.Name = value
	case "isActive":
		b, err := parseBool(value)
		if err != nil {
			return err
		}
		f.IsActive = b
	case "isADR":
		b, err := parseBool(value)
		if err != nil {
			return err
		}
		f.IsADR = b
	case "sector":
		f.Sector = value
	case "industry":
		f.Industry = value
	case "sicCode":
		i, err := parseInt(value)
		if err != nil {
			return err
		}
		f.SicCode = int32(i)
	case "sicSector":
		f.SicSector = value
	case "sicIndustry":
		f.SicIndustry = value
	case "reportingCurrency":
		f.ReportingCurrency = value
	case "location":
		f.Location = value
	case "companyWebsite":
		f.CompanyWebsite = value
	case "secFilingWebsite":
		f.SecFilingWebsite = value
	case "statementLastUpdated":
		date, err := parseTime("2006-01-02 15:04:05.9-07:00", value)
		if err != nil {
			return err
		}
		f.StatementLastUpdated = date.UTC()
	case "dailyLastUpdated":
		date, err := parseTime("2006-01-02 15:04:05.9-07:00", value)
		if err != nil {
			return err
		}
		f.DailyLastUpdated = date.UTC()
	case "dataProviderPermaTicker":
		f.DataProviderPermaTicker = value
	}

	return nil
}

// SearchResult corresponds to [Utility].4.1.2 Search Endpoint
type SearchResult struct {
	Ticker            string `json:"ticker,omitempty" csv:"ticker,omitempty"`
	Name              string `json:"name,omitempty" csv:"name,omitempty"`
	AssetType         string `json:"assetType,omitempty" csv:"assetType,omitempty"`
	IsActive          bool   `json:"isActive,omitempty" csv:"isActive,omitempty"`
	PermaTicker       string `json:"permaTicker,omitempty" csv:"permaTicker,omitempty"`
	OpenFIGIComposite string `json:"openFIGIComposite,omitempty" csv:"openFIGIComposite,omitempty"`
	CountryCode       string `json:"countryCode" csv:"countryCode"`
}

// PriceData corresponds to the [Crypto].2.3.2 priceData Response
type PriceData struct {
	Date 			time.Time `json:"date,omitempty"`
	Open			float64   `json:"open,omitempty"`
	High			float64   `json:"high,omitempty"`
	Low				float64   `json:"low,omitempty"`
	Close			float64   `json:"close,omitempty"`
	TradesDone		float64     `json:"tradesDone,omitempty"`
	Volume			float64   `json:"volume,omitempty"`
	VolumeNotional	float64   `json:"volumeNotional,omitempty"`
}

// CryptoResults corresponds to the [Crypto].2.3.2 Crypto Endpoint
type CryptoResult struct {
	Ticker 			  string 		`json:"ticker,omitempty"`
	BaseCurrency	  string 		`json:"baseCurrency,omitempty"`
	QuoteCurrency	  string 		`json:"quoteCurrency,omitempty"`
	PriceData		  []PriceData 	`json:"priceData,omitempty"`
	//ExchangeData	  ExchangeData  `json:"exchangeData,omitempty"`
}

func parseFloat(value string) (float64, error) {
	if value == "" {
		return 0, nil
	}
	return strconv.ParseFloat(value, 64)
}

func parseInt(value string) (int64, error) {
	if value == "" {
		return 0, nil
	}
	return strconv.ParseInt(value, 10, 64)
}

func parseTime(layout, value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(layout, value, time.UTC)
}

func parseBool(value string) (bool, error) {
	if value == "" {
		return false, nil
	}
	return strconv.ParseBool(value)
}

func parseTimeMultiLayout(layouts []string, value string) (time.Time, error) {
	var errs []error
	for _, layout := range layouts {
		t, err := parseTime(layout, value)
		if err == nil {
			return t, nil
		}
		errs = append(errs, err)
	}

	return time.Time{}, fmt.Errorf("multipl errors: %v", errs)
}

// Parse takes in raw bytes and parses them according to the format. Valid
// formats are CSV or JSON (an empty string defaults to JSON).
func Parse[T any](rawBytes []byte, format Format) (T, error) {
	var data T
	var err error

	fmt.Println(string(rawBytes))

	switch format {
	case JSON, "":
		err = json.Unmarshal(rawBytes, &data)
	case CSV:
		err = gocsv.UnmarshalBytes(rawBytes, &data)
	default:
		return data, fmt.Errorf("format not recognized: %s", format)
	}

	if err != nil {
		return data, fmt.Errorf("failed to parse raw bytes: %w", err)
	}

	return data, nil
}
