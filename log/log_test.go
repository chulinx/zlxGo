package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				msg: "Server start ...",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg, tt.args.fields...)
			var logParams LoggerParams
			logParams.LogLevel = zap.DebugLevel
			NewJsonFormat("user", &logParams)
			Debug(tt.args.msg, zap.Bool("registered", true))
			EmojiInfoF("Enabled Server,Port: %d", 8888)
			EmojiFatalF("ssss")
			TermSuccess("Install success,service: %s", "webserver")
			TermInfo("Install success,service: %s", "webserver")
			TermDebug("Install success,service: %s", "webserver")
		})
	}
}
