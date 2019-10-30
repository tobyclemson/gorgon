package organization

import "github.com/spf13/cobra"

func AddSubCommands(parentCommand *cobra.Command) {
	parentCommand.AddCommand(organizationListReposCommand)
	parentCommand.AddCommand(organizationSyncReposCommand)
}
