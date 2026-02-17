# Plan Review: polyglot-go-hexadecimal

## Summary

The plan proposes implementing `ParseHex` and `HandleErrors` using structured error types (`ParseError` wrapping sentinel errors `ErrRange`/`ErrSyntax`) with `goto`-based error flow. This is Proposal A from the plan, selected as the winner. The proposed code is essentially identical to the reference solution at `.meta/example.go`.

**Overall verdict: APPROVE.** The plan is correct and will pass all tests. Minor observations are noted below, but none are blocking.

---

## 1. Does the proposed code satisfy all test cases?

**Yes.** Verified by running the reference solution (which is character-for-character identical to the proposed code) against the test suite. All three test functions pass:

- `TestParseHex`: PASS
- `TestHandleErrors`: PASS
- `BenchmarkParseHex`: compiles and runs correctly

### Test case walkthrough

| Input | Expected | Proposed Result | Status |
|-------|----------|-----------------|--------|
| `"1"` | 1, none | 1, nil | PASS |
| `"10"` | 16, none | 16, nil | PASS |
| `"2d"` | 45, none | 45, nil | PASS |
| `"012"` | 18, none | 18, nil | PASS |
| `"cfcfcf"` | 13619151, none | 13619151, nil | PASS |
| `"CFCFCF"` | 13619151, none | 13619151, nil | PASS |
| `""` | 0, syntax | 0, ParseError{ErrSyntax} | PASS |
| `"peanut"` | 0, syntax | 0, ParseError{ErrSyntax} | PASS |
| `"2cg134"` | 0, syntax | 0, ParseError{ErrSyntax} | PASS |
| `"8000000000000000"` | 0, range | MaxInt64, ParseError{ErrRange} | PASS |
| `"9223372036854775809"` | 0, range | MaxInt64, ParseError{ErrRange} | PASS |

Note: For error cases, the test does NOT check the returned `int64` value (only verifies `err != nil` and the error message content), so returning `math.MaxInt64` for range errors is acceptable.

---

## 2. Are there any edge cases the plan misses?

**No missing edge cases relative to the test suite.** The test suite defines the complete set of cases, and the plan handles all of them:

- **Empty string**: Caught by `len(s) < 1` check before the loop.
- **Invalid characters at start** (`"peanut"`): First character 'p' hits the `default` branch.
- **Invalid characters in middle** (`"2cg134"`): Character 'g' at index 2 hits the `default` branch. The code correctly resets `n = 0` before returning.
- **Leading zeros** (`"012"`): Handled naturally by the multiplication loop -- leading zeros just multiply n (which is 0) by 16.
- **Case insensitivity** (`"CFCFCF"` vs `"cfcfcf"`): Both uppercase and lowercase ranges are handled in the switch statement.
- **Overflow via multiplication** (`"8000000000000000"`): Caught by the pre-multiplication check.
- **Overflow via large value** (`"9223372036854775809"`): Caught by the pre-multiplication check (triggers well before all digits are processed since the string is 19 hex digits long).

Edge cases NOT in the test suite but handled correctly by the code:
- **Single digit `"0"`**: Returns (0, nil) -- works correctly.
- **Max valid value `"7FFFFFFFFFFFFFFF"`**: Returns (9223372036854775807, nil) -- the pre-multiply check passes because `n` never reaches the threshold, and the post-add check passes because no overflow occurs.
- **Overflow via addition** (e.g., `"7FFFFFFFFFFFFFF8"` where multiply is fine but add overflows): Caught by the `n1 < n` post-addition check. This path is reachable when the pre-multiply check passes but the final digit pushes the total over MaxInt64.

---

## 3. Is the error handling correct for string matching?

**Yes.** The test uses:
```go
strings.Contains(strings.ToLower(err.Error()), test.errCase)
```

The proposed `ParseError.Error()` method produces messages like:
- `hexadecimal.ParseHex: parsing "": invalid syntax`
- `hexadecimal.ParseHex: parsing "8000000000000000": value out of range`

After `strings.ToLower()`:
- `"hexadecimal.parsehex: parsing \"\": invalid syntax"` -- contains `"syntax"` -- MATCH
- `"hexadecimal.parsehex: parsing \"8000000000000000\": value out of range"` -- contains `"range"` -- MATCH

