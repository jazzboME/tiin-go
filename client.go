package tiingo

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"sync/atomic"
	"time"
)

type Client struct {
	HttpClient  *http.Client  // default client that all tiingo requests are routed through
	RateLimiter Limiter       // rate limits request speed
	Logger      *slog.Logger  // default logger
	apiToken    string        // tiingo API key
	reqCount    atomic.Uint64 // # of requests the client sent
}

type Limiter interface {
	Wait(ctx context.Context) error
}

// NewClient initializes a new Tiingo client.
//
// Specific config options can be set by providing an options func that will
// be applied to the returned Client.
func NewClient(apiToken string, options ...func(*Client)) *Client {
	// Init default c
	c := Client{
		apiToken: apiToken,
	}

	// Apply any options
	for _, optFunc := range options {
		if optFunc == nil {
			continue
		}
		optFunc(&c)
	}

	// Set any options that are nil with defaults
	if c.HttpClient == nil {
		c.HttpClient = http.DefaultClient
	}
	if c.Logger == nil {
		c.Logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}

	return &c
}

func (c *Client) get(ctx context.Context, url string) ([]byte, error) {
	// Build request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("faled to build request for url=%s: %w", url, err)
	}

	// Set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+c.apiToken)
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Keep-Alive", "timeout=5, max=1000")

	// Log request
	l := c.Logger.With(
		slog.String("requestID", uuid()),
		slog.String("url", url),
	)
	l.Debug("request created")

	// Block on rate limiter
	if c.RateLimiter != nil {
		if err = c.RateLimiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limiting failed: %w", err)
		}
	}

	// Make request
	s := time.Now()
	l.Info("making request",
		slog.Uint64("reqCount", c.reqCount.Add(1)),
	)
	resp, err := c.HttpClient.Do(req)
	l.Info("request completed",
		slog.Int64("totalMS", time.Since(s).Milliseconds()),
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			l.Warn("failed to close request body",
				slog.String("error", err.Error()),
			)
		}
	}()

	// Validate status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status: %d", resp.StatusCode)
	}

	// Read the body so the http.RoundTripper can be closed & immediately reused
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	return data, nil
}

func uuid() string {
	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randSeq := make([]byte, 8)
	for i := range 8 {
		randSeq[i] = chars[rand.IntN(len(chars))]
	}

	return string(randSeq)
}
