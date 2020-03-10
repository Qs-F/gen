// Package gen is the packeg of internal gen
package gen

import (
	"bytes"
	"errors"
	"html/template"
)

// Variables is the type expressing front matters
type Variables map[string]interface{}

const (
	ContentIdent = "Content"
)

var (
	// ErrContentNotFound is returned when Field 'Content' is not found in Variables
	ErrContentNotFound = errors.New("Field 'Content' is not found in Variables")
)

// Gen is the struct for gen cmd
type Gen struct {
	BasePath   string
	LayoutPath string
}

// New returns new *Gen
func New(basePath string, layoutPath string) *Gen {
	return &Gen{
		BasePath:   basePath,
		LayoutPath: layoutPath,
	}
}

func Expand(v Variables) (string, error) {
	ti, ok := v[ContentIdent]
	if !ok {
		return "", ErrContentNotFound
	}
	t, ok := ti.(string)
	if !ok {
		return "", ErrContentNotFound
	}

	tmpl, err := template.New("page").Parse(t)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, v)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
