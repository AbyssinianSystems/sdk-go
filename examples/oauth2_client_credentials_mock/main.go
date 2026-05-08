package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/AbyssForge/sdk-go/auth"
	"github.com/AbyssForge/sdk-go/client"
)

func main() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2/token":
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			clientID, clientSecret, ok := r.BasicAuth()
			if !ok || clientID != "sdk-client" || clientSecret != "sdk-secret" {
				http.Error(w, "invalid client credentials", http.StatusUnauthorized)
				return
			}
			if r.Form.Get("grant_type") != "client_credentials" {
				http.Error(w, "unsupported grant type", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "mock-access-token",
				"token_type":   "Bearer",
				"expires_in":   300,
			})
		case "/readyz":
			if got := r.Header.Get("Authorization"); got != "Bearer mock-access-token" {
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"status":"ready","store":"ok"}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	tokenSource, err := auth.NewClientCredentialsSource(auth.ClientCredentialsConfig{
		TokenURL:     server.URL + "/oauth2/token",
		ClientID:     "sdk-client",
		ClientSecret: "sdk-secret",
		Scopes:       []string{"evaluation:read"},
		Audience:     "abyssforge-api",
	})
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{Transport: &auth.Transport{TokenSource: tokenSource}}
	sdk := client.New(server.URL, httpClient)

	ready, _, err := sdk.Readyz(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("readyz status=%s", ready.GetStatus())
}
