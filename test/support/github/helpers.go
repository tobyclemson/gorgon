package github

import (
	"github.com/google/go-github/v28/github"
	"sort"
)

func ToSortedRepositoryNames(repositories []*github.Repository) []string {
	var repositoryNames []string
	for _, repository := range repositories {
		repositoryNames = append(repositoryNames, *repository.Name)
	}
	sort.Strings(repositoryNames)

	return repositoryNames
}

