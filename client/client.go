package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/AbyssForge/sdk-go/auth"
	gen "github.com/AbyssForge/sdk-go/generated"
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

// BEGIN generated:client-env-factory
func NewFromEnv() (*Client, error) {
	baseURL := os.Getenv("ABYSSFORGE_BASE_URL")
	token := os.Getenv("ABYSSFORGE_TOKEN")
	if baseURL == "" {
		return nil, errors.New("ABYSSFORGE_BASE_URL is not set")
	}
	if token == "" {
		return nil, errors.New("ABYSSFORGE_TOKEN is not set")
	}
	return New(baseURL, auth.NewHTTPClient(token)), nil
}

// END generated:client-env-factory

// BEGIN generated:openapi-wrapper-methods
// CompareSubjectRuleset compares baseline and candidate ruleset outputs for a subject.
func (c *Client) CompareSubjectRuleset(ctx context.Context, subjectID string, request gen.RulesetComparisonRequest) (*gen.RulesetComparison, *http.Response, error) {
	return c.api.SubjectEvaluationsAPI.CompareSubjectRuleset(ctx, subjectID).RulesetComparisonRequest(request).Execute()
}

// Healthz checks service liveness.
func (c *Client) Healthz(ctx context.Context) (*gen.HealthStatus, *http.Response, error) {
	return c.api.SystemAPI.GetHealthz(ctx).Execute()
}

// GetLatestSubjectEvaluation returns the latest persisted evaluation bundle.
func (c *Client) GetLatestSubjectEvaluation(ctx context.Context, subjectID string) (*gen.SubjectEvaluationBundle, *http.Response, error) {
	return c.api.SubjectEvaluationsAPI.GetLatestSubjectEvaluation(ctx, subjectID).Execute()
}

// Livez checks process liveness.
func (c *Client) Livez(ctx context.Context) (*gen.HealthStatus, *http.Response, error) {
	return c.api.SystemAPI.GetLivez(ctx).Execute()
}

// GetOutcomeAnalysis returns deterministic recommendation outcome analysis.
func (c *Client) GetOutcomeAnalysis(ctx context.Context, subjectID, recommendationType, reviewLabel *string) (*gen.OutcomeAnalysis, *http.Response, error) {
	req := c.api.SubjectEvaluationsAPI.GetOutcomeAnalysis(ctx)
	if subjectID != nil {
		req = req.SubjectId(*subjectID)
	}
	if recommendationType != nil {
		req = req.RecommendationType(*recommendationType)
	}
	if reviewLabel != nil {
		req = req.ReviewLabel(*reviewLabel)
	}
	return req.Execute()
}

// Readyz checks service readiness.
func (c *Client) Readyz(ctx context.Context) (*gen.HealthStatus, *http.Response, error) {
	return c.api.SystemAPI.GetReadyz(ctx).Execute()
}

// GetSubjectInvestigation returns the investigation read model for a subject.
func (c *Client) GetSubjectInvestigation(ctx context.Context, subjectID string) (*gen.SubjectInvestigation, *http.Response, error) {
	return c.api.SubjectEvaluationsAPI.GetSubjectInvestigation(ctx, subjectID).Execute()
}

// ListSubjectEvaluations lists persisted evaluations for a subject.
func (c *Client) ListSubjectEvaluations(ctx context.Context, subjectID string) (*gen.SubjectEvaluationHistory, *http.Response, error) {
	return c.api.SubjectEvaluationsAPI.ListSubjectEvaluations(ctx, subjectID).Execute()
}

// ListSubjectSignalEvents lists canonical signal events accepted for a subject.
func (c *Client) ListSubjectSignalEvents(ctx context.Context, subjectID string) (*gen.SubjectSignalEventList, *http.Response, error) {
	return c.api.SignalEventsAPI.ListSubjectSignalEvents(ctx, subjectID).Execute()
}

// IngestSignal posts a raw signal event payload.
func (c *Client) IngestSignal(ctx context.Context, payload gen.RawSignalEventPayload) (*gen.IngestResult, *http.Response, error) {
	return c.api.SignalEventsAPI.PostSignalEvent(ctx).RawSignalEventPayload(payload).Execute()
}

// RecordReviewOutcome persists a review outcome for a subject artifact.
func (c *Client) RecordReviewOutcome(ctx context.Context, subjectID string, request gen.ReviewOutcomeWriteRequest) (*gen.ReviewOutcome, *http.Response, error) {
	return c.api.SubjectEvaluationsAPI.PostSubjectReviewOutcome(ctx, subjectID).ReviewOutcomeWriteRequest(request).Execute()
}

// RecomputeSubject recomputes and persists a subject evaluation.
func (c *Client) RecomputeSubject(ctx context.Context, subjectID string, request *gen.SubjectRecomputeRequest) (*gen.SubjectEvaluationBundle, *http.Response, error) {
	req := c.api.SubjectEvaluationsAPI.RecomputeSubject(ctx, subjectID)
	if request != nil {
		req = req.SubjectRecomputeRequest(*request)
	}
	return req.Execute()
}
// END generated:openapi-wrapper-methods

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
