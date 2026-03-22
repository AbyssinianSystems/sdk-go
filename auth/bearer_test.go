package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AbyssForge/sdk-go/auth"
)

func TestStaticToken_ReturnsConfiguredToken(t *testing.T) {
	src := auth.StaticToken("test-token-value")
	got, err := src.Token()
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	if got != "test-token-value" {
		t.Fatalf("Token() = %q, want %q", got, "test-token-value")
	}
}

func TestNewHTTPClient_InjectsBearerHeader(t *testing.T) {
	var capturedHeader string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeader = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	httpClient := auth.NewHTTPClient("my-service-token")
	resp, err := httpClient.Get(srv.URL)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	resp.Body.Close()

	want := "Bearer my-service-token"
	if capturedHeader != want {
		t.Fatalf("Authorization header = %q, want %q", capturedHeader, want)
	}
}

func TestTransport_DoesNotMutateOriginalRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)

	tr := &auth.Transport{TokenSource: auth.StaticToken("tok")}
	resp, err := tr.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip() error = %v", err)
	}
	resp.Body.Close()

	if got := req.Header.Get("Authorization"); got != "" {
		t.Fatalf("original request was mutated: Authorization = %q", got)
	}
}

func TestTransport_OverridesExistingAuthorizationHeader(t *testing.T) {
	var capturedHeader string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedHeader = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	req.Header.Set("Authorization", "Bearer old-token")

	tr := &auth.Transport{TokenSource: auth.StaticToken("new-token")}
	resp, err := tr.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip() error = %v", err)
	}
	resp.Body.Close()

	want := "Bearer new-token"
	if capturedHeader != want {
		t.Fatalf("Authorization header = %q, want %q", capturedHeader, want)
	}
}
