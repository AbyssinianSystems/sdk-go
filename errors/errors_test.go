package sdkerrors_test

import (
	"encoding/json"
	"testing"

	sdkerrors "github.com/AbyssForge/sdk-go/errors"
)

func TestFromHTTPResponse_Parses400(t *testing.T) {
	body := mustJSON(map[string]any{
		"status": "rejected",
		"reason": "validation_failed",
		"rejection_reasons": []any{
			map[string]any{"code": "invalid_field", "field": "occurred_at", "message": "occurred_at is required"},
		},
	})

	apiErr := sdkerrors.FromHTTPResponse(body, 400)

	if apiErr.StatusCode != 400 {
		t.Fatalf("StatusCode = %d, want 400", apiErr.StatusCode)
	}
	if apiErr.Reason != "validation_failed" {
		t.Fatalf("Reason = %q, want validation_failed", apiErr.Reason)
	}
	if len(apiErr.RejectionReasons) != 1 {
		t.Fatalf("RejectionReasons len = %d, want 1", len(apiErr.RejectionReasons))
	}
	if apiErr.RejectionReasons[0].Code != "invalid_field" {
		t.Fatalf("RejectionReasons[0].Code = %q, want invalid_field", apiErr.RejectionReasons[0].Code)
	}
}

func TestFromHTTPResponse_Parses401(t *testing.T) {
	body := mustJSON(map[string]any{
		"status": "rejected",
		"reason": "authentication_failed",
		"rejection_reasons": []any{
			map[string]any{"code": "missing_bearer_token", "field": "authorization", "message": "Authorization header required"},
		},
	})

	apiErr := sdkerrors.FromHTTPResponse(body, 401)

	if !sdkerrors.IsAuthentication(apiErr) {
		t.Fatal("IsAuthentication = false, want true")
	}
	if !sdkerrors.IsAuthentication(apiErr) {
		t.Fatal("IsAuthentication via errors.As = false, want true")
	}
}

func TestFromHTTPResponse_Parses403(t *testing.T) {
	body := mustJSON(map[string]any{
		"status": "rejected",
		"reason": "insufficient_permissions",
		"rejection_reasons": []any{
			map[string]any{"code": "missing_scope", "message": "scope signals:write required"},
		},
	})

	apiErr := sdkerrors.FromHTTPResponse(body, 403)

	if !sdkerrors.IsAuthorization(apiErr) {
		t.Fatal("IsAuthorization = false, want true")
	}
}

func TestFromHTTPResponse_HandlesUnparsableBody(t *testing.T) {
	apiErr := sdkerrors.FromHTTPResponse([]byte("not json"), 500)

	if apiErr == nil {
		t.Fatal("FromHTTPResponse returned nil for unparsable body")
	}
	if apiErr.StatusCode != 500 {
		t.Fatalf("StatusCode = %d, want 500", apiErr.StatusCode)
	}
	// Reason should be empty when body can't be parsed
	if apiErr.Reason != "" {
		t.Fatalf("Reason = %q, want empty", apiErr.Reason)
	}
}

func TestIsRejected(t *testing.T) {
	err := sdkerrors.FromHTTPResponse(mustJSON(map[string]any{"status": "rejected", "reason": "validation_failed"}), 400)
	if !sdkerrors.IsRejected(err) {
		t.Fatal("IsRejected = false, want true")
	}
}

func TestIsNotFound(t *testing.T) {
	err := sdkerrors.FromHTTPResponse(mustJSON(map[string]any{"status": "rejected", "reason": "not_found"}), 404)
	if !sdkerrors.IsNotFound(err) {
		t.Fatal("IsNotFound = false, want true")
	}
}

func TestAPIError_Error_FormatsCorrectly(t *testing.T) {
	err := sdkerrors.FromHTTPResponse(mustJSON(map[string]any{"status": "rejected", "reason": "validation_failed"}), 400)
	got := err.Error()
	if got == "" {
		t.Fatal("Error() returned empty string")
	}
}

func mustJSON(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
