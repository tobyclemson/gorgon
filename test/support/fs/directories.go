package fs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func GetWorkingDirectory(t *testing.T) string {
	workingDirectory, err := os.Getwd()
	assert.Nil(t, err)

	return workingDirectory
}

func ListDirectories(t *testing.T, directory string) []string {
	fileInfos, err := ioutil.ReadDir(directory)
	if err != nil {
		t.Fatal("Failed to list directory:", directory)
	}
	var directories []string
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			directories = append(directories, fileInfo.Name())
		}
	}

	return directories
}

func CreateTemporaryWorkDirectory(t *testing.T) string {
	workingDirectory := GetWorkingDirectory(t)
	temporaryDirectory, err := ioutil.TempDir(
		filepath.Join(workingDirectory, "..", "..", "work"),
		"test-")
	if err != nil {
		t.Fatalf("Failed to create work directory '%v': %v",
			temporaryDirectory,
			err)
	}

	return temporaryDirectory
}
