// Package abyssforge is the Go SDK for AbyssForge — an explainable
// abuse-signal scoring and recommendation engine.
//
// # Module layout
//
// The module is organised into focused sub-packages. Most callers only need
// the [client] package.
//
//   - [github.com/AbyssForge/sdk-go/client] — ergonomic wrapper around the
//     generated HTTP client; start here.
//   - [github.com/AbyssForge/sdk-go/generated] — generated OpenAPI types and
//     low-level API service structs; import directly when you need domain types
//     such as [generated.RawSignalEventPayload] or [generated.ScoreResult].
//
// # Quick start
//
//	import (
//	    "context"
//	    "log"
//
//	    "github.com/AbyssForge/sdk-go/client"
//	    gen "github.com/AbyssForge/sdk-go/generated"
//	)
//
//	func main() {
//	    c := client.New("https://abyssforge.example.com", nil)
//	    c.SetBearerToken("your-token")
//
//	    result, _, err := c.IngestSignal(context.Background(), gen.RawSignalEventPayload{
//	        SubjectId:  "user_123",
//	        Producer:   "login-detector",
//	        SignalType: "failed_login",
//	        OccurredAt: time.Now().UTC(),
//	    })
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    log.Printf("ingested event %s", result.GetEventId())
//	}
package abyssforge
