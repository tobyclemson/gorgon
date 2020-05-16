package git

import (
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func determineAuth(
	protocol Protocol,
) (transport.AuthMethod, error) {
	return ssh.DefaultAuthBuilder("git")
}
