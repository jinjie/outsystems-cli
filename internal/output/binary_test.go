package output

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStreamBinaryToExplicitFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.bin")
	res, err := StreamBinary(strings.NewReader("hello"), path, "ignored.oap", false)
	if err != nil {
		t.Fatalf("StreamBinary error: %v", err)
	}
	if res.Path != path {
		t.Errorf("path = %q, want %q", res.Path, path)
	}
	if res.Bytes != 5 {
		t.Errorf("bytes = %d, want 5", res.Bytes)
	}
	data, _ := os.ReadFile(path)
	if string(data) != "hello" {
		t.Errorf("file content = %q", data)
	}
}

func TestStreamBinaryUsesServerFilename(t *testing.T) {
	dir := t.TempDir()
	// Run from the temp dir so the suggested filename lands there.
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	res, err := StreamBinary(strings.NewReader("abc"), "", "MyApp.oap", false)
	if err != nil {
		t.Fatalf("StreamBinary error: %v", err)
	}
	if res.Path != "MyApp.oap" {
		t.Errorf("path = %q, want MyApp.oap", res.Path)
	}
	if _, err := os.Stat(filepath.Join(dir, "MyApp.oap")); err != nil {
		t.Errorf("expected file written: %v", err)
	}
}

func TestStreamBinaryRefusesOverwriteWithoutForce(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "exists.bin")
	if err := os.WriteFile(path, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := StreamBinary(strings.NewReader("new"), path, "", false)
	if err == nil {
		t.Fatal("expected error when file exists without --force")
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("error = %v", err)
	}
	// Original content preserved.
	data, _ := os.ReadFile(path)
	if string(data) != "old" {
		t.Errorf("file should be unchanged, got %q", data)
	}

	// With force it overwrites.
	if _, err := StreamBinary(strings.NewReader("new"), path, "", true); err != nil {
		t.Fatalf("force overwrite error: %v", err)
	}
	data, _ = os.ReadFile(path)
	if string(data) != "new" {
		t.Errorf("force should overwrite, got %q", data)
	}
}
