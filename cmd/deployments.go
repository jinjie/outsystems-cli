package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

var deploymentsCmd = &cobra.Command{
	Use:     "deployments",
	Aliases: []string{"deps", "dep"},
	Short:   "Manage deployments",
	Long:    "Commands for listing, inspecting, creating, and controlling deployments.",
}

// --- deployments list ---

var depsListFlags struct {
	minDate              string
	maxDate              string
	targetEnvironmentKey string
}

var depsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List deployments (newest first)",
	Long:  "Calls GET /deployments/ and returns deployments ordered by creation date.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("min-date") {
				q.Set("MinDate", depsListFlags.minDate)
			}
			if cmd.Flags().Changed("max-date") {
				q.Set("MaxDate", depsListFlags.maxDate)
			}
			if cmd.Flags().Changed("target-environment-key") {
				q.Set("TargetEnvironmentKey", depsListFlags.targetEnvironmentKey)
			}
			return c.Get(commandContext(cmd), "/deployments/", q)
		})
	},
}

// --- deployments get ---

var depsGetCmd = &cobra.Command{
	Use:   "get <deployment-key>",
	Short: "Get the details of a deployment",
	Long:  "Calls GET /deployments/{DeploymentKey}/, including applications and possible conflicts.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/deployments/"+url.PathEscape(args[0])+"/", nil)
		})
	},
}

// --- deployments status ---

var depsStatusCmd = &cobra.Command{
	Use:   "status <deployment-key>",
	Short: "Get the execution status of a deployment",
	Long:  "Calls GET /deployments/{DeploymentKey}/status/ and returns the status and messages.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/deployments/"+url.PathEscape(args[0])+"/status/", nil)
		})
	},
}

// --- deployments create ---

var depsCreateFlags struct {
	data string
}

var depsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a deployment",
	Long: `Calls POST /deployments/ and returns the key of the new deployment.

The request body is a deployment record (SourceEnvironmentKey,
TargetEnvironmentKey, Notes, ApplicationOperations). Provide it with --data
(inline JSON, @file, or - for stdin).`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(depsCreateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/deployments/", nil, body)
		})
	},
}

// --- deployments update ---

var depsUpdateFlags struct {
	data string
}

var depsUpdateCmd = &cobra.Command{
	Use:   "update <deployment-key>",
	Short: "Update a deployment",
	Long: `Calls PUT /deployments/{DeploymentKey}/ and returns the updated record.

Provide the deployment information to update with --data (inline JSON, @file,
or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(depsUpdateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Put(commandContext(cmd), "/deployments/"+url.PathEscape(args[0])+"/", nil, body)
		})
	},
}

// --- deployments delete ---

var depsDeleteCmd = &cobra.Command{
	Use:   "delete <deployment-key>",
	Short: "Discard a saved deployment",
	Long:  `Calls DELETE /deployments/{DeploymentKey}/. Only deployments in the "saved" state can be deleted.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Delete(commandContext(cmd), "/deployments/"+url.PathEscape(args[0])+"/", nil)
		})
	},
}

// --- deployments exec (start | continue | abort) ---

var depsExecFlags struct {
	redeployOutdated   bool
	continueWithErrors bool
	includeErrorDetail bool
}

var depsExecCmd = &cobra.Command{
	Use:       "exec <deployment-key> <command>",
	Short:     "Execute a deployment command (start, continue, abort)",
	Long:      `Calls POST /deployments/{DeploymentKey}/{Command}/. The command must be one of "start", "continue", or "abort".`,
	Args:      cobra.ExactArgs(2),
	ValidArgs: []string{"start", "continue", "abort"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "RedeployOutdated", depsExecFlags.redeployOutdated, cmd.Flags().Changed("redeploy-outdated"))
			addBoolQuery(q, "ContinueWithErrors", depsExecFlags.continueWithErrors, cmd.Flags().Changed("continue-with-errors"))
			addBoolQuery(q, "IncludeErrorDetails", depsExecFlags.includeErrorDetail, cmd.Flags().Changed("include-error-details"))
			path := "/deployments/" + url.PathEscape(args[0]) + "/" + url.PathEscape(args[1]) + "/"
			return c.Post(commandContext(cmd), path, q, nil)
		})
	},
}

// --- deployments configuration ---

var depsConfigurationFlags struct{ data string }

var depsConfigurationCmd = &cobra.Command{
	Use:   "configuration <deployment-key>",
	Short: "Set the extension database configuration for a deployment",
	Long: `Calls POST /deployments/{DeploymentKey}/configuration.

The request body is an array of ExtensionDBFullMapping. Provide it with --data
(inline JSON, @file, or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(depsConfigurationFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/deployments/"+url.PathEscape(args[0])+"/configuration", nil, body)
		})
	},
}

// --- deployments missing-db-mappings ---

var depsMissingDBMappingsCmd = &cobra.Command{
	Use:   "missing-db-mappings <deployment-key>",
	Short: "List the extension logical database mappings missing for a deployment",
	Long:  "Calls GET /deployments/{DeploymentKey}/getmissingdbmappings.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/deployments/"+url.PathEscape(args[0])+"/getmissingdbmappings", nil)
		})
	},
}

func init() {
	depsListCmd.Flags().StringVar(&depsListFlags.minDate, "min-date", "", "minimum creation date (YYYY-MM-DD)")
	depsListCmd.Flags().StringVar(&depsListFlags.maxDate, "max-date", "", "maximum creation date (YYYY-MM-DD)")
	depsListCmd.Flags().StringVar(&depsListFlags.targetEnvironmentKey, "target-environment-key", "", "filter by target environment key")

	depsCreateCmd.Flags().StringVar(&depsCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	depsUpdateCmd.Flags().StringVar(&depsUpdateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	depsConfigurationCmd.Flags().StringVar(&depsConfigurationFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	depsExecCmd.Flags().BoolVar(&depsExecFlags.redeployOutdated, "redeploy-outdated", false, "redeploy outdated applications in the target environment")
	depsExecCmd.Flags().BoolVar(&depsExecFlags.continueWithErrors, "continue-with-errors", false, "continue the deployment even if errors occur")
	depsExecCmd.Flags().BoolVar(&depsExecFlags.includeErrorDetail, "include-error-details", false, "include more error detail in returned messages")

	deploymentsCmd.AddCommand(
		depsListCmd,
		depsGetCmd,
		depsStatusCmd,
		depsCreateCmd,
		depsUpdateCmd,
		depsDeleteCmd,
		depsExecCmd,
		depsConfigurationCmd,
		depsMissingDBMappingsCmd,
	)
	rootCmd.AddCommand(deploymentsCmd)
}
