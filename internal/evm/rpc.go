package evm

import (
	"context"
	"math"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RetryConfig struct {
	MaxRetries  int           `mapstructure:"max_retries"  toml:"max_retries"`  // Maximum number of retry attempts
	BaseDelay   time.Duration `mapstructure:"base_delay"   toml:"base_delay"`   // Base delay duration between retries
	MaxDelay    time.Duration `mapstructure:"max_delay"    toml:"max_delay"`    // Maximum delay duration between retries
	BackoffRate float64       `mapstructure:"backoff_rate" toml:"backoff_rate"` // Exponential backoff multiplier
}

func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:  5,
		BaseDelay:   100 * time.Millisecond,
		MaxDelay:    5 * time.Second,
		BackoffRate: 2.0,
	}
}

func (c *RetryConfig) Validate() error {
	if c.MaxRetries < 0 {
		return errors.New("max_retries must be greater than 0")
	}
	if c.MaxDelay > 0 {
		if c.BaseDelay > c.MaxDelay {
			return errors.New("base_delay must be less than max_delay")
		}
		if c.BackoffRate <= 1 {
			return errors.New("backoff_rate must be greater than 1")
		}
	}
	return nil
}

func NewClient(ctx context.Context, url string, config *RetryConfig) (*ethclient.Client, error) {
	httpClient := &http.Client{
		Transport: newRetryTransport(http.DefaultTransport, config),
		Timeout:   5 * time.Second,
	}

	c, err := rpc.DialOptions(ctx, url, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, errors.Wrapf(err, "dial rpc url %s", url)
	}

	return ethclient.NewClient(c), nil
}

type RetryTransport struct {
	logger zerolog.Logger
	http.RoundTripper
	config *RetryConfig
}

// newRetryTransport creates a transport wrapper with retry functionality
func newRetryTransport(tripper http.RoundTripper, config *RetryConfig) *RetryTransport {
	if config == nil {
		config = DefaultRetryConfig()
	}
	return &RetryTransport{
		logger:       log.With().Str("module", "RetryTransport").Logger(),
		RoundTripper: tripper,
		config:       config,
	}
}

// RoundTrip implements the http.RoundTripper interface, providing retry functionality
func (rt *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var lastErr error
	var lastResp *http.Response

	// Log request details for debugging
	requestDump, err := httputil.DumpRequest(req, true)
	if err == nil {
		rt.logger.Debug().Str("request", string(requestDump)).Msg("HTTP Request")
	}

	for attempt := 0; attempt <= rt.config.MaxRetries; attempt++ {
		if attempt > 0 {
			delay := calculateBackoff(attempt, rt.config)
			rt.logger.Info().
				Str("url", req.URL.String()).
				Int("attempt", attempt+1).
				Int("max_attempts", rt.config.MaxRetries+1).
				Dur("delay", delay).
				Msg("Retrying HTTP request")

			select {
			case <-req.Context().Done():
				return nil, req.Context().Err()
			case <-time.After(delay):
			}
		}

		var reqToUse *http.Request
		if attempt == 0 {
			reqToUse = req
		} else {
			// Retry attempts: create a copy since the body may be consumed
			reqToUse = req.Clone(req.Context())
		}

		resp, err := rt.RoundTripper.RoundTrip(reqToUse)
		if err != nil {
			lastErr = err
			rt.logger.Warn().Err(err).
				Int("attempt", attempt+1).
				Msg("HTTP request attempt failed")
			continue
		}

		// Check response status code to decide whether to retry
		if shouldRetry(resp.StatusCode) {
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
			lastResp = resp
			lastErr = errors.Errorf("retryable status code %d", resp.StatusCode)
			rt.logger.Warn().
				Int("attempt", attempt+1).
				Int("status_code", resp.StatusCode).
				Msg("HTTP request attempt returned retryable status, will retry")
			continue
		}

		// Successful response
		if attempt > 0 {
			rt.logger.Info().
				Int("attempts", attempt+1).
				Msg("HTTP request succeeded after retries")
		}
		rt.logger.Debug().Int("status", resp.StatusCode).Msg("HTTP Response")
		return resp, nil
	}

	// All retries failed
	if lastResp != nil {
		return nil, errors.Wrapf(lastErr, "HTTP request failed after %d attempts, last status: %d", rt.config.MaxRetries+1, lastResp.StatusCode)
	}
	if lastErr == nil {
		lastErr = errors.New("request failed without response")
	}
	return nil, errors.Wrapf(lastErr, "HTTP request failed after %d attempts", rt.config.MaxRetries+1)
}

// shouldRetry determines whether a request should be retried based on HTTP status code
func shouldRetry(statusCode int) bool {
	// 5xx server errors are typically retryable
	// 4xx client errors are usually not retryable, with some exceptions
	switch {
	case statusCode >= 500:
		return true
	case statusCode == 429: // Too Many Requests
		return true
	case statusCode == 408: // Request Timeout
		return true
	default:
		return false
	}
}

// calculateBackoff calculates exponential backoff delay duration
func calculateBackoff(attempt int, config *RetryConfig) time.Duration {
	delay := time.Duration(
		float64(config.BaseDelay) * math.Pow(config.BackoffRate, float64(attempt-1)),
	)
	if delay > config.MaxDelay {
		delay = config.MaxDelay
	}
	return delay
}
