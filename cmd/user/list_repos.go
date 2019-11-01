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
	Long:  "Lists all available repositories for a given GitHub organization",
	Args:  cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		token := viper.GetString("github-token")
		credentials := github.Credentials{Token: token}
		name := args[0]

		repositories, err :=
			github.ListUserRepositories(name, credentials)
		if err == nil {
			fmt.Printf("Listing repositories for user: '%v'\n", name)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(*repository.Name)
			}
			fmt.Println()
		}

		return err
	},
}
