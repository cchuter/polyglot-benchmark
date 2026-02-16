# Goal: polyglot-go-markdown (Issue #221)

## Problem Statement

The `markdown` exercise in `go/exercises/practice/markdown/` requires implementing a `Render` function that parses a subset of Markdown syntax and returns the corresponding HTML string. The current stub file (`markdown.go`) contains only a package declaration — the `Render` function must be implemented.

This is a **refactoring exercise** by nature: the reference solution in `.meta/example.go` provides a working implementation. The task is to write a clean, readable, well-structured implementation of the `Render` function that passes all 17 test cases.

## Acceptance Criteria

1. **All tests pass**: `go test ./go/exercises/practice/markdown/` exits with status 0 and all 17 test cases in `cases_test.go` pass.
2. **Function signature**: The package exports a `Render(markdown string) string` function.
3. **Markdown features supported**:
   - Paragraphs: plain text wrapped in `<p>...</p>`
   - Headings h1–h6: lines starting with 1–6 `#` characters
   - h7+ treated as paragraphs (no heading conversion)
   - Unordered lists: lines starting with `* ` wrapped in `<ul><li>...</li></ul>`
   - Bold: `__text__` → `<strong>text</strong>`
   - Italic: `_text_` → `<em>text</em>`
   - Mixed inline formatting within paragraphs and list items
   - Markdown symbols (`#`, `*`) in header/list/paragraph text are not re-interpreted
   - Unordered lists close properly with preceding and following lines
4. **Code quality**: Clean, readable Go code following standard conventions.
5. **No modification to test files**: `markdown_test.go` and `cases_test.go` must remain unchanged.

## Key Constraints

- Must be in package `markdown`
- Must use Go 1.18+ (as specified in go.mod)
- Only the `markdown.go` file should be modified
- The solution must be self-contained (no external dependencies beyond the standard library)
