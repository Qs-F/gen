// Package html is the package for load html
package html

import (
	"bytes"
	"errors"
	"unicode"

	"github.com/Qs-F/bort"
	"github.com/Qs-F/gen/lib/gen"
)

const ext = ".html"

var (
	// ErrFileBinary is returned when the html file is binary
	ErrFileBinary = errors.New("File is binary")
)

// HTML is struct of HTML Loader
type HTML struct {
	ContentKey string
}

// New returns new instance of HTML
func New(countentKey string) *HTML {
	return &HTML{ContentKey: countentKey}
}

// Ext implements gen.Loader
func (_ *HTML) Ext() string {
	return ext
}

// Load implements gen.Loader
func (h *HTML) Load(p []byte) (gen.Variables, error) {
	b, err := bort.IsBin(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	if b {
		return nil, ErrFileBinary
	}

	for _, v := range h.ContentKey {
		if unicode.IsSpace(v) {
			return nil, nil
		}
	}
	v := make(gen.Variables)
	v[h.ContentKey] = string(p)
	return v, nil
}
