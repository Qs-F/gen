---
import: 
- content/index.md
committers:
- name: Gopher A
  url: golang.org
- name: Gopher B
  url: pkg.go.dev
---

{{ range $v := .committers }}
- {{ $v.name }} : {{ $v.url }}
{{ end }}

