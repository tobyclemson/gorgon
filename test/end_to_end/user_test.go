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

func TestUserListReposCommand(t *testing.T) {
	token := github.GetToken(t)
	binary := support.GetBinaryPath(t)
	user := "tobyclemson"

	err := os.Setenv("GORGON_GITHUB_TOKEN", token)
	assert.Nil(t, err)

	_, stdout, _ := command.Run(t,
		binary, "user", "list-repos", user)

	expectedRepos := github.ListUserRepositories(t, user, token)
	expectedRepoNames := github.ToSortedRepositoryNames(expectedRepos)

	commandOutput := command.OutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		"Listing repositories for user: 'tobyclemson'")
	assert.Len(t, commandOutput.Body, len(expectedRepos))
	assert.Equal(t, commandOutput.Body, expectedRepoNames)
}

func TestUserSyncReposCommandForFreshDirectory(t *testing.T) {
	token := github.GetToken(t)
	binary := support.GetBinaryPath(t)
	user := "tobyclemson"
	directory := fs.CreateTemporaryWorkDirectory(t)

	err := os.Setenv("GORGON_GITHUB_TOKEN", token)
	assert.Nil(t, err)

	expectedRepos := github.ListUserRepositories(t, user, token)

	_, stdout, _ := command.Run(t,
		binary, "user", "sync-repos",
		"-d", directory,
		user)

	expectedRepoNames := github.ToSortedRepositoryNames(expectedRepos)
	actualRepoNames := fs.ListDirectories(t, directory)

	commandOutput := command.OutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		fmt.Sprintf(
			"Syncing repositories for user: '%v' into directory: '%v'",
			user,
			directory))
	assert.Equal(t,
		support.NewStringSet(expectedRepoNames),
		support.NewStringSet(actualRepoNames))
	assert.Subset(t, commandOutput.Body, expectedRepoNames)
}

func TestUserSyncReposCommandForPopulatedDirectory(t *testing.T) {
	token := github.GetToken(t)
	binary := support.GetBinaryPath(t)
	user := "tobyclemson"
	directory := fs.CreateTemporaryWorkDirectory(t)

	err := os.Setenv("GORGON_GITHUB_TOKEN", token)
	assert.Nil(t, err)

	expectedRepos := github.ListUserRepositories(t, user, token)

	for i := 0; i < 3; i++ {
		repo := expectedRepos[i]
		git.CloneRepository(t, directory, *repo.Name, *repo.SSHURL)
	}

	_, stdout, _ := command.Run(t,
		binary, "user", "sync-repos",
		"-d", directory,
		user)

	expectedRepoNames := github.ToSortedRepositoryNames(expectedRepos)
	actualRepoNames := fs.ListDirectories(t, directory)

	commandOutput := command.OutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		fmt.Sprintf(
			"Syncing repositories for user: '%v' into directory: '%v'",
			user,
			directory))
	assert.Equal(t,
		support.NewStringSet(expectedRepoNames),
		support.NewStringSet(actualRepoNames))
	assert.Subset(t, commandOutput.Body, expectedRepoNames)
}
