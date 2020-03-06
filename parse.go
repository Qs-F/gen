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
