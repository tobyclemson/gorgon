package github

import (
	"context"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

func NewClient(ctx context.Context, token string) *github.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})
	tokenClient := oauth2.NewClient(ctx, tokenSource)
	githubClient := github.NewClient(tokenClient)

	return githubClient
}

func GetDependencies(token string) (context.Context, *github.Client) {
	ctx := context.Background()
	client := NewClient(ctx, token)

	return ctx, client
}
