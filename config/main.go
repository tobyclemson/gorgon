package config

import (
	"github.com/spf13/cobra"
	"github.com/tobyclemson/gorgon/git"
	"path/filepath"
)

func ensureAbsolute(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	} else {
		return filepath.Abs(path)
	}
}

func GithubToken(cmd *cobra.Command) (string, error) {
	return cmd.Flag("github-token").Value.String(), nil
}

func Protocol(cmd *cobra.Command) (git.Protocol, error) {
	return git.ToProtocol(cmd.Flag("protocol").Value.String())
}

func SSHPrivateKeyPath(cmd *cobra.Command) (string, error) {
	return ensureAbsolute(
		cmd.Flag("ssh-private-key-path").Value.String())
}

func TargetDirectory(cmd *cobra.Command, or string) (string, error) {
	targetDirectory := cmd.Flag("target-directory").Value.String()
	resolvedDirectory := ""
	if targetDirectory == "" {
		resolvedDirectory = or
	} else {
		resolvedDirectory = targetDirectory
	}
	return ensureAbsolute(resolvedDirectory)
}
