# Implementation Plan: polyglot-go-markdown

## Proposal A (Proponent)

### Approach: Line-by-line processor with inline formatting

A single-pass, line-by-line parser. Split input on `\n`, classify each line (header, list item, or paragraph), apply inline formatting (bold/italic), track list state, and build the HTML string.

### Files to modify
- `go/exercises/practice/markdown/markdown.go` — implement `Render` function

### Architecture

```go
func Render(input string) string {
    // 1. Split input into lines
    // 2. For each line, determine type: header, list item, or paragraph
    // 3. Track whether we're inside a <ul> block
    // 4. Apply inline formatting (bold before italic to handle __ vs _)
    // 5. Build result string, closing <ul> when transitioning out of list
}
```

**Key design decisions:**
- Use `strings.Builder` for efficient string concatenation
- Process bold (`__`) before italic (`_`) to avoid `__` being matched as two `_`
- Use `strings.Replace` for inline formatting (simple, no regex needed)
- Count leading `#` characters to determine header level (1-6 valid, 7+ is paragraph)
- Track `inList` boolean to manage `<ul>` open/close tags

### Rationale
- Simple, procedural code that's easy to read and debug
- No regex dependency — just string operations
- Matches the "refactored" style the exercise expects (clean, readable code)
- Single responsibility: line classification, inline formatting, and output building are clearly separated concerns

---

## Proposal B (Opponent)

### Approach: Regex-based parser with token types

Use Go's `regexp` package to classify and transform lines. Define regex patterns for each line type and for inline formatting.

### Files to modify
- `go/exercises/practice/markdown/markdown.go` — implement `Render` function

### Architecture

```go
var (
    headerRe   = regexp.MustCompile(`^(#{1,6})\s+(.*)$`)
    listItemRe = regexp.MustCompile(`^\*\s+(.*)$`)
    boldRe     = regexp.MustCompile(`__(.+?)__`)
    italicRe   = regexp.MustCompile(`_(.+?)_`)
)

func Render(input string) string {
    // Use regex matching for line classification
    // Use regex replacement for inline formatting
}
```

### Critique of Proposal A
- Manual string counting (`#` counting) is fragile — off-by-one errors are easy
- `strings.Replace` for bold/italic could have edge cases with overlapping patterns
- No formal grammar or pattern definition

### Rationale for Proposal B
- Regex provides declarative, well-defined pattern matching
- Less manual string manipulation
- Patterns are self-documenting

---

## Selected Plan (Judge)

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | Both can satisfy all tests | Both can satisfy all tests |
| Risk | Low — simple string ops are predictable | Medium — regex for nested `_`/`__` can be tricky |
| Simplicity | Very simple, no imports beyond `strings` | Adds `regexp` import, compiled patterns |
| Consistency | Matches style of other exercises (ledger uses strings.Builder, no regex) | No other exercise in this repo uses regex |

### Decision: **Proposal A wins**

**Rationale:**
1. The test cases don't have edge cases complex enough to warrant regex (no nested/overlapping formatting)
2. The bold-before-italic processing order with `strings.Replace` handles the `__` vs `_` distinction correctly and simply
3. The codebase convention favors simple string operations over regex
4. Fewer imports, less complexity, easier to understand

### Selected Plan: Line-by-line processor

**File:** `go/exercises/practice/markdown/markdown.go`

**Implementation:**

```go
package markdown

import "strings"

// Render translates markdown input to HTML.
func Render(input string) string {
    var result strings.Builder
    lines := strings.Split(input, "\n")
    inList := false

    for _, line := range lines {
        if strings.HasPrefix(line, "* ") {
            if !inList {
                result.WriteString("<ul>")
                inList = true
            }
            content := applyInline(line[2:])
            result.WriteString("<li>" + content + "</li>")
        } else {
            if inList {
                result.WriteString("</ul>")
                inList = false
            }
            if strings.HasPrefix(line, "#") {
                level := headerLevel(line)
                if level > 0 && level <= 6 {
                    content := applyInline(line[level+1:])
                    lvl := string(rune('0' + level))
                    result.WriteString("<h" + lvl + ">" + content + "</h" + lvl + ">")
                } else {
                    result.WriteString("<p>" + applyInline(line) + "</p>")
                }
            } else {
                result.WriteString("<p>" + applyInline(line) + "</p>")
            }
        }
    }

    if inList {
        result.WriteString("</ul>")
    }

    return result.String()
}
```

**Helper functions:**

1. `headerLevel(line string) int` — Count leading `#` characters. Return count if followed by a space, 0 otherwise.

```go
func headerLevel(line string) int {
    level := 0
    for _, c := range line {
        if c == '#' {
            level++
        } else {
            break
        }
    }
    if level > 0 && level < len(line) && line[level] == ' ' {
        return level
    }
    return 0
}
```

2. `applyInline(text string) string` — Use a split-and-rejoin approach to replace markers with alternating open/close tags. Process `__` (bold) before `_` (italic) to prevent `__` from being consumed as two `_`.

```go
func applyInline(text string) string {
    text = replaceMarker(text, "__", "<strong>", "</strong>")
    text = replaceMarker(text, "_", "<em>", "</em>")
    return text
}

func replaceMarker(text, marker, open, close string) string {
    parts := strings.Split(text, marker)
    var result strings.Builder
    for i, part := range parts {
        if i > 0 {
            if i%2 == 1 {
                result.WriteString(open)
            } else {
                result.WriteString(close)
            }
        }
        result.WriteString(part)
    }
    return result.String()
}
```

**Processing order:**
1. Split input into lines
2. For each line, check if it's a list item (`* ` prefix)
3. If not a list item, close any open list, then check for header (`#` prefix)
4. Apply inline formatting to content
5. After all lines, close any remaining open list
6. Return built string

**Key details:**
- Only import `"strings"` — no `fmt`, no `strconv`, no `regexp`
- Header level integer (1-6) converted to string via `string(rune('0' + level))`
- Inline formatting uses split-and-rejoin with alternating open/close tags
- Bold markers (`__`) processed before italic (`_`) to prevent interference
