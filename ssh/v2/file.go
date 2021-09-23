package ssh

import (
	"bytes"
	"fmt"
	"github.com/povsister/scp"
	"io"
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
	scpClient, err := scp.NewClientFromExistingSSH(c.Client, &clientOpt)
	if err != nil {
		fmt.Println("Error creating new SSH Session from existing connection", err)
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

func (c *Client) CopyFileFromRemoteToByte(filePath string) (error, []byte) {
	buffer := new(bytes.Buffer)
	scpClient, err := scp.NewClientFromExistingSSH(c.Client, &scp.ClientOption{})
	if err != nil {
		fmt.Println("Error creating new SSH Session from existing connection", err)
		return err, nil
	}
	err = scpClient.CopyFromRemote(filePath, buffer, &scp.FileTransferOption{})
	if err != nil {
		return err, nil
	}
	return err, buffer.Bytes()
}

func (c *Client) CopyFileFromRemote(filePath string, w io.Writer) error {
	scpClient, err := scp.NewClientFromExistingSSH(c.Client, &scp.ClientOption{})
	if err != nil {
		fmt.Println("Error creating new SSH Session from existing connection", err)
		return err
	}
	err = scpClient.CopyFromRemote(filePath, w, &scp.FileTransferOption{})
	if err != nil {
		return err
	}
	return err
}

func (c *Client) CopyFileToRemoteFromByte(filePath string, b []byte) error {
	buffer := bytes.NewReader(b)
	scpClient, err := scp.NewClientFromExistingSSH(c.Client, &scp.ClientOption{})
	if err != nil {
		return err
	}
	err = scpClient.CopyToRemote(buffer, filePath, &scp.FileTransferOption{})
	if err != nil {
		return err
	}
	return err
}
