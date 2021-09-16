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
	addrList := strings.Split(auth.addr, ":")
	if len(addrList) < 2 {
		panic("addr format ip:port")
	}
	_, err := strconv.Atoi(addrList[1])
	if err != nil {
		panic("addr format ip:port")
	}

	auths := func() ssh.AuthMethod {
		if auth.pass != "" {
			return ssh.Password(auth.pass)
		}
		if auth.privateKey != "" {
			authData, err := ioutil.ReadFile(auth.privateKey)
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
			User:            auth.user,
			Auth:            []ssh.AuthMethod{auths()},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil },
			Timeout:         time.Second * 10,
		}
		config.SetDefaults()
		client, err := ssh.Dial("tcp", auth.addr, config)
		if nil != err {
			return nil
		}
		return client
	}

	return &Client{
		SAuth: &SAuth{
			user:       auth.user,
			pass:       auth.pass,
			addr:       auth.addr,
			privateKey: auth.privateKey,
		},
		Client: client(),
	}
}

func (c *Client) Close() {
	err := c.Client.Close()
	if err != nil {
		return
	}
}
