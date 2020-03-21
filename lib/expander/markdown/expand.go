// Package markdown is the expander for markdown to html
package markdown

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/Qs-F/gen/lib/gen"
	"github.com/gohugoio/hugo/parser/pageparser"

	bf "github.com/russross/blackfriday/v2"
)

const (
	from = ".md"
	to   = ".html"
)

const (
	noValueIdent = "<no value>"
)

var (
	ErrLayoutFileNotFound = errors.New("No such layout file is found")
)

// Makrdown satisfies gen.Expander
type Markdown struct {
	LayoutKey  string
	HTMLKey    string
	ContentKey string
	List       gen.List
}

// New returns new instance of Markdown.
// layoutKey is ident of layout file in markdown file.
// htmlKey is Variables key of content for html loader.
// contentKey is content ident in html file for markdown content.
func New(layoutKey, htmlKey, contentKey string, list gen.List) *Markdown {
	return &Markdown{
		LayoutKey:  layoutKey,
		HTMLKey:    htmlKey,
		ContentKey: contentKey,
		List:       list,
	}
}

// Ext implements gen.Expander
func (_ *Markdown) Ext() (string, string) {
	return from, to
}

// Expand implements gen.Expander
func (m *Markdown) Expand(p []byte, v gen.Variables) ([]byte, error) {
	file := p
	cfm, err := pageparser.ParseFrontMatterAndContent(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	if string(cfm.Content) != "" {
		file = cfm.Content
	}

	b, err := text(file, v)
	if err != nil {
		return nil, err
	}

	c := markdown(b)

	keyI, ok := v[m.LayoutKey]
	if !ok {
		return c, nil
	}
	key, ok := keyI.(string)
	if !ok {
		return c, nil
	}
	layoutV, ok := m.List[key]
	if !ok {
		return c, nil
	}
	layoutI, ok := layoutV[m.HTMLKey]
	if !ok {
		return c, ErrLayoutFileNotFound
	}
	layout, ok := layoutI.(string)
	if !ok {
		return c, nil
	}

	nv := v.Copy()
	nv[m.ContentKey] = string(c)
	return html(layout, nv)
}

func html(page string, v gen.Variables) ([]byte, error) {
	tmpl, err := template.New("page").Parse(page)
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
