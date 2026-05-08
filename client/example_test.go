package client
package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/AbyssForge/sdk-go/client"
	gen "github.com/AbyssForge/sdk-go/generated"
)

// ExampleClient_IngestSignal demonstrates the smallest end-to-end usage of
// the convenience client: build a payload, send it, and inspect the result.
func ExampleClient_IngestSignal() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"status":"accepted","reason":"","canonical_event_ref":{"kind":"signal_event","id":"evt_demo"}}`))
	}))
	defer server.Close()

	c := client.New(server.URL, &http.Client{Timeout: 2 * time.Second})

	result, _, err := c.IngestSignal(context.Background(), gen.RawSignalEventPayload{
		SubjectId:       "acct_demo",
		Producer:        "demo-detector",
		ProducerEventId: gen.PtrString("det-demo-1"),
		SignalType:      "consistent-action-interval",
		OccurredAt:      time.Date(2026, 4, 1, 12, 0, 0, 0, time.UTC),
	})
	if err != nil {
		fmt.Println("ingest error:", err)
		return
	}
	fmt.Println("status:", result.Status)
	// Output: status: accepted
}

// ExampleClient_RecordReviewOutcome demonstrates writing a review outcome
// using the convenience client.
func ExampleClient_RecordReviewOutcome() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"rev_demo","subject_id":"acct_demo","reviewed_object":{"kind":"score_result","id":"score_1"},"reviewer":"reviewer_42","label":"true_positive","reviewed_at":"2026-04-01T12:00:00Z"}`))
	}))
	defer server.Close()

	c := client.New(server.URL, &http.Client{Timeout: 2 * time.Second})

	out, _, err := c.RecordReviewOutcome(context.Background(), "acct_demo", gen.ReviewOutcomeWriteRequest{
		ReviewedObject: gen.ReviewTargetRef{Kind: "score_result", Id: "score_1"},
		Reviewer:       "reviewer_42",
		Label:          "true_positive",
	})
	if err != nil {
		fmt.Println("review error:", err)
		return
	}
	fmt.Println("label:", out.Label)
	// Output: label: true_positive
}
