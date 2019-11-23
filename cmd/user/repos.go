package user

import (
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/cmd/user/repos"
)

var userReposCommand = &cobra.Command{
	Use:   "repos",
	Short: "Manage repos for a GitHub user",
	Long:  "List, check status of and sync repos for a GitHub user.",
}

func init() {
	repos.AddSubCommands(userReposCommand)
}
