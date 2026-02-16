# Plan Review: beer-song

## Summary

The plan is well-structured and addresses the core requirements correctly. It identifies the right file to modify, the right functions to implement, and a sound overall approach. There are a few areas worth tightening up before implementation.

## 1. Coverage of Test Cases

### `Verse(n int) (string, error)`

The plan identifies five cases handled via a switch statement:

| Plan Case | Test Case | Covered? |
|-----------|-----------|----------|
| Invalid input (n < 0 or n > 99) | `{"invalid verse", 104, "", true}` | Yes |
| n == 0 | `{"verse 0", 0, verse0, false}` | Yes |
| n == 1 | `{"verse 1", 1, verse1, false}` | Yes |
| n == 2 | `{"verse 2", 2, verse2, false}` | Yes |
| default (3-99) | `{"a typical verse", 8, ...}` and `{"another typical verse", 3, ...}` | Yes |

All `Verse` test cases are covered.

### `Verses(start, stop int) (string, error)`

| Plan Case | Test Case | Covered? |
|-----------|-----------|----------|
| Valid range | `{"multiple verses", 8, 6, ...}` and `{"a different set of verses", 7, 5, ...}` | Yes |
| Invalid start | `{"invalid start", 109, 5, "", true}` | Yes |
| Invalid stop | `{"invalid stop", 99, -20, "", true}` | Yes |
| start < stop | `{"start less than stop", 8, 14, "", true}` | Yes |

All `Verses` test cases are covered.

### `Song() string`

The test (`TestEntireSong`) calls `Verses(99, 0)` and compares its output to `Song()`. The plan's approach of delegating to `Verses(99, 0)` directly matches this expectation.

**Verdict: All test cases are addressed.**

## 2. Edge Cases and Potential Pitfalls

### Trailing newline in `Verses` output -- CRITICAL

This is the most important detail to get right. Looking at the expected test output carefully:

- `verses86` and `verses75` are defined as raw string literals (backtick-quoted) that end with `\n\n` (a blank line followed by a closing backtick on the next line). This means each multi-verse output ends with a trailing `\n\n` after the last verse.

The plan says: "Join verses with an extra newline between them (each verse already ends with `\n`, so append `\n` after each)." This means the loop appends `\n` after **every** verse, including the last one. This produces the correct trailing `\n\n` that matches the test expectations.

The reference solution confirms this -- it writes `v + "\n"` for every verse in the loop, including the last. This is correct.

**However, the plan should explicitly note that the trailing `\n` after the final verse is intentional and required, not just an inter-verse separator.** A common mistake would be to only add `\n` *between* verses (not after the last one), which would fail the tests.

### `Song()` return type

The plan says `Song() string` (returning only a string, no error). The test confirms this -- `Song()` is called with a single return value. The reference solution also returns only `string`. The plan's description ("ignoring the error") is correct.

**Make sure the implementation signature is `Song() string`, not `Song() (string, error)`.** The plan's wording is slightly ambiguous -- it could be read as returning `(string, error)` and the caller ignoring the error, but the actual signature must be just `string`.

### Negative verse numbers

The plan validates `n < 0` for `Verse` and checks `[0, 99]` range for `Verses`. The test case `{"invalid stop", 99, -20, "", true}` exercises this path. This is correctly handled.

### `Verses` when start equals stop

The plan says `start >= stop` is valid. When `start == stop`, the loop should produce a single verse followed by `\n`. This is not directly tested but would work correctly with the described loop logic.

## 3. Soundness of Approach

The approach is sound and closely mirrors the reference solution:

- **Switch-based `Verse`**: Correct. The four special cases (0, 1, 2, and invalid) plus a default for 3-99 is the right decomposition.
- **Loop-based `Verses`**: Correct. Iterating from `start` down to `stop` (inclusive) and concatenating is the right approach.
- **Delegating `Song`**: Correct and simple.

The choice of `bytes.Buffer` for string concatenation in `Verses` is appropriate and matches the reference solution. It is more efficient than repeated string concatenation with `+`.

The choice of imports (`fmt` and `bytes`) is correct and matches exactly what the reference solution uses.

## 4. Suggestions for Improvement

1. **Clarify the trailing newline behavior in `Verses`**: The plan should explicitly state that `\n` is appended after every verse including the last, not just between verses. This is a subtle but critical detail that directly affects whether tests pass. A note like "Note: the output ends with a trailing blank line (double `\n`) after the last verse" would prevent confusion during implementation.

2. **Clarify `Song()` signature explicitly**: State that `Song()` returns only `string` (not `(string, error)`). The current wording "ignoring the error" could be misread. The reference solution uses a named return: `Song() (result string)` and a bare `return`, but a simpler `func Song() string { s, _ := Verses(99, 0); return s }` would also work.

3. **Consider noting the exact string literals for verse 0 and verse 1**: The plan describes them at a high level ("No more bottles..." / "Take it down...") but the exact wording matters. For instance, verse 0's second line is "Go to the store and buy some more, 99 bottles of beer on the wall." and verse 1 uses "Take it down" (not "Take one down") and "no more bottles" (lowercase "no"). These details are easy to get wrong. Referencing the test constants or the reference solution during implementation would be wise.

4. **Minor: the plan lists case n == 2 separately, which is correct**. This is needed because the second line of verse 2 says "1 bottle" (singular), whereas the default case uses "bottles" (plural) for `n-1`. Good catch in the plan.

## Conclusion

The plan is solid and will produce a correct implementation if followed carefully. The main risk is getting the trailing newline wrong in `Verses` or getting the exact string literals wrong in the special verse cases. The suggestions above are minor clarifications rather than fundamental issues with the approach.

**Recommendation: Approve with minor clarifications noted above.**
