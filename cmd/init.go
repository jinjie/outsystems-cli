package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/jinjie/outsystems-cli/internal/config"
)

var initFlags struct {
	hostname string
	apiKey   string
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure a profile with a LifeTime hostname and API key",
	Long: `init creates or updates a profile in ~/.outsystems-cli/config.yaml.

Each profile stores the LifeTime hostname (without scheme or path; HTTPS and
the API base path are assumed) and the API access token used for Bearer
authentication.

Provide values with --hostname and --api-key, or omit them to be prompted
interactively. The API key is read without echo when prompted. Use --profile
to target a profile other than "default".`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		name := cfg.ResolveProfileName(flags.profile)

		hostname := strings.TrimSpace(initFlags.hostname)
		if hostname == "" {
			hostname, err = promptLine(fmt.Sprintf("LifeTime hostname for profile %q (e.g. lifetime.example.com): ", name))
			if err != nil {
				return err
			}
		}
		hostname = normalizeHostname(hostname)
		if hostname == "" {
			return fmt.Errorf("hostname is required")
		}

		apiKey := initFlags.apiKey
		if apiKey == "" {
			apiKey, err = promptSecret(fmt.Sprintf("API key for profile %q: ", name))
			if err != nil {
				return err
			}
		}
		apiKey = strings.TrimSpace(apiKey)
		if apiKey == "" {
			return fmt.Errorf("api key is required")
		}

		cfg.SetProfile(name, config.Profile{Hostname: hostname, APIKey: apiKey})
		if cfg.DefaultProfile == "" {
			cfg.DefaultProfile = config.DefaultProfile
		}
		if err := config.Save(cfg); err != nil {
			return err
		}

		path, _ := config.Path()
		// Diagnostics go to stderr to keep stdout pipe-clean.
		fmt.Fprintf(os.Stderr, "Saved profile %q (hostname %s) to %s\n", name, hostname, path)
		return nil
	},
}

func init() {
	initCmd.Flags().StringVar(&initFlags.hostname, "hostname", "", "LifeTime hostname (without scheme or path)")
	initCmd.Flags().StringVar(&initFlags.apiKey, "api-key", "", "API access token (prompted if omitted)")
	rootCmd.AddCommand(initCmd)
}

// normalizeHostname strips a scheme, trailing slashes, and any path so the
// client can safely prepend "https://" and the API base path.
func normalizeHostname(h string) string {
	h = strings.TrimSpace(h)
	h = strings.TrimPrefix(h, "https://")
	h = strings.TrimPrefix(h, "http://")
	if i := strings.IndexByte(h, '/'); i >= 0 {
		h = h[:i]
	}
	return strings.TrimSpace(h)
}

// promptLine prints a prompt to stderr and reads a line from stdin.
func promptLine(prompt string) (string, error) {
	fmt.Fprint(os.Stderr, prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		return "", fmt.Errorf("read input: %w", err)
	}
	return strings.TrimSpace(line), nil
}

// promptSecret prints a prompt to stderr and reads a secret without echo when
// stdin is a terminal; otherwise it falls back to a normal read (useful for
// scripted input).
func promptSecret(prompt string) (string, error) {
	fmt.Fprint(os.Stderr, prompt)
	fd := int(os.Stdin.Fd())
	if term.IsTerminal(fd) {
		b, err := term.ReadPassword(fd)
		fmt.Fprintln(os.Stderr)
		if err != nil {
			return "", fmt.Errorf("read secret: %w", err)
		}
		return string(b), nil
	}
	return promptLine("")
}
