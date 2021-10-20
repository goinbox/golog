/**
* @file logger.go
* @author ligang
* @date 2016-02-04
 */

package golog

import (
	"time"

	"github.com/goinbox/color"
)

const (
	defaultFieldKeyLevel = "level"
	defaultFieldKeyTime  = "t"
	defaultFieldKeyMsg   = "msg"
	defaultTimeLayout    = "2006-02-01 15:04:05.000"
)

var logLevels = map[int]string{
	LevelEmergency: "emergency",
	LevelAlert:     "alert",
	LevelCritical:  "critical",
	LevelError:     "error",
	LevelWarning:   "warning",
	LevelNotice:    "notice",
	LevelInfo:      "info",
	LevelDebug:     "debug",
}

type colorFunc func(msg []byte) []byte

var levelColorFuncs = map[int]colorFunc{
	LevelDebug:     color.Yellow,
	LevelInfo:      color.Blue,
	LevelNotice:    color.Cyan,
	LevelWarning:   color.Maganta,
	LevelError:     color.Red,
	LevelCritical:  color.Black,
	LevelAlert:     color.White,
	LevelEmergency: color.Green,
}

type simpleLogger struct {
	writer   Writer
	formater Formater

	glevel     int
	withFields []*Field
	timeLayout string

	enableColor bool

	fieldKeyLevel string
	fieldKeyTime  string
	fieldKeyMsg   string
}

func NewSimpleLogger(writer Writer, formater Formater) *simpleLogger {
	return &simpleLogger{
		writer:   writer,
		formater: formater,

		glevel:      LevelInfo,
		timeLayout:  defaultTimeLayout,
		enableColor: false,

		fieldKeyLevel: defaultFieldKeyLevel,
		fieldKeyTime:  defaultFieldKeyTime,
		fieldKeyMsg:   defaultFieldKeyMsg,
	}
}

func (l *simpleLogger) SetLogLevel(level int) *simpleLogger {
	_, ok := logLevels[level]
	if ok {
		l.glevel = level
	}

	return l
}

func (l *simpleLogger) SetTimeLayout(layout string) *simpleLogger {
	l.timeLayout = layout

	return l
}

func (l *simpleLogger) EnableColor() *simpleLogger {
	l.enableColor = true

	return l
}

func (l *simpleLogger) SetDefaultFieldKeyLevel(key string) *simpleLogger {
	l.fieldKeyLevel = key

	return l
}

func (l *simpleLogger) SetDefaultFieldKeyTime(key string) *simpleLogger {
	l.fieldKeyTime = key

	return l
}

func (l *simpleLogger) SetDefaultFieldKeyMsg(key string) *simpleLogger {
	l.fieldKeyMsg = key

	return l
}

func (l *simpleLogger) With(fields ...*Field) Logger {
	return &simpleLogger{
		writer:   l.writer,
		formater: l.formater,

		glevel:     l.glevel,
		withFields: append(l.withFields, fields...),
		timeLayout: l.timeLayout,

		enableColor: l.enableColor,

		fieldKeyLevel: l.fieldKeyLevel,
		fieldKeyTime:  l.fieldKeyTime,
		fieldKeyMsg:   l.fieldKeyMsg,
	}
}

func (l *simpleLogger) Debug(msg string, fields ...*Field) {
	l.log(LevelDebug, msg, fields)
}

func (l *simpleLogger) Info(msg string, fields ...*Field) {
	l.log(LevelInfo, msg, fields)
}

func (l *simpleLogger) Notice(msg string, fields ...*Field) {
	l.log(LevelNotice, msg, fields)
}

func (l *simpleLogger) Warning(msg string, fields ...*Field) {
	l.log(LevelWarning, msg, fields)
}

func (l *simpleLogger) Error(msg string, fields ...*Field) {
	l.log(LevelError, msg, fields)
}

func (l *simpleLogger) Critical(msg string, fields ...*Field) {
	l.log(LevelCritical, msg, fields)
}

func (l *simpleLogger) Alert(msg string, fields ...*Field) {
	l.log(LevelAlert, msg, fields)
}

func (l *simpleLogger) Emergency(msg string, fields ...*Field) {
	l.log(LevelEmergency, msg, fields)
}

func (l *simpleLogger) log(level int, msg string, fields []*Field) {
	if level > l.glevel {
		return
	}

	allFields := append([]*Field{
		{
			Key:   l.fieldKeyLevel,
			Value: logLevels[level],
		},
		{
			Key:   l.fieldKeyTime,
			Value: time.Now().Format(l.timeLayout),
		},
	}, l.withFields...)

	allFields = append(allFields, &Field{
		Key:   l.fieldKeyMsg,
		Value: msg,
	})

	allFields = append(allFields, fields...)

	p := l.formater.Format(allFields...)
	if l.enableColor {
		p = levelColorFuncs[level](p)
	}
	p = append(p, '\n')

	_, _ = l.writer.Write(p)
}
