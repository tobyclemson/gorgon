package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var rootCommand = &cobra.Command{
	Use:   "gorgon",
	Short: "Manage GitHub organizations or users.",
	Long: "Gorgon is a powerful command line tool to manage GitHub " +
		"organizations and users.",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
				if err := cmd.Flags().Set(f.Name, viper.GetString(f.Name)); err != nil {
					panic(err)
				}
			}
		})

		return nil
	},
	Version: "0.1",
}

func init() {
	viper.SetEnvPrefix("gorgon")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	rootCommand.PersistentFlags().
		StringP(
			"github-token",
			"t",
			"",
			"personal access or OAuth token used to "+
				"log in to GitHub")

	AddSubcommands(rootCommand)
}

func AddSubcommands(parentCommand *cobra.Command) {
	parentCommand.AddCommand(organizationCommand)
	parentCommand.AddCommand(userCommand)
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
