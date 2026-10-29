package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the LifeTime version and supported Platform Server version",
	Long: `version calls GET /version/ and prints the LifeTime version information,
including the installed LifeTime version and the supported Platform Server
version.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAndPrint(func() ([]byte, error) {
			c, err := resolveClient()
			if err != nil {
				return nil, err
			}
			return c.Get(commandContext(cmd), "/version/", nil)
		})
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
