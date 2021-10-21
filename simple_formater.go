package golog

import (
	"fmt"

	"github.com/goinbox/gomisc"
)

type simpleFormater struct {
	formater Formater
}

func NewSimpleFormater() *simpleFormater {
	return &simpleFormater{
		formater: NewJsonFormater(),
	}
}

func (f *simpleFormater) SetFormater(formater Formater) *simpleFormater {
	f.formater = formater

	return f
}

func (f *simpleFormater) Format(fields ...*Field) []byte {
	var msg []byte
	var notPresetFields []*Field
	for _, field := range fields {
		if field.preset {
			msg = gomisc.AppendBytes(msg, []byte(fmt.Sprintf("%v", field.Value)), []byte("\t"))
		} else {
			notPresetFields = append(notPresetFields, field)
		}
	}

	return gomisc.AppendBytes(msg, f.formater.Format(notPresetFields...))
}
