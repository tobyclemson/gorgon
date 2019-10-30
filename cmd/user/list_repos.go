package user

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tobyclemson/gorgon/github"
)

var userListReposCommand = &cobra.Command{
	Use:   "list-repos",
	Short: "List repositories for GitHub user",
	Long:  "Lists all available repositories for a given GitHub organisation",
	Args:  cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		token := viper.GetString("github-token")
		credentials := github.Credentials{Token: token}
		name := args[0]

		repositories, err :=
			github.ListUserRepositories(name, credentials)
		if err == nil {
			fmt.Println("Repositories for user:", name)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(repository)
			}
			fmt.Println()
		}

		return err
	},
}
