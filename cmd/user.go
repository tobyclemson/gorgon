package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/cmd/user"
)

var userCommand = &cobra.Command{
	Use:   "user",
	Short: "Manage GitHub users",
	Long:  "Perform various actions for a given GitHub user.",
}

func init() {
	user.AddSubCommands(userCommand)
}
