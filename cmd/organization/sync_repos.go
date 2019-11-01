package organization

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tobyclemson/gorgon/github"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"os"
	"path/filepath"
)

var organizationSyncReposCommand = &cobra.Command{
	Use:   "sync-repos",
	Short: "Synchronizes repositories for GitHub organization",
	Long: "Synchronises all repositories for a given GitHub organization " +
		"with a local directory, which may or may not already contain " +
		"repositories from the organization.",
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		token := viper.GetString("github-token")
		targetDirectory := viper.GetString("target-directory")
		credentials := github.Credentials{Token: token}
		name := args[0]

		err := error(nil)
		resolvedDirectory := ""
		if targetDirectory == "" {
			resolvedDirectory = name
		} else {
			resolvedDirectory = targetDirectory
		}
		if !filepath.IsAbs(targetDirectory) {
			resolvedDirectory, err = filepath.Abs(targetDirectory)
			if err != nil {
				return err
			}
		}
		repositories, err :=
			github.ListOrganizationRepositories(name, credentials)
		if err == nil {
			fmt.Printf(
				"Syncing repositories for organization: '%v' into "+
					"directory: '%v'\n",
				name,
				resolvedDirectory)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(*repository.Name)
				repositoryDirectory :=
					filepath.Join(resolvedDirectory, *repository.Name)
				repositoryOptions := &git.CloneOptions{
					URL:      *repository.GitURL,
					Progress: os.Stdout,
				}
				_, err := git.PlainClone(
					repositoryDirectory, false, repositoryOptions)
				if err != nil && err != transport.ErrEmptyRemoteRepository {
					return err
				}
				fmt.Println()
			}
			fmt.Println()
		}

		return err
	},
}

func init() {
	organizationSyncReposCommand.Flags().
		StringP(
			"target-directory",
			"d",
			"",
			"The directory into which repositories should be synced, "+
				"defaults to a directory under the working directory with "+
				"the name of the organization")
}
