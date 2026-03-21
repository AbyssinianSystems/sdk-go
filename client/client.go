package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	gen "github.com/abyssforge/sdk-go/generated"
)

// Client wraps the generated API client with a focused, ergonomic surface.
type Client struct {
	api *gen.APIClient
}

// New creates a convenience SDK client for AbyssForge.
func New(baseURL string, httpClient *http.Client) *Client {
	cfg := gen.NewConfiguration()
	if baseURL != "" {
		cfg.Servers = gen.ServerConfigurations{{URL: strings.TrimRight(baseURL, "/")}}
	}
	if httpClient != nil {
		cfg.HTTPClient = httpClient
	}
	return &Client{api: gen.NewAPIClient(cfg)}
}

// SetBearerToken sets an Authorization header for all future requests.
func (c *Client) SetBearerToken(token string) {
	c.api.GetConfig().AddDefaultHeader("Authorization", "Bearer "+token)
}

// ClearBearerToken removes the default Authorization header.
func (c *Client) ClearBearerToken() {
	delete(c.api.GetConfig().DefaultHeader, "Authorization")
}

// IngestSignal posts a raw signal event payload.
func (c *Client) IngestSignal(ctx context.Context, payload gen.RawSignalEventPayload) (*gen.IngestResult, *http.Response, error) {
	return c.api.SignalEventsAPI.PostSignalEvent(ctx).RawSignalEventPayload(payload).Execute()
}

// GetLatestSubjectEvaluation returns the latest persisted evaluation bundle.
func (c *Client) GetLatestSubjectEvaluation(ctx context.Context, subjectID string) (*gen.SubjectEvaluationBundle, *http.Response, error) {
	return c.api.SubjectEvaluationsAPI.GetLatestSubjectEvaluation(ctx, subjectID).Execute()
}

// IngestSignalAndGetLatestEvaluation ingests and then fetches latest evaluation in one call.
func (c *Client) IngestSignalAndGetLatestEvaluation(ctx context.Context, payload gen.RawSignalEventPayload) (*gen.IngestResult, *gen.SubjectEvaluationBundle, error) {
	ingest, _, err := c.IngestSignal(ctx, payload)
	if err != nil {
		return nil, nil, err
	}
	if ingest == nil {
		return nil, nil, fmt.Errorf("ingest result is nil")
	}
	eval, _, err := c.GetLatestSubjectEvaluation(ctx, payload.SubjectId)
	if err != nil {
		return ingest, nil, err
	}
	return ingest, eval, nil
}

// Healthz checks service liveness.
func (c *Client) Healthz(ctx context.Context) (*gen.HealthStatus, *http.Response, error) {
	return c.api.SystemAPI.GetHealthz(ctx).Execute()
}

// Readyz checks service readiness.
func (c *Client) Readyz(ctx context.Context) (*gen.HealthStatus, *http.Response, error) {
	return c.api.SystemAPI.GetReadyz(ctx).Execute()
}
