package ssh

import (
	"bufio"
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

// RunCmdSudoStream the c.User must have sudo permission
func (c *Client) RunCmdSudoStream(ctx context.Context, textChan chan string, shell string) error {
	if c.Pass == "" {
		return errors.New("Sudo no allow type privateKey run ")
	}
	return c.runCmdStream(ctx, textChan, shell, true)
}

func (c *Client) RunCmdStream(ctx context.Context, textChan chan string, shell string) error {
	return c.runCmdStream(ctx, textChan, shell, false)
}

// RunCmdSudo the c.User must have sudo permission
func (c *Client) RunCmdSudo(shell string) (string, error) {
	if c.Pass == "" {
		return "", errors.New("Sudo no allow type privateKey run ")
	}
	return c.runCmd(shell, true, false)
}

func (c *Client) RunCmd(shell string) (string, error) {
	return c.runCmd(shell, false, false)
}

// RunCmdWihScriptSudo the c.User must have sudo permission
func (c *Client) RunCmdWihScriptSudo(shell string) (string, error) {
	if c.Pass == "" {
		return "", errors.New("Sudo no allow type privateKey run ")
	}
	return c.runCmd(shell, true, true)
}

func (c *Client) RunCmdWihScript(shell string) (string, error) {
	return c.runCmd(shell, false, true)
}

// runCmd runs a command, return stdout or stderr after the command exec completion
func (c *Client) runCmd(shell string, sudo, scriptMode bool) (string, error) {

	if c.Client == nil {
		return "", errors.New("init ssh client failed")
	}
	session, err := c.Client.NewSession()
	if err != nil {
		return "", err
	}
	cmd, err2 := c.makeCmd(shell, sudo, scriptMode)
	if err2 != nil {
		return "", err2
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
	// input sudo password
	go c.sudoPass(in, stdoutB, passTipEn, passTipCn, ctx, err)

	err = session.Run(cmd)
	if err != nil {
		return stdoutA.String(), err
	}
	s := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(stdoutB.String(), passTipCn), passTipEn))
	return s, nil
}

func (c *Client) runCmdStream(ctx context.Context, textChan chan string, cmd string, sudo bool) error {
	context, cancel := context.WithCancel(ctx)
	defer cancel()
	cmd, err := c.makeCmd(cmd, sudo, false)
	if err != nil {
		return err
	}
	if c.Client == nil {
		return errors.New("init ssh client failed")
	}
	session, err := c.Client.NewSession()
	if err != nil {
		return err
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return err
	}
	in, _ := session.StdinPipe()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	go c.copyStdout(context, in, stdout, textChan, sudo)
	go func() {
		for {
			select {
			case <-context.Done():
				err := session.Close()
				if err != nil {
					return
				}
			default:
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	err = session.Start(cmd)
	if err != nil {
		return err
	}

	err = session.Wait()
	if err != nil {
		return err
	}
	return nil
}

// sudoPass input pass when ask sudo pass
func (c *Client) sudoPass(in io.Writer, output *bytes.Buffer, passTipEn string, passTipCn string, ctx context.Context, err error) {
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
}

// copyStdout copy session.StdoutPipe to io.Writer
func (c *Client) copyStdout(ctx context.Context, in io.Writer, stdout io.Reader, textChan chan string, sudo bool) error {
	defer close(textChan)
	passTipCn := fmt.Sprintf("[sudo] %s 的密码：", c.User)
	passTipEn := fmt.Sprintf("[sudo] password for %s:", c.User)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			// if sudo
			if sudo {
				var output []byte
				select {
				case <-ctx.Done():
					return nil
				default:
					var (
						r = bufio.NewReader(stdout)
					)
					for {
						b, err := r.ReadByte()
						if err != nil {
							break
						}
						output = append(output, b)
						if b == byte('\n') {
							continue
						}

						if strings.Contains(string(output), passTipCn) || strings.Contains(string(output), passTipEn) {
							_, err = in.Write([]byte(c.Pass + "\n"))
							if err != nil {
								continue
							}
							break
						}

					}
				}
			}
			scan := bufio.NewScanner(stdout)
			scan.Split(bufio.ScanLines)
			for scan.Scan() {
				textChan <- scan.Text()
			}
		}
	}
}

// makeCmd generate command
// if sudo is true, add sudo before cmd
// if scriptMode is true,copy shell content to remote host and exec, finally delete it
func (c *Client) makeCmd(shell string, sudo bool, scriptMode bool) (string, error) {
	var cmd string
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
	return cmd, nil
}

func SelectTest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("select test")
			time.Sleep(time.Microsecond * 1000)
		}
	}
}
