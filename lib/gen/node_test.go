package gen

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_link(t *testing.T) {
	tests := []struct {
		Target *Node
		List   List

		Output *Node
	}{
		{
			Target: &Node{
				Deps: []string{"A", "B", "C"},
			},
			List: List{
				"A": Variables{
					"test": "Gopher",
				},
				"B": Variables{
					"test": "Gopher B",
				},
			},
			Output: &Node{
				Deps: []string{"A", "B", "C"},
				children: []*Node{
					{
						Key:      "A",
						Content:  Variables{"test": "Gopher"},
						Deps:     []string{},
						children: []*Node{},
					},
					{
						Key:      "B",
						Content:  Variables{"test": "Gopher B"},
						Deps:     []string{},
						children: []*Node{},
					},
				},
				resolved: true,
			},
		},
	}

	for _, test := range tests {
		test.Target.link(test.List)

		if !reflect.DeepEqual(test.Target, test.Output) {
			t.Errorf("want: \n%s\nbut got: \n%s\n", spew.Sdump(test.Output), spew.Sdump(test.Target))
		} else {
			t.Logf("got: \n%s\n", spew.Sdump(test.Target))
		}
	}
}
