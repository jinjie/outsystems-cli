package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveProfileNamePrecedence(t *testing.T) {
	cfg := &Config{DefaultProfile: "cfgdefault", Profiles: map[string]Profile{}}

	// 1. Flag wins over everything.
	t.Setenv(EnvProfile, "envprof")
	if got := cfg.ResolveProfileName("flagprof"); got != "flagprof" {
		t.Errorf("flag precedence: got %q, want flagprof", got)
	}

	// 2. Env wins when no flag.
	if got := cfg.ResolveProfileName(""); got != "envprof" {
		t.Errorf("env precedence: got %q, want envprof", got)
	}

	// 3. Config default when no flag/env.
	t.Setenv(EnvProfile, "")
	if got := cfg.ResolveProfileName(""); got != "cfgdefault" {
		t.Errorf("config default precedence: got %q, want cfgdefault", got)
	}

	// 4. Built-in default when config default empty.
	cfg.DefaultProfile = ""
	if got := cfg.ResolveProfileName(""); got != DefaultProfile {
		t.Errorf("builtin default: got %q, want %q", got, DefaultProfile)
	}
}

func TestProfileValidation(t *testing.T) {
	cfg := &Config{Profiles: map[string]Profile{
		"ok":     {Hostname: "h", APIKey: "k"},
		"nohost": {APIKey: "k"},
		"nokey":  {Hostname: "h"},
	}}

	if _, err := cfg.Profile("ok"); err != nil {
		t.Errorf("valid profile returned error: %v", err)
	}
	if _, err := cfg.Profile("missing"); err == nil {
		t.Error("missing profile should error")
	}
	if _, err := cfg.Profile("nohost"); err == nil {
		t.Error("profile without hostname should error")
	}
	if _, err := cfg.Profile("nokey"); err == nil {
		t.Error("profile without api_key should error")
	}
}

func TestSaveWritesFileWith0600(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	// On some platforms os.UserHomeDir uses other vars; set them too.
	t.Setenv("USERPROFILE", home)

	cfg := &Config{
		DefaultProfile: DefaultProfile,
		Profiles:       map[string]Profile{"default": {Hostname: "h", APIKey: "k"}},
	}
	if err := Save(cfg); err != nil {
		t.Fatalf("Save error: %v", err)
	}

	path := filepath.Join(home, configDirName, configFileName)
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat config: %v", err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Errorf("config perms = %o, want 600", perm)
	}

	// Round-trip: Load should read it back.
	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load error: %v", err)
	}
	if loaded.Profiles["default"].Hostname != "h" {
		t.Errorf("round-trip hostname mismatch: %+v", loaded.Profiles["default"])
	}
}
