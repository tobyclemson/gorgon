package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/google/go-github/v28/github"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
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

func newStringSet(items []string) *mapset.Set {
	set := mapset.NewSet()
	for _, item := range items {
		set.Add(item)
	}
	return &set
}

func getWorkingDirectory(t *testing.T) string {
	workingDirectory, err := os.Getwd()
	assert.Nil(t, err)

	return workingDirectory
}

func listDirectories(t *testing.T, directory string) []string {
	fileInfos, err := ioutil.ReadDir(directory)
	if err != nil {
		t.Fatal("Failed to list directory:", directory)
	}
	var directories []string
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			directories = append(directories, fileInfo.Name())
		}
	}

	return directories
}

type commandOutput struct {
	Header string
	Body   []string
	Raw    []string
}

func commandOutputFrom(buffer *bytes.Buffer) *commandOutput {
	lines := strings.Split(buffer.String(), "\n")
	header := lines[0]
	body := lines[2 : len(lines)-2]
	return &commandOutput{
		Header: header,
		Body:   body,
		Raw:    lines,
	}
}

func logCommand(t *testing.T, title string, cmd *exec.Cmd) {
	t.Logf("%v: %v", title, cmd.String())
	t.Log()
}

func logLines(t *testing.T, title string, linesBuffer *bytes.Buffer) {
	t.Logf("%v:", title)
	t.Log()
	lines := strings.Split(linesBuffer.String(), "\n")
	for _, line := range lines {
		t.Log(line)
	}
}

func buildCommand(
	binary string,
	args ...string,
) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	cmd := exec.Command(binary, args...)
	var outputBuffer, errorBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer

	return cmd, &outputBuffer, &errorBuffer
}

func runCommand(
	t *testing.T,
	name string,
	args ...string,
) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	workingDirectory := getWorkingDirectory(t)
	binary := fmt.Sprint(workingDirectory, "/../", name)

	cmd, outputBuffer, errorBuffer :=
		buildCommand(binary, args...)
	logCommand(t, "Executing command", cmd)

	err := cmd.Run()
	logLines(t, "Standard Output", outputBuffer)
	logLines(t, "Standard Error", errorBuffer)
	assert.Nil(t, err)

	return cmd, outputBuffer, errorBuffer
}

func getGithubToken(t *testing.T) string {
	token, found := os.LookupEnv("TEST_GITHUB_TOKEN")
	assert.True(t, found)

	return token
}

func createTemporaryWorkDirectory(t *testing.T) string {
	workingDirectory := getWorkingDirectory(t)
	temporaryDirectory, err := ioutil.TempDir(
		fmt.Sprint(workingDirectory, "/../work"),
		"test-")
	if err != nil {
		t.Fatalf("Failed to create work directory '%v': %v",
			temporaryDirectory,
			err)
	}

	return temporaryDirectory
}

func toSortedNames(repositories []*github.Repository) []string {
	var repositoryNames []string
	for _, repository := range repositories {
		repositoryNames = append(repositoryNames, *repository.Name)
	}
	sort.Strings(repositoryNames)

	return repositoryNames
}

func listOrganizationRepositories(
	t *testing.T,
	organization string,
	token string,
) []*github.Repository {
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

	return allRepos
}

func listUserRepositories(
	t *testing.T,
	user string,
	token string,
) []*github.Repository {
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

	return allRepos
}
