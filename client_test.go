package tiingo

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

var (
	getClient = sync.OnceValue(func() *Client {
		apiToken, ok := os.LookupEnv("TIINGO_TOKEN")
		if !ok {
			panic("TIINGO_TOKEN must be set if doing a live client test")
		}
		return NewClient(apiToken, func(client *Client) {
			client.RateLimiter = rate.NewLimiter(10, 1)
			// client.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		})
	})
	startDate = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate   = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
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

func TestNewClient(t *testing.T) {
	rl := rate.NewLimiter(1, 1)
	type args struct {
		apiToken string
		options  []func(*Client)
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "nilOptions",
			args: args{
				options: nil,
			},
			want: &Client{
				HttpClient: http.DefaultClient,
				Logger:     slog.New(slog.NewTextHandler(io.Discard, nil)),
			},
		},
		{
			name: "nilSlice",
			args: args{
				options: []func(*Client){nil, nil, nil},
			},
			want: &Client{
				HttpClient: http.DefaultClient,
				Logger:     slog.New(slog.NewTextHandler(io.Discard, nil)),
			},
		},
		{
			name: "basicOptionsFunc",
			args: args{
				options: []func(*Client){func(c *Client) {
					c.Logger = slog.Default()
					c.HttpClient = http.DefaultClient
					c.RateLimiter = rl
				}},
			},
			want: &Client{
				HttpClient:  http.DefaultClient,
				Logger:      slog.Default(),
				RateLimiter: rl,
			},
		},
		{
			name: "basicMultipleOptionsFunc",
			args: args{
				options: []func(*Client){
					func(c *Client) {
						c.Logger = slog.Default()
					},
					func(c *Client) {
						c.RateLimiter = rl
					},
					func(c *Client) {
						c.HttpClient = http.DefaultClient
					},
				},
			},
			want: &Client{
				HttpClient:  http.DefaultClient,
				Logger:      slog.Default(),
				RateLimiter: rl,
			},
		},
		{
			name: "withOptionsFuncs",
			args: args{
				options: []func(*Client){
					WithLogger(slog.Default()),
					WithRateLimiter(rl),
					WithHttpClient(http.DefaultClient),
				},
			},
			want: &Client{
				HttpClient:  http.DefaultClient,
				Logger:      slog.Default(),
				RateLimiter: rl,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.apiToken, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ge_basic(t *testing.T) {
	// Init all the stuff
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Token FAKE" {
			w.WriteHeader(666)
			return
		}
		content := r.Header.Get("Content-Type")
		if content != "application/json" {
			w.WriteHeader(667)
			return
		}

		_, _ = w.Write([]byte(`{"msg":"working"}`))
	}))
	defer svr.Close()
	badSvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(666)
	}))
	defer badSvr.Close()

	// Headers are set
	c := NewClient("FAKE", nil)
	data, err := c.get(context.Background(), svr.URL)
	if err != nil {
		t.Errorf("header test failed: %s", err)
	}

	// Reading body correctly
	if string(data) != `{"msg":"working"}` {
		t.Errorf(`body test failed. got = %s. wanted = {"msg":"working"}`, data)
	}

	// Reading status code correctly
	data, err = c.get(context.Background(), badSvr.URL)
	if len(data) != 0 && err == nil {
		t.Errorf("not processing bad status codes correctly")
	}
}

func TestClient_get_rateLimiting(t *testing.T) {
	// Init server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svr.Close()

	// Init client
	c := NewClient("",
		WithRateLimiter(rate.NewLimiter(10, 1)),
	)

	s := time.Now()
	_, _ = c.get(context.Background(), svr.URL)
	_, _ = c.get(context.Background(), svr.URL)
	_, _ = c.get(context.Background(), svr.URL)
	if time.Since(s).Milliseconds() < 200 {
		t.Fatalf("rate limiter not working, 3 requests only took %d milliseconds", time.Since(s).Milliseconds())
	}
}
