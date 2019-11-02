package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"path/filepath"
	"testing"
)

func TestOrganizationListReposCommand(t *testing.T) {
	token := getGithubToken(t)
	organization := "infrablocks"

	expectedRepos := listOrganizationRepositories(t, organization, token)
	expectedRepoNames := toSortedNames(expectedRepos)

	_, stdout, _ := runCommand(t,
		"gorgon", "organization", "list-repos", "-t", token, organization)

	commandOutput := commandOutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		"Listing repositories for organization: 'infrablocks'")
	assert.Len(t, commandOutput.Body, len(expectedRepos))
	assert.Equal(t, commandOutput.Body, expectedRepoNames)
}

func TestOrganizationSyncReposCommandForFreshDirectory(t *testing.T) {
	token := getGithubToken(t)
	organization := "infrablocks"
	directory := createTemporaryWorkDirectory(t)

	expectedRepos := listOrganizationRepositories(t, organization, token)

	_, stdout, _ := runCommand(t,
		"gorgon", "organization", "sync-repos",
		"-t", token,
		"-d", directory,
		organization)

	expectedRepoNames := toSortedNames(expectedRepos)
	actualRepoNames := listDirectories(t, directory)

	commandOutput := commandOutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		fmt.Sprintf(
			"Syncing repositories for organization: '%v' into directory: '%v'",
			organization,
			directory))
	assert.Equal(t,
		newStringSet(expectedRepoNames),
		newStringSet(actualRepoNames))
	assert.Subset(t, commandOutput.Body, expectedRepoNames)
}

func TestOrganizationSyncReposCommandForPopulatedDirectory(t *testing.T) {
	token := getGithubToken(t)
	organization := "infrablocks"
	directory := createTemporaryWorkDirectory(t)

	expectedRepos := listOrganizationRepositories(t, organization, token)

	for i := 0; i < 5; i++ {
		expectedRepo := expectedRepos[i]
		expectedRepoPath := filepath.Join(directory, *expectedRepo.Name)
		_, err := git.PlainClone(expectedRepoPath, false, &git.CloneOptions{
			URL: *expectedRepo.CloneURL,
		})
		if err != nil {
			t.Fatalf("Failed to clone repository: '%v': %v",
				expectedRepo.Name,
				err)
		}
	}

	_, stdout, _ := runCommand(t,
		"gorgon", "organization", "sync-repos",
		"-t", token,
		"-d", directory,
		organization)

	expectedRepoNames := toSortedNames(expectedRepos)
	actualRepoNames := listDirectories(t, directory)

	commandOutput := commandOutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		fmt.Sprintf(
			"Syncing repositories for organization: '%v' into directory: '%v'",
			organization,
			directory))
	assert.Equal(t,
		newStringSet(expectedRepoNames),
		newStringSet(actualRepoNames))
	assert.Subset(t, commandOutput.Body, expectedRepoNames)
}

func TestUserListReposCommand(t *testing.T) {
	token := getGithubToken(t)

	user := "tobyclemson"

	_, stdout, _ := runCommand(t,
		"gorgon", "user", "list-repos", "-t", token, user)

	expectedRepos := listUserRepositories(t, user, token)
	expectedRepoNames := toSortedNames(expectedRepos)

	commandOutput := commandOutputFrom(stdout)

	assert.Equal(t, commandOutput.Header,
		"Listing repositories for user: 'tobyclemson'")
	assert.Len(t, commandOutput.Body, len(expectedRepos))
	assert.Equal(t, commandOutput.Body, expectedRepoNames)
}
