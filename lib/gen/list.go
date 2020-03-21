package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// List is all Variable list.
// map key is the key to access Variable
type List map[string]Variables

// String implements fmt.Stringer
func (list List) String() string {
	s := []string{}
	for k := range list {
		s = append(s, k, "\n")
	}
	return strings.Join(s, "")
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

// GenList makes list from given directory.
func GenList(basePath string, loaders ...Loader) (List, error) {
	list := make(List)
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
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
			key, err := filepath.Rel(basePath, path)
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
