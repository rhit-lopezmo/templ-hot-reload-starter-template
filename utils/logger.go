package utils

import (
	stdlog "log"
	"os"
	"strings"
)

type Logger struct {
	*stdlog.Logger
}

var (
	MainLogger *Logger
)

func InitLoggers() {
	MainLogger = NewLogger("MAIN")
}

func NewLogger(tag string) *Logger {
	prefix := "[" + strings.ToUpper(tag) + "] "

	return &Logger{
		Logger: stdlog.New(os.Stdout, prefix, stdlog.Lmsgprefix),
	}
}
