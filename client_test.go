package tiingo

import (
	"fmt"
	"os"
	"reflect"
	"sync"

	"golang.org/x/time/rate"
)

var (
	getClient = sync.OnceValue[*Client](func() *Client {
		return NewClient(os.Getenv("TIINGO_TOKEN"), func(client *Client) {
			client.RateLimiter = rate.NewLimiter(10, 1)
			// client.Logger = slog.Default()
		})
	})
)

func liveTest[T any](fName string, wantErr bool, execute func() (T, error)) error {
	got, err := execute()
	if wantErr {
		if err == nil {
			return fmt.Errorf("%s wanted error, got nil", fName)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("%s returned non-nil error: %s", fName, err)
	}

	var zeroValue T
	if reflect.DeepEqual(got, zeroValue) {
		return fmt.Errorf("%s returned data is the zero value", fName)
	}

	return nil
}

//
// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log/slog"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"reflect"
// 	"strings"
// 	"sync"
// 	"sync/atomic"
// 	"testing"
// 	"time"
//
// 	"github.com/gocarina/gocsv"
// 	"golang.org/x/time/rate"
// )
//
// var (
// 	start              = time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
// 	startOlder         = time.Date(2020, 12, 12, 0, 0, 0, 0, time.UTC)
// 	end                = time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
// 	client     *Client = nil
// 	m          sync.Mutex
// )
//
// func getClient() (*Client, error) {
// 	m.Lock()
// 	defer m.Unlock()
//
// 	if client != nil {
// 		return client, nil
// 	}
//
// 	apiToken, ok := os.LookupEnv("TIINGO_TOKEN")
// 	if !ok {
// 		return nil, errors.New("TIINGO_TOKEN env variable not set")
// 	}
// 	client = NewClient(apiToken, func(client *Client) {
// 		client.RateLimiter = rate.NewLimiter(30, 1)
// 		// client.Logger = slog.Default()
// 	})
//
// 	return client, nil
// }
//
// func TestNewClient(t *testing.T) {
// 	t.Parallel()
//
// 	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
// 	httpClient := http.Client{
// 		Timeout: time.Hour * 100,
// 	}
// 	nilLogger := slog.New(slog.NewTextHandler(io.Discard, nil))
//
// 	type args struct {
// 		apiToken string
// 		options  []func(*Client)
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *Client
// 	}{
// 		{
// 			name: "empty",
// 			args: args{},
// 			want: &Client{
// 				HttpClient:  http.DefaultClient,
// 				RateLimiter: nil,
// 				Logger:      nilLogger,
// 				apiToken:    "",
// 				reqCount:    atomic.Uint64{},
// 			},
// 		},
// 		{
// 			name: "nilOptions",
// 			args: args{
// 				options: nil,
// 			},
// 			want: &Client{
// 				HttpClient:  http.DefaultClient,
// 				RateLimiter: nil,
// 				Logger:      nilLogger,
// 				apiToken:    "",
// 				reqCount:    atomic.Uint64{},
// 			},
// 		},
// 		{
// 			name: "basic",
// 			args: args{
// 				apiToken: "3487gaosdiy",
// 				options: []func(*Client){func(c *Client) {
// 					c.HttpClient = &httpClient
// 					c.Logger = l
// 				}},
// 			},
// 			want: &Client{
// 				HttpClient: &httpClient,
// 				Logger:     l,
// 				apiToken:   "3487gaosdiy",
// 				reqCount:   atomic.Uint64{},
// 			},
// 		},
// 		{
// 			name: "multipleOptionsFunc",
// 			args: args{
// 				apiToken: "3487gaosdiy",
// 				options: []func(*Client){
// 					func(c *Client) {
// 						c.HttpClient = &httpClient
// 						c.Logger = l
// 					},
// 					func(c *Client) {
// 						c.HttpClient = http.DefaultClient
// 					},
// 					func(c *Client) {
// 						c.RateLimiter = rate.NewLimiter(1, 1)
// 					},
// 					func(c *Client) {
// 						c.Logger = slog.Default()
// 					},
// 				},
// 			},
// 			want: &Client{
// 				HttpClient:  http.DefaultClient,
// 				RateLimiter: rate.NewLimiter(1, 1),
// 				Logger:      slog.Default(),
// 				apiToken:    "3487gaosdiy",
// 				reqCount:    atomic.Uint64{},
// 			},
// 		},
// 		{
// 			name: "badOptionsFunc",
// 			args: args{
// 				apiToken: "3487gaosdiy",
// 				options: []func(*Client){func(c *Client) {
// 					c.HttpClient = nil
// 					c.Logger = nil
// 				}},
// 			},
// 			want: &Client{
// 				HttpClient: http.DefaultClient,
// 				Logger:     nilLogger,
// 				apiToken:   "3487gaosdiy",
// 				reqCount:   atomic.Uint64{},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			if got := NewClient(tt.args.apiToken, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewClient() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestClient_get(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		auth := r.Header.Get("Authorization")
// 		if auth != "Token FAKE" {
// 			w.WriteHeader(666)
// 			return
// 		}
// 		content := r.Header.Get("Content-Type")
// 		if content != "application/json" {
// 			w.WriteHeader(667)
// 			return
// 		}
//
// 		_, _ = w.Write([]byte(`{"msg":"working"}`))
// 	}))
// 	defer svr.Close()
// 	badSvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(666)
// 	}))
// 	defer badSvr.Close()
//
// 	// Headers are set
// 	c := NewClient("FAKE", nil)
// 	data, err := c.get(context.Background(), svr.URL)
// 	if err != nil {
// 		t.Errorf("header test failed: %s", err)
// 	}
//
// 	// Reading body correctly
// 	if string(data) != `{"msg":"working"}` {
// 		t.Errorf(`body test failed. got = %s. wanted = {"msg":"working"}`, data)
// 	}
//
// 	// Reading status code correctly
// 	data, err = c.get(context.Background(), badSvr.URL)
// 	if len(data) != 0 && err == nil {
// 		t.Errorf("not processing bad status codes correctly")
// 	}
// }
//
// func TestClient_RateLimiter(t *testing.T) {
// 	t.Parallel()
//
// 	// Init server
// 	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	defer svr.Close()
//
// 	// Init client
// 	c := NewClient("", func(c *Client) {
// 		c.RateLimiter = rate.NewLimiter(10, 1)
// 	})
//
// 	s := time.Now()
// 	_, _ = c.get(context.Background(), svr.URL)
// 	_, _ = c.get(context.Background(), svr.URL)
// 	_, _ = c.get(context.Background(), svr.URL)
// 	if time.Since(s).Milliseconds() < 200 {
// 		t.Fatalf("rate limiter not working, 3 requests only took %d milliseconds", time.Since(s).Milliseconds())
// 	}
// }
//
// func unmarshal[T any](bytes []byte, k kind) (T, error) {
// 	var data T
// 	if k == json_ {
// 		if err := json.Unmarshal(bytes, &data); err != nil {
// 			return data, err
// 		}
// 	} else if k == csv_ {
// 		if err := gocsv.UnmarshalBytes(bytes, &data); err != nil {
// 			return data, err
// 		}
// 	} else {
// 		return data, fmt.Errorf("kind not recognized: %v", k)
// 	}
//
// 	return data, nil
// }
//
// func unmarshalAndValidateStruct[T any](bytes []byte, k kind) error {
// 	data, err := unmarshal[T](bytes, k)
// 	if err != nil {
// 		return fmt.Errorf("unmarshall of %v bytes failed: %w", k, err)
// 	}
//
// 	var zeroValue T
// 	if reflect.DeepEqual(data, zeroValue) {
// 		return fmt.Errorf("data=%v is the zero value of %T", data, zeroValue)
// 	}
//
// 	return nil
// }
//
// func unmarshalAndValidateSlice[T any](bytes []byte, k kind) error {
// 	data, err := unmarshal[[]T](bytes, k)
// 	if err != nil {
// 		return fmt.Errorf("unmarshall of %v bytes failed: %w", k, err)
// 	}
//
// 	if len(data) == 0 {
// 		return errors.New("response is empty: len(data) == 0")
// 	}
//
// 	var zeroValue T
// 	for i := range data {
// 		if reflect.DeepEqual(data[i], zeroValue) {
// 			return fmt.Errorf("data[%d]=%v is the zero value of %T", i, data[i], zeroValue)
// 		}
// 	}
//
// 	return nil
// }
//
// func TestClient_DefaultEndpoints(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	tests := []struct {
// 		name         string
// 		fetchFunc    func() ([]byte, error)
// 		validateFunc func([]byte, kind) error
// 		kind         kind
// 	}{
// 		{
// 			name: "EodPriceDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultEodPrice(ctx, "AAPL")
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[EodPrice](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "EodMetadataDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultEodMetadata(ctx, "AAPL")
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateStruct[EodMetadata](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "SymbolListDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultSymbolList(ctx)
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				if len(bytes) == 0 {
// 					return errors.New("response is empty")
// 				}
// 				return nil
// 			},
// 			kind: 0,
// 		},
// 		{
// 			name: "IexTopOfBookDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultIexTopOfBook(ctx)
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[IexTopOfBook](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "IexHistoryDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultIexHistory(ctx, "AAPL")
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[IexPrice](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "StmtDefsDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultStmtDefs(ctx)
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[StmtDef](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "StmtValsDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultStmtData(ctx, "AAPL")
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[StmtDataNested](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "DailyFundamentalDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultDailyFundamental(ctx, "AAPL")
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[DailyFundamental](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "FundamentalMetadataDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultFundamentalMetadata(ctx)
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[FundamentalMetadata](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 		{
// 			name: "SearchDefault",
// 			fetchFunc: func() ([]byte, error) {
// 				return c.DefaultSearch(ctx, "AAPL")
// 			},
// 			validateFunc: func(bytes []byte, k kind) error {
// 				return unmarshalAndValidateSlice[SearchResult](bytes, k)
// 			},
// 			kind: json_,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response, tt.kind); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// type clientTest struct {
// 	name         string
// 	fetchFunc    func() ([]byte, error)
// 	validateFunc func([]byte) error
// }
//
// func TestClient_EodPrice(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	// Build the tests
// 	freqs := []EodFreq{Daily, Weekly, Monthly, Annually}
// 	sorts := []Sort{DateAsc, DateDesc, OpenAsc, OpenDesc, HighAsc, HighDesc,
// 		LowAsc, LowDesc, CloseAsc, CloseDesc, VolumeAsc, VolumeDesc, AdjOpenAsc,
// 		AdjOpenDesc, AdjHighAsc, AdjHighDesc, AdjLowAsc, AdjLowDesc, AdjCloseAsc,
// 		AdjCloseDesc, AdjVolumeAsc, AdjVolumeDesc, DivCashAsc, DivCashDesc,
// 		SplitFactorAsc, SplitFactorDesc}
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, freq := range freqs {
// 		for _, sort := range sorts {
// 			for _, format := range formats {
// 				var k kind
// 				if format == CSV {
// 					k = csv_
// 				} else {
// 					k = json_
// 				}
//
// 				tests = append(tests, clientTest{
// 					name: fmt.Sprintf("resampleFreq='%s'__sort='%s'__format='%s'", freq, sort, format),
// 					fetchFunc: func() ([]byte, error) {
// 						return c.EodPrice(ctx, "AAPL", start, end, freq, sort, format, nil)
// 					},
// 					validateFunc: func(b []byte) error {
// 						return unmarshalAndValidateSlice[EodPrice](b, k)
// 					},
// 				})
// 			}
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_EodMetadata(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	// Build the tests
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, format := range formats {
// 		tests = append(tests, clientTest{
// 			name: fmt.Sprintf("format=%s", format),
// 			fetchFunc: func() ([]byte, error) {
// 				return c.EodMetadata(ctx, "AAPL", format)
// 			},
// 			validateFunc: func(b []byte) error {
// 				if format == CSV {
// 					return unmarshalAndValidateSlice[EodMetadata](b, csv_)
// 				} else {
// 					return unmarshalAndValidateStruct[EodMetadata](b, json_)
// 				}
// 			},
// 		})
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_IexTopOfBook(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	// Build the tests
// 	formats := []Format{CSV, JSON}
// 	tickers := [][]string{{"AAPL"}, {"AAPL", "GOOG", "MSFT"}, nil}
// 	var tests []clientTest
// 	for _, symbols := range tickers {
// 		for _, format := range formats {
// 			tests = append(tests, clientTest{
// 				name: fmt.Sprintf("symbols=[%s]__format=%s", strings.Join(symbols, ","), format),
// 				fetchFunc: func() ([]byte, error) {
// 					return c.IexTopOfBook(ctx, symbols, format)
// 				},
// 				validateFunc: func(b []byte) error {
// 					if format == CSV {
// 						return unmarshalAndValidateSlice[IexTopOfBook](b, csv_)
// 					} else {
// 						return unmarshalAndValidateSlice[IexTopOfBook](b, json_)
// 					}
// 				},
// 			})
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_IexHistory(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	// Build the tests
// 	freqs := []IexFreq{OneMin, FiveMin, FifteenMin, ThirtyMin, OneHour, TwoHour, FourHour}
// 	afterHours := []bool{true, false}
// 	forceFill := []bool{true, false}
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, freq := range freqs {
// 		for _, ah := range afterHours {
// 			for _, ff := range forceFill {
// 				for _, format := range formats {
// 					tests = append(tests, clientTest{
// 						name: fmt.Sprintf("freq=%v__afterHours=%t__forceFill=%t__format=%v", freq, ah, ff, format),
// 						fetchFunc: func() ([]byte, error) {
// 							return c.IexHistory(ctx, "AAPL", start, end, freq, ah, ff, format)
// 						},
// 						validateFunc: func(b []byte) error {
// 							if format == CSV {
// 								return unmarshalAndValidateSlice[IexPrice](b, csv_)
// 							} else {
// 								return unmarshalAndValidateSlice[IexPrice](b, json_)
// 							}
// 						},
// 					})
// 				}
// 			}
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_StmtDefs(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	// Build the tests
// 	tickers := [][]string{{"AAPL"}, {"AAPL", "GOOG", "MSFT"}, nil}
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, symbols := range tickers {
// 		for _, format := range formats {
// 			tests = append(tests, clientTest{
// 				name: fmt.Sprintf("tickers=%v__format=%v", symbols, format),
// 				fetchFunc: func() ([]byte, error) {
// 					return c.StmtDefs(ctx, symbols, format)
// 				},
// 				validateFunc: func(b []byte) error {
// 					if format == CSV {
// 						return unmarshalAndValidateSlice[StmtDef](b, csv_)
// 					} else {
// 						return unmarshalAndValidateSlice[StmtDef](b, json_)
// 					}
// 				},
// 			})
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_StmtData(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	// Build the tests
// 	asReported := []bool{true, false}
// 	sorts := []Sort{DateAsc, DateDesc}
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, r := range asReported {
// 		for _, sort := range sorts {
// 			for _, format := range formats {
// 				tests = append(tests, clientTest{
// 					name: fmt.Sprintf("asReported=%t__sort=%v__format=%v", r, sort, format),
// 					fetchFunc: func() ([]byte, error) {
// 						return c.StmtData(ctx, "AAPL", r, startOlder, end, sort, format)
// 					},
// 					validateFunc: func(b []byte) error {
// 						if format == CSV {
// 							return unmarshalAndValidateSlice[StmtDataFlat](b, csv_)
// 						} else {
// 							return unmarshalAndValidateSlice[StmtDataNested](b, json_)
// 						}
// 					},
// 				})
// 			}
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_DailyFundamental(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	sorts := []Sort{DateAsc, DateDesc}
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, sort := range sorts {
// 		for _, format := range formats {
// 			tests = append(tests, clientTest{
// 				name: fmt.Sprintf("sort=%v__format=%v", sort, format),
// 				fetchFunc: func() ([]byte, error) {
// 					return c.DailyFundamental(ctx, "AAPL", start, end, sort, format)
// 				},
// 				validateFunc: func(b []byte) error {
// 					if format == CSV {
// 						return unmarshalAndValidateSlice[DailyFundamental](b, csv_)
// 					} else {
// 						return unmarshalAndValidateSlice[DailyFundamental](b, json_)
// 					}
// 				},
// 			})
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_FundamentalMetadata(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	tickers := [][]string{
// 		{"AAPL"},
// 		{"AAPL", "GOOG", "MSFT"},
// 		nil,
// 	}
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, symbols := range tickers {
// 		for _, format := range formats {
// 			tests = append(tests, clientTest{
// 				name: fmt.Sprintf("tickers=[%s]__format=%v", strings.Join(symbols, ","), format),
// 				fetchFunc: func() ([]byte, error) {
// 					return c.FundamentalMetadata(ctx, symbols, format)
// 				},
// 				validateFunc: func(b []byte) error {
// 					if format == CSV {
// 						return unmarshalAndValidateSlice[FundamentalMetadata](b, csv_)
// 					} else {
// 						return unmarshalAndValidateSlice[FundamentalMetadata](b, json_)
// 					}
// 				},
// 			})
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
//
// func TestClient_Search(t *testing.T) {
// 	t.Parallel()
//
// 	// Init all the stuff
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to init Tiingo client: %s", err)
// 	}
// 	ctx := context.Background()
//
// 	exactMatches := []bool{true, false}
// 	includeDelisted := []bool{true, false}
// 	formats := []Format{CSV, JSON}
// 	var tests []clientTest
// 	for _, match := range exactMatches {
// 		for _, d := range includeDelisted {
// 			for _, format := range formats {
// 				tests = append(tests, clientTest{
// 					name: fmt.Sprintf("exactMatch=%t__includeDelited=%t__format=%v", match, d, format),
// 					fetchFunc: func() ([]byte, error) {
// 						return c.Search(ctx, "AAPL", match, d, 100, format, nil)
// 					},
// 					validateFunc: func(b []byte) error {
// 						if format == CSV {
// 							return unmarshalAndValidateSlice[SearchResult](b, csv_)
// 						} else {
// 							return unmarshalAndValidateSlice[SearchResult](b, json_)
// 						}
// 					},
// 				})
// 			}
// 		}
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			// Fetch the data
// 			var response []byte
// 			response, err = tt.fetchFunc()
// 			if err != nil {
// 				t.Errorf("failed to fetch data: %s", err)
// 			}
//
// 			// Validate response data
// 			if err = tt.validateFunc(response); err != nil {
// 				t.Errorf("validation failed: %s", err)
// 			}
// 		})
// 	}
// }
