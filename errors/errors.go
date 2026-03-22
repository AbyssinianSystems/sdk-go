// Package sdkerrors provides structured error types for the AbyssForge Go SDK.
//
// The AbyssForge API returns structured rejection bodies on 400, 401, 403, 404,
// and 500 responses. This package unpacks those into Go error values that
// preserve the rejection reason and per-field codes.
//
// # Checking error kind
//
//	result, _, err := c.IngestSignal(ctx, payload)
//	if sdkerrors.IsAuthentication(err) {
//	    // 401: token is missing or invalid
//	}
//	if sdkerrors.IsAuthorization(err) {
//	    // 403: token lacks the required scope
//	}
//
// # Accessing rejection details
//
//	if apiErr, ok := sdkerrors.FromError(err); ok {
//	    for _, r := range apiErr.RejectionReasons {
//	        log.Printf("field=%s code=%s msg=%s", r.GetField(), r.GetCode(), r.GetMessage())
//	    }
//	}
package sdkerrors

import (
	"encoding/json"
	"errors"
	"fmt"

	gen "github.com/AbyssForge/sdk-go/generated"
)

// APIError represents a structured failure returned by the AbyssForge API.
type APIError struct {
	// StatusCode is the HTTP status code (400, 401, 403, 404, 500, …).
	StatusCode int

	// Reason is the machine-readable top-level reason from the response body.
	Reason string

	// RejectionReasons lists per-field rejection details when present.
	RejectionReasons []gen.RejectionReason
}

func (e *APIError) Error() string {
	return fmt.Sprintf("abyssforge: API error %d reason=%s", e.StatusCode, e.Reason)
}

// IsAuthentication reports whether err is a 401 Unauthorized API error,
// meaning the bearer token was missing or could not be verified.
func IsAuthentication(err error) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.StatusCode == 401
}

// IsAuthorization reports whether err is a 403 Forbidden API error,
// meaning the bearer token is valid but lacks the required scope.
func IsAuthorization(err error) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.StatusCode == 403
}

// IsRejected reports whether err is a 400 validation rejection from the API.
func IsRejected(err error) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.StatusCode == 400
}

// IsNotFound reports whether err is a 404 response from the API.
func IsNotFound(err error) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.StatusCode == 404
}

// FromError extracts an [*APIError] from a [gen.GenericOpenAPIError] returned
// by the generated SDK client. Returns (nil, false) when err is not an API
// error (e.g. a network failure).
func FromError(err error) (*APIError, bool) {
	if err == nil {
		return nil, false
	}

	var genErr *gen.GenericOpenAPIError
	if !errors.As(err, &genErr) {
		return nil, false
	}

	apiErr := &APIError{}
	parseRejectionBody(genErr.Body(), apiErr)
	return apiErr, true
}

// FromHTTPResponse parses an [*APIError] from a raw response body and HTTP
// status code. Use this when working directly with [*http.Response] rather
// than the generated client.
func FromHTTPResponse(body []byte, statusCode int) *APIError {
	apiErr := &APIError{StatusCode: statusCode}
	parseRejectionBody(body, apiErr)
	return apiErr
}

type rejectionBody struct {
	Status           string                `json:"status"`
	Reason           string                `json:"reason"`
	RejectionReasons []gen.RejectionReason `json:"rejection_reasons"`
}

func parseRejectionBody(body []byte, dst *APIError) {
	var rb rejectionBody
	if err := json.Unmarshal(body, &rb); err == nil {
		dst.Reason = rb.Reason
		dst.RejectionReasons = rb.RejectionReasons
	}
}
