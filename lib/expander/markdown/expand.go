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
	// ErrLayoutFileNotFound is retunred when layout file is not found in list
	ErrLayoutFileNotFound = errors.New("No such layout file is found")
)

// Markdown satisfies gen.Expander
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
	cfm, err := pageparser.ParseFrontMatterAndContent(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	file := cfm.Content
	if string(cfm.Content) == "" && isEmptyMap(cfm.FrontMatter) {
		file = p
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
	var buf bytes.Buffer
	ret := page
	for {
		tmpl, err := template.New("page").Parse(ret)
		if err != nil {
			return nil, err
		}

		tmpl = tmpl.Option("missingkey=error")

		buf.Reset()
		err = tmpl.Execute(&buf, v)
		if err != nil {
			return nil, err
		}
		if ret == buf.String() {
			break
		}
		ret = buf.String()
	}

	return []byte(ret), nil
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

func isEmptyMap(m map[string]interface{}) bool {
	if m == nil {
		return true
	}
	for _ = range m {
		return false
	}
	return true
}
