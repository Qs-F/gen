package gen_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Qs-F/gen/lib/gen"
	"github.com/Qs-F/gen/lib/loader/markdown"
	"github.com/davecgh/go-spew/spew"
)

var (
	base = filepath.Join("..", "..", "testdata")
)

func TestUnit_GenList_ResolveKey(t *testing.T) {
	tests := []struct {
		Path    string
		Loaders []gen.Loader
		Root    string

		Output gen.Variables
	}{
		{
			Path:    filepath.Join(base, "imports", "md"),
			Loaders: []gen.Loader{markdown.New()},
			Root:    "index.md",

			Output: gen.Variables{
				"title":      "Index Page",
				"committers": map[string]interface{}{"gopher": "Gopher Google"},
			},
		},
	}

	for _, test := range tests {
		list, err := gen.GenList(test.Path, test.Loaders...)
		if err != nil {
			t.Error(err)
			continue
		}
		v, err := gen.ResolveKey(list, test.Root)
		if err != nil {
			t.Error(err)
			continue
		}
		if !reflect.DeepEqual(v, test.Output) {
			t.Errorf("want: \n%s\n, but got: \n%s\n", spew.Sdump(test.Output), spew.Sdump(v))
		} else {
			t.Logf("got: \n%s\n", spew.Sdump(v))
		}
	}
}
