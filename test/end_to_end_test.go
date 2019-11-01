package test

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestOrganizationListReposCommand(t *testing.T) {
	token := getGithubToken(t)
	organization := "infrablocks"

	_, stdout, _ := runCommand(t,
		"gorgon", "organization", "list-repos", "-t", token, organization)

	expectedRepos := listOrganizationRepositories(t, organization, token)
	sort.Strings(expectedRepos)

	outputLines := strings.Split(stdout.String(), "\n")
	introLine := outputLines[0]
	repoLines := outputLines[2 : len(outputLines)-2]

	assert.Equal(t, introLine,
		"Listing repositories for organization: 'infrablocks'")
	assert.Len(t, repoLines, len(expectedRepos))
	assert.Equal(t, repoLines, expectedRepos)
}

func TestOrganizationSyncReposCommand(t *testing.T) {
	token := getGithubToken(t)
	organization := "infrablocks"
	directory := createTemporaryWorkDirectory(t)

	_, stdout, _ := runCommand(t,
		"gorgon", "organization", "sync-repos",
		"-t", token,
		"-d", directory,
		organization)

	expectedRepos := listOrganizationRepositories(t, organization, token)
	actualRepos := listDirectories(t, directory)

	expectedRepoSet := mapset.NewSet()
	for _, expectedRepo := range expectedRepos {
		expectedRepoSet.Add(expectedRepo)
	}
	actualRepoSet := mapset.NewSet()
	for _, actualRepo := range actualRepos {
		actualRepoSet.Add(actualRepo)
	}

	outputLines := strings.Split(stdout.String(), "\n")
	introLine := outputLines[0]
	repoLines := outputLines[2 : len(outputLines)-2]

	assert.Equal(t, introLine,
		fmt.Sprintf(
			"Syncing repositories for organization: '%v' into directory: '%v'",
			organization,
			directory))
	assert.Equal(t, expectedRepoSet, actualRepoSet)
	assert.Subset(t, repoLines, expectedRepos)
}

func TestUserListReposCommand(t *testing.T) {
	token := getGithubToken(t)

	user := "tobyclemson"

	_, stdout, _ := runCommand(t,
		"gorgon", "user", "list-repos", "-t", token, user)

	expectedRepos := listUserRepositories(t, user, token)
	sort.Strings(expectedRepos)

	outputLines := strings.Split(stdout.String(), "\n")
	introLine := outputLines[0]
	repoLines := outputLines[2 : len(outputLines)-2]

	assert.Equal(t, introLine,
		"Listing repositories for user: 'tobyclemson'")
	assert.Len(t, repoLines, len(expectedRepos))
	assert.Equal(t, repoLines, expectedRepos)
}
