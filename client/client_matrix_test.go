package client

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	gen "github.com/AbyssForge/sdk-go/generated"
)

// TestClientHTTPMatrix exercises the convenience client against an
// httptest.Server covering the canonical response shapes a deployed
// AbyssForge service can return: 2xx happy path, 400 validation error, 401
// auth challenge, 429 rate limit, 5xx server error, and a network timeout.
//
// For each case we assert that the typed gen.* error returned by the
// underlying client carries the upstream status so SDK consumers can branch
// on it without parsing the error string.
func TestClientHTTPMatrix(t *testing.T) {
	type wantErr struct {
		expectErr  bool
		statusCode int
	}
	tests := []struct {
		name        string
		handler     http.HandlerFunc
		clientHTTP  *http.Client
		baseURLOnly bool
		want        wantErr
	}{
		{
			name: "happy path 202",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusAccepted)
				_, _ = w.Write([]byte(`{"status":"accepted","reason":"","canonical_event_ref":{"kind":"signal_event","id":"evt_ok"}}`))
			},
			want: wantErr{expectErr: false},
		},
		{
			name: "validation error 400",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"status":"rejected","reason":"validation_failed","rejection_reasons":[{"code":"invalid_field","field":"signal_type","message":"unsupported"}]}`))
			},
			want: wantErr{expectErr: true, statusCode: http.StatusBadRequest},
		},
		{
			name: "auth challenge 401",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("WWW-Authenticate", `Bearer realm="abyssforge"`)
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
			},
			want: wantErr{expectErr: true, statusCode: http.StatusUnauthorized},
		},
		{
			name: "forbidden 403",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"error":"forbidden"}`))
			},
			want: wantErr{expectErr: true, statusCode: http.StatusForbidden},
		},
		{
			name: "rate limited 429",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Retry-After", "30")
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(`{"error":"rate_limited"}`))
			},
			want: wantErr{expectErr: true, statusCode: http.StatusTooManyRequests},
		},
		{
			name: "server error 500",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"error":"internal"}`))
			},
			want: wantErr{expectErr: true, statusCode: http.StatusInternalServerError},
		},
		{
			name: "service unavailable 503",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusServiceUnavailable)
			},
			want: wantErr{expectErr: true, statusCode: http.StatusServiceUnavailable},
		},
		{
			name: "network timeout",
			handler: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(2 * time.Second)
				w.WriteHeader(http.StatusOK)
			},
			clientHTTP: &http.Client{Timeout: 200 * time.Millisecond},
			want:       wantErr{expectErr: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.handler)
			defer server.Close()

			httpClient := tt.clientHTTP
			if httpClient == nil {
				httpClient = shortTimeoutHTTPClient()
			}
			client := New(server.URL, httpClient)

			payload := gen.RawSignalEventPayload{
				SubjectId:       "test_subject_matrix",
				Producer:        "matrix_detector",
				ProducerEventId: gen.PtrString("matrix_event"),
				SignalType:      "test_signal",
				OccurredAt:      time.Now().UTC(),
			}

			ingestResult, response, err := client.IngestSignal(context.Background(), payload)

			if tt.want.expectErr {
				if err == nil {
					t.Fatalf("err = nil, want non-nil")
				}
				if tt.want.statusCode != 0 {
					if response == nil {
						t.Fatalf("response is nil, want HTTP status %d", tt.want.statusCode)
					}
					if response.StatusCode != tt.want.statusCode {
						t.Fatalf("response.StatusCode = %d, want %d", response.StatusCode, tt.want.statusCode)
					}
					var apiErr *gen.GenericOpenAPIError
					if !errors.As(err, &apiErr) {
						t.Fatalf("error type = %T, want *gen.GenericOpenAPIError", err)
					}
				}
				return
			}

			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}
			if ingestResult == nil {
				t.Fatalf("ingestResult is nil; want non-nil")
			}
			if response == nil || response.StatusCode != http.StatusAccepted {
				t.Fatalf("response status = %v, want 202", response)
			}
		})
	}
}

// TestRecordReviewOutcomeMatrix covers the review-outcome endpoint with the
// same response taxonomy. This locks in the SDK's treatment of upstream
// errors for the audit-write path.
func TestRecordReviewOutcomeMatrix(t *testing.T) {
	tests := []struct {
		name       string
		status     int
		body       string
		expectErr  bool
		statusCode int
	}{
		{name: "201 accepted", status: http.StatusCreated, body: `{"id":"rev_1","subject_id":"acct_x","reviewed_object":{"kind":"score_result","id":"score_1"},"reviewer":"rev_42","label":"true_positive","reviewed_at":"2026-04-01T12:00:00Z"}`, expectErr: false},
		{name: "400 validation", status: http.StatusBadRequest, body: `{"error":"invalid_label"}`, expectErr: true, statusCode: http.StatusBadRequest},
		{name: "404 unknown subject", status: http.StatusNotFound, body: `{"error":"subject_not_found"}`, expectErr: true, statusCode: http.StatusNotFound},
		{name: "409 idempotency conflict", status: http.StatusConflict, body: `{"error":"idempotency_conflict"}`, expectErr: true, statusCode: http.StatusConflict},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.body))
			}))
			defer server.Close()

			client := New(server.URL, shortTimeoutHTTPClient())
			_, response, err := client.RecordReviewOutcome(context.Background(), "acct_x", gen.ReviewOutcomeWriteRequest{
				ReviewedObject: gen.ReviewTargetRef{Kind: "score_result", Id: "score_1"},
				Reviewer:       "rev_42",
				Label:          "true_positive",
			})

			if tt.expectErr {
				if err == nil {
					t.Fatalf("err = nil, want non-nil")
				}
				if response == nil || response.StatusCode != tt.statusCode {
					t.Fatalf("status = %v, want %d", response, tt.statusCode)
				}
				return
			}
			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}
		})
	}
}
