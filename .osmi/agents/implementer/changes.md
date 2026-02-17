# Changes: Implement Render function for markdown parser

## File Modified
- `go/exercises/practice/markdown/markdown.go`

## What Was Done
Implemented the `Render` function that translates markdown input to HTML using a line-by-line processor with inline formatting support.

### Functions Implemented
1. **`Render(input string) string`** — Main entry point. Splits input into lines, classifies each as header, list item, or paragraph, tracks list state (`<ul>` open/close), applies inline formatting, and returns the built HTML string.
2. **`headerLevel(line string) int`** — Counts leading `#` characters and returns the level (1-6) if followed by a space, 0 otherwise.
3. **`applyInline(text string) string`** — Applies bold (`__` → `<strong>`) then italic (`_` → `<em>`) formatting.
4. **`replaceMarker(text, marker, open, close string) string`** — Generic split-and-rejoin helper that replaces markers with alternating open/close HTML tags.

### Design Decisions
- Uses `strings.Builder` for efficient string concatenation
- Processes bold (`__`) before italic (`_`) to prevent marker interference
- Only imports `"strings"` — no regex or other dependencies
- Header level 7+ treated as paragraph (per spec)

## Test Results
All 17 tests pass.
