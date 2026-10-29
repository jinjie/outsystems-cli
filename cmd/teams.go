package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

var teamsCmd = &cobra.Command{
	Use:     "teams",
	Aliases: []string{"team"},
	Short:   "Manage teams",
	Long:    "Commands for listing, inspecting, creating, updating, and deleting teams, and managing their users and applications.",
}

// --- teams list ---

var teamsListFlags struct {
	includeUsers        bool
	includeApplications bool
}

var teamsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List teams",
	Long:  "Calls GET /teams and returns all teams.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeUsers", teamsListFlags.includeUsers, cmd.Flags().Changed("include-users"))
			addBoolQuery(q, "IncludeApplications", teamsListFlags.includeApplications, cmd.Flags().Changed("include-applications"))
			return c.Get(commandContext(cmd), "/teams", q)
		})
	},
}

// --- teams get ---

var teamsGetFlags struct {
	includeUsers        bool
	includeApplications bool
}

var teamsGetCmd = &cobra.Command{
	Use:   "get <team-key>",
	Short: "Get the details of a team",
	Long:  "Calls GET /teams/{TeamKey}, with its users and applications.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			q := url.Values{}
			addBoolQuery(q, "IncludeUsers", teamsGetFlags.includeUsers, cmd.Flags().Changed("include-users"))
			addBoolQuery(q, "IncludeApplications", teamsGetFlags.includeApplications, cmd.Flags().Changed("include-applications"))
			return c.Get(commandContext(cmd), "/teams/"+url.PathEscape(args[0]), q)
		})
	},
}

// --- teams create ---

var teamsCreateFlags struct{ data string }

var teamsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a team",
	Long:  "Calls POST /teams. Provide a Team body with --data (inline JSON, @file, or - for stdin).",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(teamsCreateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/teams", nil, body)
		})
	},
}

// --- teams update ---

var teamsUpdateFlags struct{ data string }

var teamsUpdateCmd = &cobra.Command{
	Use:   "update <team-key>",
	Short: "Update a team",
	Long:  "Calls PUT /teams/{TeamKey}. Provide a Team body with --data (inline JSON, @file, or - for stdin).",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(teamsUpdateFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Put(commandContext(cmd), "/teams/"+url.PathEscape(args[0]), nil, body)
		})
	},
}

// --- teams delete ---

var teamsDeleteCmd = &cobra.Command{
	Use:   "delete <team-key>",
	Short: "Delete a team",
	Long:  "Calls DELETE /teams/{TeamKey}.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Delete(commandContext(cmd), "/teams/"+url.PathEscape(args[0]), nil)
		})
	},
}

// --- teams add-application ---

var teamsAddAppFlags struct{ data string }

var teamsAddAppCmd = &cobra.Command{
	Use:   "add-application <team-key>",
	Short: "Assign an application to a team",
	Long: `Calls POST /teams/{TeamKey}/applications/.

Provide an ApplicationShortInfo body (Name, Key) with --data (inline JSON,
@file, or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(teamsAddAppFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/teams/"+url.PathEscape(args[0])+"/applications/", nil, body)
		})
	},
}

// --- teams remove-application ---

var teamsRemoveAppCmd = &cobra.Command{
	Use:   "remove-application <team-key> <application-key>",
	Short: "Remove an application from a team",
	Long:  "Calls DELETE /teams/{TeamKey}/applications/{ApplicationKey}.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/teams/" + url.PathEscape(args[0]) + "/applications/" + url.PathEscape(args[1])
			return c.Delete(commandContext(cmd), path, nil)
		})
	},
}

// --- teams add-user ---

var teamsAddUserFlags struct{ data string }

var teamsAddUserCmd = &cobra.Command{
	Use:   "add-user <team-key>",
	Short: "Add a user to a team with a role",
	Long: `Calls POST /teams/{TeamKey}/users/.

Provide a UserRoleShortInfo body (userKey, roleKey) with --data (inline JSON,
@file, or - for stdin).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			body, err := readBodyInput(teamsAddUserFlags.data)
			if err != nil {
				return nil, err
			}
			return c.Post(commandContext(cmd), "/teams/"+url.PathEscape(args[0])+"/users/", nil, body)
		})
	},
}

// --- teams remove-user ---

var teamsRemoveUserCmd = &cobra.Command{
	Use:   "remove-user <team-key> <user-key>",
	Short: "Remove a user from a team",
	Long:  "Calls DELETE /teams/{TeamKey}/users/{UserKey}.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			path := "/teams/" + url.PathEscape(args[0]) + "/users/" + url.PathEscape(args[1])
			return c.Delete(commandContext(cmd), path, nil)
		})
	},
}

func init() {
	teamsListCmd.Flags().BoolVar(&teamsListFlags.includeUsers, "include-users", false, "include users in each team")
	teamsListCmd.Flags().BoolVar(&teamsListFlags.includeApplications, "include-applications", false, "include applications in each team")
	teamsGetCmd.Flags().BoolVar(&teamsGetFlags.includeUsers, "include-users", false, "include users in the team")
	teamsGetCmd.Flags().BoolVar(&teamsGetFlags.includeApplications, "include-applications", false, "include applications in the team")

	teamsCreateCmd.Flags().StringVar(&teamsCreateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	teamsUpdateCmd.Flags().StringVar(&teamsUpdateFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	teamsAddAppCmd.Flags().StringVar(&teamsAddAppFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")
	teamsAddUserCmd.Flags().StringVar(&teamsAddUserFlags.data, "data", "", "request body: inline JSON, @file, or - for stdin")

	teamsCmd.AddCommand(
		teamsListCmd,
		teamsGetCmd,
		teamsCreateCmd,
		teamsUpdateCmd,
		teamsDeleteCmd,
		teamsAddAppCmd,
		teamsRemoveAppCmd,
		teamsAddUserCmd,
		teamsRemoveUserCmd,
	)
	rootCmd.AddCommand(teamsCmd)
}
