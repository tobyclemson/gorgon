package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestOrganizationListReposCommand(t *testing.T) {
	token, found := os.LookupEnv("TEST_GITHUB_TOKEN")
	assert.True(t, found)

	organization := "infrablocks"

	_, stdout, _ := runCommand(t,
		"gorgon", "organization", "list-repos", "-t", token, organization)

	expectedRepos := listOrganizationRepositories(t, organization, token)
	sort.Strings(expectedRepos)

	outputLines := strings.Split(stdout.String(), "\n")
	fmt.Println(outputLines)
	fmt.Println(len(outputLines))
	introLine := outputLines[0]
	repoLines := outputLines[2 : len(outputLines)-2]

	assert.Equal(t, introLine, "Repositories for organization: infrablocks")
	assert.Len(t, repoLines, len(expectedRepos))
	assert.Equal(t, repoLines, expectedRepos)
}

func TestUserListReposCommand(t *testing.T) {
	token, found := os.LookupEnv("TEST_GITHUB_TOKEN")
	assert.True(t, found)

	user := "tobyclemson"

	_, stdout, _ := runCommand(t,
		"gorgon", "user", "list-repos", "-t", token, user)

	expectedRepos := listUserRepositories(t, user, token)
	sort.Strings(expectedRepos)

	outputLines := strings.Split(stdout.String(), "\n")
	introLine := outputLines[0]
	repoLines := outputLines[2 : len(outputLines)-2]

	assert.Equal(t, introLine, "Repositories for user: tobyclemson")
	assert.Len(t, repoLines, len(expectedRepos))
	assert.Equal(t, repoLines, expectedRepos)
}
