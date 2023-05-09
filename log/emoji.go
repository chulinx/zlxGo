package log

import "fmt"

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
