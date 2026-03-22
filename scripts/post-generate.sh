#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)"
SDK_DIR="$ROOT_DIR/sdks/go"
GEN_DIR="$SDK_DIR/generated"

ARTIFACT_REF_FILE="$GEN_DIR/model_artifact_ref.go"
COMPAT_TEST_FILE="$GEN_DIR/model_artifact_ref_compat_test.go"

if [[ ! -f "$ARTIFACT_REF_FILE" ]]; then
  echo "[error] expected file not found: $ARTIFACT_REF_FILE" >&2
  exit 1
fi

# Keep response decoding tolerant for ArtifactRef.kind while still requiring id.
perl -0pi -e 's/requiredProperties := \[\]string\{\s*"kind",\s*"id",\s*\}/requiredProperties := []string{\n\t\t"id",\n\t}/s' "$ARTIFACT_REF_FILE"

# Replace placeholder imports in generated integration test stubs.
for file in "$GEN_DIR"/test/*.go; do
  [[ -f "$file" ]] || continue
  sed -i 's|github.com/GIT_USER_ID/GIT_REPO_ID|github.com/AbyssForge/sdk-go/generated|g' "$file"
done

cat > "$COMPAT_TEST_FILE" <<'EOF'
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
EOF

if command -v gofmt >/dev/null 2>&1; then
  gofmt -w "$ARTIFACT_REF_FILE" "$COMPAT_TEST_FILE" "$GEN_DIR"/test/*.go
fi

if [[ -f "$SDK_DIR/go.mod" ]] && command -v go >/dev/null 2>&1; then
  (cd "$SDK_DIR" && go mod tidy)
fi

echo "[ok] Go SDK post-generation patch applied"
