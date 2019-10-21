package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tobyclemson/gorgon/github"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List repositories for GitHub organisation or user",
	Long: "Lists all available repositories for a given GitHub organisation or user.",
	Args: cobra.ExactArgs(1),
	Run: func(command *cobra.Command, args []string) {
		githubToken := viper.GetString("github-token")
		githubType := viper.GetString("github-type")
		githubCredentials := github.Credentials{Token: githubToken}
		organisationOrUser := args[0]

		var repositories []string
		var	err error
		if githubType == "organization" {
			repositories, err = github.ListOrganisationRepositories(
				organisationOrUser,
				githubCredentials)
		} else {
			repositories, err = github.ListUserRepositories(
				organisationOrUser,
				githubCredentials)
		}

		if err != nil {
			panic(err)
		}

		fmt.Println("Repositories for", githubType, organisationOrUser)
		fmt.Println()
		for _, repository := range repositories {
			fmt.Println(repository)
		}
		fmt.Println()
	},
}

func init() {
	rootCommand.AddCommand(listCommand)
}
