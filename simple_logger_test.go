package golog

import (
	"strconv"
	"testing"
	"time"
)

func TestSimpleLogger(t *testing.T) {
	fw, _ := NewFileWriter("/tmp/test_simple_logger.log", 1024)
	aw := NewAsyncWriter(fw, 1024)
	defer aw.Free()

	f := NewSimpleFormater()
	logger := NewSimpleLogger(aw, f).SetLogLevel(LEVEL_DEBUG)

	for i := 0; i < 1000; i++ {
		msg := []byte("test simple logger " + strconv.Itoa(i))

		logger.Debug(msg)
		logger.Info(msg)
		logger.Notice(msg)
		logger.Warning(msg)
		logger.Error(msg)
		logger.Critical(msg)
		logger.Alert(msg)
		logger.Emergency(msg)
	}

	time.Sleep(time.Second * 10)
}
