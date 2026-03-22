// Package main demonstrates ingesting a signal event with the AbyssForge Go SDK,
// using the auth and retry packages for bearer token injection and resilience.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbyssForge/sdk-go/auth"
	"github.com/AbyssForge/sdk-go/client"
	sdkerrors "github.com/AbyssForge/sdk-go/errors"
	gen "github.com/AbyssForge/sdk-go/generated"
	"github.com/AbyssForge/sdk-go/retry"
)

func main() {
	token := os.Getenv("ABYSSFORGE_TOKEN")
	if token == "" {
		log.Fatal("ABYSSFORGE_TOKEN env var required")
	}

	baseURL := os.Getenv("ABYSSFORGE_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// auth.NewHTTPClient creates an *http.Client that automatically injects the
	// bearer token into every request — no manual header management needed.
	c := client.New(baseURL, auth.NewHTTPClient(token))

	payload := gen.RawSignalEventPayload{
		SubjectId:       "user_42",
		Producer:        "login-detector",
		ProducerEventId: gen.PtrString("evt-001"),
		SignalType:      "failed_login",
		OccurredAt:      time.Now().UTC(),
		Severity:        gen.PtrString("high"),
		Confidence:      gen.PtrFloat32(0.95),
	}

	// retry.Default retries on 429 and 5xx up to 3 times with exponential backoff.
	var result *gen.IngestResult
	err := retry.Default.Do(context.Background(), func(ctx context.Context) (int, error) {
		var resp *http.Response
		var callErr error
		result, resp, callErr = c.IngestSignal(ctx, payload)
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}
		return statusCode, callErr
	})

	if err != nil {
		if sdkerrors.IsAuthentication(err) {
			log.Fatal("authentication failed: check ABYSSFORGE_TOKEN")
		}
		if sdkerrors.IsAuthorization(err) {
			log.Fatal("authorization failed: token lacks signals:write scope")
		}
		if apiErr, ok := sdkerrors.FromError(err); ok {
			for _, r := range apiErr.RejectionReasons {
				log.Printf("rejection: field=%s code=%s msg=%s", r.GetField(), r.GetCode(), r.GetMessage())
			}
		}
		log.Fatalf("IngestSignal: %v", err)
	}

	log.Printf("status=%s reason=%s", result.GetStatus(), result.GetReason())
}
