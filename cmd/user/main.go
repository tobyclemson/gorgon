package user

import "github.com/spf13/cobra"

func AddSubCommands(parentCommand *cobra.Command) {
	parentCommand.AddCommand(userListReposCommand)
	parentCommand.AddCommand(userSyncReposCommand)
}
