package github

import (
	"github.com/pkg/errors"
)

func ListRepositories(
	entity Entity, name string, credentials Credentials) ([]string, error) {
	switch entity {
	case Organization:
		return ListOrganisationRepositories(name, credentials)
	case User:
		return ListUserRepositories(name, credentials)
	default:
		return nil, errors.Errorf(
			"Unknown entity type: %w", entity.String())
	}
}
