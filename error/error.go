package error

import (
	"log"
	"os"
)

type Error struct {
	Prefix string
}

func NewError() *Error {
	return &Error{}
}

func (e Error) SetPrefix(prefix string) *Error {
	if prefix == "" {
		prefix = "zlxGo"
	}
	e.Prefix = prefix
	return &e
}

func (e Error) FatalError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s--%s", e.Prefix, msg, err)
	}
}

func (e Error) PrintfError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s--%s", e.Prefix, msg, err)
	}
}

func (e Error) FatalErrorExit(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s--%s", e.Prefix, msg, err)
		os.Exit(1)
	}
}
