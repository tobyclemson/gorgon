package github

import (
	"github.com/google/go-github/v31/github"
	"testing"
)

func ListOrganizationRepositories(
	t *testing.T,
	organization string,
	token string,
) []*github.Repository {
	ctx, client := GetDependencies(token)

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

func ListUserRepositories(
	t *testing.T,
	user string,
	token string,
) []*github.Repository {
	ctx, client := GetDependencies(token)

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
