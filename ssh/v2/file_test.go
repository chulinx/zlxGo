package ssh

import (
	"github.com/chulinx/zlxGo/assert"
	"github.com/chulinx/zlxGo/stringfile"
	"os"
	"testing"
)

var (
	localFile  = "/Users/zhangxiang/index.html"
	remoteFile = "/tmp/index.html"
	content    = "hello world"
)

func TestClient_ScpFileWithPass(t *testing.T) {
	scpFile(t, pwdAuth)
}

func TestClient_ScpFileWithKey(t *testing.T) {
	scpFile(t, pubKeyAuth)
}

func TestClient_CopyFileWithByte(t *testing.T) {
	c := NewSSHClient(pwdAuth)
	defer c.Close()
	err := c.CopyFileToRemoteFromByte(remoteFile, []byte("hello world1"))
	assert.AssertError(err, t)
	err, i := c.CopyFileFromRemoteToByte(remoteFile)
	assert.AssertEqualExpect(string(i), "hello world1", t)
	assert.AssertError(err, t)
}

func scpFile(t *testing.T, sAuth *SAuth) {
	err := os.WriteFile(localFile, []byte(content), 0755)
	if err != nil {
		return
	}
	err = stringfile.RewriteFile(content, localFile)
	assert.AssertError(err, t)
	c := NewSSHClient(sAuth)
	defer c.Close()
	file, err := stringfile.ReadFile(localFile)
	assert.AssertError(err, t)
	err = c.ScpFile(localFile, remoteFile)
	assert.AssertError(err, t)
	cmd, err1 := c.RunCmd("cat " + remoteFile)
	assert.AssertError(err1, t)
	assert.AssertEqualExpect(file, cmd+"\n", t)
}
