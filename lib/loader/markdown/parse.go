// Package markdown is the loader of markdown
package markdown

import (
	"bytes"

	"github.com/Qs-F/gen/lib/gen"
	"github.com/gohugoio/hugo/parser/pageparser"
)

const ext = ".md"

// Markdown is the type of markdown loader
type Markdown struct{}

// New returns new instance of Markdown
func New() *Markdown {
	return &Markdown{}
}

// Ext implements gen.Loader
func (_ *Markdown) Ext() string {
	return ext
}

// Load implements gen.Loader
func (_ *Markdown) Load(p []byte) (gen.Variables, error) {
	cfm, err := pageparser.ParseFrontMatterAndContent(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	return cfm.FrontMatter, nil
}
