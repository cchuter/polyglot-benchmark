# Scope: polyglot-go-markdown (Issue #264)

## In Scope

- Implement `Render` function in `go/exercises/practice/markdown/markdown.go`
- Handle all Markdown features tested: paragraphs, headings (h1-h6), bold, italic, unordered lists
- Handle edge cases: h7+ as paragraphs, markdown symbols in text, list open/close boundaries
- Pass all existing tests in `cases_test.go` and `markdown_test.go`

## Out of Scope

- Modifying test files (`cases_test.go`, `markdown_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` files
- Supporting Markdown features not tested (e.g., links, images, ordered lists, code blocks)
- Adding external dependencies
- Nested lists or multi-level headings within lists

## Dependencies

- Go standard library only (`fmt`, `strings`)
- No external packages
