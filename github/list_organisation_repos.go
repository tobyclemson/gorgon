package github

import (
	"context"
	"github.com/google/go-github/v28/github"
	"sort"
)

func ListOrganisationRepositories(
	organisation string, credentials Credentials) ([]string, error) {
	ctx := context.Background()
	githubClient := newClient(ctx, credentials)

	repositoryListOptions := &github.RepositoryListByOrgOptions{}

	var allRepos []*github.Repository
	for {
		repos, resp, err := githubClient.Repositories.ListByOrg(
			ctx, organisation, repositoryListOptions)
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
