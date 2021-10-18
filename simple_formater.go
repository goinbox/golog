package golog

import (
	"fmt"
	"strings"
)

type simpleFormater struct {
}

func NewSimpleFormater() *simpleFormater {
	return &simpleFormater{}
}

func (f *simpleFormater) Format(fields ...*Field) []byte {
	list := make([]string, len(fields))
	for i, field := range fields {
		list[i] = fmt.Sprintf("%s:%v", field.Key, field.Value)
	}

	return []byte(strings.Join(list, "\t"))
}
