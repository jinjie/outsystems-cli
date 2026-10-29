package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

// --- environments module-sourcecode (request) ---

var envsModSourceCodeCmd = &cobra.Command{
	Use:   "module-sourcecode <environment-key> <module-key>",
	Short: "Request the source code package of a module",
	Long: `Calls POST /environments/{EnvironmentKey}/modules/{ModuleKey}/sourcecodeaccess
and returns the operation key (PackageKey) and status. Poll status with
'module-sourcecode-status' and fetch the link with 'module-sourcecode-download'.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/modules/" + url.PathEscape(args[1]) + "/sourcecodeaccess"
			return c.Post(commandContext(cmd), path, nil, nil)
		})
	},
}

// --- environments module-sourcecode-status ---

var envsModSourceCodeStatusCmd = &cobra.Command{
	Use:   "module-sourcecode-status <environment-key> <module-key> <package-key>",
	Short: "Get the status of a module source code package",
	Long:  "Calls GET /environments/{EnvironmentKey}/modules/{ModuleKey}/sourcecodeaccess/{PackageKey}/status.",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/modules/" + url.PathEscape(args[1]) +
				"/sourcecodeaccess/" + url.PathEscape(args[2]) + "/status"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- environments module-sourcecode-download ---

var envsModSourceCodeDownloadCmd = &cobra.Command{
	Use:   "module-sourcecode-download <environment-key> <module-key> <package-key>",
	Short: "Get a download link for a module source code package",
	Long: `Calls GET /environments/{EnvironmentKey}/modules/{ModuleKey}/sourcecodeaccess/{PackageKey}/download
and returns a link that expires in 60 minutes.`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/modules/" + url.PathEscape(args[1]) +
				"/sourcecodeaccess/" + url.PathEscape(args[2]) + "/download"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- environments csp-get ---

var envsCSPGetCmd = &cobra.Command{
	Use:   "csp-get <environment-key>",
	Short: "Get the Content Security Policies of an environment",
	Long:  "Calls GET /environments/{EnvironmentKey}/content_security_policies.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/content_security_policies", nil)
		})
	},
}

// --- environments csp-set ---

var envsCSPSetFlags struct{ data string }

var envsCSPSetCmd = &cobra.Command{
	Use:   "csp-set <environment-key>",
	Short: "Set the Content Security Policies of an environment",
	Long: `Calls PUT /environments/{EnvironmentKey}/content_security_policies.

Provide a CSPDirectives body with --data (inline JSON, @file, or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(envsCSPSetFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Put(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/content_security_policies", nil, body)
		})
	},
}

// --- environments solution-create ---

var envsSolutionCreateFlags struct{ data string }

var envsSolutionCreateCmd = &cobra.Command{
	Use:   "solution-create <environment-key>",
	Short: "Create a solution in an environment",
	Long: `Calls POST /environments/{EnvironmentKey}/solution and returns the solution
key.

Provide a SolutionCreate body with --data (inline JSON, @file, or - for stdin).
Required field: ApplicationKeys.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(envsSolutionCreateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/solution", nil, body)
		})
	},
}

// --- environments solution-download ---

var envsSolutionDownloadCmd = &cobra.Command{
	Use:   "solution-download <environment-key> <solution-key>",
	Short: "Get a download link for a solution",
	Long:  "Calls GET /environments/{EnvironmentKey}/solution/{SolutionKey}.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/solution/" + url.PathEscape(args[1])
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- environments solution-status ---

var envsSolutionStatusCmd = &cobra.Command{
	Use:   "solution-status <environment-key> <solution-key>",
	Short: "Get the status of a solution package",
	Long:  "Calls GET /environments/{EnvironmentKey}/solutionstatus/{SolutionKey}.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/solutionstatus/" + url.PathEscape(args[1])
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- environments deployment-create-with-binary ---

var envsDeploymentBinaryFlags struct{ data string }

var envsDeploymentBinaryCmd = &cobra.Command{
	Use:   "deployment-create-with-binary <environment-key>",
	Short: "Create a deployment from a staging binary",
	Long: `Calls POST /environments/{EnvironmentKey}/deployment and returns the key of
the new deployment.

The body is the staging binary payload. Provide it with --data (inline string,
@file, or - for stdin). Use @file to point at the binary content.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(envsDeploymentBinaryFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/deployment", nil, body)
		})
	},
}

// --- environments blocked-ips ---

var envsBlockedIPsFlags struct {
	environmentKey string
}

var envsBlockedIPsCmd = &cobra.Command{
	Use:   "blocked-ips",
	Short: "List IP addresses with blocked login attempts",
	Long: `Calls GET /environments/blockedipaddresses/.

Use --environment-key to scope the query to a single environment; if omitted,
all environments are returned.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("environment-key") {
				q.Set("EnvironmentKey", envsBlockedIPsFlags.environmentKey)
			}
			return c.Get(commandContext(cmd), "/environments/blockedipaddresses/", q)
		})
	},
}

// --- environments unblock-ip ---

var envsUnblockIPFlags struct {
	environmentKey string
}

var envsUnblockIPCmd = &cobra.Command{
	Use:   "unblock-ip <ip-address>",
	Short: "Unblock an IP address",
	Long: `Calls DELETE /environments/blockedipaddresses/{ipaddress}/.

Use --environment-key to scope the unblock to a single environment; if omitted,
the IP is unblocked in all environments.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("environment-key") {
				q.Set("EnvironmentKey", envsUnblockIPFlags.environmentKey)
			}
			return c.Delete(commandContext(cmd), "/environments/blockedipaddresses/"+url.PathEscape(args[0])+"/", q)
		})
	},
}

func init() {
	envsCSPSetCmd.Flags().StringVar(&envsCSPSetFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	envsSolutionCreateCmd.Flags().StringVar(&envsSolutionCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	envsDeploymentBinaryCmd.Flags().StringVar(&envsDeploymentBinaryFlags.data, "data", "", "request body: inline string, @file, or - for stdin")

	envsBlockedIPsCmd.Flags().StringVar(&envsBlockedIPsFlags.environmentKey, "environment-key", "", "scope to a single environment (all if empty)")
	envsUnblockIPCmd.Flags().StringVar(&envsUnblockIPFlags.environmentKey, "environment-key", "", "scope to a single environment (all if empty)")

	environmentsCmd.AddCommand(
		envsModSourceCodeCmd,
		envsModSourceCodeStatusCmd,
		envsModSourceCodeDownloadCmd,
		envsCSPGetCmd,
		envsCSPSetCmd,
		envsSolutionCreateCmd,
		envsSolutionDownloadCmd,
		envsSolutionStatusCmd,
		envsDeploymentBinaryCmd,
		envsBlockedIPsCmd,
		envsUnblockIPCmd,
	)
}
