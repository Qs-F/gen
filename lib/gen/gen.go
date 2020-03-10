// Package gen is the packeg of internal gen
package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Variables is the type expressing front matters
type Variables map[string]interface{}

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

// Loader is the interface that provides the way to get Variables from document.
//
// Ext returns file extension, e.g. md, html.
// If the given file is matched to this, then the loader will be used to resolve.
//
// Load is the function to get Variables from document.
type Loader interface {
	Ext() string
	Load(p []byte) (Variables, error)
}

func (gen *Gen) List(loaders ...Loader) (List, error) {
	list := make(List)
	err := filepath.Walk(gen.BasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		for _, loader := range loaders {
			if filepath.Ext(path) != loader.Ext() {
				continue
			}
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			v, err := loader.Load(b)
			if err != nil {
				return err
			}
			key, err := filepath.Rel(gen.BasePath, path)
			if err != nil {
				return err
			}
			list[key] = v
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}
