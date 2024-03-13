package dets

import (
	"log/slog"
)

type LogInterface interface {
	Error(msg string, args ...any)
}

type defaultLogger struct {
}

func (l *defaultLogger) Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

var logger LogInterface = &defaultLogger{}
