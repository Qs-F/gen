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
				"_": Variables{
					"test": "Gopher OK",
				},
			},
			Output: Variables{
				"test": "Gopher OK",
			},
		},
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
		{
			Input: map[string]Variables{
				"_": Variables{
					"import": []string{"A", "B"},
					"test":   "Gopher X",
				},
				"A": Variables{
					"test": "Gopher A",
				},
				"B": Variables{
					"test": "Gopher B",
				},
			},
			Output: Variables{
				"test": "Gopher X",
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

func TestGenTree(t *testing.T) {
	tests := []struct {
		List List
		Root string

		Tree Tree
	}{
		{
			List: List{
				"A": Variables{
					"test":   "Hello Gopher",
					"import": []string{"B", "C"},
				},
				"B": Variables{
					"b": "Hello B",
				},
				"C": Variables{
					"test": "Hello C",
				},
			},
			Root: "A",
			Tree: Tree{
				0: &Depth{
					Nodes: []*Node{
						0: {
							Key: "A",
							Content: Variables{
								"test":   "Hello Gopher",
								"import": []string{"B", "C"},
							},
							Deps: []string{"B", "C"},
							children: []*Node{
								0: {
									Key: "B",
									Content: Variables{
										"b": "Hello B",
									},
									Deps:     []string{},
									children: []*Node{},
								},
								1: {
									Key: "C",
									Content: Variables{
										"test": "Hello C",
									},
									Deps:     []string{},
									children: []*Node{},
								},
							},
							resolved: true,
						},
					},
				},
				1: &Depth{
					Nodes: []*Node{
						0: {
							Key: "B",
							Content: Variables{
								"b": "Hello B",
							},
							Deps:     []string{},
							children: []*Node{},
						},
						1: {
							Key: "C",
							Content: Variables{
								"test": "Hello C",
							},
							Deps:     []string{},
							children: []*Node{},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		tree := test.List.GenTree(test.List[test.Root].ToNode(test.Root, nil))
		if !reflect.DeepEqual(tree, test.Tree) {
			t.Errorf("want: \n%s\n but got: \n%s\n", spew.Sdump(test.Tree), spew.Sdump(tree))
		} else {
			t.Logf("got: \n%s\n", spew.Sdump(tree))
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		Tree      Tree
		Variables Variables
	}{
		{
			Tree: Tree{
				0: &Depth{
					Nodes: []*Node{
						0: {
							Key: "A",
							Content: Variables{
								"test":   "Hello Gopher",
								"import": []string{"B", "C"},
							},
							Deps: []string{"B", "C"},
							children: []*Node{
								0: {
									Key: "B",
									Content: Variables{
										"b": "Hello B",
									},
									Deps:     []string{},
									children: []*Node{},
								},
								1: {
									Key: "C",
									Content: Variables{
										"test": "Hello C",
									},
									Deps:     []string{},
									children: []*Node{},
								},
							},
							resolved: true,
						},
					},
				},
				1: &Depth{
					Nodes: []*Node{
						0: {
							Key: "B",
							Content: Variables{
								"b": "Hello B",
							},
							Deps:     []string{},
							children: []*Node{},
						},
						1: {
							Key: "C",
							Content: Variables{
								"test": "Hello C",
							},
							Deps:     []string{},
							children: []*Node{},
						},
					},
				},
			},
			Variables: Variables{
				"test": "Hello Gopher",
				"b":    "Hello B",
			},
		},
	}

	for _, test := range tests {
		if v := test.Tree.Reduce(); !reflect.DeepEqual(v, test.Variables) {
			t.Errorf("want: \n%s\nbut got:\n%s\n", spew.Sdump(test.Variables), spew.Sdump(v))
		} else {
			t.Logf("got: \n%s\n", spew.Sdump(v))
		}
	}
}
