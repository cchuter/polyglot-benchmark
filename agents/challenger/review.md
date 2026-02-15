# Challenger Review: bottle-song capitalize fix

## Verdict: PASS

The implementation correctly replaces the deprecated `strings.Title` with a manual `capitalize` helper, exactly as specified in the plan.

## Detailed Review

### 1. `strings.Title` completely removed?

**PASS.** No references to `strings.Title` exist in `bottle_song.go`. The only mentions are in `bottle_song_test.go` as comments documenting the exercism test infrastructure's own `Title` copy -- these are not production code and are not our concern.

### 2. Correctness of `capitalize` function

**PASS.** The function:

```go
func capitalize(s string) string {
    if s == "" {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

- Correctly uppercases the first character and preserves the rest.
- All values in `numberToWord` ("one", "two", ..., "ten") are simple ASCII lowercase strings. `s[:1]` safely captures one byte = one rune for all of them.
- Verified all 10 number words produce correct output: "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten".

### 3. Empty string handling

**PASS.** The guard `if s == "" { return s }` prevents a panic from `s[:1]` on empty input. While `capitalize` is never called with an empty string in practice (the default case only fires for n >= 3, where `numberToWord[n]` is always populated), this is still good defensive coding.

### 4. Imports clean?

**PASS.** Two imports remain:
- `"fmt"` -- used by `fmt.Sprintf` in `verse()`
- `"strings"` -- used by `strings.ToUpper` in `capitalize()`

Both are necessary. No unused imports.

### 5. Adherence to the plan

**PASS.** The implementation follows the plan exactly:
1. Added `capitalize` helper function (matches plan's code block exactly)
2. Replaced `strings.Title(numberToWord[n])` with `capitalize(numberToWord[n])` in the default case
3. Kept necessary imports

### 6. Control flow / edge case analysis

**PASS.** Traced all paths through `verse()`:
- `n == 1`: hardcoded strings, no `capitalize` call. Correct output verified against test cases.
- `n == 2`: hardcoded strings, no `capitalize` call. Correct output verified against test cases.
- `n >= 3` (default): uses `capitalize(numberToWord[n])` for the first two lines and `numberToWord[n-1]` (no capitalize) for the last line. This matches expected test output (e.g., "Ten green bottles" with capital T, "nine green bottles" lowercase for the "There'll be..." line).

### 7. No regressions introduced

**PASS.** The `Recite` function and `verse` function structure are unchanged. Only the internal implementation of how capitalization is done was modified. The public API remains identical.

## Summary

No issues found. The change is minimal, correct, and follows the plan precisely. Ready for testing.
