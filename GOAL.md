# Goal: polyglot-go-markdown

## Problem Statement

The `markdown` exercise in `go/exercises/practice/markdown/` has an empty stub file (`markdown.go`) that needs a working implementation. The exercise is a "refactoring exercise" — the task is to provide a clean, readable Markdown-to-HTML parser that exports a `Render(markdown string) string` function.

The parser must handle:
- Headings (`#` through `######`, with `#######` treated as a paragraph)
- Paragraphs (plain text lines)
- Unordered lists (`* ` prefixed lines)
- Bold text (`__text__` → `<strong>text</strong>`)
- Italic text (`_text_` → `<em>text</em>`)
- Mixed inline formatting within paragraphs and list items

## Acceptance Criteria

1. All 17 test cases in `cases_test.go` pass (`go test ./...`)
2. `go vet ./...` reports no issues
3. The `Render` function is exported from the `markdown` package
4. Code is clean, readable, and well-structured (this is a refactoring exercise)
5. The solution handles edge cases: `#######` as paragraph, `#` and `*` in text that should not be interpreted as markdown syntax

## Key Constraints

- The solution file is `markdown.go` in the `markdown` package
- Must use Go 1.18+ (per `go.mod`)
- Must pass `go test` and `go vet` cleanly
- The reference solution in `.meta/example.go` provides a working implementation to guide the approach
