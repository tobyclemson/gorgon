package support

import (
	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func NewStringSet(items []string) *mapset.Set {
	set := mapset.NewSet()
	for _, item := range items {
		set.Add(item)
	}
	return &set
}

func GetBinaryPath(t *testing.T) string {
	binary, found := os.LookupEnv("TEST_BINARY_PATH")
	assert.True(t, found)

	return binary
}

func GetSSHPrivateKeyPath(t *testing.T) string {
	token, found := os.LookupEnv("TEST_SSH_PRIVATE_KEY_PATH")
	assert.True(t, found)

	return token
}
