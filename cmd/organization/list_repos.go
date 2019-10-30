package organization

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tobyclemson/gorgon/github"
)

var organizationListReposCommand = &cobra.Command{
	Use:   "list-repos",
	Short: "List repositories for GitHub organization",
	Long:  "Lists all available repositories for a given GitHub organization",
	Args:  cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		token := viper.GetString("github-token")
		credentials := github.Credentials{Token: token}
		name := args[0]

		repositories, err :=
			github.ListOrganisationRepositories(name, credentials)
		if err == nil {
			fmt.Println("Repositories for organization:", name)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(repository)
			}
			fmt.Println()
		}

		return err
	},
}
