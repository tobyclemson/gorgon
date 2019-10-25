package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tobyclemson/gorgon/github"
)

var organizationCommand = &cobra.Command{
	Use:   "organization",
	Short: "Manage GitHub organizations",
	Long:  "Perform various actions for a given GitHub organization.",
}

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

func init() {
	rootCommand.AddCommand(organizationCommand)
	organizationCommand.AddCommand(organizationListReposCommand)
	organizationCommand.AddCommand(organizationSyncReposCommand)
}
