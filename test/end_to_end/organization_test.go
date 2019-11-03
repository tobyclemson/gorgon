package end_to_end

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tobyclemson/gorgon/test/support"
	"github.com/tobyclemson/gorgon/test/support/command"
	"github.com/tobyclemson/gorgon/test/support/fs"
	"github.com/tobyclemson/gorgon/test/support/git"
	"github.com/tobyclemson/gorgon/test/support/github"
	"testing"
)

func TestOrganizationListReposCommand(t *testing.T) {
	token := github.GetToken(t)
	organization := "infrablocks"

	expectedRepos :=
		github.ListOrganizationRepositories(t, organization, token)
	expectedRepoNames := github.ToSortedRepositoryNames(expectedRepos)

	_, stdout, _ := command.Run(t,
		"gorgon", "organization", "list-repos", "-t", token, organization)

	commandOutput := command.OutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		"Listing repositories for organization: 'infrablocks'")
	assert.Len(t, commandOutput.Body, len(expectedRepos))
	assert.Equal(t, commandOutput.Body, expectedRepoNames)
}

func TestOrganizationSyncReposCommandForFreshDirectory(t *testing.T) {
	token := github.GetToken(t)
	organization := "infrablocks"
	directory := fs.CreateTemporaryWorkDirectory(t)

	expectedRepos :=
		github.ListOrganizationRepositories(t, organization, token)

	_, stdout, _ := command.Run(t,
		"gorgon", "organization", "sync-repos",
		"-t", token,
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

func TestOrganizationSyncReposCommandForPopulatedDirectory(t *testing.T) {
	token := github.GetToken(t)
	organization := "infrablocks"
	directory := fs.CreateTemporaryWorkDirectory(t)

	expectedRepos :=
		github.ListOrganizationRepositories(t, organization, token)

	for i := 0; i < 5; i++ {
		repo := expectedRepos[i]
		git.CloneRepository(t, directory, *repo.Name, *repo.CloneURL)
	}

	_, stdout, _ := command.Run(t,
		"gorgon", "organization", "sync-repos",
		"-t", token,
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

