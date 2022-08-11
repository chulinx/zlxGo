package endecryption

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_confusionString(t *testing.T) {
	type args struct {
		s         string
		confusion string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s:         "zlx",
				confusion: "xs",
			},
			want: "xzlsx",
		},
		{
			name: "test2",
			args: args{
				s:         "zlx",
				confusion: "xss",
			},
			want: "xszlsx",
		},
		{
			name: "test3",
			args: args{
				s:         "ccwork",
				confusion: "xss",
			},
			want: "xsccworsk",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confusionString(tt.args.s, tt.args.confusion); got != tt.want {
				t.Errorf("confusionString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clarifyString(t *testing.T) {
	type args struct {
		s         string
		confusion string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "test1",
			args: args{
				s:         "xzlsx",
				confusion: "xs",
			},
			want: "zlx",
		},
		{
			name: "test2",
			args: args{
				s:         "xszlsx",
				confusion: "xss",
			},
			want: "zlx",
		},
		{
			name: "test3",
			args: args{
				s:         "xsxszlxssss",
				confusion: "xss",
			},
			want: "xszlxsss",
		},
		{
			name: "test4",
			args: args{
				s:         "xsccworsk",
				confusion: "xss",
			},
			want: "ccwork",
		},

		{
			name: "test5",
			args: args{
				s:         "^Z*x5lGccwork202o2",
				confusion: "^Z*x5lGo",
			},
			want: "ccwork2022",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clarifyString(tt.args.s, tt.args.confusion); got != tt.want {
				t.Errorf("clarifyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64_Base64Encode(t *testing.T) {
	type fields struct {
		encoding base64.Encoding
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test1",
			args: args{s: ",Qwedq2132xsa"},
			want: "YzJGemMyUjNMRkYzWldSeE1qRXpNbmh6WkdF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBase64("sassdwd")
			fmt.Println(b.Base64Decode(tt.want))
			if got := b.Base64Encode(tt.args.s); got != tt.want {
				t.Errorf("Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
