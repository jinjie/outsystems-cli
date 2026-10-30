// Package cmd defines the cobra command tree for the OutSystems LifeTime v2 CLI.
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jinjie/outsystems-cli/internal/client"
	"github.com/jinjie/outsystems-cli/internal/config"
	"github.com/jinjie/outsystems-cli/internal/output"
)

// globalFlags holds the values of the persistent flags shared by all commands.
type globalFlags struct {
	profile string
	format  string
	output  string
}

var flags globalFlags

// Build information, injected at build time via -ldflags by GoReleaser (or a
// plain `go build`). Defaults are used for `go run` / source builds.
var (
	buildVersion = "dev"
	buildCommit  = "none"
	buildDate    = "unknown"
)

// SetVersionInfo wires the build metadata into the root command's --version
// output. It is called from main() with values set via -ldflags.
func SetVersionInfo(version, commit, date string) {
	if version != "" {
		buildVersion = version
	}
	if commit != "" {
		buildCommit = commit
	}
	if date != "" {
		buildDate = date
	}
	rootCmd.Version = buildVersion
	rootCmd.SetVersionTemplate(
		"outsystems-cli {{.Version}} (commit " + buildCommit + ", built " + buildDate + ")\n",
	)
}

// rootCmd is the base command invoked as "outsystems-cli".
var rootCmd = &cobra.Command{
	Use:     "outsystems-cli",
	Short:   "CLI for the OutSystems LifeTime API v2",
	Version: buildVersion,
	Long: `outsystems-cli interacts with the OutSystems LifeTime REST API v2.

Configure a profile with 'outsystems-cli init', then run resource commands
such as 'applications list' or 'environments list'. Output is JSON by default
and written to stdout so it can be piped into tools like jq. Use --output to
write the result to a file instead.

Note: the 'version' subcommand reports the remote LifeTime server version;
use '--version' to see this CLI's own build version.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command and maps errors to a non-zero exit code.
// All diagnostics are written to stderr so stdout carries only payloads.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// RootCmd returns the root cobra command. It is exported so tooling (such as
// the Markdown docs generator in tools/gendocs) can walk the command tree.
func RootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	pf := rootCmd.PersistentFlags()
	pf.StringVar(&flags.profile, "profile", "",
		"profile to use (overrides $"+config.EnvProfile+" and the config default)")
	pf.StringVar(&flags.format, "format", string(output.FormatJSON),
		"output format (only 'json' is supported)")
	pf.StringVarP(&flags.output, "output", "o", "",
		"write output to a file instead of stdout")
}

// newWriter builds an output.Writer from the global flags, validating format.
func newWriter() (*output.Writer, error) {
	format, err := output.ParseFormat(flags.format)
	if err != nil {
		return nil, err
	}
	return output.NewWriter(format, flags.output), nil
}

// resolveClient loads config, resolves the active profile (flag > env >
// default), and returns a ready API client.
func resolveClient() (*client.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	name := cfg.ResolveProfileName(flags.profile)
	profile, err := cfg.Profile(name)
	if err != nil {
		return nil, err
	}
	return client.New(profile), nil
}

// commandContext returns the context used for API calls. It is tied to the
// cobra command so cancellation (e.g. Ctrl-C) propagates.
func commandContext(cmd *cobra.Command) context.Context {
	if ctx := cmd.Context(); ctx != nil {
		return ctx
	}
	return context.Background()
}
