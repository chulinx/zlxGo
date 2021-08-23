package ssh

import (
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"os"
	"strings"
)

func (c *Client) ScpFile(srcPath, destPath string) error {
	sshClient, err := c.Connect()
	defer sshClient.client.Close()
	if err != nil {
		return err
	}
	scpClient, err := scp.NewClientBySSH(sshClient.client)
	if err != nil {
		fmt.Println("Error creating new SSH session from existing connection", err)
		return err
	}
	f, _ := os.Open(srcPath)
	defer f.Close()
	// complete dest file path
	srcPathSplit := strings.Split(srcPath, "/")
	onlyFileName := srcPathSplit[len(srcPathSplit)-1]
	if !strings.Contains(destPath, onlyFileName) {
		destPath = fmt.Sprintf("%s/%s", destPath, onlyFileName)
	}
	err = scpClient.CopyFile(f, destPath, "0655")
	if err != nil {
		return err
	}
	return nil
}
