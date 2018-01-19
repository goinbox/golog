package golog

import (
	"os"
	"sync"
)

type ConsoleWriter struct {
	lock *sync.Mutex

	*os.File
}

func NewStdoutWriter() *ConsoleWriter {
	return &ConsoleWriter{
		lock: new(sync.Mutex),

		File: os.Stdout,
	}
}

func NewStderrWriter() *ConsoleWriter {
	return &ConsoleWriter{
		lock: new(sync.Mutex),

		File: os.Stderr,
	}
}

func (this *ConsoleWriter) Write(msg []byte) (int, error) {
	this.lock.Lock()
	n, err := this.File.Write(msg)
	this.lock.Unlock()

	return n, err
}

func (this *ConsoleWriter) Flush() error {
	return nil
}

func (this *ConsoleWriter) Free() {
}
