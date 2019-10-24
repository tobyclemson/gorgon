package github

import (
	"github.com/pkg/errors"
)

type Entity int

const (
	Unknown Entity = iota
	Organization
	User
)

var all = []Entity{Organization, User}
var names = []string{"organization", "user"}

func (e Entity) Is(other Entity) bool {
	return e == other
}

func (e Entity) String() string {
	if e < Organization || e > User {
		return "unknown"
	}
	return names[e-1]
}

func EntityFromString(entity string) (Entity, error) {
	for _, target := range all {
		if target.String() == entity {
			return target, nil
		}
	}
	return Unknown, errors.Errorf(
		"No entity matches: %w", entity)
}
