package ssh

import (
	"fmt"
)

// sAuth ssh config
type sAuth struct {
	user       string
	privateKey string
	pass       string
	addr       string // address format ip:port
}

func NewAuthPass(user, pass, addr string) *sAuth {
	return &sAuth{
		user: user,
		pass: pass,
		addr: addr,
	}
}

func NewAuthPrivateKey(user, privateKey, addr string) *sAuth {
	return &sAuth{
		user:       user,
		privateKey: privateKey,
		addr:       addr,
	}
}

func (c Client) Run(shell string) (string, error) {
	if c.client == nil {
		if _, err := c.Connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	cmd := fmt.Sprintf("sh -c \"%s\"", shell)
	buf, err := session.CombinedOutput(cmd)

	c.lastResult = string(buf)
	return c.lastResult, err
}
