package ssh

import (
	"fmt"
	"testing"
)

func TestClient_RunWithPass(t *testing.T) {
	type fields struct {
		User string
		Pwd  string
		Addr string
	}
	type args struct {
		shell string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "echo hello world",
			fields: fields{"root", "www.51imo.com", "192.168.1.141:22"},
			args:   args{"echo hello world"},
			want:   "hello world\n",
		},
	}
	fmt.Println("Start pass ssh command")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwdAuth := NewAuthPass("zhangsan", "xxxxxx", "10.229.3.217:22")
			c := NewSSHClient(pwdAuth)
			got, err := c.Run(tt.args.shell)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
	fmt.Println("Start pubkey ssh command")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKeyAuth := NewAuthPrivateKey("zhangsan", "/Users/zhangsan/.ssh/id_rsa.pub", "10.229.3.217:22")
			c := NewSSHClient(pubKeyAuth)
			got, err := c.Run(tt.args.shell)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RunWithPubKey(t *testing.T) {
	type fields struct {
		User string
		Pwd  string
		Addr string
	}
	type args struct {
		shell string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "echo hello world",
			fields: fields{"root", "www.51imo.com", "192.168.1.141:22"},
			args:   args{"echo hello world"},
			want:   "hello world\n",
		},
	}
	fmt.Println("Start pubkey ssh command")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKeyAuth := NewAuthPrivateKey("zhangxiang1", "/Users/zhangxiang/.ssh/id_rsa", "10.229.3.217:22")
			c := NewSSHClient(pubKeyAuth)
			got, err := c.Run(tt.args.shell)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
