package ssh

import (
	"github.com/chulinx/zlxGo/assert"
	"github.com/chulinx/zlxGo/stringfile"
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

func scpFile(t *testing.T, sAuth *SAuth) {
	err := stringfile.RewriteFile(content, localFile)
	assert.AssertError(err, t)
	c := NewSSHClient(sAuth)
	file, err := stringfile.ReadFile(localFile)
	assert.AssertError(err, t)
	err = c.ScpFile(localFile, remoteFile)
	assert.AssertError(err, t)
	cmd, err1 := c.RunCmd("cat " + remoteFile)
	assert.AssertError(err1, t)
	assert.AssertEqualExpect(file, cmd+"\n", t)
}
