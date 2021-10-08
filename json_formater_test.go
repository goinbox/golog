package golog

import "testing"

func TestJsonFormater(t *testing.T) {
	f := NewJsonFormater()
	fields := []*Field{
		{
			Key:   "kint",
			Value: 1,
		},
		{
			Key:   "kstring",
			Value: "abc",
		},
		{
			Key: "struct",
			Value: struct {
				Name string
				Age  int
			}{Name: "aaa", Age: 10},
		},
	}

	msg := f.Format(fields...)

	t.Log(string(msg))
}
