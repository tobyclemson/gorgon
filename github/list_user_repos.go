package github

import (
	"context"
	"github.com/google/go-github/v28/github"
	"sort"
)

func ListUserRepositories(
	user string, credentials Credentials) ([]string, error) {
	ctx := context.Background()
	githubClient := newClient(ctx, credentials)

	repositoryListOptions := &github.RepositoryListOptions{}

	var allRepos []*github.Repository
	for {
		repos, resp, err := githubClient.Repositories.List(
			ctx, user, repositoryListOptions)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		repositoryListOptions.Page = resp.NextPage
	}

	repoNames := ToRepoName(allRepos)
	sort.Strings(repoNames)

	return repoNames, nil
}
