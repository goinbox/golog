package golog

import (
	"strconv"
	"time"

	//     "fmt"
	"testing"
)

func TestAsyncWriter(t *testing.T) {
	path := "/tmp/test_async_writer.log"

	fw, _ := NewFileWriter(path, 1024)
	aw := NewAsyncWriter(fw, 1024)

	for i := 0; i < 1000; i++ {
		s := "test async writer " + strconv.Itoa(i) + " \n"
		aw.Write([]byte(s))
	}

	time.Sleep(time.Second * 10)

	aw.Free()
}
