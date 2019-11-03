package user

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/github"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
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
		token := cmd.Flag("github-token").Value.String()
		targetDirectory := cmd.Flag("target-directory").Value.String()
		credentials := github.Credentials{Token: token}
		name := args[0]

		err := error(nil)
		resolvedDirectory := ""
		if targetDirectory == "" {
			resolvedDirectory = name
		} else {
			resolvedDirectory = targetDirectory
		}
		if !filepath.IsAbs(resolvedDirectory) {
			resolvedDirectory, err = filepath.Abs(resolvedDirectory)
			if err != nil {
				return err
			}
		}
		repositories, err :=
			github.ListUserRepositories(name, credentials)
		if err == nil {
			fmt.Printf(
				"Syncing repositories for user: '%v' into "+
					"directory: '%v'\n",
				name,
				resolvedDirectory)
			fmt.Println()
			for _, repository := range repositories {
				fmt.Println(*repository.Name)
				repositoryDirectory :=
					filepath.Join(resolvedDirectory, *repository.Name)
				if _, err := os.Stat(repositoryDirectory); os.IsNotExist(err) {
					cloneOptions := &git.CloneOptions{
						URL:      *repository.GitURL,
						Progress: os.Stdout,
					}
					_, err := git.PlainClone(
						repositoryDirectory, false, cloneOptions)
					if err != nil && err != transport.ErrEmptyRemoteRepository {
						return err
					}
				} else {
					repository, err := git.PlainOpen(repositoryDirectory)
					if err != nil {
						return err
					}
					worktree, err := repository.Worktree()
					if err != nil {
						return err
					}
					pullOptions := &git.PullOptions{
						RemoteName: "origin",
						Progress:   os.Stdout,
					}
					err = worktree.Pull(pullOptions)
					if err != nil && err != git.NoErrAlreadyUpToDate {
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
