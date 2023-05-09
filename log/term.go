package log

import (
	"fmt"
	"github.com/pterm/pterm"
)

func TermSuccess(format string, a ...interface{}) {
	msg := fmt.Sprintf("✌️ %s\n", format)
	pterm.Success.Printf(msg, a)
}

func TermInfo(format string, a ...interface{}) {
	msg := fmt.Sprintf("ℹ️ ️%s\n", format)
	pterm.Info.Printf(msg, a)
}

func TermDebug(format string, a ...interface{}) {
	pterm.EnableDebugMessages()
	msg := fmt.Sprintf("🤔️ ️%s\n", format)
	pterm.Debug.Printf(msg, a)
}
