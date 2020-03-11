// Package html is the package for load html
package html

import (
	"bytes"
	"errors"

	"github.com/Qs-F/bort"
	"github.com/Qs-F/gen/lib/gen"
)

const ext = ".html"

var (
	// ErrFileBinary is returned when the html file is binary
	ErrFileBinary = errors.New("File is binary")
)

// HTML is struct of HTML Loader
type HTML struct{}

// New returns new instance of HTML
func New() *HTML {
	return &HTML{}
}

// Ext implements gen.Loader
func (_ *HTML) Ext() string {
	return ext
}

// Load implements gen.Loader
func (_ *HTML) Load(p []byte) (gen.Variables, error) {
	b, err := bort.IsBin(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	if b {
		return nil, ErrFileBinary
	}

	v := make(gen.Variables)
	v[gen.ContentIdent] = string(p)
	return v, nil
}
