package cmd

import (
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication operations",
	Long:  "Commands related to authentication tokens.",
}

// --- auth renew ---

var authRenewFlags struct{ data string }

var authRenewCmd = &cobra.Command{
	Use:   "renew",
	Short: "Renew the token of the calling user",
	Long: `Calls POST /auth/tokens/ and returns the renewed token and its expiry.

The request body is a Duration ({"timePeriod": <int>, "unit": "M"}). Provide it
with --data (inline JSON, @file, or - for stdin). The only documented unit is
"M" (months).

Note: this renews the token server-side and returns a new token; it does not
automatically update the stored profile. Save the returned token with
'outsystems-cli init' if you want to persist it.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(authRenewFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/auth/tokens/", nil, body)
		})
	},
}

func init() {
	authRenewCmd.Flags().StringVar(&authRenewFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	authCmd.AddCommand(authRenewCmd)
	rootCmd.AddCommand(authCmd)
}
