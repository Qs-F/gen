package markdown

import (
	"reflect"
	"testing"

	"github.com/Qs-F/gen/lib/gen"
)

func TestLoad(t *testing.T) {
	m := New()

	tests := []struct {
		Input string

		Variables gen.Variables
	}{
		{
			Input: `---
hello: Gopher
---
# test doc
あいうえお😄
`,
			Variables: gen.Variables{
				"hello": "Gopher",
				"Content": `# test doc
あいうえお😄
`,
			},
		},
	}

	for _, test := range tests {
		v, err := m.Load([]byte(test.Input))
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(v, test.Variables) {
			t.Errorf("want v: %s\n\nbut got v: %s\n\n", test.Variables, v)
		} else {
			t.Logf("got v: %s\n", v)
		}
	}
}
