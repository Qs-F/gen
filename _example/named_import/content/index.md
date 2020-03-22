---
layout: layout/index.html
author: たふみ
a:
  import:
  - content/another.md
---

This page is written by {{ .author }}

Another page was written by {{ .a.author }}
