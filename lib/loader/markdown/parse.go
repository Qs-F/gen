package markdown

import (
	"bytes"

	"github.com/Qs-F/gen/lib/gen"
	"github.com/gohugoio/hugo/parser/pageparser"
)

const (
	ContentIdent = "Content"
)

const ext = ".md"

type Markdown struct{}

func NewMarkdownLoader() *Markdown {
	return &Markdown{}
}

func (_ *Markdown) Ext() string {
	return ext
}

func (_ *Markdown) Load(p []byte) (gen.Variables, error) {
	var v gen.Variables
	cfm, err := pageparser.ParseFrontMatterAndContent(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	v = cfm.FrontMatter
	v[ContentIdent] = string(cfm.Content)
	return v, nil
}
