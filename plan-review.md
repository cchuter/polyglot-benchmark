# Plan Review: polyglot-go-markdown

## Summary

The plan proposes a line-by-line processor using `strings.Builder` and simple string operations (no regex) to convert markdown to HTML. I reviewed the plan against all 16 test cases in `cases_test.go`. Below is a detailed analysis of each review area, followed by a per-test-case trace.

---

## Area 1: `headerLevel` function and the h7 case

**Test case:** `"####### This will not be an h7"` should produce `"<p>####### This will not be an h7</p>"`

**Verdict: PASS (with caveat)**

The plan states: "Count leading `#` characters. Return count if followed by a space, 0 otherwise." and "level > 0 && level <= 6" guards the header branch. For 7 leading `#` characters, `headerLevel` would return 7, the guard `level <= 6` would fail, and the code falls through to the paragraph branch, producing `<p>####### This will not be an h7</p>`.

However, there is a subtle issue. The line `####### This will not be an h7` starts with `#`, so it enters the `if strings.HasPrefix(line, "#")` branch. Inside that branch, `headerLevel` returns 7. The condition `level > 0 && level <= 6` is false, so it falls to the else clause: `result.WriteString("<p>" + applyInline(line) + "</p>")`. This correctly wraps the **entire original line** (including the `#######` prefix) in a `<p>` tag.

This is correct. The plan handles h7 properly.

**Caveat:** The `headerLevel` function description says "Return count if followed by a space, 0 otherwise." This means for a line like `###noSpace`, `headerLevel` would return 0 (no space after `#`s), and since `level > 0` is false, it falls to the else (paragraph). That is correct behavior, but the plan does not explicitly test or discuss this edge case. It is not tested by the test suite either, so this is a non-issue for correctness.

---

## Area 2: Inline formatting -- bold (`__`) before italic (`_`)

**Test case:** `"This will _be_ __mixed__"` should produce `"<p>This will <em>be</em> <strong>mixed</strong></p>"`

**Verdict: PASS (with important implementation detail to get right)**

The plan says: "Replace `__` with `<strong>` / `</strong>` first (alternating open/close), then replace `_` with `<em>` / `</em>` (alternating open/close)."

For input `This will _be_ __mixed__`:

1. **Bold pass:** Find `__` markers. The string has two `__` pairs: at positions around "mixed". First `__` becomes `<strong>`, second `__` becomes `</strong>`. Result: `This will _be_ <strong>mixed</strong>`
2. **Italic pass:** Find `_` markers. The string has two `_` markers: around "be". First `_` becomes `<em>`, second `_` becomes `</em>`. Result: `This will <em>be</em> <strong>mixed</strong>`

This is correct. The critical insight the plan gets right is that processing `__` first prevents `__` from being consumed as two separate `_` markers.

**However, the plan says to use `strings.Replace` but also says to use "alternating open/close replacement."** The standard `strings.Replace` function replaces all occurrences with the same string -- it cannot alternate between `<strong>` and `</strong>`. The plan needs a custom replacement function that iterates through the string and alternates between the opening and closing tags. The plan alludes to this with "alternating open/close" but does not show the implementation of `applyInline`. This is a **critical implementation detail** that must be correct.

**Recommendation:** The `applyInline` function should be implemented as something like:

```go
func applyInline(text string) string {
    text = replaceAlternating(text, "__", "<strong>", "</strong>")
    text = replaceAlternating(text, "_", "<em>", "</em>")
    return text
}

func replaceAlternating(text, marker, open, close string) string {
    // Split on marker, rejoin with alternating open/close tags
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

The split-and-rejoin approach naturally alternates. **If the implementer uses `strings.Replace` or `strings.ReplaceAll` directly, it will NOT work** because those functions replace all occurrences with the same string.

**Verdict: Conceptually PASS, but the lack of explicit `applyInline` implementation is a risk. The implementer must use an alternating approach, not simple `strings.Replace`.**

---

## Area 3: List closing when a non-list line follows

**Test case:** `"# Start a list\n* Item 1\n* Item 2\nEnd a list"` should produce `"<h1>Start a list</h1><ul><li>Item 1</li><li>Item 2</li></ul><p>End a list</p>"`

**Verdict: PASS**

The plan handles this correctly. Tracing through the logic:

1. Line `# Start a list`: Not a list item. `inList` is false, no `</ul>` to emit. Starts with `#`, `headerLevel` returns 1, within 1-6 range. Output: `<h1>Start a list</h1>`.
2. Line `* Item 1`: List item. `inList` is false, so emit `<ul>`, set `inList = true`. Output: `<ul><li>Item 1</li>`.
3. Line `* Item 2`: List item. `inList` is true, no `<ul>` to emit. Output: `<li>Item 2</li>`.
4. Line `End a list`: Not a list item. `inList` is true, so emit `</ul>`, set `inList = false`. Does not start with `#`. Output: `</ul><p>End a list</p>`.

