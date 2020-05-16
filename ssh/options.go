package ssh

type Options struct {
	AuthenticationMethod AuthenticationMethod
	PrivateKeyPath string
	PrivateKeyPassword string
}
