package retry_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/AbyssForge/sdk-go/retry"
)

func TestRetryPolicy_SucceedsFirstAttempt(t *testing.T) {
	calls := 0
	policy := &retry.Policy{MaxAttempts: 3, InitialWait: time.Millisecond, MaxWait: time.Millisecond, Multiplier: 2}

	err := policy.Do(context.Background(), func(_ context.Context) (int, error) {
		calls++
		return http.StatusOK, nil
	})

	if err != nil {
		t.Fatalf("Do() error = %v, want nil", err)
	}
	if calls != 1 {
		t.Fatalf("calls = %d, want 1", calls)
	}
}

func TestRetryPolicy_RetriesOnTransient500(t *testing.T) {
	calls := 0
	policy := &retry.Policy{MaxAttempts: 3, InitialWait: time.Millisecond, MaxWait: time.Millisecond, Multiplier: 2}
	sentinelErr := errors.New("internal server error")

	err := policy.Do(context.Background(), func(_ context.Context) (int, error) {
		calls++
		if calls < 3 {
			return http.StatusInternalServerError, sentinelErr
		}
		return http.StatusOK, nil
	})

	if err != nil {
		t.Fatalf("Do() error = %v, want nil", err)
	}
	if calls != 3 {
		t.Fatalf("calls = %d, want 3", calls)
	}
}

func TestRetryPolicy_DoesNotRetryOn400(t *testing.T) {
	calls := 0
	policy := &retry.Policy{MaxAttempts: 3, InitialWait: time.Millisecond, MaxWait: time.Millisecond, Multiplier: 2}
	clientErr := errors.New("bad request")

	err := policy.Do(context.Background(), func(_ context.Context) (int, error) {
		calls++
		return http.StatusBadRequest, clientErr
	})

	if err == nil {
		t.Fatal("Do() error = nil, want non-nil")
	}
	if calls != 1 {
		t.Fatalf("calls = %d, want 1 (no retry on 400)", calls)
	}
}

func TestRetryPolicy_ExhaustsAttemptsAndReturnsLastError(t *testing.T) {
	calls := 0
	policy := &retry.Policy{MaxAttempts: 3, InitialWait: time.Millisecond, MaxWait: time.Millisecond, Multiplier: 2}
	sentinelErr := errors.New("gateway timeout")

	err := policy.Do(context.Background(), func(_ context.Context) (int, error) {
		calls++
		return http.StatusGatewayTimeout, sentinelErr
	})

	if !errors.Is(err, sentinelErr) {
		t.Fatalf("Do() error = %v, want %v", err, sentinelErr)
	}
	if calls != 3 {
		t.Fatalf("calls = %d, want 3", calls)
	}
}

func TestRetryPolicy_RetriesOnRateLimit(t *testing.T) {
	calls := 0
	policy := &retry.Policy{MaxAttempts: 2, InitialWait: time.Millisecond, MaxWait: time.Millisecond, Multiplier: 2}

	policy.Do(context.Background(), func(_ context.Context) (int, error) { //nolint:errcheck
		calls++
		return http.StatusTooManyRequests, errors.New("rate limited")
	})

	if calls != 2 {
		t.Fatalf("calls = %d, want 2 (429 is retried)", calls)
	}
}

func TestRetryPolicy_RespectsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	calls := 0
	policy := &retry.Policy{MaxAttempts: 5, InitialWait: 50 * time.Millisecond, MaxWait: time.Second, Multiplier: 2}

	cancel()

	err := policy.Do(ctx, func(_ context.Context) (int, error) {
		calls++
		return http.StatusInternalServerError, errors.New("error")
	})

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Do() error = %v, want context.Canceled", err)
	}
}

func TestIsTransient(t *testing.T) {
	cases := []struct {
		code      int
		transient bool
	}{
		{0, true},
		{200, false},
		{400, false},
		{401, false},
		{403, false},
		{404, false},
		{429, true},
		{500, true},
		{502, true},
		{503, true},
		{504, true},
	}

	for _, tc := range cases {
		got := retry.IsTransient(tc.code)
		if got != tc.transient {
			t.Errorf("IsTransient(%d) = %v, want %v", tc.code, got, tc.transient)
		}
	}
}
