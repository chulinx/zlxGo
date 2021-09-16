package ssh

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	*SAuth
	Client     *ssh.Client
	Session    *ssh.Session
	lastResult string
}

func NewSSHClient(auth *SAuth) *Client {
	addrList := strings.Split(auth.Addr, ":")
	if len(addrList) < 2 {
		panic("addr format ip:port")
	}
	_, err := strconv.Atoi(addrList[1])
	if err != nil {
		panic("addr format ip:port")
	}

	auths := func() ssh.AuthMethod {
		if auth.Pass != "" {
			return ssh.Password(auth.Pass)
		}
		if auth.PrivateKey != "" {
			authData, err := ioutil.ReadFile(auth.PrivateKey)
			if err != nil {
				return nil
			}
			key, err := ssh.ParsePrivateKey(authData)
			if err != nil {
				return nil
			}
			return ssh.PublicKeys(key)
		}
		return nil
	}

	client := func() *ssh.Client {
		config := &ssh.ClientConfig{
			User:            auth.User,
			Auth:            []ssh.AuthMethod{auths()},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil },
			Timeout:         time.Second * 10,
		}
		config.SetDefaults()
		client, err := ssh.Dial("tcp", auth.Addr, config)
		if nil != err {
			return nil
		}
		return client
	}

	return &Client{
		SAuth:  auth,
		Client: client(),
	}
}

func (c *Client) Close() {
	err := c.Client.Close()
	if err != nil {
		return
	}
}
