package git

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/google/go-github/v31/github"
	"os"
)

func determineURL(
	protocol Protocol,
	repository *github.Repository,
) (string, error) {
	switch protocol {
	case SSH:
		return *repository.SSHURL, nil
	case HTTPS:
		return *repository.CloneURL, nil
	case Git:
		return *repository.GitURL, nil
	}
	return "", errors.New("Unknown protocol: " + protocol.String())
}

func determineAuth(
	protocol Protocol,
) (transport.AuthMethod, error) {
	return ssh.DefaultAuthBuilder("git")
}

func Clone(
	repository *github.Repository,
	destination string,
	protocol Protocol,
) error {
	url, err := determineURL(protocol, repository)
	if err != nil {
		return err
	}
	auth, err := determineAuth(protocol)
	if err != nil {
		return err
	}

	cloneOptions := &git.CloneOptions{
		URL:      url,
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
