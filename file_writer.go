/**
* @file file.go
* @brief writer msg to file
* @author ligang
* @date 2016-02-03
 */

package golog

import (
	"github.com/goinbox/gomisc"

	"os"
	"sync"
	"time"
)

type fileWriter struct {
	*os.File

	path           string
	lock           *sync.Mutex
	lastTimeSecond int64

	buf     []byte
	bufsize int
	bufpos  int
}

func NewFileWriter(path string, bufsize int) (*fileWriter, error) {
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}

	return &fileWriter{
		File: file,

		path:           path,
		lock:           new(sync.Mutex),
		lastTimeSecond: time.Now().Unix(),

		buf:     make([]byte, bufsize),
		bufsize: bufsize,
		bufpos:  0,
	}, nil
}

func (w *fileWriter) Write(msg []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	err := w.ensureFileExist()
	if err != nil {
		return 0, err
	}

	if w.bufsize == 0 {
		return w.File.Write(msg)
	}

	if w.appendToBuffer(msg) {
		return len(msg), nil
	}

	err = w.flushBuffer()
	if err != nil {
		return 0, err
	}

	if w.appendToBuffer(msg) {
		return len(msg), nil
	}

	return w.File.Write(msg)
}

func (w *fileWriter) Flush() error {
	if w.bufsize == 0 {
		return nil
	}

	w.lock.Lock()
	defer w.lock.Unlock()

	err := w.ensureFileExist()
	if err != nil {
		return err
	}

	return w.flushBuffer()
}

func (w *fileWriter) Free() {
	_ = w.ensureFileExist()

	_ = w.flushBuffer()
	_ = w.File.Close()
}

func openFile(path string) (*os.File, error) {
	switch path {
	case "/dev/stderr":
		return os.Stderr, nil
	case "/dev/stdout":
		return os.Stdout, nil
	}

	return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

func (w *fileWriter) ensureFileExist() error {
	nowTimeSecond := time.Now().Unix()
	if w.lastTimeSecond == nowTimeSecond {
		return nil
	}

	if gomisc.FileExist(w.path) {
		return nil
	}

	_ = w.Close()

	var err error
	w.File, err = openFile(w.path)
	if err != nil {
		return err
	}

	w.lastTimeSecond = nowTimeSecond
	return nil
}

func (w *fileWriter) appendToBuffer(msg []byte) bool {
	after := w.bufpos + len(msg)
	if after >= w.bufsize {
		return false
	}

	copy(w.buf[w.bufpos:], msg)
	w.bufpos = after

	return true
}

func (w *fileWriter) flushBuffer() error {
	if w.bufpos == 0 {
		return nil
	}

	_, err := w.File.Write(w.buf[:w.bufpos])
	w.bufpos = 0

	return err
}
