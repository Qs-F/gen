package markdown

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Qs-F/gen/lib/gen"
)

func Test_text(t *testing.T) {
	tests := []struct {
		Content   []byte
		Variables gen.Variables

		Output   []byte
		MustFail bool
	}{
		{
			Content: []byte("Test: {{ .title }}"),
			Variables: gen.Variables{
				"title": "Title Expansion",
			},
			Output: []byte("Test: Title Expansion"),
		},
		{
			Content:   []byte("Test: {{ .title }}"),
			Variables: gen.Variables{},
			MustFail:  true,
		},
	}

	for _, test := range tests {
		b, err := text(test.Content, test.Variables)
		if err != nil {
			if !test.MustFail {
				t.Error(err)
			} else {
				t.Log(err)
			}
			continue
		}
		if !bytes.Equal(b, test.Output) {
			t.Errorf("want: \n%s\n but got: \n%s\n", string(test.Output), string(b))
		} else {
			t.Logf("got: \n%s\n", string(b))
		}
	}
}

func TestExpand(t *testing.T) {
	tests := []struct {
		Content    []byte
		Variables  gen.Variables
		ContentKey string

		Output       []byte
		SideEffect   bool
		ModVariables gen.Variables
		MustFail     bool
	}{
		{
			Content: []byte("Test: {{ .title }}"),
			Variables: gen.Variables{
				"title": "Title Expansion",
			},
			Output: []byte("<p>Test: Title Expansion</p>\n"),
		},
		{
			Content:   []byte("Test: {{ .title }}"),
			Variables: gen.Variables{},
			MustFail:  true,
		},
		{
			Content: []byte("Test: {{ .title }}"),
			Variables: gen.Variables{
				"title": "<span>Test</span>",
			},
			Output: []byte("<p>Test: <span>Test</span></p>\n"),
		},
		{
			Content: []byte("Test: {{ .title }}"),
			Variables: gen.Variables{
				"title": "Title Expansion",
			},
			ContentKey: "__content__",
			Output:     []byte("<p>Test: Title Expansion</p>\n"),
			SideEffect: true,
			ModVariables: gen.Variables{
				"title":       "Title Expansion",
				"__content__": "<p>Test: Title Expansion</p>\n",
			},
		},
	}

	md := New("__content__")

	for _, test := range tests {
		b, err := md.Expand(test.Content, test.Variables)
		if err != nil {
			if !test.MustFail {
				t.Error(err)
			} else {
				t.Log(err)
			}
			continue
		}
		if !bytes.Equal(b, test.Output) {
			t.Errorf("want: \n%s\n but got: \n%s\n", string(test.Output), string(b))
		} else {
			t.Logf("got: \n%s\n", string(b))
		}
		if test.SideEffect && !reflect.DeepEqual(test.Variables, test.ModVariables) {
			t.Errorf("want: \n%v\n but got: \n%v\n", test.ModVariables, test.Variables)
		} else {
			t.Logf("got: \n%v\n", test.Variables)
		}
	}
}

func TestSatisfy(t *testing.T) {
	var _ gen.Expander = (*Markdown)(nil)
}