package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

var environmentsCmd = &cobra.Command{
	Use:     "environments",
	Aliases: []string{"envs", "env"},
	Short:   "Manage and inspect environments",
	Long:    "Commands for listing and inspecting environments and their running applications.",
}

// --- environments list ---

var envsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments in the infrastructure",
	Long:  "Calls GET /environments/ and returns all environments.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/", nil)
		})
	},
}

// --- environments get ---

var envsGetCmd = &cobra.Command{
	Use:   "get <environment-key>",
	Short: "Get the details of an environment",
	Long:  "Calls GET /environments/{EnvironmentKey}/ and returns the environment details.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/", nil)
		})
	},
}

// --- environments applications (running apps) ---

var envsAppsFlags struct {
	includeModules   bool
	includeEnvStatus bool
}

var envsAppsCmd = &cobra.Command{
	Use:   "applications <environment-key>",
	Short: "List running applications in an environment",
	Long:  "Calls GET /environments/{EnvironmentKey}/applications/ and returns the running applications.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeModules", envsAppsFlags.includeModules, cmd.Flags().Changed("include-modules"))
			addBoolQuery(q, "IncludeEnvStatus", envsAppsFlags.includeEnvStatus, cmd.Flags().Changed("include-env-status"))
			return c.Get(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/applications/", q)
		})
	},
}

// --- environments deployment-zones ---

var envsZonesCmd = &cobra.Command{
	Use:   "deployment-zones <environment-key>",
	Short: "List deployment zones of an environment",
	Long:  "Calls GET /environments/{EnvironmentKey}/deploymentzones/.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/deploymentzones/", nil)
		})
	},
}

// --- environments templates ---

var envsTemplatesCmd = &cobra.Command{
	Use:   "templates <environment-key>",
	Short: "List application templates available in an environment",
	Long:  "Calls GET /environments/{EnvironmentKey}/templates/.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/templates/", nil)
		})
	},
}

// --- environments maintenance-mode ---

var envsMaintenanceFlags struct {
	data string
}

var envsMaintenanceCmd = &cobra.Command{
	Use:   "maintenance-mode <environment-key>",
	Short: "Set the maintenance mode of an environment",
	Long: `Calls PUT /environments/{EnvironmentKey}/maintenancemode.

The request body is a record with IsMaintenanceModeEnabled and Reason. Provide
it with --data (inline JSON, @file, or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(envsMaintenanceFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Put(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/maintenancemode", nil, body)
		})
	},
}

func init() {
	envsAppsCmd.Flags().BoolVar(&envsAppsFlags.includeModules, "include-modules", false, "include module details")
	envsAppsCmd.Flags().BoolVar(&envsAppsFlags.includeEnvStatus, "include-env-status", false, "include per-environment status")

	envsMaintenanceCmd.Flags().StringVar(&envsMaintenanceFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	environmentsCmd.AddCommand(
		envsListCmd,
		envsGetCmd,
		envsAppsCmd,
		envsZonesCmd,
		envsTemplatesCmd,
		envsMaintenanceCmd,
	)
	rootCmd.AddCommand(environmentsCmd)
}