Final result: `<h1>Start a list</h1><ul><li>Item 1</li><li>Item 2</li></ul><p>End a list</p>` -- matches expected.

The end-of-input `</ul>` safety check at the bottom is also correct: if the input ended with list items, the `</ul>` would still be emitted.

---

## Area 4: Markdown symbols that should NOT be interpreted

**Test cases:**
- `"# This is a header with # and * in the text"` -> `"<h1>This is a header with # and * in the text</h1>"`
- `"* Item 1 with a # in the text\n* Item 2 with * in the text"` -> `"<ul><li>Item 1 with a # in the text</li><li>Item 2 with * in the text</li></ul>"`
- `"This is a paragraph with # and * in the text"` -> `"<p>This is a paragraph with # and * in the text</p>"`

**Verdict: PASS**

The plan's approach correctly handles this because:

1. **Header detection** only looks at the **beginning** of the line (`strings.HasPrefix(line, "#")`), then `headerLevel` counts leading `#` characters. Once the content portion is extracted (everything after `# `), the `#` and `*` characters within the content are just plain text -- they are never re-parsed.

2. **List detection** only checks for `* ` at the start of the line (`strings.HasPrefix(line, "* ")`). A `*` embedded in the middle of text (like `Item 2 with * in the text`) is never at the start, so it is not matched. After extracting the content (everything after `* `), the `#` and `*` are just plain text.

3. **Paragraph text** with `#` and `*` in the middle is not at the start of the line, so `strings.HasPrefix(line, "#")` is false and `strings.HasPrefix(line, "* ")` is false. The whole line is wrapped in `<p>...</p>`.

4. **Inline formatting** (`applyInline`) only looks for `__` and `_` markers, not `#` or `*`. So `#` and `*` in the content pass through unchanged.

This is all correct.

---

## Area 5: Import list -- `fmt` vs `strings`

**Verdict: FAIL (bug in the plan)**

The plan contains a contradiction:

- Line 129 of the plan shows: `result.WriteString(fmt.Sprintf("<h%d>%s</h%d>", level, content, level))`
- The import block on line 103 shows: `import "strings"` (no `fmt`)
- The note at line 165 says: "avoid `fmt.Sprintf` for header tags -- use string concatenation with `strconv.Itoa` or direct string building to avoid importing `fmt`. Actually, since header levels are 1-6, we can just use string concatenation with the digit character."

The plan explicitly uses `fmt.Sprintf` in the code but then says not to use it. The code as written will not compile because `fmt` is not imported.

**Fix:** Replace the `fmt.Sprintf` call with string concatenation. Since header levels are 1-6, convert the level to a string using `strconv.Itoa(level)` (requires importing `strconv`) or by using a simple expression like `string(rune('0' + level))`. Example:

```go
lvl := strconv.Itoa(level)
result.WriteString("<h" + lvl + ">" + content + "</h" + lvl + ">")
```

Or to avoid importing `strconv`:
```go
lvl := string(rune('0' + level))
result.WriteString("<h" + lvl + ">" + content + "</h" + lvl + ">")
```

**This is a compilation error and must be fixed before implementation.**

---

## Area 6: `applyInline` alternating open/close replacement

**Verdict: CONDITIONAL PASS -- correct concept, risky description**

As discussed in Area 2, the plan describes the right approach (process `__` before `_`, use alternating open/close) but:

1. **Does not show the `applyInline` implementation.** The plan says "Use `strings.Replace` for inline formatting" in one place, and "alternating open/close" in another. Standard `strings.Replace` cannot alternate. This is a contradiction.

2. **The alternating approach works for all test cases** because:
   - `_This will be italic_` -- one pair of `_`, alternates to `<em>...</em>`.
   - `__This will be bold__` -- one pair of `__`, alternates to `<strong>...</strong>`.
   - `This will _be_ __mixed__` -- bold pass handles `__` pair, italic pass handles `_` pair, no interference.
   - `__Bold Item__` and `_Italic Item_` -- single pairs, straightforward.
   - All other test cases have no inline formatting markers.

