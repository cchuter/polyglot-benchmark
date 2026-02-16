# Scope: polyglot-go-markdown (Issue #221)

## In Scope

- Implementing the `Render` function in `go/exercises/practice/markdown/markdown.go`
- Supporting all Markdown features tested by the 17 test cases:
  - Paragraph wrapping
  - Headings h1–h6 (h7+ as paragraph)
  - Unordered lists with `* ` prefix
  - Bold (`__...__`) and italic (`_..._`) inline formatting
  - Proper list open/close across mixed line types
- Ensuring all tests pass with `go test`

## Out of Scope

- Modifying test files (`markdown_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Supporting Markdown features not covered by the test cases (e.g., ordered lists, links, images, code blocks, nested lists)
- Performance optimization beyond reasonable correctness
- Adding new test cases

## Dependencies

- Go standard library only (`fmt`, `strings`)
- No external packages required
- No dependencies on other exercises or packages in the repository
