package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/cmd/organization"
)

var organizationCommand = &cobra.Command{
	Use:   "organization",
	Short: "Manage GitHub organizations",
	Long:  "Perform various actions for a given GitHub organization.",
}

func AddOrganizationSubCommand(cmd *cobra.Command) {
	organizationCommand.AddCommand(cmd)
}

func init() {
	organization.AddSubCommands(organizationCommand)
}
