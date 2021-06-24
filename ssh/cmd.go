package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

type Client struct {
	User    string
	Pwd     string
	Addr    string
	client  *ssh.Client
	session *ssh.Session
	LastResult string
}

func (c *Client) Connect() (*Client, error) {
	config := &ssh.ClientConfig{}
	config.SetDefaults()
	config.User = c.User
	config.Auth = []ssh.AuthMethod{ssh.Password(c.Pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }
	client, err := ssh.Dial("tcp", c.Addr, config)
	if nil != err {
		return c, err
	}
	c.client = client
	return c, nil
}

func (c Client) Run(shell string) (string, error) {
	if c.client == nil {
		if _,err := c.Connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	cmd := fmt.Sprintf("sh -c \"%s\"",shell)
	buf, err := session.CombinedOutput(cmd)

	c.LastResult = string(buf)
	return c.LastResult, err
}
