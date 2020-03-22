---
name: "{{ .another.name }}"
another:
  import:
  - content/another.md
---

{{ .name }}
