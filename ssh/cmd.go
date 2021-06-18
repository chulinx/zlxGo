package ssh

import (
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
	config.User = c.user
	config.Auth = []ssh.AuthMethod{ssh.Password(c.pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }
	client, err := ssh.Dial("tcp", c.addr, config)
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
	buf, err := session.CombinedOutput("sh -c "+shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}
