package main

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
				"_": Variables{
					"import": []string{"A", "B"},
				},
				"A": Variables{
					"test": "Gopher A",
				},
				"B": Variables{
					"test": "Gopher B",
				},
			},
			Output: Variables{
				"test": "Gopher A",
			},
		},
	}

	for _, test := range tests {
		if r, err := Resolve(test.Input); err != nil {
			t.Errorf("err: %v\nwant %s\nbut got %s\nfor %s\n", err, spew.Sdump(test.Output), spew.Sdump(r), spew.Sdump(test.Input))
		} else if !reflect.DeepEqual(r, test.Output) {
			t.Errorf("want %s\nbut got %s\nfor %s\n", spew.Sdump(test.Output), spew.Sdump(r), spew.Sdump(test.Input))
		} else {
			t.Errorf("got %s\nfor %s\n", spew.Sdump(r), spew.Sdump(test.Input))
		}
	}
}
