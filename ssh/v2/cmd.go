package ssh

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"strings"
	"time"
)

// SAuth ssh config
type SAuth struct {
	User       string
	PrivateKey string
	Pass       string
	Addr       string // address format ip:port
}

func NewAuthPass(user, pass, addr string) *SAuth {
	return &SAuth{
		User: user,
		Pass: pass,
		Addr: addr,
	}
}

func NewAuthPrivateKey(user, privateKey, addr string) *SAuth {
	return &SAuth{
		User:       user,
		PrivateKey: privateKey,
		Addr:       addr,
	}
}

func (c *Client) RunCmdSudo(shell string) (string, error) {
	if c.Pass == "" {
		return "", errors.New("Sudo no allow type privateKey run ")
	}
	return c.runCmd(shell, true, false)
}

func (c *Client) RunCmdWihScriptSudo(shell string) (string, error) {
	return c.runCmd(shell, true, true)
}

func (c *Client) RunCmd(shell string) (string, error) {
	return c.runCmd(shell, false, false)
}

func (c *Client) RunCmdWihScript(shell string) (string, error) {
	return c.runCmd(shell, false, true)
}

func (c *Client) runCmd(shell string, sudo, scriptMode bool) (string, error) {
	var cmd string
	session, err := c.Client.NewSession()
	if err != nil {
		return "", err
	}
	if scriptMode {
		scriptFileName := fmt.Sprintf("/tmp/%d.sh", time.Now().Unix())
		err := c.CopyFileToRemoteFromByte(scriptFileName, []byte(shell))
		if err != nil {
			return "", err
		}
		if scriptFileName == "/" || scriptFileName == "/*" {
			return "", errors.New("file name not allow / or /*")
		}
		shell = fmt.Sprintf("sh %s && rm -f %s", scriptFileName, scriptFileName)
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
	stdoutA := new(bytes.Buffer)
	session.Stdout = stdoutB
	session.Stderr = stdoutA
	in, _ := session.StdinPipe()

	passTipCn := fmt.Sprintf("[sudo] %s 的密码：", c.User)
	passTipEn := fmt.Sprintf("[sudo] password for %s:", c.User)
	// exit goroutine when execute function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(in io.Writer, output *bytes.Buffer, passTipEn, passTipCn string) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if strings.Contains(string(output.Bytes()), passTipCn) || strings.Contains(string(output.Bytes()), passTipEn) {
					_, err = in.Write([]byte(c.Pass + "\n"))
					if err != nil {
						break
					}
					break
				}
			}
			time.Sleep(time.Microsecond * 100)
		}
	}(in, stdoutB, passTipEn, passTipCn)

	err = session.Run(cmd)
	if err != nil {
		return stdoutA.String(), err
	}
	s := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(stdoutB.String(), passTipCn), passTipEn))
	return s, nil
}
