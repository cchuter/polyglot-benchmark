# Plan Review: polyglot-go-connect

## Overall Assessment

The plan is **well-structured and correct**. It closely mirrors the reference solution at `/go/exercises/practice/connect/.meta/example.go` and covers all essential aspects of the Connect/Hex board game exercise. The approach will pass all 10 test cases.

---

## 1. Correctness of the Approach

**Verdict: Correct.**

- The DFS flood-fill with bit-flag visited tracking is a sound algorithm for this problem. It is the same approach used by the reference solution.
- The hex adjacency directions listed -- `(x+1,y), (x-1,y), (x,y+1), (x,y-1), (x-1,y+1), (x+1,y-1)` -- are correct for a hex board with the offset/indentation layout used by the test cases.
- The player-to-direction mapping is correct: X (black) connects left-to-right (start at x=0, target x=width-1), O (white) connects top-to-bottom (start at y=0, target y=height-1).
- Checking black first, then white, then returning `""` matches the expected behavior and the reference solution's ordering.

**One clarification needed:** The plan mentions `colorFlags` in function signatures (e.g., `at(coord, colorFlags)`) but defines the constants section as "bit flags for stone colors and connected state" using `iota`. It does not explicitly state that a `colorFlags` *struct* (with `.color` and `.connected` fields) should be created, as the reference solution does:

```go
type colorFlags struct {
    color     int8
    connected int8
}
```

This is important because the functions reference `colorFlags` as a parameter type, and without the struct definition, the implementation would need to pass raw `int8` values and handle the distinction between color and connected flags manually. The plan should explicitly specify this struct and the two predefined instances (`flagsBlack`, `flagsWhite`).

---

## 2. Potential Issues and Edge Cases

### 2a. Input Parsing -- Space Stripping

The test file's `prepare()` function strips all spaces from each input line before passing them to `ResultOf()`:

```go
func prepare(lines []string) []string {
    newLines := make([]string, len(lines))
    for i, l := range lines {
        newLines[i] = strings.ReplaceAll(l, " ", "")
    }
    return newLines
}
```

The plan does not mention this. The `newBoard` parser should expect **pre-stripped input** (no spaces, no leading indentation). This is actually handled naturally if the parser just iterates over characters and maps `'X'` / `'O'` / anything else, but it is worth calling out explicitly in the plan since:
- The `width` is derived from `len(lines[0])`, which must be the count of actual cell characters (not including spaces).
- All rows after stripping should have the same length. The plan does not mention validating consistent row lengths. The reference solution does not validate this either, but it is a latent assumption.

### 2b. Error Handling

The plan mentions `newBoard([]string) (board, error)` and the main function signature returns `(string, error)`, which is correct. However, the plan does not specify what error conditions to check. The reference solution checks:
- Empty input (`len(lines) < 1`)
- Empty first line (`len(lines[0]) < 1`)

The plan should enumerate these checks since the function signature requires error handling.

### 2c. The `at()` Method Return Signature

The plan describes `at(coord, colorFlags) (bool, bool)` but does not explain what the two booleans represent. The reference solution returns `(isCorrectColor, isConnected)`. This is a critical detail for the DFS logic: the algorithm must only recurse into cells that have the correct color AND have not yet been marked as connected (visited). Misunderstanding these return values would produce incorrect results.

### 2d. Board as Value Type vs. Pointer

The reference solution passes `board` by value throughout (not as a pointer). This works because the `fields` slice header is copied, but the underlying array is shared. The plan does not discuss this, and it is a subtle Go detail. If an implementer were to use a separate `visited` map instead of in-place bit flags, they would need pointer receivers. The plan should note that slice-of-slices as a field allows mutation through value receivers.

### 2e. 1x1 Board

The 1x1 board is an important edge case where a single cell is simultaneously at the start edge and the target edge. The DFS handles this correctly because `evaluate` first checks if the cell has the right color, marks it connected, then checks `isTargetCoord` -- and on a 1x1 board, the start coord IS the target coord. The plan does not explicitly call this out, but the algorithm handles it implicitly. It would be good to note this.

