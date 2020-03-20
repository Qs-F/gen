// Package gen is the packeg of internal gen
package gen

import (
	"errors"
)

// Variables is the type expressing front matters
type Variables map[string]interface{}

// Copy copies and returns new Variables
func (vs Variables) Copy() Variables {
	ret := make(Variables)
	for k, v := range vs {
		ret[k] = v
	}
	return ret
}

var (
	// ErrContentNotFound is returned when Field 'Content' is not found in Variables
	ErrContentNotFound = errors.New("Field 'Content' is not found in Variables")
)

// Gen is the struct for gen cmd
type Gen struct {
	BasePath string
	SrcPath  string
	DstPath  string
}

// New returns new *Gen
func New(basePath string, srcPath string, dstPath string) *Gen {
	return &Gen{
		BasePath: basePath,
		SrcPath:  srcPath,
		DstPath:  dstPath,
	}
}
