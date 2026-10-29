package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

// dbConnectionsCmd groups the DB connection endpoints under the environments
// resource, e.g. "environments dbconnections list <environment-key>".
var dbConnectionsCmd = &cobra.Command{
	Use:     "dbconnections",
	Aliases: []string{"dbconnection", "dbconn"},
	Short:   "Manage environment database connections",
	Long:    "Commands for listing, creating, editing, deleting, and testing environment DB connections and their permissions.",
}

// --- dbconnections providers ---

var dbProvidersCmd = &cobra.Command{
	Use:   "providers <environment-key>",
	Short: "List DB connection providers for an environment",
	Long:  "Calls GET /environments/{environmentKey}/dbconnection/dbproviders/.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/dbconnection/dbproviders/", nil)
		})
	},
}

// --- dbconnections list ---

var dbListCmd = &cobra.Command{
	Use:   "list <environment-key>",
	Short: "List DB connections of an environment",
	Long:  "Calls GET /environments/{environmentKey}/dbconnections/.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/dbconnections/", nil)
		})
	},
}

// --- dbconnections get ---

var dbGetCmd = &cobra.Command{
	Use:   "get <environment-key> <connection-name>",
	Short: "Get a DB connection's details",
	Long:  "Calls GET /environments/{environmentKey}/dbconnections/{dbConnectionName}/.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) + "/"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- dbconnections create ---

var dbCreateFlags struct{ data string }

var dbCreateCmd = &cobra.Command{
	Use:   "create <environment-key>",
	Short: "Create a DB connection",
	Long: `Calls POST /environments/{environmentKey}/dbconnections/.

Provide a DBConnection_Request_POST body with --data (inline JSON, @file, or -
for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(dbCreateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/dbconnections/", nil, body)
		})
	},
}

// --- dbconnections edit ---

var dbEditFlags struct{ data string }

var dbEditCmd = &cobra.Command{
	Use:   "edit <environment-key> <connection-name>",
	Short: "Update a DB connection",
	Long: `Calls PUT /environments/{environmentKey}/dbconnections/{dbConnectionName}/.

Provide a DBConnection_Request_POST body with --data (inline JSON, @file, or -
for stdin).`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(dbEditFlags.data)
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) + "/"
			return c.Put(commandContext(cmd), path, nil, body)
		})
	},
}

// --- dbconnections delete ---

var dbDeleteCmd = &cobra.Command{
	Use:   "delete <environment-key> <connection-name>",
	Short: "Delete a DB connection",
	Long:  "Calls DELETE /environments/{environmentKey}/dbconnections/{dbConnectionName}/.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) + "/"
			return c.Delete(commandContext(cmd), path, nil)
		})
	},
}

// --- dbconnections test ---

var dbTestFlags struct{ data string }

var dbTestCmd = &cobra.Command{
	Use:   "test <environment-key>",
	Short: "Test a DB connection",
	Long: `Calls POST /environments/{environmentKey}/testdbconnection/.

Provide a DBConnection_Request_POST body with --data (inline JSON, @file, or -
for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(dbTestFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/environments/"+url.PathEscape(args[0])+"/testdbconnection/", nil, body)
		})
	},
}

// --- dbconnections permission-levels ---

var dbPermissionLevelsCmd = &cobra.Command{
	Use:   "permission-levels",
	Short: "List DB connection permission levels",
	Long:  "Calls GET /environments/dbconnection/permissionlevels/.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/environments/dbconnection/permissionlevels/", nil)
		})
	},
}

// --- dbconnections role-permission-get ---

var dbRolePermGetCmd = &cobra.Command{
	Use:   "role-permission-get <environment-key> <connection-name> <role>",
	Short: "Get a role's DB connection permission level",
	Long:  "Calls GET /environments/{environmentKey}/dbconnections/{dbConnectionName}/rolepermissionlevel/{role}/.",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) +
				"/rolepermissionlevel/" + url.PathEscape(args[2]) + "/"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- dbconnections role-permission-set ---

var dbRolePermSetFlags struct{ data string }

