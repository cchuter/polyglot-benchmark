# Plan Review: polyglot-go-markdown

## Summary

The plan is well-structured and closely mirrors the reference solution. It correctly identifies the three-function architecture (`Render`, `getHeadingWeight`, `renderHTML`) and the key design decisions. However, there are several specific issues and gaps worth noting before implementation.

---

## 1. Test Case Coverage

The plan's described logic covers all 17 test cases:

| # | Test Description | Covered? | Notes |
|---|-----------------|----------|-------|
| 1 | Normal text as paragraph | Yes | Default path in `Render` |
| 2 | Italics | Yes | `renderHTML` handles `_` pairs |
| 3 | Bold text | Yes | `renderHTML` handles `__` pairs |
| 4 | Mixed normal, italics, bold | Yes | Both substitutions in `renderHTML` |
| 5-10 | Headings h1-h6 | Yes | `getHeadingWeight` returns 1-6 |
| 11 | h7 treated as paragraph | Yes | `getHeadingWeight` returns -1 |
| 12 | Unordered lists | Yes | List item accumulation logic |
| 13 | Header + bold/italic list items | Yes | Combined heading and list path |
| 14 | Markdown symbols in header text | Yes | Only first char checked for `#` |
| 15 | Markdown symbols in list item text | Yes | Only first char checked for `*` |
| 16 | Markdown symbols in paragraph text | Yes | Default paragraph path, no reinterpretation |
| 17 | List close with preceding/following lines | Yes | List flush on non-list line + end-of-input |

**Verdict**: All 17 test cases are accounted for by the described architecture.

---

## 2. Edge Cases Analysis

### Correctly handled edge cases

- **h7 fallback to paragraph**: The plan explicitly calls out returning -1 for 7+ `#` characters, matching test case 11.
- **List flushing at end of input**: The plan mentions closing any open list at the end, matching test case 17.
- **Bold before italic processing order**: The plan correctly identifies that `__` must be replaced before `_` to avoid misinterpretation.

### Potential edge cases not explicitly discussed (but handled by the reference approach)

- **Heading line with no renderHTML call**: In the reference solution, heading content (`line[headerWeight+1:]`) is inserted directly into the `<h>` tag without passing through `renderHTML`. This means bold/italic markers inside headings would NOT be converted. None of the 17 test cases test for bold/italic inside headings, so this is fine for passing the tests, but the plan does not explicitly call this out. This is a non-issue for correctness against the test suite but worth noting.

- **h7 paragraph also skips renderHTML**: When `getHeadingWeight` returns -1, the reference solution wraps the entire raw line (including the `#######` prefix) in `<p>` tags and does NOT call `renderHTML`. The plan's description of this path says `fmt.Sprintf("<p>%s</p>", line)` which matches the reference, but the plan does not explicitly note that inline formatting is intentionally skipped here. Again, no test case exercises this combination, so it is a non-issue.

- **List item marker parsing `line[2:]`**: The reference solution strips the first two characters (`* `) from list items. This assumes the format is always `* ` (asterisk + space). The plan mentions `listItemMarker = '*'` but does not explicitly document the `[2:]` slice. If implemented correctly per the reference, this works. If someone misreads and uses `[1:]`, the leading space would leak into the output.

---

## 3. Architecture Assessment

The three-function architecture is sound:

- **`Render`** handles line-level parsing and list state management. The state machine (tracking whether we are inside a list) is the right approach for this problem.
- **`getHeadingWeight`** cleanly separates heading detection. The loop-up-to-6 approach with a -1 sentinel for overflow is correct.
- **`renderHTML`** handles inline formatting with a simple iterative replacement strategy.

**The architecture directly mirrors the reference solution**, which is both a strength (known-correct) and a limitation (no room for alternative designs that might be more robust). For the purposes of this exercise, mirroring the reference is the right call.

### strings.Builder usage

The plan correctly identifies `strings.Builder` for efficient concatenation. This is appropriate.

### Separation of concerns

Line-level processing (block elements) is cleanly separated from inline processing (`renderHTML`). This is good design.

---

## 4. Potential Issues

### Issue 1: `getHeadingWeight` returns 0 for non-heading lines

The plan says: "Loop up to index 6; if all are `#`, return -1." However, examining the reference implementation more carefully:

```go
func getHeadingWeight(line string) int {
    for i := 0; i <= 6; i++ {
        if line[i] != headingMarker {
            return i
        }
    }
    return -1
}
```

This function is only called when `line[0] == headingMarker` is already true (checked in `Render`), so `i=0` will never trigger the early return. The plan does not explicitly note this precondition. If `getHeadingWeight` were called on a line not starting with `#`, it would return 0, which would produce malformed HTML (`<h0>...</h0>`). This is not a bug in the reference solution since the precondition is always met, but the plan should ideally document this assumption.

### Issue 2: The plan does not show the full implementation

The plan shows function signatures but not the full function bodies. While it says to follow the reference solution, the actual implementation details are left implicit. This is acceptable for a plan but means the implementer needs to carefully study the reference solution to get the details right. Specifically:

- The exact slicing logic for list items (`line[2:]`)
- The exact slicing logic for headings (`line[headerWeight+1:]`)
- The fact that headings do NOT call `renderHTML`
- The `else if` structure that flushes the list only when transitioning from list to non-list

### Issue 3: No mention of the `renderHTML` not being called for headings

As noted in the edge cases section, the reference solution does not call `renderHTML` on heading content. The plan's description of `renderHTML` as an "inline formatting" helper might lead an implementer to also apply it to headings. The plan should clarify that `renderHTML` is only applied to paragraph text and list item text.

### Issue 4: Imports

The plan correctly lists `fmt` and `strings` as required imports. No issues here.

---

## 5. Overall Assessment

**The plan is solid and will produce a correct implementation if followed carefully alongside the reference solution.** The architecture is sound, all test cases are covered, and the key design decisions (bold-before-italic ordering, list state tracking) are correctly identified.

**Minor gaps to address before implementation:**

1. Explicitly note that `renderHTML` is NOT applied to heading content or h7-fallback paragraphs.
2. Explicitly document the `line[2:]` slicing for list items (asterisk + space).
3. Note the precondition that `getHeadingWeight` is only called when the line starts with `#`.

These are documentation gaps, not correctness issues -- the plan will produce correct code if the reference solution is followed faithfully.

**Recommendation: Approve with minor notes. Proceed to implementation.**
