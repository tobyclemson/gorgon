package github

import "github.com/google/go-github/v28/github"

func ToRepoName(repos []*github.Repository) []string {
	var names []string
	for _, org := range repos {
		names = append(names, *org.Name)
	}
	return names
}

