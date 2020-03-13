// Package markdown is the expander for markdown to html
package markdown

import (
	"bytes"
	"text/template"

	"github.com/Qs-F/gen/lib/gen"

	bf "github.com/russross/blackfriday/v2"
)

const (
	from = "md"
	to   = "html"
)

const (
	noValueIdent = "<no value>"
)

type Markdown struct{}

func New() *Markdown {
	return &Markdown{}
}

func (_ *Markdown) Ext() (string, string) {
	return from, to
}

func (_ *Markdown) Expand(p []byte, v gen.Variables) ([]byte, error) {
	b, err := text(p, v)
	if err != nil {
		return nil, err
	}
	return markdown(b), nil
}

func text(p []byte, v gen.Variables) ([]byte, error) {
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

func markdown(p []byte) []byte {
	return bf.Run(p, bf.WithExtensions(bf.CommonExtensions))
}
