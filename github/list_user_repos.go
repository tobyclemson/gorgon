package github

import (
	"context"
	"github.com/google/go-github/v28/github"
	"sort"
)

func ListUserRepositories(
	user string,
	credentials Credentials,
) ([]*github.Repository, error) {
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

	sort.Slice(allRepos, func(i, j int) bool {
		return *allRepos[i].Name < *allRepos[j].Name
	})

	return allRepos, nil
}
