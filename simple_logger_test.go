package golog

import (
	"strconv"
	"testing"
)

func TestSimpleLogger(t *testing.T) {
	w, _ := NewFileWriter("/dev/stdout", 0)
	f := NewSimpleFormater()
	logger := NewSimpleLogger(w, f).SetLogLevel(LevelDebug).EnableColor().
		With(&Field{
			Key:   "test_with",
			Value: "test with value",
		})

	for i := 0; i < 1000; i++ {
		msg := "test simple logger " + strconv.Itoa(i)
		field := &Field{
			Key:   "i",
			Value: i,
		}

		logger.Debug(msg, field)
		logger.Info(msg, field)
		logger.Notice(msg, field)
		logger.Warning(msg, field)
		logger.Error(msg, field)
		logger.Critical(msg, field)
		logger.Alert(msg, field)
		logger.Emergency(msg, field)
	}
}
