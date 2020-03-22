# cmd `gen`

![Gen logo](.github/logo/logo-unified.svg)

cmd `gen` is a tool to output static htmls from markdown.

![Test](https://github.com/Qs-F/gen/workflows/test/badge.svg)
[![GoDoc](https://godoc.org/github.com/Qs-F/gen?status.svg)](https://godoc.org/github.com/Qs-F/gen)
[![Go Report Card](https://goreportcard.com/badge/github.com/Qs-F/gen)](https://goreportcard.com/report/github.com/Qs-F/gen)

## Design

Variable resolve -> text/template expansion to markdown content -> blackfriday convert (-> html/template if needed)

## Installation

```
go get github.com/Qs-F/gen
```

## Usage

```
# Example - see _example directory
gen -base _example/ -src _example/content -dst _example/dist
```

## Features

### `import` directive

in front matter, you can import other markdown.

```markdown
---
import:
- content/a.md
---
```

Likewise, write the path to the file from **base** path **without** first slash.

`import` is useful, like you can import with namespace. For more cases, please see `_example/*/content` directory.

```markdown
---
named: 
  import:
  - content/a.md
---

you can access with {{ .named.hogehoge }} in a.md.
```

### `layout` directive

layout is **optional**. If you want to output whole html with embedded markdown content in html file, this helps.

```markdown
---
title: hello
layout: layout/index.html
---

## Hoge page

This is new page.
```

write the path of layout file from `base` without slash in `layout`.

And in layout file, you have to write special variable `{{ .__content__ }}` to show markdown content.

```html
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
</html>
```

### embedded value

in your markdown file, you can even write following:

```markdown
---
field1: "{{ .field2 }} is the best"
field2: golang
---

{{ field1 }}
```

This will be formed out as

```html 
<p>golang is the best</p>
```

more complex example:

`content/a.md`

```markdown
---
who: "{{ .name }}"
name: "{{ .b.name }}"
b:
  import:
  - content/b.md
display: "{{ .who }} wrote this file"
---

{{ .display }}
```

`content/b.md`

```markdown
---
name: たふみ
---
```

then

`dist/a.html`

```html
<p>たふみ wrote this file</p>
```

Note:

- You **CANNOT** use variable for the yaml key. e.g. `{{ .key }}: hello` is invalid.
- No DFS implmented, so cyclic import can cause undefined behavior.

## License

MIT License

## Copyright

Copyright 2020 de-liKeR / たふみ @CreatorQsF
