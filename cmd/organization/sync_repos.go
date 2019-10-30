package organization

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var organizationSyncReposCommand = &cobra.Command{
	Use:   "sync-repos",
	Short: "Synchronizes repositories for GitHub organization",
	Long: "Synchronises all repositories for a given GitHub organization " +
		"with a local directory, which may or may not already contain " +
		"repositories from the organization.",
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		token := viper.GetString("github-token")
		//credentials := github.Credentials{Token: token}
		name := args[0]

		fmt.Println("Token:", token)
		fmt.Println("Name:", name)

		return nil
	},
}
