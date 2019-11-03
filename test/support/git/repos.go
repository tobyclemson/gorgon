package git

import (
	"gopkg.in/src-d/go-git.v4"
	"path/filepath"
	"testing"
)

func CloneRepository(t *testing.T, directory string, name string, url string) {
	path := filepath.Join(directory, name)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		t.Fatalf("Failed to clone repository: '%v': %v", name, err)
	}
}
