package ssh

import (
	"fmt"
	"github.com/chulinx/zlxGo/assert"
	"testing"
)

var (
	user           = "zhangxiang1"
	pass           = "_79QP13zm5g"
	pubKeyAuthPath = "/Users/xxx/.ssh/id_rsa"
	addr           = "10.229.3.217:36000"
	addr1          = "10.16.88.7:36000"

	pwdAuth    = NewAuthPass(user, pass, addr)
	pwdAuth1   = NewAuthPass(user, pass, addr1)
	pubKeyAuth = NewAuthPrivateKey(user, pubKeyAuthPath, addr)
)

var tests = [][]string{
	{"hello world", "echo hello world", "hello world"},
	{"pipeline", "echo hello world|awk '{print \\$1}'", "hello"},
	//{"foreach",`for i in "1 2 3";do echo ${i};done`,"1 2 3"},
}

func TestClient_RunWithPass(t *testing.T) {
	c := NewSSHClient(pwdAuth)
	for _, test := range tests {
		name, cmd, result := test[0], test[1], test[2]
		fmt.Printf("Test ssh run cmd %s\n", name)
		runCmd, err := c.RunCmd(cmd)
		fmt.Println(runCmd, err)
		assert.AssertEqualExpect(runCmd, result, t)
		fmt.Printf("Test sudo ssh run cmd %s\n", name)
		sudo, err := c.RunCmdSudo(cmd)
		fmt.Println(sudo, err)
		assert.AssertEqualExpect(sudo, result, t)
	}
}
func TestClient_RunWithPassPassPassTipEn(t *testing.T) {
	c := NewSSHClient(pwdAuth)
	cmd, result := "echo hello world", "hello world"
	runCmd, _ := c.RunCmd(cmd)
	assert.AssertEqualExpect(runCmd, result, t)
	sudo, _ := c.RunCmdSudo(cmd)
	assert.AssertEqualExpect(sudo, result, t)
}

func run() {
	c := NewSSHClient(pwdAuth)
	cmd := "echo hello world"
	defer c.Close()
	c.RunCmd(cmd)
	c.RunCmdSudo(cmd)
}

func BenchmarkRunWithPassPassPassTipEn(b *testing.B) {
	for i := 0; i < 10; i++ {
		run()
	}
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
