package user

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userSyncReposCommand = &cobra.Command{
	Use:   "sync-repos",
	Short: "Synchronizes repositories for GitHub user",
	Long: "Synchronises all repositories for a given GitHub user " +
		"with a local directory, which may or may not already contain " +
		"repositories from the user.",
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
