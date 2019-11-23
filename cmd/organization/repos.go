package organization

import (
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/cmd/organization/repos"
)

var organizationReposCommand = &cobra.Command{
	Use:   "repos",
	Short: "Manage repos for a GitHub organizations",
	Long:  "List, check status of and sync repos for a GitHub organization.",
}

func init() {
	repos.AddSubCommands(organizationReposCommand)
}
