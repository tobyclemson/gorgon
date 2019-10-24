package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var rootCommand = &cobra.Command{
	Use:   "gorgon",
	Short: "Manage local clone of a GitHub organisation or user.",
	Long: "Gorgon clones and updates all repositories for a GitHub " +
		"organisation or user.",
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
			"The personal access or OAuth token used to " +
				"log in to GitHub")
	rootCommand.PersistentFlags().
		StringP(
			"github-type",
			"o",
			"organization",
			"The type of account to manage, one of 'user' or " +
				"'organization'")

	if err := viper.BindPFlag(
		"github-token",
		rootCommand.PersistentFlags().Lookup("github-token"));
		err != nil {
		panic(err)
	}
	if err := viper.BindPFlag(
		"github-type",
		rootCommand.PersistentFlags().Lookup("github-type"));
		err != nil {
		panic(err)
	}
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
