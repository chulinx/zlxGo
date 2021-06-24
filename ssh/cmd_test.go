package ssh

import (
	"testing"
)


func TestClient_Run(t *testing.T) {
	type fields struct {
		User       string
		Pwd        string
		Addr       string
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
			name: "echo hello world",
			fields: fields{"root","www.51imo.com","192.168.1.141:22"},
			args: args{"echo hello world"},
			want: "hello world\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				User:       tt.fields.User,
				Pwd:        tt.fields.Pwd,
				Addr:       tt.fields.Addr,
			}
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

