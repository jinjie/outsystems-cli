package cmd

import (
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

var modulesCmd = &cobra.Command{
	Use:     "modules",
	Aliases: []string{"module", "mod"},
	Short:   "Inspect modules",
	Long:    "Commands for listing and inspecting modules and their versions.",
}

// --- modules list ---

var modsListFlags struct {
	includeEnvStatus bool
}

var modsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List modules in the infrastructure",
	Long:  "Calls GET /modules/ and returns all modules.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeEnvStatus", modsListFlags.includeEnvStatus, cmd.Flags().Changed("include-env-status"))
			return c.Get(commandContext(cmd), "/modules/", q)
		})
	},
}

// --- modules get ---

var modsGetFlags struct {
	includeEnvStatus bool
}

var modsGetCmd = &cobra.Command{
	Use:   "get <module-key>",
	Short: "Get the details of a module",
	Long:  "Calls GET /modules/{ModuleKey}/ and returns the module details.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeEnvStatus", modsGetFlags.includeEnvStatus, cmd.Flags().Changed("include-env-status"))
			return c.Get(commandContext(cmd), "/modules/"+url.PathEscape(args[0])+"/", q)
		})
	},
}

// --- modules versions ---

var modsVersionsFlags struct {
	maxVersions int
}

var modsVersionsCmd = &cobra.Command{
	Use:   "versions <module-key>",
	Short: "List versions of a module",
	Long:  "Calls GET /modules/{ModuleKey}/versions/ and returns the version list.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("max-versions") {
				q.Set("MaximumVersionsToReturn", strconv.Itoa(modsVersionsFlags.maxVersions))
			}
			return c.Get(commandContext(cmd), "/modules/"+url.PathEscape(args[0])+"/versions/", q)
		})
	},
}

// --- modules version-get ---

var modsVersionGetCmd = &cobra.Command{
	Use:   "version-get <module-key> <module-version-key>",
	Short: "Get the details of a specific module version",
	Long:  "Calls GET /modules/{ModuleKey}/versions/{ModuleVersionKey}/.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/modules/" + url.PathEscape(args[0]) + "/versions/" + url.PathEscape(args[1]) + "/"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

func init() {
	modsListCmd.Flags().BoolVar(&modsListFlags.includeEnvStatus, "include-env-status", false, "include per-environment status")
	modsGetCmd.Flags().BoolVar(&modsGetFlags.includeEnvStatus, "include-env-status", false, "include per-environment status")
	modsVersionsCmd.Flags().IntVar(&modsVersionsFlags.maxVersions, "max-versions", 0, "maximum number of versions to return (API default is 5)")

	modulesCmd.AddCommand(
		modsListCmd,
		modsGetCmd,
		modsVersionsCmd,
		modsVersionGetCmd,
	)
	rootCmd.AddCommand(modulesCmd)
}
