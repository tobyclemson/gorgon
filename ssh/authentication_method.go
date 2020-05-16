package ssh

import (
	"errors"
)

type AuthenticationMethod int

const (
	Agent AuthenticationMethod = iota + 1
	Key
)

var authenticationMethodNames = [...]string{
	"agent",
	"key",
}

func (authenticationMethod AuthenticationMethod) String() string {
	if authenticationMethod < Agent || authenticationMethod > Key {
		return "unknown"
	}

	return authenticationMethodNames[authenticationMethod-1]
}

func ToAuthenticationMethod(value string) (AuthenticationMethod, error) {
	for _, authenticationMethod := range [...]AuthenticationMethod{Agent, Key} {
		if authenticationMethod.String() == value {
			return authenticationMethod, nil
		}
	}
	return -1, errors.New("Unknown authentication method name: " + value)
}
