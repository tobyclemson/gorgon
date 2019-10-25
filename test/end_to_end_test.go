package test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestOrganizationListReposCommand(t *testing.T) {
	workingDirectory, err := os.Getwd()
	assert.Nil(t, err)

	token, found := os.LookupEnv("TEST_GITHUB_TOKEN")
	assert.True(t, found)

	cmd := exec.Command(
		fmt.Sprint(workingDirectory, "/../gorgon"),
		"organization",
		"list-repos",
		"-t", token,
		"infrablocks")
	var outputBuffer, errorBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer

	err = cmd.Run()
	assert.Nil(t, err)

	assert.Contains(t, outputBuffer.String(),
		"Repositories for organization: infrablocks")
}

func TestUserListReposCommand(t *testing.T) {
	workingDirectory, err := os.Getwd()
	assert.Nil(t, err)

	token, found := os.LookupEnv("TEST_GITHUB_TOKEN")
	assert.True(t, found)

	cmd := exec.Command(
		fmt.Sprint(workingDirectory, "/../gorgon"),
		"user",
		"list-repos",
		"-t", token,
		"tobyclemson")
	var outputBuffer, errorBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer

	err = cmd.Run()
	assert.Nil(t, err)

	assert.Contains(t, outputBuffer.String(),
		"Repositories for user: tobyclemson")
}
