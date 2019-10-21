package github

import (
	"context"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

func newClient(ctx context.Context, credentials Credentials) *github.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: credentials.Token},
	)
	tokenClient := oauth2.NewClient(ctx, tokenSource)
	githubClient := github.NewClient(tokenClient)

	return githubClient
}
