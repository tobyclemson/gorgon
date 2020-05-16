package git

import (
	"github.com/go-git/go-git/v5/plumbing/transport"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/tobyclemson/gorgon/ssh"
)

func determineAuth(
	protocol Protocol,
	sshOptions ssh.Options,
) (transport.AuthMethod, error) {
	authenticationMethod := sshOptions.AuthenticationMethod
	username := "git"
	if authenticationMethod == ssh.Agent {
		return gitssh.NewSSHAgentAuth(username)
	} else {
		return gitssh.NewPublicKeysFromFile(
			username,
			sshOptions.PrivateKeyPath,
			sshOptions.PrivateKeyPassword)
	}
}