var dbRolePermSetCmd = &cobra.Command{
	Use:   "role-permission-set <environment-key> <connection-name>",
	Short: "Set a role's DB connection permission level",
	Long: `Calls PUT /environments/{environmentKey}/dbconnections/{dbConnectionName}/rolepermissionlevel/.

Provide a DBConnection_RolePermissionLevel body with --data (inline JSON,
@file, or - for stdin). Permission levels: Full Control, No Access, Read-Only.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(dbRolePermSetFlags.data)
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) + "/rolepermissionlevel/"
			return c.Put(commandContext(cmd), path, nil, body)
		})
	},
}

// --- dbconnections role-permission-revoke ---

var dbRolePermRevokeFlags struct{ data string }

var dbRolePermRevokeCmd = &cobra.Command{
	Use:   "role-permission-revoke <environment-key> <connection-name>",
	Short: "Revoke a role's DB connection permissions",
	Long: `Calls DELETE /environments/{environmentKey}/dbconnections/{dbConnectionName}/rolepermissionlevel/.

Provide a DBConnectionRole body ({"role":"..."}) with --data (inline JSON,
@file, or - for stdin).`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(dbRolePermRevokeFlags.data)
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) + "/rolepermissionlevel/"
			return c.Do(commandContext(cmd), "DELETE", path, nil, bytesBody(body))
		})
	},
}

// --- dbconnections user-permission-get ---

var dbUserPermGetCmd = &cobra.Command{
	Use:   "user-permission-get <environment-key> <connection-name> <username>",
	Short: "Get a user's DB connection permission level",
	Long:  "Calls GET /environments/{environmentKey}/dbconnections/{dbConnectionName}/userpermissionlevel/{username}/.",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) +
				"/userpermissionlevel/" + url.PathEscape(args[2]) + "/"
			return c.Get(commandContext(cmd), path, nil)
		})
	},
}

// --- dbconnections user-permission-set ---

var dbUserPermSetFlags struct{ data string }

var dbUserPermSetCmd = &cobra.Command{
	Use:   "user-permission-set <environment-key> <connection-name>",
	Short: "Set a user's DB connection permission level",
	Long: `Calls PUT /environments/{environmentKey}/dbconnections/{dbConnectionName}/userpermissionlevel/.

Provide a DBConnection_UserPermissionLevel body with --data (inline JSON,
@file, or - for stdin). Permission levels: Full Control, No Access, Read-Only.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(dbUserPermSetFlags.data)
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) + "/userpermissionlevel/"
			return c.Put(commandContext(cmd), path, nil, body)
		})
	},
}

// --- dbconnections user-permission-revoke ---

var dbUserPermRevokeFlags struct{ data string }

var dbUserPermRevokeCmd = &cobra.Command{
	Use:   "user-permission-revoke <environment-key> <connection-name>",
	Short: "Revoke a user's DB connection permissions",
	Long: `Calls DELETE /environments/{environmentKey}/dbconnections/{dbConnectionName}/userpermissionlevel/.

Provide a DBConnectionUser body ({"username":"..."}) with --data (inline JSON,
@file, or - for stdin).`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(dbUserPermRevokeFlags.data)
			if err != nil {
				return nil, err
			}
			path := "/environments/" + url.PathEscape(args[0]) + "/dbconnections/" + url.PathEscape(args[1]) + "/userpermissionlevel/"
			return c.Do(commandContext(cmd), "DELETE", path, nil, bytesBody(body))
		})
	},
}

func init() {
	dbCreateCmd.Flags().StringVar(&dbCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	dbEditCmd.Flags().StringVar(&dbEditFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	dbTestCmd.Flags().StringVar(&dbTestFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	dbRolePermSetCmd.Flags().StringVar(&dbRolePermSetFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	dbRolePermRevokeCmd.Flags().StringVar(&dbRolePermRevokeFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	dbUserPermSetCmd.Flags().StringVar(&dbUserPermSetFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	dbUserPermRevokeCmd.Flags().StringVar(&dbUserPermRevokeFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	dbConnectionsCmd.AddCommand(
		dbProvidersCmd,
		dbListCmd,
		dbGetCmd,
		dbCreateCmd,
		dbEditCmd,
		dbDeleteCmd,
		dbTestCmd,
		dbPermissionLevelsCmd,
		dbRolePermGetCmd,
		dbRolePermSetCmd,
		dbRolePermRevokeCmd,
		dbUserPermGetCmd,
		dbUserPermSetCmd,
		dbUserPermRevokeCmd,
	)
	environmentsCmd.AddCommand(dbConnectionsCmd)
}
