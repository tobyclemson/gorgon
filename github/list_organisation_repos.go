package github

import (
	"context"
	"github.com/google/go-github/v28/github"
)

func ListOrganisationRepositories(
	organisation string, credentials Credentials) ([]string, error) {
	ctx := context.Background()
	githubClient := newClient(ctx, credentials)

	repos, _, err := githubClient.Repositories.ListByOrg(
		ctx,
		organisation,
		&github.RepositoryListByOrgOptions{})
	if err != nil {
		return nil, err
	}

	repoNames := ToRepoName(repos)

	return repoNames, nil
}
