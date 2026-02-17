# Scope: polyglot-go-markdown

## In Scope

- Implementing the `Render` function in `go/exercises/practice/markdown/markdown.go`
- Supporting: paragraphs, headers (h1-h6), bold, italic, unordered lists
- Handling mixed content across multiple lines
- Proper list open/close semantics (lists close when a non-list line follows)
- Inline formatting (bold/italic) within any block element

## Out of Scope

- Ordered lists, blockquotes, code blocks, links, images, or any Markdown features not tested
- Nested lists
- HTML escaping
- Modifying any test files or go.mod
- Modifying any files outside the `go/exercises/practice/markdown/` directory

## Dependencies

- Go standard library only (specifically `strings` package)
- No external packages
- Existing test infrastructure (`go test`)