---

## 3. Test Case Coverage

The plan's algorithm will correctly handle all 10 test cases:

| # | Test Case | Covered? | Notes |
|---|-----------|----------|-------|
| 1 | Empty board has no winner | Yes | No X or O cells, DFS finds nothing |
| 2 | X can win on a 1x1 board | Yes | Single X cell is both start and target |
| 3 | O can win on a 1x1 board | Yes | Single O cell is both start and target |
| 4 | Only edges does not make a winner | Yes | Edge cells without a connected path |
| 5 | Illegal diagonal does not make a winner | Yes | Tests that square-grid diagonals are not hex neighbors |
| 6 | Nobody wins crossing adjacent angles | Yes | Tests hex adjacency rules more thoroughly |
| 7 | X wins crossing from left to right | Yes | Standard X win |
| 8 | O wins crossing from top to bottom | Yes | Standard O win |
| 9 | X wins using a convoluted path | Yes | Non-obvious winding path |
| 10 | X wins using a spiral path | Yes | Large 9x9 board with spiral |

The algorithm is sound for all cases. The hex adjacency definition is correct and will properly reject square-grid diagonal moves (test cases 5 and 6 specifically test this).

---

## 4. Suggestions for Improvement

### 4.1 Explicitly Define the `colorFlags` Struct

The plan references `colorFlags` in function signatures but never defines it as a struct. Add:

```
### Data Structures (addition)
- **`colorFlags` struct**: holds `color int8` and `connected int8`
- **Pre-defined instances**: `flagsBlack = colorFlags{black, connectedBlack}` and `flagsWhite = colorFlags{white, connectedWhite}`
```

Without this, the function signatures in the plan are ambiguous.

### 4.2 Document Input Assumptions

Add a note that the input arrives pre-stripped (no spaces) because the test harness calls `prepare()` before `ResultOf()`. This affects how `newBoard` calculates board dimensions.

### 4.3 Note Error Conditions

List the specific error cases:
- Zero lines provided
- First line is empty string

### 4.4 Call Out the 1x1 Board Edge Case

Explicitly note that on a 1x1 board, the start coordinate is also the target coordinate, so the DFS must check `isTargetCoord` before recursing into neighbors (which is what the algorithm does).

### 4.5 Note the Return Values of `at()`

Specify that `at()` returns `(isCorrectColor bool, isConnected bool)` and that the DFS only recurses when `isCorrectColor && !isConnected`. This is the core visited-tracking mechanism.

### 4.6 Consider Stack Depth for Spiral Test

The spiral test case (test 10) is a 9x9 board where X follows a spiral path, visiting up to ~40+ cells. With recursive DFS, this creates ~40 stack frames, which is fine. However, for very large boards, an iterative BFS/DFS with an explicit stack might be more robust. For this exercise's test cases, recursive DFS is perfectly adequate, but a brief note acknowledging this trade-off would demonstrate awareness.

### 4.7 Consider Mentioning the `dump()` Helper

The reference solution includes a `dump()` debug helper. While not required to pass tests, it is useful during development. The plan could mention it as an optional debugging aid.

---

## Summary

The plan is fundamentally sound and will produce a correct, passing implementation. The main gaps are:

1. **Missing `colorFlags` struct definition** -- the most significant omission, as function signatures reference a type that is not defined in the plan.
2. **No mention of input pre-processing** -- the test harness strips spaces before calling `ResultOf`.
3. **Error conditions not enumerated** -- the plan should specify what errors `newBoard` checks for.
4. **Subtle Go details omitted** -- slice sharing through value receivers, dual return values from `at()`.

None of these gaps would prevent an experienced Go developer from producing a correct implementation, but they could cause confusion or bugs for someone following the plan literally. Addressing them would make the plan self-contained and unambiguous.
