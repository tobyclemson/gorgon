package config

import (
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/git"
	"path/filepath"
)

func GithubToken(cmd *cobra.Command) (string, error) {
	return cmd.Flag("github-token").Value.String(), nil
}

func Protocol(cmd *cobra.Command) (git.Protocol, error) {
	return git.ToProtocol(cmd.Flag("protocol").Value.String())
}

func TargetDirectory(cmd *cobra.Command, or string) (string, error) {
	err := error(nil)
	targetDirectory := cmd.Flag("target-directory").Value.String()
	resolvedDirectory := ""
	if targetDirectory == "" {
		resolvedDirectory = or
	} else {
		resolvedDirectory = targetDirectory
	}
	if !filepath.IsAbs(resolvedDirectory) {
		resolvedDirectory, err = filepath.Abs(resolvedDirectory)
		if err != nil {
			return "", err
		}
	}
	return resolvedDirectory, nil
}