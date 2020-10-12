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

var LogLevels = map[int][]byte{
	LevelEmergency: []byte("emergency"),
	LevelAlert:     []byte("alert"),
	LevelCritical:  []byte("critical"),
	LevelError:     []byte("error"),
	LevelWarning:   []byte("warning"),
	LevelNotice:    []byte("notice"),
	LevelInfo:      []byte("info"),
	LevelDebug:     []byte("debug"),
}

type ILogger interface {
	Debug(msg []byte)
	Info(msg []byte)
	Notice(msg []byte)
	Warning(msg []byte)
	Error(msg []byte)
	Critical(msg []byte)
	Alert(msg []byte)
	Emergency(msg []byte)

	Log(level int, msg []byte) error
}

type IFormater interface {
	Format(level int, msg []byte) []byte
}

type IWriter interface {
	io.Writer

	Flush() error
	Free()
}
