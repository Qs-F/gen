package gen

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestContainsImports(t *testing.T) {
	tests := []struct {
		Input Variables

		Result bool
	}{
		{
			Input: Variables{
				"import": []string{
					"doc/doc/doc.md",
				},
			},
			Result: true,
		},
		{
			Input: Variables{
				"import": "must fail",
			},
			Result: false,
		},
		{
			Input: Variables{
				"import": []interface{}{
					"doc/doc/doc.md",
				},
			},
			Result: true,
		},
	}

	for _, test := range tests {
		if test.Input.ContainsImports() != test.Result {
			t.Errorf("error on %s\n", test.Input)
		} else {
			t.Logf("ok %s\n", test.Input)
		}
	}
}

func TestResolve(t *testing.T) {
	tests := []struct {
		Input  map[string]Variables
		Output Variables
	}{
		{
			Input: map[string]Variables{
				"_": {
					"test": "Gopher OK",
				},
			},
			Output: Variables{
				"test": "Gopher OK",
			},
		},
		{
			Input: map[string]Variables{
				"_": {
					"import": []string{"A", "B"},
				},
				"A": {
					"test": "Gopher A",
				},
				"B": {
					"test": "Gopher B",
				},
			},
			Output: Variables{
				"test": "Gopher A",
			},
		},
		{
			Input: map[string]Variables{
				"_": {
					"import": []string{"A", "B"},
					"test":   "Gopher X",
				},
				"A": {
					"test": "Gopher A",
				},
				"B": {
					"test": "Gopher B",
				},
			},
			Output: Variables{
				"test": "Gopher X",
			},
		},
		// recurrent named imports
		{
			Input: map[string]Variables{
				"_": {
					"import": []string{"A", "B"},
					"test":   "Gopher X",
					"named": Variables{
						"import":         []string{"A", "B"},
						"notoverwritten": "not",
					},
				},
				"A": {
					"test": "Gopher A",
				},
				"B": {
					"test": "Gopher B",
				},
			},
			Output: Variables{
				"test": "Gopher X",
				"named": Variables{
					"test":           "Gopher A",
					"notoverwritten": "not",
				},
			},
		},
		{
			Input: map[string]Variables{
				"_": {
					"import": []string{"A", "B"},
					"test":   "Gopher X",
					"named": map[string]interface{}{
						"import":         []string{"A", "B"},
						"notoverwritten": "not",
					},
				},
				"A": {
					"test": "Gopher A",
				},
				"B": {
					"test": "Gopher B",
				},
			},
			Output: Variables{
				"test": "Gopher X",
				"named": Variables{
					"test":           "Gopher A",
					"notoverwritten": "not",
				},
			},
		},
	}

	for _, test := range tests {
		if r, err := Resolve(test.Input); err != nil {
			t.Errorf("err: %v\nwant %s\nbut got %s\nfor %s\n", err, spew.Sdump(test.Output), spew.Sdump(r), spew.Sdump(test.Input))
		} else if !reflect.DeepEqual(r, test.Output) {
			t.Errorf("want %s\nbut got %s\nfor %s\n", spew.Sdump(test.Output), spew.Sdump(r), spew.Sdump(test.Input))
		} else {
			t.Logf("got %s\nfor %s\n", spew.Sdump(r), spew.Sdump(test.Input))
		}
	}
}
