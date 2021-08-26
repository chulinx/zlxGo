package ssh

import (
	"fmt"
	"github.com/povsister/scp"
	"os"
)

func (c *Client) ScpFile(srcPath, destPath string) error {
	return c.scpFile(srcPath, destPath, false)
}

func (c *Client) scpFile(srcPath, destPath string, isRoot bool) error {
	var (
		fileInfo  os.FileInfo
		clientOpt scp.ClientOption
	)
	clientOpt.Sudo = isRoot
	sshClient, err := c.Connect()
	if err != nil {
		return err
	}
	scpClient, err := scp.NewClientFromExistingSSH(sshClient.client, &clientOpt)
	if err != nil {
		fmt.Println("Error creating new SSH session from existing connection", err)
		return err
	}
	// file not exits return
	if fileInfo, err = os.Stat(srcPath); err != nil && os.IsNotExist(err) {
		return err
	} else {
		if fileInfo.IsDir() {
			err = scpClient.CopyDirToRemote(srcPath, destPath, &scp.DirTransferOption{})
			if err != nil {
				return err
			}
			return nil
		}
	}
	err = scpClient.CopyFileToRemote(srcPath, destPath, &scp.FileTransferOption{})
	if err != nil {
		return err
	}
	return nil
}
