package ssh

import (
	"github.com/chulinx/zlxGo/assert"
	"testing"
)

var (
	user           = "xxx"
	pass           = "xxx"
	pubKeyAuthPath = "/Users/xxx/.ssh/id_rsa"
	addr           = "10.229.3.217:22"
	addr1          = "10.16.88.7:22"

	pwdAuth    = NewAuthPass(user, pass, addr)
	pwdAuth1   = NewAuthPass(user, pass, addr1)
	pubKeyAuth = NewAuthPrivateKey(user, pubKeyAuthPath, addr)
)

func TestClient_RunWithPassPassPassTipEn(t *testing.T) {
	c := NewSSHClient(pwdAuth)
	cmd, result := "echo hello world", "hello world"
	runCmd, _ := c.RunCmd(cmd)
	assert.AssertEqualExpect(runCmd, result, t)
	sudo, _ := c.RunCmdSudo(cmd)
	assert.AssertEqualExpect(sudo, result, t)
}

func TestClient_RunWithPassPassTipCn(t *testing.T) {
	c := NewSSHClient(pwdAuth1)
	cmd, result := "echo hello world", "hello world"
	runCmd, _ := c.RunCmd(cmd)
	assert.AssertEqualExpect(runCmd, result, t)
	sudo, _ := c.RunCmdSudo(cmd)
	assert.AssertEqualExpect(sudo, result, t)
}

func TestClient_RunWithPubKey(t *testing.T) {
	c := NewSSHClient(pubKeyAuth)
	cmd, result := "echo hello world", "hello world"
	runCmd, _ := c.RunCmd(cmd)
	assert.AssertEqualExpect(runCmd, result, t)
	_, err := c.RunCmdSudo(cmd)
	assert.AssertEqualExpect(err.Error(), "Sudo no allow type privateKey run ", t)
}
