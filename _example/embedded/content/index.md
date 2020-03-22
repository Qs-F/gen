---
name: "{{ .another.name }}"
another:
  import:
  - content/another.md
field: "{{ .field2 }} is great"
field2: golang
---

{{ .name }}

{{ .field }}
