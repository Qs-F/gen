# cmd `gen`

cmd `gen` is a tool to output static htmls from markdown.

## Design

Variable resolve -> text/template expansion to markdown content -> blackfriday convert (-> html/template if needed)

## Default values

- import directive

in front matter, you can import other markdown
