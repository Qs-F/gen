package markdown

import (
	"bytes"
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

func TestExpandWithoutLayout(t *testing.T) {
	tests := []struct {
		Content    []byte
		Variables  gen.Variables
		ContentKey string

		Output   []byte
		MustFail bool
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
	}

	md := New("layout", "__content__", "content", gen.List{})

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
	}
}

func TestExpand(t *testing.T) {
	tests := []struct {
		Markdown   []byte
		Var        gen.Variables
		LayoutKey  string
		HTMLKey    string
		ContentKey string
		List       gen.List

		Output []byte
	}{
		{
			Markdown: []byte("## Title\nI'm **{{ .name }}**"),
			Var: gen.Variables{
				"name":   "Gopher",
				"title":  "New Page",
				"layout": "default.html",
			},
			LayoutKey:  "layout",
			HTMLKey:    "content",
			ContentKey: "__markdown__",
			List: gen.List{
				"default.html": gen.Variables{
					"content": `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>{{ .title }}</title>
</head>
<body>
	<h1>Hello World</h1>
	<article>
{{ .__markdown__ }}
	</article>
</body>
</html>`,
				},
			},
			Output: []byte(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>New Page</title>
</head>
<body>
	<h1>Hello World</h1>
	<article>
<h2>Title</h2>

<p>I&rsquo;m <strong>Gopher</strong></p>

	</article>
</body>
</html>`),
		},
	}

	for _, test := range tests {
		md := New(test.LayoutKey, test.HTMLKey, test.ContentKey, test.List)
		b, err := md.Expand(test.Markdown, test.Var)
		if err != nil {
			t.Error(err)
			continue
		}
		if !bytes.Equal(b, test.Output) {
			t.Errorf("want: \n%s\nbut got: \n%s\n", string(test.Output), string(b))
		} else {
			t.Logf("got: \n%s\n", string(b))
		}
	}
}

func TestSatisfy(t *testing.T) {
	var _ gen.Expander = (*Markdown)(nil)
}
