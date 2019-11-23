package git

import (
	"errors"
)

type Protocol int

const (
	SSH Protocol = iota + 1
	HTTPS
	Git
)

var protocolNames = [...]string{
	"ssh",
	"https",
	"git",
}

func (protocol Protocol) String() string {
	if protocol < SSH || protocol > Git {
		return "unknown"
	}

	return protocolNames[protocol-1]
}

func ToProtocol(value string) (Protocol, error) {
	for _, protocol := range [...]Protocol{SSH, HTTPS, Git} {
		if protocol.String() == value {
			return protocol, nil
		}
	}
	return -1, errors.New("Unknown protocol name: " + value)
}
