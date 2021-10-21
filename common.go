/**
* @file logger.go
* @author ligang
* @date 2016-02-04
 */

package golog

import "io"

const (
	LevelEmergency = 0
	LevelAlert     = 1
	LevelCritical  = 2
	LevelError     = 3
	LevelWarning   = 4
	LevelNotice    = 5
	LevelInfo      = 6
	LevelDebug     = 7
)

type Field struct {
	Key   string
	Value interface{}

	preset bool
}

type Logger interface {
	Debug(msg string, fields ...*Field)
	Info(msg string, fields ...*Field)
	Notice(msg string, fields ...*Field)
	Warning(msg string, fields ...*Field)
	Error(msg string, fields ...*Field)
	Critical(msg string, fields ...*Field)
	Alert(msg string, fields ...*Field)
	Emergency(msg string, fields ...*Field)

	With(fields ...*Field) Logger
}

type Formater interface {
	Format(fields ...*Field) []byte
}

type Writer interface {
	io.Writer

	Flush() error
	Free()
}
