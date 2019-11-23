package repos

import (
	"github.com/spf13/cobra"
)

func AddSubCommands(parentCommand *cobra.Command) {
	parentCommand.AddCommand(userReposListCommand)
	parentCommand.AddCommand(userReposSyncCommand)
}
