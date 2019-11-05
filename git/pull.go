package git

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"os"
)

func Pull(repositoryDirectory string) error {
	auth, err := ssh.DefaultAuthBuilder("git")
	if err != nil {
		return err
	}
	repository, err := git.PlainOpen(repositoryDirectory)
	if err != nil {
		return err
	}
	worktree, err := repository.Worktree()
	if err != nil {
		if err == git.ErrIsBareRepository {
			fmt.Println("Repository is empty.")
			fmt.Println()
			return nil
		} else {
			fmt.Println("Got error:", err)
			return err
		}
	}
	pullOptions := &git.PullOptions{
		Auth:       auth,
		RemoteName: "origin",
		Progress:   os.Stdout,
	}
	err = worktree.Pull(pullOptions)
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Println("Already up to date")
			fmt.Println()
			return nil
		} else {
			fmt.Println("Got error:", err)
			return err
		}
	}

	return nil
}
