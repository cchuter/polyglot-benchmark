# Goal: polyglot-go-markdown

## Problem Statement

Implement a `Render` function in Go that parses a subset of Markdown syntax and returns the corresponding HTML string. The exercise is framed as a "refactoring" exercise from Exercism, but the stub file is empty, so the full implementation must be written.

The supported Markdown features are:
- **Paragraphs**: Plain text is wrapped in `<p>` tags
- **Headers**: Lines starting with 1-6 `#` characters become `<h1>` through `<h6>`. Lines starting with 7+ `#` are treated as paragraphs (including the `#` symbols).
- **Bold**: Text wrapped in `__` becomes `<strong>`
- **Italic**: Text wrapped in `_` becomes `<em>`
- **Unordered lists**: Lines starting with `* ` become list items wrapped in `<ul><li>...</li></ul>`
- **Mixed content**: Headers, lists, and paragraphs can be intermixed. Lists must close properly when followed by non-list lines.
- **Markdown symbols in text**: `#` and `*` appearing in header text, list item text, or paragraph text (not at the start of a line in their special positions) should NOT be interpreted as markdown.

## Acceptance Criteria

1. All 17 test cases in `cases_test.go` pass
2. The `Render` function accepts a string and returns a string (`func Render(input string) string`)
3. The function lives in `package markdown` in the file `markdown.go`
4. The code compiles without errors
5. The benchmark test runs without errors

## Key Constraints

- Must be in Go, using `go 1.18` module
- Only standard library imports allowed
- The function signature must be `Render(string) string`
- No modification of test files (`markdown_test.go`, `cases_test.go`, `go.mod`)
