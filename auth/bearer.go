// Package auth provides bearer token injection for the AbyssForge Go SDK.
//
// For most consumers, [NewHTTPClient] is the only function needed:
//
//	httpClient := auth.NewHTTPClient(os.Getenv("ABYSSFORGE_TOKEN"))
//	c := client.New("https://abyssforge.example.com", httpClient)
//
// Advanced consumers who need dynamic tokens (e.g. automatic refresh) can
// implement [TokenSource] and construct a [Transport] directly.
package auth

import (
	"fmt"
	"net/http"
)

// TokenSource supplies bearer tokens for outgoing requests.
type TokenSource interface {
	Token() (string, error)
}

// staticToken holds a fixed bearer token.
type staticToken struct{ value string }

func (s *staticToken) Token() (string, error) { return s.value, nil }

// StaticToken returns a [TokenSource] that always yields the same token.
// Use this when the token is long-lived or externally refreshed.
func StaticToken(token string) TokenSource {
	return &staticToken{value: token}
}

// Transport is an [http.RoundTripper] that injects an Authorization Bearer
// header into every request using the provided [TokenSource].
//
// Set Base to nil to fall back to [http.DefaultTransport].
type Transport struct {
	TokenSource TokenSource
	Base        http.RoundTripper
}

// RoundTrip implements [http.RoundTripper]. It clones the request, adds the
// Authorization header, and delegates to the Base transport.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.TokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("auth: token source error: %w", err)
	}

	// Clone before mutation so the caller's Request is never modified.
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+token)

	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(clone)
}

// NewHTTPClient returns an [*http.Client] pre-configured to inject the given
// static bearer token into every request.
func NewHTTPClient(token string) *http.Client {
	return &http.Client{
		Transport: &Transport{TokenSource: StaticToken(token)},
	}
}
