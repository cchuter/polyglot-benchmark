# Context: polyglot-go-markdown (Issue #305)

## Status: DONE

## Branch: issue-305

## Commit: "Implement Render function for markdown parser"

## Implementation

File: `go/exercises/practice/markdown/markdown.go`

The `Render` function parses a subset of Markdown and returns HTML:
- **Lines** split on `\n`, classified as header (`# `), list item (`* `), or paragraph
- **Headers** h1-h6 via `headerLevel()` counting leading `#`; 7+ `#` treated as paragraph
- **Lists** tracked with `inList` boolean for proper `<ul>` open/close
- **Inline formatting** via `applyInline()` → `replaceMarker()` (split-and-rejoin with alternating open/close tags)
- Bold (`__`) processed before italic (`_`) to prevent `__` matching as two `_`

## Test Results

17/17 tests pass. Build and vet clean. Only `markdown.go` modified.

## Key Files

| File | Purpose |
|------|---------|
| `go/exercises/practice/markdown/markdown.go` | Implementation |
| `go/exercises/practice/markdown/cases_test.go` | 17 test cases |
| `go/exercises/practice/markdown/markdown_test.go` | Test runner |
| `.osmi/agents/verifier/report.md` | Final verification: PASS |
