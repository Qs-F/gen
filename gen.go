package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

// Variables is the type expressing front matters
type Variables map[string]interface{}

type Gen struct {
	BasePath   string
	LayoutPath string
}

func New(basePath string, layoutPath string) *Gen {
	return &Gen{
		BasePath:   basePath,
		LayoutPath: layoutPath,
	}
}

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

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return
	}
	gen := New(wd, "")
	list, err := gen.List(NewMarkdownLoader())
	if err != nil {
		logrus.Error(err)
		return
	}
	v, err := ResolveKey(list, "testdata/md/basic.md")
	if err != nil {
		logrus.Error(err)
		return
	}
	spew.Dump(v)
}
