package end_to_end

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tobyclemson/gorgon/test/support"
	"github.com/tobyclemson/gorgon/test/support/command"
	"github.com/tobyclemson/gorgon/test/support/fs"
	"github.com/tobyclemson/gorgon/test/support/git"
	"github.com/tobyclemson/gorgon/test/support/github"
	"os"
	"testing"
)

func TestOrganizationReposListCommand(t *testing.T) {
	token := github.GetToken(t)
	binary := support.GetBinaryPath(t)
	organization := "javafunk"

	err := os.Setenv("GORGON_GITHUB_TOKEN", token)
	assert.Nil(t, err)

	expectedRepos :=
		github.ListOrganizationRepositories(t, organization, token)
	expectedRepoNames := github.ToSortedRepositoryNames(expectedRepos)

	_, stdout, _, err := command.Run(t,
		binary, "organization", "repos", "list", organization)
	assert.Nil(t, err)

	commandOutput := command.OutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		"Listing repositories for organization: 'javafunk'")
	assert.Len(t, commandOutput.Body, len(expectedRepos))
	assert.Equal(t, commandOutput.Body, expectedRepoNames)
}

func TestOrganizationSyncReposCommandForFreshDirectory(t *testing.T) {
	binary := support.GetBinaryPath(t)
	sshPrivateKeyPath := support.GetSSHPrivateKeyPath(t)
	token := github.GetToken(t)
	organization := "javafunk"
	directory := fs.CreateTemporaryWorkDirectory(t)

	err := os.Setenv(
		"GORGON_GITHUB_TOKEN", token)
	assert.Nil(t, err)
	err = os.Setenv(
		"GORGON_SSH_PRIVATE_KEY_PATH", sshPrivateKeyPath)
	assert.Nil(t, err)

	expectedRepos :=
		github.ListOrganizationRepositories(t, organization, token)

	_, stdout, _, err := command.Run(t,
		binary, "organization", "repos", "sync",
		"-d", directory,
		organization)

	expectedRepoNames := github.ToSortedRepositoryNames(expectedRepos)
	actualRepoNames := fs.ListDirectories(t, directory)

	commandOutput := command.OutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		fmt.Sprintf(
			"Syncing repositories for organization: '%v' into directory: '%v'",
			organization,
			directory))
	assert.Equal(t,
		support.NewStringSet(expectedRepoNames),
		support.NewStringSet(actualRepoNames))
	assert.Subset(t, commandOutput.Body, expectedRepoNames)
}

func TestOrganizationReposSyncCommandForPopulatedDirectory(t *testing.T) {
	binary := support.GetBinaryPath(t)
	sshPrivateKeyPath := support.GetSSHPrivateKeyPath(t)
	token := github.GetToken(t)
	organization := "javafunk"
	directory := fs.CreateTemporaryWorkDirectory(t)

	err := os.Setenv(
		"GORGON_GITHUB_TOKEN", token)
	assert.Nil(t, err)
	err = os.Setenv(
		"GORGON_SSH_PRIVATE_KEY_PATH", sshPrivateKeyPath)
	assert.Nil(t, err)

	expectedRepos :=
		github.ListOrganizationRepositories(t, organization, token)

	for i := 0; i < 3; i++ {
		repo := expectedRepos[i]
		git.CloneRepository(t,
			directory,
			*repo.Name,
			*repo.SSHURL,
			sshPrivateKeyPath)
	}

	_, stdout, _, err := command.Run(t,
		binary, "organization", "repos", "sync",
		"-d", directory,
		organization)
	assert.Nil(t, err)

	expectedRepoNames := github.ToSortedRepositoryNames(expectedRepos)
	actualRepoNames := fs.ListDirectories(t, directory)

	commandOutput := command.OutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		fmt.Sprintf(
			"Syncing repositories for organization: '%v' into directory: '%v'",
			organization,
			directory))
	assert.Equal(t,
		support.NewStringSet(expectedRepoNames),
		support.NewStringSet(actualRepoNames))
	assert.Subset(t, commandOutput.Body, expectedRepoNames)
}