3. **Potential edge case (not in test suite):** If there is an odd number of markers (e.g., `_hello`), the split-based approach would leave a dangling opening tag with no closing tag. This is not tested, so it does not affect correctness for this exercise, but it is worth noting.

**Recommendation:** The implementer must use a split-and-rejoin approach or a manual scan with alternating state, NOT `strings.Replace` or `strings.ReplaceAll`. The plan description should be clarified to avoid confusion.

---

## Per-Test-Case Trace

| # | Description | Expected Result | Plan Handles Correctly? |
|---|------------|----------------|------------------------|
| 1 | parses normal text as a paragraph | `<p>This will be a paragraph</p>` | PASS -- no `#` or `*` prefix, falls to paragraph branch |
| 2 | parsing italics | `<p><em>This will be italic</em></p>` | PASS -- paragraph + `applyInline` handles `_..._` |
| 3 | parsing bold text | `<p><strong>This will be bold</strong></p>` | PASS -- paragraph + `applyInline` handles `__...__` |
| 4 | mixed normal, italics and bold text | `<p>This will <em>be</em> <strong>mixed</strong></p>` | PASS -- bold before italic ordering is correct |
| 5 | with h1 header level | `<h1>This will be an h1</h1>` | PASS (after fixing `fmt.Sprintf` bug) |
| 6 | with h2 header level | `<h2>This will be an h2</h2>` | PASS (after fixing `fmt.Sprintf` bug) |
| 7 | with h3 header level | `<h3>This will be an h3</h3>` | PASS (after fixing `fmt.Sprintf` bug) |
| 8 | with h4 header level | `<h4>This will be an h4</h4>` | PASS (after fixing `fmt.Sprintf` bug) |
| 9 | with h5 header level | `<h5>This will be an h5</h5>` | PASS (after fixing `fmt.Sprintf` bug) |
| 10 | with h6 header level | `<h6>This will be an h6</h6>` | PASS (after fixing `fmt.Sprintf` bug) |
| 11 | h7 header level is a paragraph | `<p>####### This will not be an h7</p>` | PASS -- level 7 > 6, falls to paragraph |
| 12 | unordered lists | `<ul><li>Item 1</li><li>Item 2</li></ul>` | PASS -- list tracking works correctly |
| 13 | With a little bit of everything | `<h1>Header!</h1><ul><li>...</li><li>...</li></ul>` | PASS -- header then list transition is correct |
| 14 | markdown symbols in header text | `<h1>This is a header with # and * in the text</h1>` | PASS -- only leading `#` is parsed |
| 15 | markdown symbols in list item text | `<ul><li>Item 1 with a # in the text</li>...</ul>` | PASS -- only leading `* ` is parsed |
| 16 | markdown symbols in paragraph text | `<p>This is a paragraph with # and * in the text</p>` | PASS -- no leading markers, plain paragraph |
| 17 | unordered lists close properly | `<h1>...</h1><ul>...</ul><p>End a list</p>` | PASS -- `inList` tracking closes list correctly |

---

## Overall Assessment

**The plan is fundamentally sound but has two issues that must be addressed before implementation:**

### Bug (must fix):
1. **`fmt.Sprintf` used without importing `fmt` (Area 5).** The code in the plan uses `fmt.Sprintf` on line 129 but only imports `strings`. The note at line 165 acknowledges this but does not update the code. The implementer must use string concatenation instead. This is a **compilation error**.

### Ambiguity (should clarify):
2. **`applyInline` implementation is not shown and the description is contradictory (Areas 2 and 6).** The plan says "use `strings.Replace`" in one place and "alternating open/close" in another. Standard `strings.Replace` cannot alternate between opening and closing tags. The implementer must use a split-based or state-based approach. The plan should either show the full `applyInline` implementation or clearly state that a custom alternating replacement function is needed (not `strings.Replace`).

### All test cases pass (conceptually):
With the `fmt.Sprintf` bug fixed and `applyInline` implemented using an alternating strategy, all 17 test cases would be handled correctly by the plan.

### Recommendation:
Fix the `fmt.Sprintf` usage and clarify the `applyInline` implementation before proceeding to coding. The overall architecture (line-by-line processing, `inList` state tracking, bold-before-italic ordering) is correct and well-suited to the exercise requirements.
