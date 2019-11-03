package github

import "github.com/google/go-github/v28/github"

func ToRepoName(repos []*github.Repository) []string {
	var names []string
	for _, repo := range repos {
		names = append(names, *repo.Name)
	}
	return names
}
