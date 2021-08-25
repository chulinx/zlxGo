package ssh

import (
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
				srcPath:  "/Users/zhangxiang/docker-compose.yml",
				destPath: "/tmp/docker-compose.yml",
			},
			wantErr: false,
		},
	}
	auth := NewAuthPass("zhangsan", "xxx", "10.229.3.217:22")
	c := NewSSHClient(auth)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.ScpFile(tt.args.srcPath, tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("ScpFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
