# Implementation Plan: polyglot-go-markdown

## Branch 1: Clean Imperative (Minimal Refactor of Reference)

Take the reference solution in `.meta/example.go` and write a clean version that follows the same imperative, line-by-line approach but with improved readability.

### Files to Modify
- `go/exercises/practice/markdown/markdown.go` — write the full solution

### Approach
1. Keep the same overall structure: split on newlines, process each line
2. Use named constants for marker characters
3. Keep `strings.Builder` for output assembly
4. Helper functions: `parseHeading`, `parseInlineMarkup` (replaces `renderHTML`)
5. Track list state with a boolean `inList` flag instead of accumulating items in a slice
6. Process bold before italic in `parseInlineMarkup`

### Evaluation
- **Feasibility**: High — mirrors the working reference with clearer naming
- **Risk**: Low — minimal deviation from proven approach
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~60 lines, 1 file, straightforward

---

## Branch 2: Structured with Line Types (Extensible)

Parse each line into a typed structure first, then render to HTML in a second pass. This separates parsing from rendering.

### Files to Modify
- `go/exercises/practice/markdown/markdown.go` — write the full solution

### Approach
1. Define a `lineType` enum (heading, listItem, paragraph)
2. Define a `parsedLine` struct with type, content, and heading level
3. First pass: parse all lines into `[]parsedLine`
4. Second pass: iterate `[]parsedLine`, managing list grouping, and render to HTML
5. Separate `parseInline` function for bold/italic substitution

### Evaluation
- **Feasibility**: High — but more code for no additional test coverage
- **Risk**: Medium — over-engineering for the current requirements
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~90-100 lines, 1 file, more types and abstractions

---

## Branch 3: Regex-Based Inline Parsing (Performance/Unconventional)

Use `regexp` for inline markup (bold, italic) instead of iterative `strings.Replace`. Keep the line-by-line structure.

### Files to Modify
- `go/exercises/practice/markdown/markdown.go` — write the full solution

### Approach
1. Compile regexes at package init: `__(.+?)__` for bold, `_(.+?)_` for italic
2. Line processing same as Branch 1 (split, detect type, build HTML)
3. `renderInline` uses `regexp.ReplaceAllString` for bold then italic
4. Heading detection uses a simple loop or `strings.TrimLeft`

### Evaluation
- **Feasibility**: High — regexes are well-suited for pattern matching
- **Risk**: Medium — regex compilation cost, harder to debug for simple cases
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~55 lines, 1 file, but adds `regexp` dependency

---

## Selected Plan

**Branch 1: Clean Imperative** is selected.

### Rationale
- **Simplest approach** that directly addresses the requirements without over-engineering
- **No additional dependencies** — only `fmt` and `strings` from stdlib
- **Proven pattern** — follows the reference solution's logic which passes all tests
- **Most readable** — clear, linear control flow that's easy to follow
- **Lowest risk** — minimal deviation from working code means fewer surprises
- Branch 2 adds unnecessary abstraction layers for 17 test cases. Branch 3 introduces regex compilation overhead and a new import for simple string replacements.

### Detailed Implementation Plan

**File**: `go/exercises/practice/markdown/markdown.go`

```go
package markdown

import (
    "fmt"
    "strings"
)
```

**Constants:**
- `headingMarker = '#'`
- `listItemPrefix = "* "`

**Function `Render(markdown string) string`:**
1. Initialize `strings.Builder` for output and `inList bool` for list state tracking
2. Split input on `"\n"`, iterate each line:
   - **List item** (`strings.HasPrefix(line, "* ")`): If not already `inList`, write `<ul>`. Write `<li>` + renderInline(content) + `</li>`. Set `inList = true`.
   - **Not a list item**: If `inList`, write `</ul>` and set `inList = false`.
   - **Heading** (line starts with `#`): Count consecutive `#` chars. If 1-6, write `<hN>content</hN>`. If 7+, treat as paragraph.
   - **Paragraph** (default): Write `<p>` + renderInline(content) + `</p>`.
3. After loop: if still `inList`, write `</ul>`.
4. Return `html.String()`.

**Function `headingLevel(line string) int`:**
- Count leading `#` characters, return the count if 1-6, else return 0 (meaning "not a heading").

**Function `renderInline(text string) string`:**
- Replace `__` pairs with `<strong>`/`</strong>` (bold first)
- Replace `_` pairs with `<em>`/`</em>` (italic second)
- Return result

### Order of Implementation
1. Write `renderInline` helper
2. Write `headingLevel` helper
3. Write main `Render` function
4. Run `go test -v` to verify all 17 tests pass
5. Run `go vet` to check for issues
