package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

// --- environments register ---

var envsRegisterFlags struct{ data string }

var envsRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register an environment in LifeTime",
	Long: `Calls POST /environments/ to register an environment.

Provide an EnvironmentRegistration body with --data (inline JSON, @file, or -
for stdin). Required fields include EnvironmentName, EnvironmentAddress,
ServiceCenterUsername, ServiceCenterPassword, LifeTimeAddress, and
EnvironmentPosition.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(envsRegisterFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/environments/", nil, body)
		})
	},
}

// --- environments unregister ---

var envsUnregisterCmd = &cobra.Command{
	Use:   "unregister <environment-key>",
	Short: "Unregister an environment from LifeTime",
	Long:  "Calls DELETE /environments/{EnvironmentKey}.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			// Note: this endpoint has no trailing slash in swagger.
			return c.Delete(commandContext(cmd), "/environments/"+url.PathEscape(args[0]), nil)
		})
	},
}

// --- environments app-get (running app detail) ---

var envsAppGetFlags struct {
	includeEnvStatus bool
	includeModules   bool
}

var envsAppGetCmd = &cobra.Command{
	Use:   "app-get <environment-key> <application-key>",
	Short: "Get the running version of an application in an environment",
	Long:  "Calls GET /environments/{EnvironmentKey}/applications/{ApplicationKey}/.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeEnvStatus", envsAppGetFlags.includeEnvStatus, cmd.Flags().Changed("include-env-status"))
			addBoolQuery(q, "IncludeModules", envsAppGetFlags.includeModules, cmd.Flags().Changed("include-modules"))
			path := "/environments/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1]) + "/"
			return c.Get(commandContext(cmd), path, q)
		})
	},
}

// --- environments app-download (download running app binary link) ---

var envsAppDownloadFlags struct {
	binaryType string
}

var envsAppDownloadCmd = &cobra.Command{
	Use:   "app-download <environment-key> <application-key>",
	Short: "Get a download link for a running application's binary",
	Long: `Calls GET /environments/{EnvironmentKey}/applications/{ApplicationKey}/content/
and returns a link that expires in 60 minutes. Use --type to pick the binary
type when applicable (one of "oap", "apk", or "ipa").`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("type") {
				q.Set("Type", envsAppDownloadFlags.binaryType)
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1]) + "/content/"
			return c.Get(commandContext(cmd), path, q)
		})
	},
}

// --- environments app-create ---

var envsAppCreateFlags struct{ data string }

var envsAppCreateCmd = &cobra.Command{
	Use:   "app-create <environment-key>",
	Short: "Create a new application in an environment",
	Long: `Calls POST /environments/{EnvironmentKey}/applications/ and returns the key
of the new application.

Provide an ApplicationCreate body with --data (inline JSON, @file, or - for
stdin). Required fields: TemplateKey, TeamKey, Name, Color.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(envsAppCreateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/applications/", nil, body)
		})
	},
}

// --- environments app-version-create ---

var envsAppVersionCreateFlags struct{ data string }

var envsAppVersionCreateCmd = &cobra.Command{
	Use:   "app-version-create <environment-key> <application-key>",
	Short: "Create a new application version from the running application",
	Long: `Calls POST /environments/{EnvironmentKey}/applications/{ApplicationKey}/versions/
and returns the key of the new version.

Provide an ApplicationVersionCreate body with --data (inline JSON, @file, or -
for stdin).`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(envsAppVersionCreateFlags.data)
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1]) + "/versions/"
			return c.Post(commandContext(cmd), path, nil, body)
		})
	},
}

// --- environments app-sourcecode (request) ---

var envsAppSourceCodeCmd = &cobra.Command{
	Use:   "app-sourcecode <environment-key> <application-key>",
	Short: "Request the source code package of an application",
	Long: `Calls POST /environments/{EnvironmentKey}/applications/{ApplicationKey}/sourcecodeaccess
and returns the operation key (PackageKey) and status. Poll status with
'app-sourcecode-status' and fetch the link with 'app-sourcecode-download'.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1]) + "/sourcecodeaccess"
			return c.Post(commandContext(cmd), path, nil, nil)
		})
	},
}

// --- environments app-sourcecode-status ---

var envsAppSourceCodeStatusCmd = &cobra.Command{
	Use:   "app-sourcecode-status <environment-key> <application-key> <package-key>",
	Short: "Get the status of an application source code package",
	Long:  "Calls GET /environments/{EnvironmentKey}/applications/{ApplicationKey}/sourcecodeaccess/{PackageKey}/status.",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1]) +
				"/sourcecodeaccess/" + url.PathEscape(args[2]) + "/status"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- environments app-sourcecode-download ---

var envsAppSourceCodeDownloadCmd = &cobra.Command{
	Use:   "app-sourcecode-download <environment-key> <application-key> <package-key>",
	Short: "Get a download link for an application source code package",
	Long: `Calls GET /environments/{EnvironmentKey}/applications/{ApplicationKey}/sourcecodeaccess/{PackageKey}/download
and returns a link that expires in 60 minutes.`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1]) +
				"/sourcecodeaccess/" + url.PathEscape(args[2]) + "/download"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

func init() {
	envsRegisterCmd.Flags().StringVar(&envsRegisterFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	envsAppGetCmd.Flags().BoolVar(&envsAppGetFlags.includeEnvStatus, "include-env-status", false, "include per-environment status")
	envsAppGetCmd.Flags().BoolVar(&envsAppGetFlags.includeModules, "include-modules", false, "include module details")

	envsAppDownloadCmd.Flags().StringVar(&envsAppDownloadFlags.binaryType, "type", "", `binary type: "oap", "apk", or "ipa"`)

	envsAppCreateCmd.Flags().StringVar(&envsAppCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	envsAppVersionCreateCmd.Flags().StringVar(&envsAppVersionCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	environmentsCmd.AddCommand(
		envsRegisterCmd,
		envsUnregisterCmd,
		envsAppGetCmd,
		envsAppDownloadCmd,
		envsAppCreateCmd,
		envsAppVersionCreateCmd,
		envsAppSourceCodeCmd,
		envsAppSourceCodeStatusCmd,
		envsAppSourceCodeDownloadCmd,
	)
}
