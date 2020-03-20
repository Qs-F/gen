// Package html is implementation of gen/lib/gen.Expander for html
package html

import (
	"bytes"
	"html/template"

	"github.com/Qs-F/gen/lib/gen"
)

const (
	from = "phtml"
	to   = "html"
)

// HTML is struct for html package to implement Expander
type HTML struct {
	ContentKey string
}

// New returns new instance of HTML
func New(key string) *HTML {
	return &HTML{ContentKey: key}
}

// Ext implements gen/lib/gen.Expander
func (_ *HTML) Ext() (string, string) {
	return from, to
}

// Expand implements gen/lib/gen.Expander
func (_ *HTML) Expand(p []byte, v gen.Variables) ([]byte, error) {
	tmpl, err := template.New("page").Parse(string(p))
	if err != nil {
		return nil, err
	}

	tmpl = tmpl.Option("missingkey=error")

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, v)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
