package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"os"
	"os/exec"
	"testing"
)

func newGithubClient(ctx context.Context, token string) *github.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})
	tokenClient := oauth2.NewClient(ctx, tokenSource)
	githubClient := github.NewClient(tokenClient)

	return githubClient
}

func newGithubDependencies(token string) (context.Context, *github.Client) {
	ctx := context.Background()
	client := newGithubClient(ctx, token)

	return ctx, client
}

func runCommand(
	t *testing.T,
	name string,
	args ...string,
) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	workingDirectory, err := os.Getwd()
	assert.Nil(t, err)

	binaryPath := fmt.Sprint(workingDirectory, "/../", name)

	fmt.Print("Executing command: ", binaryPath)
	for _, arg := range args {
		fmt.Print(" ", arg)
	}
	fmt.Println()

	cmd := exec.Command(binaryPath, args...)
	var outputBuffer, errorBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer

	err = cmd.Run()
	assert.Nil(t, err)

	return cmd, &outputBuffer, &errorBuffer
}

func toNames(repositories []*github.Repository) []string {
	var repositoryNames []string
	for _, repository := range repositories {
		repositoryNames = append(repositoryNames, *repository.Name)
	}

	return repositoryNames
}

func listOrganizationRepositories(
	t *testing.T,
	organization string,
	token string,
) []string {
	ctx, client := newGithubDependencies(token)

	options := &github.RepositoryListByOrgOptions{}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(
			ctx, organization, options)
		if err != nil {
			t.Fatalf(
				"Failed to fetch organization repositories for '%v': %v",
				organization,
				err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	return toNames(allRepos)
}

func listUserRepositories(
	t *testing.T,
	user string,
	token string,
) []string {
	ctx, client := newGithubDependencies(token)

	options := &github.RepositoryListOptions{}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(
			ctx, user, options)
		if err != nil {
			t.Fatalf(
				"Failed to fetch user repositories for '%v': %v",
				user,
				err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	return toNames(allRepos)
}
