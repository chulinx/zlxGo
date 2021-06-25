package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	User    string
	Pwd     string
	Addr    string	// address format ip:port
	client  *ssh.Client
	session *ssh.Session
	lastResult string
}

func NewSSHClient(user,pass,addr string) *Client {
	addrList := strings.Split(addr,":")
	if len(addrList) < 2 {
		panic("addr format ip:port")
	}
	_,err := strconv.Atoi(addrList[1])
	if err != nil {
		panic("addr format ip:port")
	}
	return &Client{
		User: user,
		Pwd: pass,
		Addr: addr,
	}
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

	c.lastResult = string(buf)
	return c.lastResult, err
}
