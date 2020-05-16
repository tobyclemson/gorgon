package repos

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/config"
	"github.com/tobyclemson/gorgon/git"
	"github.com/tobyclemson/gorgon/github"
	"github.com/tobyclemson/gorgon/ssh"
	"os"
	"path/filepath"
)

var userReposSyncCommand = &cobra.Command{
	Use:   "sync",
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

		protocol, err := config.Protocol(cmd)
		if err != nil {
			return err
		}

		sshPrivateKeyPath, err := config.SSHPrivateKeyPath(cmd)
		if err != nil {
			return err
		}

		targetDirectory, err := config.TargetDirectory(cmd, name)
		if err != nil {
			return err
		}

		credentials := github.Credentials{Token: token}
		sshOptions := ssh.Options{PrivateKeyPath: sshPrivateKeyPath}

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
					err = git.Clone(
						repository,
						repositoryDirectory,
						protocol,
						sshOptions)
					if err != nil {
						return err
					}
				} else {
					err = git.Pull(
						repositoryDirectory,
						sshOptions)
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
	userReposSyncCommand.Flags().
		StringP(
			"target-directory",
			"d",
			"",
			"directory into which repositories should be synced, defaults "+
				"to a directory under the working directory with the name of "+
				"the user")
	userReposSyncCommand.Flags().
		StringP(
			"protocol",
			"p",
			"ssh",
			"protocol to use when cloning repositories, one of "+
				"\"ssh\", \"git\" or \"https\"")
	userReposSyncCommand.Flags().
		StringP("ssh-private-key-path",
			"i",
			"~/.ssh/id_rsa",
			"path to SSH private key to use when using SSH private "+
				"key authentication, defaults to \"~/.ssh/id_rsa\"")
}
