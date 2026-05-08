package auth_test

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AbyssForge/sdk-go/auth"
)

func TestNewClientCredentialsSource_ValidatesConfig(t *testing.T) {
	_, err := auth.NewClientCredentialsSource(auth.ClientCredentialsConfig{})
	if err == nil {
		t.Fatal("NewClientCredentialsSource() error = nil, want error")
	}
}

func TestClientCredentialsSource_TokenFetchAndCache(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if got := r.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
			t.Fatalf("Content-Type = %q, want application/x-www-form-urlencoded", got)
		}
		wantAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("sdk-client:sdk-secret"))
		if got := r.Header.Get("Authorization"); got != wantAuth {
			t.Fatalf("Authorization = %q, want %q", got, wantAuth)
		}

		if got := r.FormValue("grant_type"); got != "client_credentials" {
			t.Fatalf("grant_type = %q, want client_credentials", got)
		}
		if got := r.FormValue("scope"); got != "signals:write evaluation:read" {
			t.Fatalf("scope = %q, want joined scopes", got)
		}
		if got := r.FormValue("audience"); got != "abyssforge-api" {
			t.Fatalf("audience = %q, want abyssforge-api", got)
		}

		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "token-1",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer srv.Close()

	source, err := auth.NewClientCredentialsSource(auth.ClientCredentialsConfig{
		TokenURL:     srv.URL,
		ClientID:     "sdk-client",
		ClientSecret: "sdk-secret",
		Scopes:       []string{"signals:write", "evaluation:read"},
		Audience:     "abyssforge-api",
		EarlyExpiry:  time.Second,
	})
	if err != nil {
		t.Fatalf("NewClientCredentialsSource() error = %v", err)
	}

	first, err := source.Token()
	if err != nil {
		t.Fatalf("Token() first error = %v", err)
	}
	second, err := source.Token()
	if err != nil {
		t.Fatalf("Token() second error = %v", err)
	}

	if first != "token-1" || second != "token-1" {
		t.Fatalf("tokens = (%q, %q), want token-1", first, second)
	}
	if calls != 1 {
		t.Fatalf("token endpoint calls = %d, want 1", calls)
	}
}

func TestClientCredentialsSource_TokenEndpointError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad credentials", http.StatusUnauthorized)
	}))
	defer srv.Close()

	source, err := auth.NewClientCredentialsSource(auth.ClientCredentialsConfig{
		TokenURL:     srv.URL,
		ClientID:     "sdk-client",
		ClientSecret: "sdk-secret",
	})
	if err != nil {
		t.Fatalf("NewClientCredentialsSource() error = %v", err)
	}

	_, err = source.Token()
	if err == nil {
		t.Fatal("Token() error = nil, want non-nil")
	}
}
