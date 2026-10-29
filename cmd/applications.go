package cmd

import (
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

var applicationsCmd = &cobra.Command{
	Use:     "applications",
	Aliases: []string{"apps", "app"},
	Short:   "Manage and inspect applications",
	Long:    "Commands for listing and inspecting applications and their versions.",
}

// --- applications list ---

var appsListFlags struct {
	includeModules   bool
	includeEnvStatus bool
}

var appsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications in the infrastructure",
	Long:  "Calls GET /applications/ and returns all applications.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeModules", appsListFlags.includeModules, cmd.Flags().Changed("include-modules"))
			addBoolQuery(q, "IncludeEnvStatus", appsListFlags.includeEnvStatus, cmd.Flags().Changed("include-env-status"))
			return c.Get(commandContext(cmd), "/applications/", q)
		})
	},
}

// --- applications get ---

var appsGetFlags struct {
	includeModules   bool
	includeEnvStatus bool
}

var appsGetCmd = &cobra.Command{
	Use:   "get <application-key>",
	Short: "Get the details of an application",
	Long:  "Calls GET /applications/{ApplicationKey}/ and returns the application details.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeModules", appsGetFlags.includeModules, cmd.Flags().Changed("include-modules"))
			addBoolQuery(q, "IncludeEnvStatus", appsGetFlags.includeEnvStatus, cmd.Flags().Changed("include-env-status"))
			return c.Get(commandContext(cmd), "/applications/"+url.PathEscape(args[0])+"/", q)
		})
	},
}

// --- applications versions list ---

var appsVersionsListFlags struct {
	changeLogFilter string
	maxVersions     int
}

var appsVersionsListCmd = &cobra.Command{
	Use:   "versions <application-key>",
	Short: "List versions of an application",
	Long:  "Calls GET /applications/{ApplicationKey}/versions/ and returns the version list.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("change-log-filter") {
				q.Set("ChangeLogFilter", appsVersionsListFlags.changeLogFilter)
			}
			if cmd.Flags().Changed("max-versions") {
				q.Set("MaximumVersionsToReturn", strconv.Itoa(appsVersionsListFlags.maxVersions))
			}
			return c.Get(commandContext(cmd), "/applications/"+url.PathEscape(args[0])+"/versions/", q)
		})
	},
}

// --- applications version-get ---

var appsVersionGetFlags struct {
	includeModules bool
}

var appsVersionGetCmd = &cobra.Command{
	Use:   "version-get <application-key> <version-key>",
	Short: "Get the details of a specific application version",
	Long:  "Calls GET /applications/{ApplicationKey}/versions/{VersionKey}/.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			// IncludeModules is required by the API for this endpoint.
			q := url.Values{}
			q.Set("IncludeModules", strconv.FormatBool(appsVersionGetFlags.includeModules))
			path := "/applications/" + url.PathEscape(args[0]) + "/versions/" + url.PathEscape(args[1]) + "/"
			return c.Get(commandContext(cmd), path, q)
		})
	},
}

// --- applications version-delete ---

var appsVersionDeleteCmd = &cobra.Command{
	Use:   "version-delete <application-key> <version-key>",
	Short: "Discard an application version",
	Long: `Calls DELETE /applications/{ApplicationKey}/versions/{VersionKey}/.

Running versions, or versions used in deployments, cannot be deleted.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/applications/" + url.PathEscape(args[0]) + "/versions/" + url.PathEscape(args[1]) + "/"
			return c.Delete(commandContext(cmd), path, nil)
		})
	},
}

// --- applications version-download ---

var appsVersionDownloadCmd = &cobra.Command{
	Use:   "version-download <application-key> <version-key>",
	Short: "Get a download link for an application version binary",
	Long: `Calls GET /applications/{ApplicationKey}/versions/{VersionKey}/content/
and returns a link to download the binary. The link expires in 60 minutes.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/applications/" + url.PathEscape(args[0]) + "/versions/" + url.PathEscape(args[1]) + "/content/"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- applications configurations get ---

var appsConfigGetCmd = &cobra.Command{
	Use:   "configurations <application-key>",
	Short: "Get the application configurations per environment",
	Long:  "Calls GET /applications/{ApplicationKey}/configurations/.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/applications/" + url.PathEscape(args[0]) + "/configurations/"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- applications configurations-update ---

var appsConfigUpdateFlags struct {
	data string
}

var appsConfigUpdateCmd = &cobra.Command{
	Use:   "configurations-update <application-key>",
	Short: "Update the application configurations per environment",
	Long: `Calls PUT /applications/{ApplicationKey}/configurations/.

The request body is an ApplicationConfigurationListRecord. Provide it with
--data, which accepts an inline JSON string, @file to read a file, or - to
read from stdin.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(appsConfigUpdateFlags.data)
			if err != nil {
				return nil, err
			}
			path := "/applications/" + url.PathEscape(args[0]) + "/configurations/"
			return c.Put(commandContext(cmd), path, nil, body)
		})
	},
}

func init() {
	appsListCmd.Flags().BoolVar(&appsListFlags.includeModules, "include-modules", false, "include module details")
	appsListCmd.Flags().BoolVar(&appsListFlags.includeEnvStatus, "include-env-status", false, "include per-environment status")

	appsGetCmd.Flags().BoolVar(&appsGetFlags.includeModules, "include-modules", false, "include module details")
	appsGetCmd.Flags().BoolVar(&appsGetFlags.includeEnvStatus, "include-env-status", false, "include per-environment status")

	appsVersionsListCmd.Flags().StringVar(&appsVersionsListFlags.changeLogFilter, "change-log-filter", "", "only versions whose change log contains this string")
	appsVersionsListCmd.Flags().IntVar(&appsVersionsListFlags.maxVersions, "max-versions", 0, "maximum number of versions to return (API default is 5)")

	appsVersionGetCmd.Flags().BoolVar(&appsVersionGetFlags.includeModules, "include-modules", false, "include module details (sent as IncludeModules)")

	appsConfigUpdateCmd.Flags().StringVar(&appsConfigUpdateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	applicationsCmd.AddCommand(
		appsListCmd,
		appsGetCmd,
		appsVersionsListCmd,
		appsVersionGetCmd,
		appsVersionDeleteCmd,
		appsVersionDownloadCmd,
		appsConfigGetCmd,
		appsConfigUpdateCmd,
	)
	rootCmd.AddCommand(applicationsCmd)
}
