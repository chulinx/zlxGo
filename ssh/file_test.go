package ssh

import (
	"fmt"
	"testing"
)

func TestClient_ScpFile(t *testing.T) {
	type args struct {
		srcPath  string
		destPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "scp file",
			args: args{
				srcPath:  "/Users/zhangsan/zlxgo/go.mod",
				destPath: "/tmp/",
			},
			wantErr: false,
		},
	}
	auth := NewAuthPass("zhangsan", "sadazx", "10.229.3.217:22")
	c := NewSSHClient(auth)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.ScpFile(tt.args.srcPath, tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("ScpFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ScpFileExecute(t *testing.T) {
	type args struct {
		srcPath  string
		destPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "scp file and run",
			args: args{
				srcPath:  "/Users/zhangsan/Desktop/Work/Code/ccwork/sretools/a.sh",
				destPath: "/tmp/a.sh",
			},
			wantErr: false,
		},
	}
	auth := NewAuthPass("zhangsan", "xxzssa", "10.229.3.217:22")
	c := NewSSHClient(auth)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.ScpFile(tt.args.srcPath, tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("ScpFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if s, err := c.Run("sh " + tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Println(s)
			}
		})
	}
}
