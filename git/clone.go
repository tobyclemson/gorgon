package git

import (
	"fmt"
	"github.com/google/go-github/v28/github"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"os"
)

func Clone(repository *github.Repository, destination string) error {
	auth, err := ssh.DefaultAuthBuilder("git")
	if err != nil {
		return err
	}
	cloneOptions := &git.CloneOptions{
		URL:      *repository.SSHURL,
		Auth:     auth,
		Progress: os.Stdout,
	}
	_, err = git.PlainClone(destination, false, cloneOptions)
	if err != nil {
		if err == transport.ErrEmptyRemoteRepository {
			fmt.Println("Repository is empty.")
			fmt.Println()
			return nil
		} else {
			fmt.Println("Got error:", err)
			return err
		}
	}

	return nil
}
