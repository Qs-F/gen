package gen_test

import (
	"bytes"
	"testing"

	"github.com/Qs-F/gen/lib/expander/markdown"
	"github.com/Qs-F/gen/lib/gen"
)

func Test_expand(t *testing.T) {
	tests := []struct {
		List     gen.List
		Expander gen.Expander
		File     string
		Content  []byte

		Dst    string
		Output []byte
	}{
		{
			List: gen.List{
				"a.md": gen.Variables{
					"layout": "a.html",
					"title":  "a page",
					"name":   "gopher",
				},
				"a.html": gen.Variables{
					"content": `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>{{ .title }}</title>
</head>
<body>
	<article>
{{ .__content__ }}
	</article>
</body>
</html>`,
				},
			},
			File:    "a.md",
			Content: []byte("# Title\nhello {{ .name }}"),
			Dst:     "a.html",
			Output: []byte(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>a page</title>
</head>
<body>
	<article>
<h1>Title</h1>

<p>hello gopher</p>

	</article>
</body>
</html>`),
		},
	}

	for _, test := range tests {
		test.Expander = markdown.New("layout", "content", "__content__", test.List)

		dst, w, err := gen.ExpandEach(test.List, test.Expander, test.File, test.Content)
		if err != nil {
			t.Error(err)
		}
		if dst != test.Dst {
			t.Errorf("dst is not satisfied: want %s but got %s", test.Dst, dst)
		}
		if !bytes.Equal(w, test.Output) {
			t.Errorf("want: \n%s\nbut got: \n%s\n", string(test.Output), string(w))
		} else {
			t.Logf("got:\n%s\n", string(w))
		}
	}
}
