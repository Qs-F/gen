// Package html is implementation of gen/lib/gen.Expander for html
package html

import "github.com/Qs-F/gen/lib/gen"

const (
	from = "html"
	to   = "html"
)

// HTML is struct for html package to implement Expander
type HTML struct{}

// New returns new instance of HTML
func New() *HTML {
	return &HTML{}
}

// Ext implements gen/lib/gen.Expander
func (_ *HTML) Ext() (string, string) {
	return from, to
}

// Expand implements gen/lib/gen.Expander
func (_ *HTML) Expand(p []byte, v gen.Variables) ([]byte, error) {
	return nil, nil
}
