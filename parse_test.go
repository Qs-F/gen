package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Input string

		Content   []byte
		Variables Variables
	}{
		{
			Input: `---
hello: Gopher
---
# test doc
あいうえお😄
`,
			Content: []byte(`# test doc
あいうえお😄
`),
			Variables: Variables{
				"hello": "Gopher",
			},
		},
	}

	for _, test := range tests {
		c, fm, err := Parse([]byte(test.Input))
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(c, test.Content) || !reflect.DeepEqual(fm, test.Variables) {
			t.Errorf("want c: %s\n\nfm: %s\n but got c: %s\n\nfm: %s\n", test.Content, test.Variables, c, fm)
		} else {
			t.Logf("got c: %s\n\nfm: %s\n", c, fm)
		}
	}
}
