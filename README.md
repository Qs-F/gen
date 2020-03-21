# cmd `gen`

cmd `gen` is a tool to output static htmls from markdown.

![Test](https://github.com/Qs-F/gen/workflows/test/badge.svg)
[![GoDoc](https://godoc.org/github.com/Qs-F/gen?status.svg)](https://godoc.org/github.com/Qs-F/gen)
[![Go Report Card](https://goreportcard.com/badge/github.com/Qs-F/gen)](https://goreportcard.com/report/github.com/Qs-F/gen)

## Design

Variable resolve -> text/template expansion to markdown content -> blackfriday convert (-> html/template if needed)

## Default values

- `import` directive

in front matter, you can import other markdown
