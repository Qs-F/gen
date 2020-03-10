package gen_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Qs-F/gen/lib/gen"
	"github.com/Qs-F/gen/lib/loader/markdown"
	"github.com/davecgh/go-spew/spew"
)

func TestGenList(t *testing.T) {
	tests := []struct {
		Path    string
		Loaders []gen.Loader
		Root    string

		Output gen.List
	}{
		{
			Path:    filepath.Join(base, "imports", "md"),
			Loaders: []gen.Loader{markdown.NewMarkdownLoader()},
			Root:    "index.md",

			Output: gen.List{
				"index.md": gen.Variables{
					"title":   "Index Page",
					"import":  []interface{}{"list.md"},
					"Content": "\n# {{ .title }}\n\n## this article is for test\n\n{{ .committers.gopher }} wrote.\n",
				},
				"list.md": gen.Variables{
					"committers": map[string]interface{}{"gopher": "Gopher Google"},
					"Content":    "",
				},
			},
		},
	}

	for _, test := range tests {
		list, err := gen.GenList(test.Path, test.Loaders...)
		if err != nil {
			t.Error(err)
			continue
		}
		if !reflect.DeepEqual(list, test.Output) {
			t.Errorf("want: \n%s\n, but got: \n%s\n", spew.Sdump(test.Output), spew.Sdump(list))
		} else {
			t.Logf("got: \n%s\n", spew.Sdump(list))
		}
	}
}
