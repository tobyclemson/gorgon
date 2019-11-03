package support

import (
	"github.com/deckarep/golang-set"
)

func NewStringSet(items []string) *mapset.Set {
	set := mapset.NewSet()
	for _, item := range items {
		set.Add(item)
	}
	return &set
}
