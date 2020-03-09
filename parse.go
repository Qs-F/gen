package main

import (
	"bytes"

	"github.com/gohugoio/hugo/parser/pageparser"
)

func Parse(p []byte) (content []byte, frontmatter Variables, err error) {
	var cfm pageparser.ContentFrontMatter
	cfm, err = pageparser.ParseFrontMatterAndContent(bytes.NewReader(p))
	if err != nil {
		return
	}
	content = cfm.Content
	frontmatter = cfm.FrontMatter
	return
}

type Markdown struct{}

func (_ *Markdown) Ext() string {
	return ".md"
}

func (_ *Markdown) Load(p []byte) (Variables, error) {
	_, fm, err := Parse(p)
	if err != nil {
		return nil, err
	}
	return fm, nil
}

func NewMarkdownLoader() *Markdown {
	return &Markdown{}
}
