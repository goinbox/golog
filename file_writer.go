/**
* @file file.go
* @brief writer msg to file
* @author ligang
* @date 2016-02-03
 */

package golog

import (
	"errors"
	"os"
	"sync"
	"time"
)

/**
* @name file writer
* @{ */

type FileWriter struct {
	path        string
	lock        *sync.Mutex
	closeOnFree bool

	*os.File
}

func NewFileWriter(path string) (*FileWriter, error) {
	f, err := openFile(path)
	if err != nil {
		return nil, err
	}

	return &FileWriter{
		path:        path,
		lock:        new(sync.Mutex),
		closeOnFree: false,

		File: f,
	}, nil
}

func (this *FileWriter) CloseOnFree(closeOneFree bool) *FileWriter {
	this.closeOnFree = closeOneFree

	return this
}

func (this *FileWriter) Write(msg []byte) (int, error) {
	// file may be deleted when doing logrotate
	if !FileExist(this.path) {
		this.Close()
		this.File, _ = openFile(this.path)
	}

	this.lock.Lock()
	n, err := this.File.Write(msg)
	this.lock.Unlock()

	return n, err
}

func (this *FileWriter) Flush() error {
	return nil
}

func (this *FileWriter) Free() {
	if this.closeOnFree {
		this.File.Close()
	}
}

func openFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

/**  @} */

/**
* @name file writer with split
* @{ */

const (
	SPLIT_BY_DAY  = 1
	SPLIT_BY_HOUR = 2
)

type FileWithSplitWriter struct {
	path   string
	split  int
	suffix string

	*FileWriter
}

func NewFileWriterWithSplit(path string, split int) (*FileWithSplitWriter, error) {
	suffix := makeFileSuffix(split)
	if suffix == "" {
		return nil, errors.New("Split not support")
	}

	f, err := NewFileWriter(path + "." + suffix)
	if err != nil {
		return nil, err
	}

	this := &FileWithSplitWriter{
		path:   path,
		split:  split,
		suffix: suffix,

		FileWriter: f,
	}

	return this, nil
}

func (this *FileWithSplitWriter) Write(msg []byte) (int, error) {
	suffix := makeFileSuffix(this.split)

	//need split
	if suffix != this.suffix {
		this.Free()
		this.FileWriter, _ = NewFileWriter(this.path + "." + suffix)
		this.suffix = suffix
	}

	return this.File.Write(msg)
}

func makeFileSuffix(split int) string {
	switch split {
	case SPLIT_BY_DAY:
		return time.Now().Format(TIME_FMT_STR_YEAR + TIME_FMT_STR_MONTH + TIME_FMT_STR_DAY)
	case SPLIT_BY_HOUR:
		return time.Now().Format(TIME_FMT_STR_YEAR + TIME_FMT_STR_MONTH + TIME_FMT_STR_DAY + TIME_FMT_STR_HOUR)
	default:
		return ""
	}
}

/**  @} */
