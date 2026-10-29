// Package config handles loading and saving the CLI configuration file,
// which stores one or more named profiles. Each profile holds the LifeTime
// hostname and the API access token used for Bearer authentication.
//
// The configuration file lives at ~/.outsystems-cli/config.yaml on all
// platforms (the home directory is resolved per-OS). Because it contains
// secrets, the file is written with 0600 permissions.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// DefaultProfile is the profile name used when none is specified.
const DefaultProfile = "default"

// EnvProfile is the environment variable that overrides the active profile.
const EnvProfile = "OUTSYSTEMS_CLI_PROFILE"

const (
	configDirName  = ".outsystems-cli"
	configFileName = "config.yaml"
)

// Profile holds the connection settings for a single LifeTime environment.
type Profile struct {
	// Hostname is the LifeTime host, e.g. "lifetime.example.com".
	// HTTPS and the API base path are assumed and added by the client.
	Hostname string `yaml:"hostname"`
	// APIKey is the service account access token used as a Bearer token.
	APIKey string `yaml:"api_key"`
}

// Config is the top-level configuration persisted to disk.
type Config struct {
	// DefaultProfile is the profile used when neither --profile nor the
	// OUTSYSTEMS_CLI_PROFILE environment variable is set.
	DefaultProfile string `yaml:"default_profile"`
	// Profiles maps a profile name to its settings.
	Profiles map[string]Profile `yaml:"profiles"`
}

// Dir returns the directory that holds the configuration file.
func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home directory: %w", err)
	}
	return filepath.Join(home, configDirName), nil
}

// Path returns the full path to the configuration file.
func Path() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}

// Load reads the configuration from disk. If the file does not exist, an
// empty Config (with no profiles) is returned without error so callers can
// treat "not initialized yet" uniformly.
func Load() (*Config, error) {
	path, err := Path()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				DefaultProfile: DefaultProfile,
				Profiles:       map[string]Profile{},
			}, nil
		}
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config %s: %w", path, err)
	}
	if cfg.Profiles == nil {
		cfg.Profiles = map[string]Profile{}
	}
	if cfg.DefaultProfile == "" {
		cfg.DefaultProfile = DefaultProfile
	}
	return &cfg, nil
}

// Save writes the configuration to disk, creating the directory if needed.
// The directory is created with 0700 and the file with 0600 because the
// configuration contains API tokens.
func Save(cfg *Config) error {
	dir, err := Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create config dir %s: %w", dir, err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	path := filepath.Join(dir, configFileName)
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write config %s: %w", path, err)
	}
	return nil
}

// ResolveProfileName determines the active profile name using this precedence:
//  1. the explicit flag value (if non-empty),
//  2. the OUTSYSTEMS_CLI_PROFILE environment variable,
//  3. the config's default profile,
//  4. the built-in DefaultProfile constant.
func (c *Config) ResolveProfileName(flagValue string) string {
	if flagValue != "" {
		return flagValue
	}
	if env := os.Getenv(EnvProfile); env != "" {
		return env
	}
	if c.DefaultProfile != "" {
		return c.DefaultProfile
	}
	return DefaultProfile
}

// Profile returns the named profile, or an error if it does not exist or is
// incomplete. This is the entry point commands use to obtain credentials.
func (c *Config) Profile(name string) (Profile, error) {
	p, ok := c.Profiles[name]
	if !ok {
		return Profile{}, fmt.Errorf("profile %q not found; run 'outsystems-cli init --profile %s' first", name, name)
	}
	if p.Hostname == "" {
		return Profile{}, fmt.Errorf("profile %q has no hostname configured; run 'outsystems-cli init --profile %s'", name, name)
	}
	if p.APIKey == "" {
		return Profile{}, fmt.Errorf("profile %q has no api_key configured; run 'outsystems-cli init --profile %s'", name, name)
	}
	return p, nil
}

// SetProfile inserts or replaces a profile by name.
func (c *Config) SetProfile(name string, p Profile) {
	if c.Profiles == nil {
		c.Profiles = map[string]Profile{}
	}
	c.Profiles[name] = p
}
