package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

var rolesCmd = &cobra.Command{
	Use:     "roles",
	Aliases: []string{"role"},
	Short:   "Manage platform roles",
	Long:    "Commands for listing, inspecting, creating, updating, and deleting platform roles.",
}

// --- roles list ---

var rolesListFlags struct {
	includeEnvPermissions bool
}

var rolesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List platform roles",
	Long:  "Calls GET /roles/ and returns all roles.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeEnvPermissions", rolesListFlags.includeEnvPermissions, cmd.Flags().Changed("include-env-permissions"))
			return c.Get(commandContext(cmd), "/roles/", q)
		})
	},
}

// --- roles get ---

var rolesGetFlags struct {
	includeEnvPermissions bool
}

var rolesGetCmd = &cobra.Command{
	Use:   "get <role-key>",
	Short: "Get the details of a role",
	Long:  "Calls GET /roles/{RoleKey} and returns the role details.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeEnvPermissions", rolesGetFlags.includeEnvPermissions, cmd.Flags().Changed("include-env-permissions"))
			return c.Get(commandContext(cmd), "/roles/"+url.PathEscape(args[0]), q)
		})
	},
}

// --- roles permission-levels ---

var rolesPermLevelsCmd = &cobra.Command{
	Use:   "permission-levels",
	Short: "List permission levels that can be assigned to roles",
	Long:  "Calls GET /roles/permissionlevels.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/roles/permissionlevels", nil)
		})
	},
}

// --- roles create ---

var rolesCreateFlags struct {
	data string
}

var rolesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a role",
	Long: `Calls POST /roles/ and returns the key of the new role.

The request body is a PlatformRole. Provide it with --data (inline JSON,
@file, or - for stdin).`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(rolesCreateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/roles/", nil, body)
		})
	},
}

// --- roles update ---

var rolesUpdateFlags struct {
	data string
}

var rolesUpdateCmd = &cobra.Command{
	Use:   "update <role-key>",
	Short: "Update a role",
	Long: `Calls PUT /roles/{RoleKey} and returns the key of the updated role.

Provide the PlatformRole body with --data (inline JSON, @file, or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(rolesUpdateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Put(commandContext(cmd), "/roles/"+url.PathEscape(args[0]), nil, body)
		})
	},
}

// --- roles delete ---

var rolesDeleteFlags struct {
	usersNewRoleKey string
}

var rolesDeleteCmd = &cobra.Command{
	Use:   "delete <role-key>",
	Short: "Delete a role",
	Long: `Calls DELETE /roles/{RoleKey}.

Use --users-new-role-key to grant an alternative role to users who currently
have the role being deleted.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			if cmd.Flags().Changed("users-new-role-key") {
				q.Set("UsersNewRoleKey", rolesDeleteFlags.usersNewRoleKey)
			}
			return c.Delete(commandContext(cmd), "/roles/"+url.PathEscape(args[0]), q)
		})
	},
}

func init() {
	rolesListCmd.Flags().BoolVar(&rolesListFlags.includeEnvPermissions, "include-env-permissions", false, "include per-environment permissions")
	rolesGetCmd.Flags().BoolVar(&rolesGetFlags.includeEnvPermissions, "include-env-permissions", false, "include per-environment permissions")
	rolesCreateCmd.Flags().StringVar(&rolesCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	rolesUpdateCmd.Flags().StringVar(&rolesUpdateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	rolesDeleteCmd.Flags().StringVar(&rolesDeleteFlags.usersNewRoleKey, "users-new-role-key", "", "role to grant to users who had the deleted role")

	rolesCmd.AddCommand(
		rolesListCmd,
		rolesGetCmd,
		rolesPermLevelsCmd,
		rolesCreateCmd,
		rolesUpdateCmd,
		rolesDeleteCmd,
	)
	rootCmd.AddCommand(rolesCmd)
}
