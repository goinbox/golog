package golog

import "encoding/json"

type jsonFormater struct {
}

func NewJsonFormater() *jsonFormater {
	return &jsonFormater{}
}

func (f *jsonFormater) Format(fields ...*Field) []byte {
	m := make(map[string]interface{})
	for _, field := range fields {
		m[field.Key] = field.Value
	}

	msg, _ := json.Marshal(m)
	return msg
}
