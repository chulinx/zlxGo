package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	Logger      *zap.Logger
	EmojiLogger *zap.Logger
)

type LoggerParams struct {
	lumberjack.Logger
	LogLevel zapcore.Level
}

func init() {
	defaultLog()
}

// defaultLog default log config
func defaultLog() {
	var logParams LoggerParams
	logParams.LogLevel = zapcore.InfoLevel
	Logger = newLog(&logParams)
	EmojiLogger = simpleLog(&logParams)
}

// NewJsonFormat defined params serviceName,log level and log output file
func NewJsonFormat(serviceName string, params *LoggerParams) {
	Logger = configLog(serviceName, params)
}

func newLog(params *LoggerParams) *zap.Logger {
	return configLog("", params)
}

func configLog(serviceName string, params *LoggerParams) *zap.Logger {
	hook := parseParams(params)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line_num",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		FunctionKey:    "func",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(params.LogLevel)

	var writerSyncer zapcore.WriteSyncer
	if params.Filename == "" {
		writerSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	} else {
		writerSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)) // 打印到控制台和文件
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		writerSyncer,
		atomicLevel, // 日志级别
	)

	// 获取上层调用函数行号
	skip := zap.AddCallerSkip(1)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	var filed zap.Option
	if serviceName != "" {
		filed = zap.Fields(zap.String("service", serviceName))
		// 构造日志
		return zap.New(core, caller, skip, development, filed)
	}
	// 构造日志
	return zap.New(core, caller, skip, development)
}

func simpleLog(params *LoggerParams) *zap.Logger {
	hook := parseParams(params)

	encoderConfig := zapcore.EncoderConfig{
		//TimeKey:        "time",
		//LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line_num",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		FunctionKey:    "func",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(params.LogLevel)

	var writerSyncer zapcore.WriteSyncer
	if params.Filename == "" {
		writerSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	} else {
		writerSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)) // 打印到控制台和文件
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 编码器配置
		writerSyncer,
		atomicLevel, // 日志级别
	)
	return zap.New(core)
}

func EmojiFatalF(msg string, fields ...interface{}) {
	printf("[❌] "+msg, fields...)
}

func EmojiSuccessF(msg string, fields ...interface{}) {
	printf("[😉] "+msg, fields...)
}

func EmojiErrorF(msg string, fields ...interface{}) {
	printf("[💔] "+msg, fields...)
}

func EmojiInfoF(msg string, fields ...interface{}) {
	printf("[🙂] "+msg, fields...)
}

func printf(msg string, fields ...interface{}) {
	s := fmt.Sprintf(msg, fields...)
	EmojiLogger.Info(s)
}

func parseParams(params *LoggerParams) *LoggerParams {
	if params.MaxSize == 0 {
		params.MaxSize = 1024
	}
	if !params.Compress {
		params.Compress = true
	}
	if params.MaxAge == 0 {
		params.MaxAge = 7
	}
	if params.MaxBackups == 0 {
		params.MaxBackups = 7
	}
	if params.LogLevel == 0 {
		params.LogLevel = zapcore.InfoLevel
	}
	return params
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Level(mod string) zapcore.Level {
	switch mod {
	case "Info", "info":
		return zapcore.InfoLevel
	case "Error", "error":
		return zapcore.ErrorLevel
	case "Debug", "debug":
		return zapcore.DebugLevel
	case "Warn", "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.InfoLevel
	}
}
