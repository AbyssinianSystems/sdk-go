package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// ClientCredentialsConfig configures OAuth2 client-credentials token fetching.
type ClientCredentialsConfig struct {
	TokenURL     string
	ClientID     string
	ClientSecret string
	Scopes       []string
	Audience     string
	HTTPClient   *http.Client
	EarlyExpiry  time.Duration
}

// ClientCredentialsSource fetches and caches OAuth2 bearer tokens.
type ClientCredentialsSource struct {
	tokenURL     string
	clientID     string
	clientSecret string
	scopes       []string
	audience     string
	httpClient   *http.Client
	earlyExpiry  time.Duration

	mu      sync.Mutex
	token   string
	expires time.Time
}

// NewClientCredentialsSource builds a [TokenSource] backed by OAuth2
// client-credentials flow.
func NewClientCredentialsSource(cfg ClientCredentialsConfig) (*ClientCredentialsSource, error) {
	if strings.TrimSpace(cfg.TokenURL) == "" {
		return nil, fmt.Errorf("auth: token URL is required")
	}
	if strings.TrimSpace(cfg.ClientID) == "" {
		return nil, fmt.Errorf("auth: client ID is required")
	}
	if strings.TrimSpace(cfg.ClientSecret) == "" {
		return nil, fmt.Errorf("auth: client secret is required")
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	earlyExpiry := cfg.EarlyExpiry
	if earlyExpiry < 0 {
		earlyExpiry = 0
	}
	if earlyExpiry == 0 {
		earlyExpiry = 30 * time.Second
	}

	return &ClientCredentialsSource{
		tokenURL:     strings.TrimSpace(cfg.TokenURL),
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		scopes:       append([]string(nil), cfg.Scopes...),
		audience:     strings.TrimSpace(cfg.Audience),
		httpClient:   httpClient,
		earlyExpiry:  earlyExpiry,
	}, nil
}

// Token returns a cached bearer token, refreshing it on expiry.
func (s *ClientCredentialsSource) Token() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	if s.token != "" && now.Add(s.earlyExpiry).Before(s.expires) {
		return s.token, nil
	}

	token, expires, err := s.fetchToken(now)
	if err != nil {
		return "", err
	}

	s.token = token
	s.expires = expires
	return s.token, nil
}

func (s *ClientCredentialsSource) fetchToken(now time.Time) (string, time.Time, error) {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	if len(s.scopes) > 0 {
		form.Set("scope", strings.Join(s.scopes, " "))
	}
	if s.audience != "" {
		form.Set("audience", s.audience)
	}

	req, err := http.NewRequest(http.MethodPost, s.tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("auth: build token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(s.clientID, s.clientSecret)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("auth: fetch token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", time.Time{}, fmt.Errorf("auth: token endpoint status %d", resp.StatusCode)
	}

	var body struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", time.Time{}, fmt.Errorf("auth: decode token response: %w", err)
	}
	if strings.TrimSpace(body.AccessToken) == "" {
		return "", time.Time{}, fmt.Errorf("auth: token response missing access_token")
	}
	if body.TokenType != "" && !strings.EqualFold(body.TokenType, "bearer") {
		return "", time.Time{}, fmt.Errorf("auth: unsupported token_type %q", body.TokenType)
	}

	expiresIn := body.ExpiresIn
	if expiresIn <= 0 {
		expiresIn = 300
	}
	return body.AccessToken, now.Add(time.Duration(expiresIn) * time.Second), nil
}
