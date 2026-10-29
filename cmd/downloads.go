package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jinjie/outsystems-cli/internal/client"
	"github.com/jinjie/outsystems-cli/internal/output"
)

var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Download binary files produced by LifeTime",
	Long: `Commands for fetching binary files (application packages, source code
packages, solutions) that the *-download endpoints expose.

The *-download endpoints (e.g. 'applications version-download') return a
DownloadLink JSON object with a "url" and an "expires" timestamp. The "url"
points at a binary file. Use 'downloads get' to fetch that binary, passing
either a download key or the full --url from the DownloadLink.`,
}

var downloadsGetFlags struct {
	url   string
	force bool
}

var downloadsGetCmd = &cobra.Command{
	Use:   "get [download-key]",
	Short: "Download a binary by download key or full URL",
	Long: `Fetches a binary and writes it to a file or stdout.

Provide the source one of two ways:
  - a download key argument: fetched from the LifeTime
    /downloads/{DownloadKey}/ endpoint on the active profile host, with the
    Bearer token attached; or
  - --url <url>: fetched from an absolute URL, typically the "url" field of a
    DownloadLink returned by a *-download command. The Bearer token is sent
    only when the URL host matches the profile host, so pre-signed links to
    other hosts are not sent your LifeTime credential.

Destination (in order of precedence):
  - --output/-o <file> : write to that path
  - --output/-o -      : write to stdout (pipe-friendly, e.g. "-o - > app.oap")
  - otherwise          : use the server-suggested filename in the current
                         directory (from Content-Disposition)
  - if no name is suggested and no --output is given, write to stdout

Existing files are not overwritten unless --force is given.

Examples:
  # Two-step: get the link, then download it.
  LINK=$(outsystems-cli applications version-download <app> <ver> | jq -r .url)
  outsystems-cli downloads get --url "$LINK" -o app.oap

  # One-liner piping the key through jq (when the link is a plain key).
  outsystems-cli downloads get "$KEY" -o app.oap`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		useURL := cmd.Flags().Changed("url")
		if useURL == (len(args) == 1) {
			// Both or neither were provided.
			return fmt.Errorf("provide exactly one of a download-key argument or --url")
		}

		c, err := resolveClient()
		if err != nil {
			return err
		}

		var dl *client.Download
		if useURL {
			dl, err = c.DownloadURL(commandContext(cmd), downloadsGetFlags.url)
		} else {
			dl, err = c.DownloadKey(commandContext(cmd), args[0])
		}
		if err != nil {
			return err
		}
		defer dl.Close()

		res, err := output.StreamBinary(dl.Body, flags.output, dl.Filename, downloadsGetFlags.force)
		if err != nil {
			return err
		}

		// Diagnostics go to stderr so stdout stays pipe-clean.
		if res.Path == "-" {
			fmt.Fprintf(os.Stderr, "Downloaded %s to stdout\n", humanBytes(res.Bytes))
		} else {
			name := dl.Filename
			suffix := ""
			if name != "" && name != res.Path {
				suffix = fmt.Sprintf(" (server name: %s)", name)
			}
			fmt.Fprintf(os.Stderr, "Downloaded %s to %s%s\n", humanBytes(res.Bytes), res.Path, suffix)
		}
		return nil
	},
}

func init() {
	downloadsGetCmd.Flags().StringVar(&downloadsGetFlags.url, "url", "", "absolute URL to download (e.g. the DownloadLink 'url' field)")
	downloadsGetCmd.Flags().BoolVar(&downloadsGetFlags.force, "force", false, "overwrite the destination file if it already exists")
	downloadsCmd.AddCommand(downloadsGetCmd)
	rootCmd.AddCommand(downloadsCmd)
}

// humanBytes renders a byte count in a compact human-readable form.
func humanBytes(n int64) string {
	const unit = 1024
	if n < 0 {
		return "an unknown number of bytes"
	}
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for m := n / unit; m >= unit; m /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(n)/float64(div), "KMGTPE"[exp])
}
