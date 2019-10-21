package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Synchronise GitHub organisation or user with local copy",
	Long: "Clones missing repositories and updates existing repositories " +
		"locally for a given GitHub organisation or user.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var githubUsername = viper.GetString("github-username")
		var githubToken = viper.GetString("github-token")
		var organisationOrUser = args[0]

		fmt.Println("Username:", githubUsername)
		fmt.Println("Token:", githubToken)
		fmt.Println("Organisation or user:", organisationOrUser)
	},
}

func init() {
	rootCommand.AddCommand(syncCommand)
}
