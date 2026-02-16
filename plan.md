# Implementation Plan: polyglot-go-markdown

## File to Modify

- `go/exercises/practice/markdown/markdown.go` — replace the stub with a full implementation

## Architecture

The implementation follows the reference solution's clean architecture with three concerns:

### 1. Main `Render` function
- Split input by newlines
- Track list state (whether we're inside a `<ul>` block)
- For each line, determine its type (list item, heading, or paragraph)
- Build HTML output using `strings.Builder`
- Close any open list at the end

### 2. `getHeadingWeight` helper
- Count leading `#` characters (up to 6)
- Return the heading level (1-6) or -1 if 7+ `#` characters

### 3. `renderHTML` helper (inline formatting)
- Replace `__...__` pairs with `<strong>...</strong>`
- Replace `_..._` pairs with `<em>...</em>`
- Process bold before italic to avoid conflicts

## Implementation Details

```go
package markdown

import (
    "fmt"
    "strings"
)

const (
    headingMarker  = '#'
    listItemMarker = '*'
)

func Render(markdown string) string
func getHeadingWeight(line string) int
func renderHTML(markdownLine string) string
```

### Key Design Decisions

1. **Bold before italic** — Process `__` pairs first, then `_` pairs, so `_` inside `__...__` doesn't get misinterpreted
2. **List state tracking** — Accumulate list items; flush `<ul>` block when a non-list line is encountered or at end of input
3. **strings.Builder** — Efficient string concatenation
4. **Heading level validation** — Loop up to index 6; if all are `#`, return -1 (treated as paragraph)

## Ordering

1. Write the complete `markdown.go` file
2. Run `go test` to verify all 17 tests pass
3. Done — single file change
