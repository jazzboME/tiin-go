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
	RateLimiter RateLimiter   // rate limits request speed
	Logger      *slog.Logger  // default logger
	apiToken    string        // tiingo API key
	reqCount    atomic.Uint64 // # of requests the client sent
}

// RateLimiter rate limits requests. It takes in a context, that if canceled,
// will stop the current waiting and return an error. If the context is not
// canceled, and the limiting time is reached, the returned error should be nil.
//
// NOTE: this most seamlessly integrates with a rate.Limiter.
type RateLimiter interface {
	Wait(ctx context.Context) error
}

// WithRateLimiter sets the *Client's RateLimiter as the provided limiter.
func WithRateLimiter(l RateLimiter) func(*Client) {
	return func(c *Client) {
		c.RateLimiter = l
	}
}

// WithLogger sets the *Client's *slog.Logger as the provided logger.
func WithLogger(l *slog.Logger) func(*Client) {
	return func(c *Client) {
		c.Logger = l
	}
}

// WithHttpClient sets the *Client's *http.Client's all requests get routed
// through as provided client.
func WithHttpClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.HttpClient = client
	}
}

// NewClient initializes a new Tiingo client.
//
// Specific config options can be set by providing options funcs that will be
// applied to the new client. The following are already pre-defined:
//   - WithRateLimiter
//   - WithLogger
//   - WithHttpClient
//
// If no options are specified, these are the default options:
//   - No RateLimiter
//   - *slog.Logger that sends all logs to io.Discard
//   - *http.DefaultClient
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

// get is the main request executioner for the *Client. It builds the request based
// on url, applies headers (including auth), blocks on any rate limiter, and then
// does the request. When the request is finished, the body is fully read and
// returned.
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
		slog.String("requestID", newID()),
		slog.String("url", url),
	)
	l.DebugContext(ctx, "request created")

	// Block on rate limiter
	if c.RateLimiter != nil {
		s := time.Now()
		l.DebugContext(ctx, "request blocking on rate limiter")
		if err = c.RateLimiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limiting failed: %w", err)
		}
		l.DebugContext(ctx, "request unblocked from rate limiter",
			slog.Duration("blockedDuration", time.Since(s)),
		)
	}

	// Make request
	s := time.Now()
	l.InfoContext(ctx, "making request",
		slog.Uint64("reqCount", c.reqCount.Add(1)),
	)
	resp, err := c.HttpClient.Do(req)
	l.InfoContext(ctx, "request completed",
		slog.Int64("totalMS", time.Since(s).Milliseconds()),
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			l.WarnContext(ctx, "failed to close request body",
				slog.String("error", err.Error()),
			)
		}
	}()

	// Validate status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status: %d", resp.StatusCode)
	}

	// Read the body so the http.RoundTripper can be immediately reused
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	return data, nil
}

// newID generate a new pseudorandom id of length 8
func newID() string {
	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randSeq := make([]byte, 8)
	for i := range 8 {
		randSeq[i] = chars[rand.IntN(len(chars))]
	}

	return string(randSeq)
}
