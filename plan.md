# Implementation Plan: polyglot-go-markdown (Issue #221)

## Branch 1: Minimal Procedural Approach (Close to Reference)

Write a clean, procedural implementation closely following the reference solution's structure but with improved naming and readability.

### Files to Modify
- `go/exercises/practice/markdown/markdown.go`

### Approach
1. Define constants for heading and list item markers
2. Implement `Render(markdown string) string` that:
   - Splits input on newlines
   - Tracks list state (whether we're inside a `<ul>`)
   - For each line: detect heading, list item, or paragraph
   - Apply inline formatting (bold, italic) to text content
3. Implement helper `getHeadingLevel(line string) int` to count leading `#` chars (return -1 if >6)
4. Implement helper `applyInlineFormatting(text string) string` to convert `__...__` and `_..._`

### Rationale
Minimal deviation from the reference. Simple, direct, easy to verify correctness.

### Evaluation
- **Feasibility**: Excellent — proven to work via reference solution
- **Risk**: Very low — straightforward translation
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: 1 file, ~60 lines, 3 functions

---

## Branch 2: Structured Line-Type Approach

Introduce explicit line type classification for extensibility.

### Files to Modify
- `go/exercises/practice/markdown/markdown.go`

### Approach
1. Define a `lineType` enum (heading, listItem, paragraph)
2. Implement `classifyLine(line string) lineType` function
3. Implement `Render` using a state machine with explicit open/close list tracking
4. Each line type has its own render function: `renderHeading`, `renderListItem`, `renderParagraph`
5. Separate `applyInlineFormatting` for bold/italic

### Rationale
Better separation of concerns. Each Markdown element has a dedicated handler, making it easy to add new element types in the future.

### Evaluation
- **Feasibility**: Good — standard approach
- **Risk**: Low, but more code than needed for the test cases
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: 1 file, ~90 lines, 6+ functions — more than necessary

---

## Branch 3: Regex-Based Approach

Use regular expressions for inline formatting and line classification.

### Files to Modify
- `go/exercises/practice/markdown/markdown.go`

### Approach
1. Use `regexp.MustCompile` for bold (`__(.+?)__`) and italic (`_(.+?)_`) patterns
2. Use regex to match heading lines (`^(#{1,6})\s(.*)`)
3. Use regex to match list items (`^\*\s(.*)`)
4. Process lines sequentially, using regex replacements for inline formatting

### Rationale
Regex provides concise pattern matching. Potentially more robust for edge cases.

### Evaluation
- **Feasibility**: Good — Go's regexp package works fine
- **Risk**: Medium — regex ordering matters (bold before italic), potential for subtle bugs. Adds `regexp` dependency. Regex can be harder to debug.
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: 1 file, ~50 lines, but regex patterns add cognitive overhead

---

## Selected Plan

**Branch 1: Minimal Procedural Approach** is selected.

### Rationale for Selection

Branch 1 is superior because:
1. **Proven correctness**: The reference solution uses this exact approach and passes all tests
2. **Simplicity**: Fewest lines of code, fewest functions, easiest to understand
3. **No over-engineering**: Branch 2 adds unnecessary abstraction for 17 test cases. Branch 3 adds regex complexity without benefit.
4. **Maintainability**: Clear, linear control flow that's easy to follow
5. **No extra dependencies**: Only uses `fmt` and `strings` from stdlib

### Detailed Implementation Plan

#### File: `go/exercises/practice/markdown/markdown.go`

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

// Render translates markdown to HTML.
func Render(markdown string) string {
    var itemList []string
    var html = strings.Builder{}
    for _, line := range strings.Split(markdown, "\n") {
        if line[0] == listItemMarker {
            itemList = append(itemList, fmt.Sprintf("<li>%s</li>", renderInline(line[2:])))
            continue
        } else if len(itemList) != 0 {
            html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
            itemList = []string{}
        }
        if line[0] == headingMarker {
            level := headingLevel(line)
            if level != -1 {
                html.WriteString(fmt.Sprintf("<h%d>%s</h%d>", level, line[level+1:], level))
            } else {
                html.WriteString(fmt.Sprintf("<p>%s</p>", line))
            }
            continue
        }
        html.WriteString(fmt.Sprintf("<p>%s</p>", renderInline(line)))
    }
    if len(itemList) != 0 {
        html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
    }
    return html.String()
}

// headingLevel counts leading '#' characters and returns the heading level (1-6).
// Returns -1 if there are more than 6 '#' characters (not a valid heading).
func headingLevel(line string) int {
    for i := 0; i <= 6; i++ {
        if line[i] != headingMarker {
            return i
        }
    }
    return -1
}

// renderInline converts markdown inline formatting to HTML.
// Handles bold (__text__) and italic (_text_) markers.
func renderInline(text string) string {
    result := text
    for strings.Contains(result, "__") {
        result = strings.Replace(result, "__", "<strong>", 1)
        result = strings.Replace(result, "__", "</strong>", 1)
    }
    for strings.Contains(result, "_") {
        result = strings.Replace(result, "_", "<em>", 1)
        result = strings.Replace(result, "_", "</em>", 1)
    }
    return result
}
```

### Implementation Steps (ordered)
1. Write the complete `markdown.go` file with `Render`, `headingLevel`, and `renderInline` functions
2. Run `go test` to verify all 17 test cases pass
3. Commit with descriptive message
