package client

import (
	"encoding/json"
	"testing"
)

func TestArtifactRef_UnmarshalJSON_AllowsMissingKind(t *testing.T) {
	var ref ArtifactRef
	if err := json.Unmarshal([]byte(`{"id":"evt_123"}`), &ref); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if ref.Id != "evt_123" {
		t.Fatalf("id = %q, want evt_123", ref.Id)
	}
	if ref.Kind != "" {
		t.Fatalf("kind = %q, want empty", ref.Kind)
	}
}

func TestArtifactRef_UnmarshalJSON_RequiresID(t *testing.T) {
	var ref ArtifactRef
	err := json.Unmarshal([]byte(`{"kind":"signal_event"}`), &ref)
	if err == nil {
		t.Fatal("expected error when id is missing")
	}
}
