package ssh

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"strings"
)

// SAuth ssh config
type SAuth struct {
	user       string
	privateKey string
	pass       string
	addr       string // address format ip:port
}

func NewAuthPass(user, pass, addr string) *SAuth {
	return &SAuth{
		user: user,
		pass: pass,
		addr: addr,
	}
}

func NewAuthPrivateKey(user, privateKey, addr string) *SAuth {
	return &SAuth{
		user:       user,
		privateKey: privateKey,
		addr:       addr,
	}
}

func (c *Client) RunCmdSudo(shell string) (string, error) {
	if c.pass == "" {
		return "", errors.New("Sudo no allow type privateKey run ")
	}
	return c.runCmd(shell, true)
}

func (c *Client) RunCmd(shell string) (string, error) {
	return c.runCmd(shell, false)
}

func (c *Client) runCmd(shell string, sudo bool) (string, error) {
	var cmd string
	if c.client == nil {
		if _, err := c.Connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	cmd = fmt.Sprintf("sh -c \"%s\"", shell)
	if sudo {
		cmd = fmt.Sprintf("sudo sh -c \"%s\"", shell)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return "", err
	}

	stdoutB := new(bytes.Buffer)
	session.Stdout = stdoutB
	in, _ := session.StdinPipe()

	passTipCn := fmt.Sprintf("[sudo] %s 的密码：", c.user)
	passTipEn := fmt.Sprintf("[sudo] password for %s:", c.user)
	go func(in io.Writer, output *bytes.Buffer, passTipEn, passTipCn string) {
		for {
			if strings.Contains(string(output.Bytes()), passTipCn) || strings.Contains(string(output.Bytes()), passTipEn) {
				_, err = in.Write([]byte(c.pass + "\n"))
				if err != nil {
					break
				}
				break
			}
		}
	}(in, stdoutB, passTipEn, passTipCn)

	err = session.Run(cmd)
	if err != nil {
		return "", err
	}
	s := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(stdoutB.String(), passTipCn), passTipEn))
	return s, nil
}
