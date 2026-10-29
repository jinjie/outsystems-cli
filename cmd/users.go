package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"user"},
	Short:   "Manage users",
	Long:    "Commands for listing, inspecting, creating, and updating users, and managing their application roles and blocked logins.",
}

// --- users list ---

var usersListFlags struct {
	includeInactive        bool
	includeTeams           bool
	includeApplicationRole bool
}

var usersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users",
	Long:  "Calls GET /users/ and returns all users.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeInactive", usersListFlags.includeInactive, cmd.Flags().Changed("include-inactive"))
			addBoolQuery(q, "IncludeTeams", usersListFlags.includeTeams, cmd.Flags().Changed("include-teams"))
			addBoolQuery(q, "IncludeApplicationRoles", usersListFlags.includeApplicationRole, cmd.Flags().Changed("include-application-roles"))
			return c.Get(commandContext(cmd), "/users/", q)
		})
	},
}

// --- users get ---

var usersGetFlags struct {
	includeTeams           bool
	includeApplicationRole bool
}

var usersGetCmd = &cobra.Command{
	Use:   "get <user-key>",
	Short: "Get the details of a user",
	Long:  "Calls GET /users/{UserKey} and returns the user, teams, and apps.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeTeams", usersGetFlags.includeTeams, cmd.Flags().Changed("include-teams"))
			addBoolQuery(q, "IncludeApplicationRoles", usersGetFlags.includeApplicationRole, cmd.Flags().Changed("include-application-roles"))
			return c.Get(commandContext(cmd), "/users/"+url.PathEscape(args[0]), q)
		})
	},
}

// --- users create ---

var usersCreateFlags struct{ data string }

var usersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Long: `Calls POST /users/ and returns the created user.

Provide a User body with --data (inline JSON, @file, or - for stdin). Teams
and applications must be associated using the dedicated endpoints.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(usersCreateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/users/", nil, body)
		})
	},
}

// --- users update ---

var usersUpdateFlags struct{ data string }

var usersUpdateCmd = &cobra.Command{
	Use:   "update <user-key>",
	Short: "Update a user",
	Long:  "Calls PUT /users/{UserKey}/. Provide a User body with --data (inline JSON, @file, or - for stdin).",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(usersUpdateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Put(commandContext(cmd), "/users/"+url.PathEscape(args[0])+"/", nil, body)
		})
	},
}

// --- users set-password ---

var usersSetPasswordFlags struct{ data string }

var usersSetPasswordCmd = &cobra.Command{
	Use:   "set-password <user-key>",
	Short: "Set a user's password",
	Long: `Calls POST /users/{UserKey}/setpassword/.

Provide a Password body ({"password": "..."}) with --data (inline JSON, @file,
or - for stdin). Prefer @file or - so the secret is not kept in shell history.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(usersSetPasswordFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/users/"+url.PathEscape(args[0])+"/setpassword/", nil, body)
		})
	},
}

// --- users add-application ---

var usersAddAppFlags struct{ data string }

var usersAddAppCmd = &cobra.Command{
	Use:   "add-application <user-key>",
	Short: "Add or update a user's role for an application",
	Long: `Calls POST /users/{UserKey}/applications/.

Provide an ApplicationRole body with --data (inline JSON, @file, or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(usersAddAppFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/users/"+url.PathEscape(args[0])+"/applications/", nil, body)
		})
	},
}

// --- users remove-application ---

var usersRemoveAppCmd = &cobra.Command{
	Use:   "remove-application <user-key> <application-key>",
	Short: "Remove a user's application permissions",
	Long:  "Calls DELETE /users/{UserKey}/applications/{ApplicationKey}.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/users/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1])
			return c.Delete(commandContext(cmd), path, nil)
		})
	},
}

// --- users blocked-logins ---

var usersBlockedLoginsFlags struct {
	environmentKey string
}

var usersBlockedLoginsCmd = &cobra.Command{
	Use:   "blocked-logins <user-key>",
	Short: "Get a user's blocked login attempts",
	Long:  "Calls GET /users/{UserKey}/blockedlogins/.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("environment-key") {
				q.Set("EnvironmentKey", usersBlockedLoginsFlags.environmentKey)
			}
			return c.Get(commandContext(cmd), "/users/"+url.PathEscape(args[0])+"/blockedlogins/", q)
		})
	},
}

// --- users unblock ---

var usersUnblockFlags struct {
	environmentKey string
	ipAddress      string
}

var usersUnblockCmd = &cobra.Command{
	Use:   "unblock <user-key>",
	Short: "Unblock a user's login attempts",
	Long:  "Calls DELETE /users/{UserKey}/blockedlogins/.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("environment-key") {
				q.Set("EnvironmentKey", usersUnblockFlags.environmentKey)
			}
			if cmd.Flags().Changed("ip-address") {
				q.Set("IpAddress", usersUnblockFlags.ipAddress)
			}
			return c.Delete(commandContext(cmd), "/users/"+url.PathEscape(args[0])+"/blockedlogins/", q)
		})
	},
}

func init() {
	usersListCmd.Flags().BoolVar(&usersListFlags.includeInactive, "include-inactive", false, "include inactive users")
	usersListCmd.Flags().BoolVar(&usersListFlags.includeTeams, "include-teams", false, "include teams each user belongs to")
	usersListCmd.Flags().BoolVar(&usersListFlags.includeApplicationRole, "include-application-roles", false, "include application roles")

	usersGetCmd.Flags().BoolVar(&usersGetFlags.includeTeams, "include-teams", false, "include teams the user belongs to")
	usersGetCmd.Flags().BoolVar(&usersGetFlags.includeApplicationRole, "include-application-roles", false, "include application roles")

	usersCreateCmd.Flags().StringVar(&usersCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	usersUpdateCmd.Flags().StringVar(&usersUpdateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	usersSetPasswordCmd.Flags().StringVar(&usersSetPasswordFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	usersAddAppCmd.Flags().StringVar(&usersAddAppFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	usersBlockedLoginsCmd.Flags().StringVar(&usersBlockedLoginsFlags.environmentKey, "environment-key", "", "filter by environment key (all if empty)")
	usersUnblockCmd.Flags().StringVar(&usersUnblockFlags.environmentKey, "environment-key", "", "environment to unblock (all if empty)")
	usersUnblockCmd.Flags().StringVar(&usersUnblockFlags.ipAddress, "ip-address", "", "IP address to unblock (all if empty)")

	usersCmd.AddCommand(
		usersListCmd,
		usersGetCmd,
		usersCreateCmd,
		usersUpdateCmd,
		usersSetPasswordCmd,
		usersAddAppCmd,
		usersRemoveAppCmd,
		usersBlockedLoginsCmd,
		usersUnblockCmd,
	)
	rootCmd.AddCommand(usersCmd)
}
