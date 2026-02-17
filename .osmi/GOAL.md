# Goal: polyglot-go-markdown (Issue #264)

## Problem Statement

Implement a Markdown-to-HTML parser in Go. The exercise provides an empty `markdown.go` stub and a reference solution in `.meta/example.go`. The task is to write a clean, readable `Render` function that converts a subset of Markdown syntax into HTML.

The supported Markdown features are:
- **Paragraphs**: Plain text lines become `<p>` elements
- **Headings**: Lines starting with `#` (h1) through `######` (h6); `#######` is treated as a paragraph
- **Bold**: Text wrapped in `__` becomes `<strong>`
- **Italic**: Text wrapped in `_` becomes `<em>`
- **Unordered lists**: Lines starting with `* ` become `<li>` items wrapped in `<ul>`

## Acceptance Criteria

1. All 17 test cases in `cases_test.go` pass when running `go test`
2. The `Render(markdown string) string` function is exported from package `markdown`
3. The code is clean, readable, and well-structured
4. Markdown symbols (`#`, `*`) appearing in text content (not as syntax) are preserved literally
5. Lists open and close properly with preceding and following non-list lines
6. Bold (`__`) is processed before italic (`_`) to avoid conflicts

## Key Constraints

- Package name must be `markdown`
- Module is `markdown` with `go 1.18`
- Must export a single `Render` function matching the test signature
- No external dependencies allowed (stdlib only)
