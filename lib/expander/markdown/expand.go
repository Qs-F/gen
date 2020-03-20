// Package markdown is the expander for markdown to html
package markdown

import (
	"bytes"
	"text/template"
	"unicode"

	"github.com/Qs-F/gen/lib/gen"

	bf "github.com/russross/blackfriday/v2"
)

const (
	from = "md"
	to   = "phtml"
)

const (
	noValueIdent = "<no value>"
)

type Markdown struct {
	ContentKey string
}

func New(key string) *Markdown {
	return &Markdown{ContentKey: key}
}

func (_ *Markdown) Ext() (string, string) {
	return from, to
}

func (m *Markdown) Expand(p []byte, v gen.Variables) ([]byte, error) {
	b, err := text(p, v)
	if err != nil {
		return nil, err
	}
	c := markdown(b)
	for _, r := range m.ContentKey {
		if unicode.IsSpace(r) {
			return c, nil
		}
	}
	v[m.ContentKey] = string(c)
	return c, nil
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
