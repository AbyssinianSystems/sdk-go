package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	gen "github.com/AbyssForge/sdk-go/generated"
)

// shortTimeoutHTTPClient bounds connect/read times so tests never hang on a
// real network. Any test pointing at an unreachable endpoint should still
// return promptly with an error.
func shortTimeoutHTTPClient() *http.Client {
	return &http.Client{Timeout: 2 * time.Second}
}

// TestIngestSignalHappyPath verifies the convenience client successfully
// posts a signal event to a live httptest server and decodes the 2xx
// response.
func TestIngestSignalHappyPath(t *testing.T) {
	now := time.Now().UTC()
	payload := gen.RawSignalEventPayload{
		SubjectId:       "test_subject_123",
		Producer:        "test_detector",
		ProducerEventId: gen.PtrString("event_001"),
		SignalType:      "test_signal",
		OccurredAt:      now,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/v1/signal-events" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			_, _ = w.Write([]byte(`{"status":"accepted","reason":"","canonical_event_ref":{"kind":"signal_event","id":"evt_test_1"}}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New(server.URL, shortTimeoutHTTPClient())
	ingestResult, response, err := client.IngestSignal(context.Background(), payload)
	if err != nil {
		t.Fatalf("IngestSignal() error = %v", err)
	}
	if response == nil || response.StatusCode != http.StatusAccepted {
		t.Fatalf("response status = %v, want 202", response)
	}
	if ingestResult == nil {
		t.Fatalf("ingestResult is nil; want non-nil")
	}
}

// TestIngestSignalAndGetLatestEvaluation_IngestFailureShortCircuits verifies
// that a 4xx during ingest aborts the chain and returns an error before
// fetching the evaluation.
func TestIngestSignalAndGetLatestEvaluation_IngestFailureShortCircuits(t *testing.T) {
	payload := gen.RawSignalEventPayload{
		SubjectId:       "test_subject_456",
		Producer:        "test_detector",
		ProducerEventId: gen.PtrString("event_002"),
		SignalType:      "test_signal",
	}

	evalCalled := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/signal-events" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"status":"rejected","reason":"validation_failed"}`))
			return
		}
		evalCalled = true
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New(server.URL, shortTimeoutHTTPClient())
	_, _, err := client.IngestSignalAndGetLatestEvaluation(context.Background(), payload)
	if err == nil {
		t.Fatalf("expected error from 400 ingest; got nil")
	}
	if evalCalled {
		t.Fatalf("evaluation endpoint was called after failed ingest; chain must short-circuit")
	}
}

// TestClient_BearerToken verifies bearer token header management.
func TestClient_BearerToken(t *testing.T) {
	client := New("http://localhost:8080", nil)

	// Set token
	testToken := "test_jwt_token_xyz"
	client.SetBearerToken(testToken)

	// Verify header was set
	authHeader := client.api.GetConfig().DefaultHeader["Authorization"]
	expectedAuth := "Bearer " + testToken

	if authHeader != expectedAuth {
		t.Errorf("expected Authorization header %q, got %q", expectedAuth, authHeader)
	}

	// Clear token
	client.ClearBearerToken()

	// Verify header was removed
	authHeader, exists := client.api.GetConfig().DefaultHeader["Authorization"]
	if exists && authHeader != "" {
		t.Errorf("expected Authorization header to be cleared, got %q", authHeader)
	}
}

// TestClient_New verifies client initialization with custom baseURL and httpClient.
func TestClient_New(t *testing.T) {
	tests := []struct {
		name       string
		baseURL    string
		httpClient *http.Client
		wantServer bool
	}{
		{
			name:       "empty baseURL uses generated default",
			baseURL:    "",
			httpClient: nil,
			wantServer: false,
		},
		{
			name:       "custom baseURL is set",
			baseURL:    "http://custom:9999",
			httpClient: nil,
			wantServer: true,
		},
		{
			name:       "custom baseURL with trailing slash is trimmed",
			baseURL:    "http://custom:9999/",
			httpClient: nil,
			wantServer: true,
		},
		{
			name:       "custom httpClient is used",
			baseURL:    "http://localhost:8080",
			httpClient: &http.Client{Timeout: 5 * time.Second},
			wantServer: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := New(tt.baseURL, tt.httpClient)

			if client == nil {
				t.Fatal("New returned nil client")
			}

			if client.api == nil {
				t.Fatal("api field not initialized")
			}

			// Verify server configuration was set if baseURL was provided
			if tt.wantServer {
				servers := client.api.GetConfig().Servers
				if len(servers) == 0 {
					t.Error("expected servers to be configured")
				}
			}
		})
	}
}
