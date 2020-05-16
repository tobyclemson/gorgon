package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"path/filepath"
	"testing"
)

func CloneRepository(
	t *testing.T,
	directory string,
	name string,
	url string,
	sshPrivateKeyPath string) {
	auth, err := ssh.NewPublicKeysFromFile("git", sshPrivateKeyPath, "")
	if err != nil {
		t.Fatalf("Failed to build SSH authentication: %v", err)
	}
	path := filepath.Join(directory, name)
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		Auth: auth,
		URL:  url,
	})
	if err != nil {
		t.Fatalf("Failed to clone repository: '%v': %v", name, err)
	}
}
