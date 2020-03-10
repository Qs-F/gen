package gen

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

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
