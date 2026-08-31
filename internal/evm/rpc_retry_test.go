package evm

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

type sequenceRoundTripper struct {
	attempt int32
	fn      func(attempt int) (*http.Response, error)
}

func (s *sequenceRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	attempt := int(atomic.AddInt32(&s.attempt, 1))
	return s.fn(attempt)
}

type closeCounterBody struct {
	io.Reader
	closed *int32
}

func (b *closeCounterBody) Close() error {
	atomic.AddInt32(b.closed, 1)
	return nil
}

func TestRetryTransport_RetryableStatusReturnsError(t *testing.T) {
	t.Parallel()

	var closeCount int32
	rt := NewRetryTransport(&sequenceRoundTripper{
		fn: func(_ int) (*http.Response, error) {
			return &http.Response{
				StatusCode: 503,
				Body: &closeCounterBody{
					Reader: strings.NewReader("retry"),
					closed: &closeCount,
				},
			}, nil
		},
	}, &RetryConfig{
		MaxRetries:  1,
		BaseDelay:   1,
		MaxDelay:    1,
		BackoffRate: 2,
	})

	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	require.NoError(t, err)

	resp, err := rt.RoundTrip(req)
	require.Error(t, err)
	require.Nil(t, resp)
	require.Equal(t, int32(2), atomic.LoadInt32(&closeCount))
}

func TestRetryTransport_ErrorWithoutResponseReturnsWrappedError(t *testing.T) {
	t.Parallel()

	rt := NewRetryTransport(&sequenceRoundTripper{
		fn: func(_ int) (*http.Response, error) {
			return nil, errors.New("dial failed")
		},
	}, &RetryConfig{
		MaxRetries:  1,
		BaseDelay:   1,
		MaxDelay:    1,
		BackoffRate: 2,
	})

	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	require.NoError(t, err)

	resp, err := rt.RoundTrip(req)
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "HTTP request failed after 2 attempts")
}
