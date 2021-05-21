package golog

import (
	"strconv"
	"testing"
)

func TestConsoleWriter(t *testing.T) {
	writer := NewConsoleWriter()
	for i := 0; i < 1000; i++ {
		s := "test console writer " + strconv.Itoa(i) + " \n"
		_, _ = writer.Write([]byte(s))
	}
}
