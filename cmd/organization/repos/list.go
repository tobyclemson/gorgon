package repos

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/github"
)

var organizationReposListCommand = &cobra.Command{
	Use:   "list",
	Short: "List repositories for GitHub organization",
	Long:  "Lists all available repositories for a given GitHub organization",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := cmd.Flag("github-token").Value.String()
		credentials := github.Credentials{Token: token}
		name := args[0]

		repositories, err :=
			github.ListOrganizationRepositories(name, credentials)
		if err == nil {
			fmt.Printf("Listing repositories for organization: '%v'\n", name)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(*repository.Name)
			}
			fmt.Println()
		}

		return err
	},
}
