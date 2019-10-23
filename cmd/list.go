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
	Long: "Lists all available repositories for a given GitHub organisation " +
		"or user.",
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		token := viper.GetString("github-token")
		credentials := github.Credentials{Token: token}
		entityType := viper.GetString("github-type")
		name := args[0]

		entity, err := github.EntityFromString(entityType)
		if err != nil {
			return err
		}

		repositories, err :=
			github.ListRepositories(entity, name, credentials)
		if err == nil {
			fmt.Println("Repositories for", entity, name)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(repository)
			}
			fmt.Println()
		}

		return err
	},
}

func init() {
	rootCommand.AddCommand(listCommand)
}
