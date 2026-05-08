# AbyssForge Go SDK

Go client for the [AbyssForge](https://github.com/AbyssForge/AbyssForge) explainable abuse-signal scoring and recommendation engine.

## Requirements

Go 1.24.3 or later.

## Installation

```sh
go get github.com/AbyssForge/sdk-go
```

## Packages

| Import path | Purpose |
|---|---|
| `github.com/AbyssForge/sdk-go/client` | Ergonomic wrapper — start here |
| `github.com/AbyssForge/sdk-go/generated` | Generated OpenAPI types and low-level service clients |

## Quick start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/AbyssForge/sdk-go/client"
    gen "github.com/AbyssForge/sdk-go/generated"
)

func main() {
    c := client.New("https://abyssforge.example.com", nil)
    c.SetBearerToken("your-token")

    result, _, err := c.IngestSignal(context.Background(), gen.RawSignalEventPayload{
        SubjectId:  "user_123",
        Producer:   "login-detector",
        SignalType: "failed_login",
        OccurredAt: time.Now().UTC(),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ingested event %s\n", result.GetEventId())
}
```

## Client API

### Construction

```go
// Use the service default (http://localhost:8080).
c := client.New("", nil)

// Custom base URL and HTTP client.
c := client.New("https://abyssforge.example.com", &http.Client{Timeout: 10 * time.Second})
```

### Authentication

```go
c.SetBearerToken("your-token")

// Remove the header when no longer needed.
c.ClearBearerToken()
```

For machine-to-machine OAuth2, create an auth transport backed by
client-credentials token exchange:

```go
src, err := auth.NewClientCredentialsSource(auth.ClientCredentialsConfig{
    TokenURL:     "https://auth.example.com/oauth2/token",
    ClientID:     os.Getenv("ABYSSFORGE_CLIENT_ID"),
    ClientSecret: os.Getenv("ABYSSFORGE_CLIENT_SECRET"),
    Scopes:       []string{"signals:write", "evaluation:read"},
    Audience:     "abyssforge-api",
})
if err != nil {
    log.Fatal(err)
}

httpClient := &http.Client{Transport: &auth.Transport{TokenSource: src}}
c := client.New("https://abyssforge.example.com", httpClient)
```

### Ingest a signal event

```go
res, httpResp, err := c.IngestSignal(ctx, gen.RawSignalEventPayload{
    SubjectId:       "user_123",
    Producer:        "login-detector",
    ProducerEventId: gen.PtrString("evt-001"),
    SignalType:      "failed_login",
    OccurredAt:      time.Now().UTC(),
})
```

### Fetch the latest subject evaluation

```go
bundle, httpResp, err := c.GetLatestSubjectEvaluation(ctx, "user_123")
```

### Ingest and immediately fetch evaluation

```go
ingestResult, bundle, err := c.IngestSignalAndGetLatestEvaluation(ctx, payload)
```

### Health checks

```go
status, _, err := c.Healthz(ctx)
status, _, err := c.Readyz(ctx)
```

## Low-level generated client

When you need full control over request options or access to endpoints not yet
wrapped in the `client` package, import the generated layer directly:

```go
import gen "github.com/AbyssForge/sdk-go/generated"

cfg := gen.NewConfiguration()
cfg.Servers = gen.ServerConfigurations{{URL: "https://abyssforge.example.com"}}
api := gen.NewAPIClient(cfg)

resp, httpResp, err := api.SubjectEvaluationsAPI.
    ListSubjectEvaluations(ctx, "user_123").
    Execute()
```

## Domain types

All domain types live in `github.com/AbyssForge/sdk-go/generated`:

| Type | Description |
|---|---|
| `RawSignalEventPayload` | Input payload for signal ingestion |
| `IngestResult` | Ingestion acknowledgement |
| `SubjectEvaluationBundle` | Latest evaluation snapshot for a subject |
| `ScoreResult` | Aggregate score and per-signal contributions |
| `Recommendation` | Advisory recommendation derived from score |
| `ReviewOutcome` | Human review decision |
| `OutcomeAnalysis` | Bulk analysis of past review outcomes |
| `HealthStatus` | Service liveness / readiness status |

## Running the tests

```sh
go test ./...
```
