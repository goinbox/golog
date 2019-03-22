package golog

import (
	"testing"
)

func TestSimpleFormater(t *testing.T) {
	format(simpleFormaterForTest(), []byte("test simple formater"), t)
}

func TestConsoleFormater(t *testing.T) {
	cf := NewConsoleFormater(simpleFormaterForTest())

	format(cf, []byte("test console formater"), t)
}

func simpleFormaterForTest() *simpleFormater {
	return NewSimpleFormater().
		SetAddress([]byte("127.0.0.1")).
		SetTraceId([]byte("123456"))
}

func format(f IFormater, msg []byte, t *testing.T) {
	for level, _ := range LogLevels {
		b := f.Format(level, append(msg, '\n'))
		t.Log(string(b))
	}
}
