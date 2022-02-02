package core

import (
	"bytes"
	"html/template"
	"os"
)

/**
 * View Engine
 */

// LoadView file
func LoadView(view string, data ...interface{}) string {
	// Open view file
	var viewText, err = os.ReadFile(view)
	CheckErrorNotPanic(err)
	parse, err := template.New(view).Parse(string(viewText))
	CheckErrorNotPanic(err)
	buf := &bytes.Buffer{}
	var i = 0
	var parsed = false
	for _, dataToParse := range data {
		if i < 1 {
			err = parse.Execute(buf, dataToParse)
			CheckErrorNotPanic(err)
			parsed = true
			break
		}
		i++
	}
	if !parsed {
		err = parse.Execute(buf, data)
		CheckErrorNotPanic(err)
		parsed = true
	}
	s := buf.String()
	return s
}
