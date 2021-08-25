package ssh

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	*SAuth
	client     *ssh.Client
	session    *ssh.Session
	lastResult string
}

func NewSSHClient(auth *SAuth) *Client {
	addrList := strings.Split(auth.addr, ":")
	if len(addrList) < 2 {
		panic("addr format ip:port")
	}
	_, err := strconv.Atoi(addrList[1])
	if err != nil {
		panic("addr format ip:port")
	}
	return &Client{
		SAuth: &SAuth{
			user:       auth.user,
			pass:       auth.pass,
			addr:       auth.addr,
			privateKey: auth.privateKey,
		},
	}
}

func (c *Client) Connect() (*Client, error) {
	authMethod, err2 := c.auth()
	if err2 != nil {
		return nil, err2
	}
	config := &ssh.ClientConfig{
		User:            c.user,
		Auth:            []ssh.AuthMethod{authMethod},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil },
	}
	config.SetDefaults()
	client, err := ssh.Dial("tcp", c.addr, config)
	if nil != err {
		return nil, err
	}
	c.client = client
	return c, nil
}

func (c *Client) auth() (ssh.AuthMethod, error) {
	// Sync support pass and pubkey auth
	if c.pass != "" {
		return ssh.Password(c.pass), nil
	}
	if c.privateKey != "" {
		authData, err := ioutil.ReadFile(c.privateKey)
		if err != nil {
			return nil, err
		}
		key, err := ssh.ParsePrivateKey(authData)
		if err != nil {
			return nil, err
		}
		return ssh.PublicKeys(key), nil
	}
	return nil, errors.New("Auth type not support ")
}

func (c *Client) Close() {
	err := c.session.Close()
	if err != nil {
		return
	}
}
