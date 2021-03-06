// Package gen is the packeg of internal gen
package gen

import (
	"path/filepath"
	"strings"
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

// String implements fmt.Stringer
func (vs Variables) String() string {
	s := []string{}
	for k := range vs {
		s = append(s, k, "\n")
	}
	return strings.Join(s, "")
}

func MapToVariables(m map[string]interface{}) Variables {
	if m == nil {
		return nil
	}
	ret := make(Variables)
	for k, v := range m {
		ret[k] = v
	}
	return ret
}

// Gen is the struct for gen cmd
type Gen struct {
	BasePath string // must be absolute
	SrcPath  string // must be relative
	DstPath  string // must be relative
}

// New returns new *Gen
func New(base, src, dst string) *Gen {
	return &Gen{
		BasePath: base,
		SrcPath:  src,
		DstPath:  dst,
	}
}

func (gen *Gen) Set(basePath, srcPath, dstPath string) error {
	base, err := filepath.Abs(basePath)
	if err != nil {
		return err
	}
	srcAbs, err := filepath.Abs(srcPath)
	if err != nil {
		return err
	}
	dstAbs, err := filepath.Abs(dstPath)
	if err != nil {
		return err
	}
	src, err := filepath.Rel(base, srcAbs)
	if err != nil {
		return err
	}
	dst, err := filepath.Rel(base, dstAbs)
	if err != nil {
		return err
	}
	gen.BasePath = base
	gen.SrcPath = src
	gen.DstPath = dst
	return nil
}
