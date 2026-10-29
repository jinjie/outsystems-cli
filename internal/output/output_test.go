package output

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteJSONReindents(t *testing.T) {
	var buf bytes.Buffer
	w := &Writer{Format: FormatJSON, Stdout: &buf}
	if err := w.WriteJSON([]byte(`{"a":1,"b":[2,3]}`)); err != nil {
		t.Fatalf("WriteJSON error: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "\n  \"a\": 1") {
		t.Errorf("expected reindented JSON, got:\n%s", got)
	}
	if !strings.HasSuffix(got, "\n") {
		t.Errorf("expected trailing newline")
	}
}

func TestWriteJSONPassthroughForNonJSON(t *testing.T) {
	var buf bytes.Buffer
	w := &Writer{Format: FormatJSON, Stdout: &buf}
	// Some endpoints return a bare key as plain text.
	if err := w.WriteJSON([]byte("some-plain-key")); err != nil {
		t.Fatalf("WriteJSON error: %v", err)
	}
	if got := buf.String(); got != "some-plain-key\n" {
		t.Errorf("passthrough mismatch: %q", got)
	}
}

func TestWriteToFileKeepsStdoutQuiet(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.json")
	var buf bytes.Buffer
	w := &Writer{Format: FormatJSON, File: path, Stdout: &buf}

	if err := w.WriteJSON([]byte(`{"x":1}`)); err != nil {
		t.Fatalf("WriteJSON error: %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("stdout should be quiet when --output set, got: %q", buf.String())
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read output file: %v", err)
	}
	if !strings.Contains(string(data), `"x": 1`) {
		t.Errorf("file content unexpected: %s", data)
	}
}

func TestParseFormat(t *testing.T) {
	if _, err := ParseFormat("json"); err != nil {
		t.Errorf("json should be valid: %v", err)
	}
	if _, err := ParseFormat("yaml"); err == nil {
		t.Error("yaml should be rejected")
	}
}
