package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Input string

		Content     []byte
		FrontMatter map[string]interface{}
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
			FrontMatter: map[string]interface{}{
				"hello": "Gopher",
			},
		},
	}

	for _, test := range tests {
		c, fm, err := Parse([]byte(test.Input))
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(c, test.Content) || !reflect.DeepEqual(fm, test.FrontMatter) {
			t.Errorf("want c: %s\n\nfm: %s\n but got c: %s\n\nfm: %s\n", test.Content, test.FrontMatter, c, fm)
		} else {
			t.Logf("got c: %s\n\nfm: %s\n", c, fm)
		}
	}
}
