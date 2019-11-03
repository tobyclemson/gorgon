package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show gorgon version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootCommand.Use + " " + Version)
	},
}
