package golog

import (
	"strconv"
	"testing"
	"time"
)

func TestFileWriter(t *testing.T) {
	path := "/tmp/test_file_writer.log"

	writer, _ := NewFileWriter(path, 1024)
	for i := 0; i < 1000; i++ {
		s := "test file writer " + strconv.Itoa(i) + " \n"
		writer.Write([]byte(s))
	}

	time.Sleep(time.Second * 10)
	writer.Free()
}
