package github

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func GetToken(t *testing.T) string {
	token, found := os.LookupEnv("TEST_GITHUB_TOKEN")
	assert.True(t, found)

	return token
}