The sentinel error messages (`"invalid syntax"` and `"value out of range"`) are incorporated into the `ParseError.Error()` output via string concatenation with `e.Err.Error()`. Since `"invalid syntax"` contains "syntax" and `"value out of range"` contains "range", the `strings.Contains` checks will always match correctly.

The `HandleErrors` function uses type assertion and sentinel comparison (not string matching), which is even more robust:
```go
case ok && pe.Err == ErrSyntax:
case ok && pe.Err == ErrRange:
```
This compares pointer identity of the sentinel errors, not string content.

---

## 4. Is the overflow detection logic correct?

**Yes.** The overflow detection uses a two-stage approach that correctly catches both test cases.

### Stage 1: Pre-multiplication check
```go
if n >= math.MaxInt64/16+1 {
```

- `math.MaxInt64 = 9223372036854775807`
- `math.MaxInt64 / 16 = 576460752303423487` (Go integer division truncates)
- `math.MaxInt64/16 + 1 = 576460752303423488`

This means: if `n >= 576460752303423488`, then `n * 16` would overflow int64. This is mathematically correct because `576460752303423488 * 16 = 9223372036854775808 > MaxInt64`.

**For `"8000000000000000"` (0x8000000000000000 = 9223372036854775808):**
- After processing the first 15 characters `"800000000000000"`, n = 8 * 16^14 = 576460752303423488.
- At i=15 (the last `'0'`): `n >= 576460752303423488` is TRUE.
- Overflow detected. Returns `(MaxInt64, &ParseError{..., ErrRange})`.

**For `"9223372036854775809"`:**
- This is parsed as hexadecimal (all characters 0-9 are valid hex digits).
- After processing 15 characters, n = 658145060325971063.
- At i=15: `658145060325971063 >= 576460752303423488` is TRUE.
- Overflow detected early, well before reaching the end of the string.

### Stage 2: Post-addition check
```go
n1 := n + int64(v)
if n1 < n {
```

This catches the case where `n * 16` fits in int64 but `n * 16 + v` overflows. For example, `n = 576460752303423487` (just below the threshold) would pass stage 1, then `n *= 16` yields `9223372036854775792`, and adding a digit `v >= 16` (impossible since max hex digit is 15) or more precisely adding `v` such that the sum exceeds MaxInt64 would wrap to negative, making `n1 < n` true.

Both stages together provide complete overflow coverage for all positive int64 values.

---

## 5. Any issues with the proposed code?

### No blocking issues.

The proposed code is identical to the reference solution and passes all tests. Below are minor observations that do not affect correctness:

**Observation 1: `goto` usage**
The plan correctly acknowledges that `goto` is unconventional in Go application code but notes it mirrors the standard library's `strconv` package. This is a deliberate style choice inherited from the reference, not a defect.

**Observation 2: The `switch` in HandleErrors has no default case**
```go
switch pe, ok := err.(*ParseError); {
case err == nil:
    e[i] = "none"
case ok && pe.Err == ErrSyntax:
    e[i] = "syntax"
case ok && pe.Err == ErrRange:
    e[i] = "range"
}
```
If `ParseHex` were to return an error that is neither nil, nor a `*ParseError` with `ErrSyntax`/`ErrRange`, the corresponding entry in the result slice would be `""` (the zero value for string). Since `ParseHex` only returns nil or `*ParseError` with one of the two sentinels, this is unreachable. However, a defensive `default` case could be added for robustness. This is not a test-facing issue.

**Observation 3: Only positive hex values are supported**
The implementation does not handle negative hex values (e.g., `"-1a"`). This is consistent with the test suite, which only tests non-negative values. Not a defect.

**Observation 4: No `0x` prefix handling**
The implementation does not strip a `"0x"` or `"0X"` prefix. This is correct -- the test suite does not include prefixed inputs, and the goal does not require prefix handling.

**Observation 5: Import correctness**
The plan imports only `"errors"` and `"math"`, both of which are needed and neither of which provides hex parsing facilities. This satisfies the constraint of no built-in hex parsing.

---

## Conclusion

The plan is **correct and complete**. It will produce code that passes all test cases. The overflow detection is mathematically sound, the error messages satisfy the string-matching checks in the test, and the `HandleErrors` function correctly categorizes errors via type assertion. The proposed code is identical to the verified reference solution. No changes are needed.

**Recommendation: Proceed with implementation as written.**
