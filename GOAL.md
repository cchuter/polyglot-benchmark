# Goal: polyglot-go-markdown

## Problem Statement

Implement the `Render` function in `go/exercises/practice/markdown/markdown.go` that parses a Markdown string and returns the corresponding HTML. The current stub file only contains the package declaration; the full implementation is needed.

This is a **refactoring exercise** — the reference solution in `.meta/example.go` demonstrates clean, readable code. The implementation should be well-structured and maintainable.

## What the Function Must Do

`Render(markdown string) string` takes a Markdown-formatted string and returns HTML. It must support:

1. **Paragraphs** — plain text lines become `<p>...</p>`
2. **Headings** — `#` through `######` become `<h1>` through `<h6>`; 7+ `#` characters are treated as a paragraph
3. **Bold** — `__text__` becomes `<strong>text</strong>`
4. **Italic** — `_text_` becomes `<em>text</em>`
5. **Unordered lists** — lines starting with `* ` become `<li>` items wrapped in `<ul>`
6. **Mixed content** — headings, lists, and paragraphs can appear in sequence; lists must properly open/close `<ul>` tags

## Acceptance Criteria

1. All 17 test cases in `cases_test.go` pass (`go test ./...` exits 0)
2. The `Render` function is exported from the `markdown` package
3. Code is clean, readable, and well-structured
4. No external dependencies beyond the Go standard library

## Key Constraints

- Package name must be `markdown`
- Module is `go 1.18`
- Only file to modify: `markdown.go`
- Tests are read-only (auto-generated from exercism)
