// Package retry provides exponential back-off retry logic for the AbyssForge Go SDK.
//
// Only transient failures (HTTP 429, 5xx, and network errors) are retried.
// Client errors (4xx) are returned immediately without retry.
//
// # Basic usage
//
//	var result *gen.IngestResult
//	err := retry.Default.Do(ctx, func(ctx context.Context) (int, error) {
//	    var resp *http.Response
//	    result, resp, err = c.IngestSignal(ctx, payload)
//	    if resp != nil {
//	        return resp.StatusCode, err
//	    }
//	    return 0, err
//	})
package retry

import (
	"context"
	"math/rand/v2"
	"net/http"
	"time"
)

// Policy controls retry behaviour.
type Policy struct {
	// MaxAttempts is the total number of attempts (including the first).
	MaxAttempts int

	// InitialWait is the base delay before the second attempt.
	InitialWait time.Duration

	// MaxWait caps the delay between any two attempts.
	MaxWait time.Duration

	// Multiplier is applied to the wait duration after each failure.
	Multiplier float64
}

// Default is the recommended retry policy for transient failures: 3 attempts
// with exponential backoff starting at 100 ms and capped at 5 s.
// BEGIN generated:retry-defaults
var Default = &Policy{
	MaxAttempts: 3,
	InitialWait: 100 * time.Millisecond,
	MaxWait:     5000 * time.Millisecond,
	Multiplier:  2,
}

// END generated:retry-defaults

// Do executes fn up to p.MaxAttempts times, retrying when fn returns a
// transient HTTP status code (429, 5xx) or a network-level error (statusCode 0).
//
// The context is forwarded to fn on every attempt; a cancelled context causes
// immediate return.
func (p *Policy) Do(ctx context.Context, fn func(ctx context.Context) (statusCode int, err error)) error {
	wait := p.InitialWait
	var lastErr error

	for attempt := 0; attempt < p.MaxAttempts; attempt++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		statusCode, err := fn(ctx)
		if err == nil {
			return nil
		}

		lastErr = err

		if !IsTransient(statusCode) {
			return err
		}

		if attempt == p.MaxAttempts-1 {
			break
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(jitter(wait)):
		}

		wait = p.nextWait(wait)
	}

	return lastErr
}

// IsTransient reports whether an HTTP status code indicates a retryable error.
// Status 0 is treated as a network-level failure (also retryable).
func IsTransient(statusCode int) bool {
	return statusCode == 0 ||
		statusCode == http.StatusTooManyRequests ||
		statusCode >= http.StatusInternalServerError
}

// nextWait returns the next back-off duration capped at p.MaxWait.
func (p *Policy) nextWait(current time.Duration) time.Duration {
	next := time.Duration(float64(current) * p.Multiplier)
	if next > p.MaxWait {
		return p.MaxWait
	}
	return next
}

// jitter adds ±10 % random variation to d to avoid thundering-herd patterns
// when multiple clients all back off at the same time.
func jitter(d time.Duration) time.Duration {
	spread := float64(d) * 0.1
	offset := (rand.Float64()*2 - 1) * spread // [-spread, +spread]
	return d + time.Duration(offset)
}
