package user

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/config"
	"github.com/tobyclemson/gorgon/git"
	"github.com/tobyclemson/gorgon/github"
	"os"
	"path/filepath"
)

var userSyncReposCommand = &cobra.Command{
	Use:   "sync-repos",
	Short: "Synchronizes repositories for GitHub user",
	Long: "Synchronises all repositories for a given GitHub user " +
		"with a local directory, which may or may not already contain " +
		"repositories from the user.",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		token, err := config.GithubToken(cmd)
		if err != nil {
			return err
		}

		targetDirectory, err := config.TargetDirectory(cmd, name)
		if err != nil {
			return err
		}

		credentials := github.Credentials{Token: token}

		repositories, err :=
			github.ListUserRepositories(name, credentials)
		if err == nil {
			fmt.Printf(
				"Syncing repositories for user: '%v' into "+
					"directory: '%v'\n",
				name,
				targetDirectory)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(*repository.Name)
				repositoryDirectory :=
					filepath.Join(targetDirectory, *repository.Name)
				if _, err := os.Stat(repositoryDirectory); os.IsNotExist(err) {
					err = git.Clone(repository, repositoryDirectory)
					if err != nil {
						return err
					}
				} else {
					err = git.Pull(repositoryDirectory)
					if err != nil {
						return err
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}

		return err
	},
}

func init() {
	userSyncReposCommand.Flags().
		StringP(
			"target-directory",
			"d",
			"",
			"directory into which repositories should be synced, defaults "+
				"to a directory under the working directory with the name of "+
				"the user")
}
