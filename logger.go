package dets

import "log"

type LogInterface interface {
	Error(v ...interface{})
}

type defaultLogger struct {
}

func (l *defaultLogger) Error(v ...interface{}) {
	log.Println(v...)
}

var logger LogInterface = &defaultLogger{}
