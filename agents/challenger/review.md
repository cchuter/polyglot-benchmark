# Challenger Review: hexadecimal.go

## Verdict: PASS — Implementation is correct

---

## 1. Hex Digit Conversion (0-9, a-f, A-F)

**Status: Correct**

The switch statement at lines 27-36 correctly handles all three ranges:
- `'0'-'9'` → `v = d - '0'` (0-9)
- `'a'-'f'` → `v = d - 'a' + 10` (10-15)
- `'A'-'F'` → `v = d - 'A' + 10` (10-15)
- Default → returns `ErrSyntax`

All test cases verified:
- `"1"` → 1 ✓
- `"10"` → 16 ✓
- `"2d"` → 45 ✓
- `"012"` → 18 ✓
- `"cfcfcf"` → 13619151 ✓
- `"CFCFCF"` → 13619151 (case insensitive) ✓

## 2. Overflow Detection Logic

**Status: Correct**

The overflow check has two guards (lines 37-44):

### Guard 1 (line 37): `n >= math.MaxInt64/16+1`

This prevents `n *= 16` from overflowing. The threshold is:
- `math.MaxInt64 / 16 + 1 = 576460752303423488`
- If `n >= 576460752303423488`, then `n * 16 >= 9223372036854775808` which overflows int64.

### Guard 2 (lines 41-43): `n1 < n`

This detects signed overflow when adding the digit value. Note: this guard is technically unreachable given Guard 1, since after Guard 1 passes, `n * 16` is at most `9223372036854775792` and `v` is at most 15, so `n1` maxes out at `9223372036854775807` (MaxInt64). Not a bug, just belt-and-suspenders.

### Traced test case: `"8000000000000000"`

- After processing "800000000000000" (15 chars): `n = 8 * 16^14 = 576460752303423488`
- At i=15 (the last `'0'`): Guard 1 triggers because `576460752303423488 >= 576460752303423488` → returns `ErrRange` ✓
- Correct: `0x8000000000000000 = 2^63 = MaxInt64 + 1`, which overflows int64.

### Traced test case: `"9223372036854775809"`

This is a 19-character hex string. The value grows exponentially and Guard 1 triggers well before all digits are processed. Returns `ErrRange` ✓

### Verified: MaxInt64 (`"7FFFFFFFFFFFFFFF"`) parses successfully

- Before last digit: `n = 0x7FFFFFFFFFFFFFF = 576460752303423487`
- Guard 1: `576460752303423487 < 576460752303423488` → passes
- `n *= 16` → `9223372036854775792`
- `n1 = 9223372036854775792 + 15 = 9223372036854775807 = MaxInt64` → no overflow ✓

## 3. Error Message Format

**Status: Correct**

`ParseError.Error()` (line 17) returns:
```
hexadecimal.ParseHex: parsing "X": invalid syntax
hexadecimal.ParseHex: parsing "X": value out of range
```

The test checks `strings.Contains(strings.ToLower(err.Error()), test.errCase)`:
- "syntax" is present in "invalid syntax" ✓
- "range" is present in "value out of range" ✓

## 4. HandleErrors Classification

**Status: Correct**

The switch-on-type-assertion pattern (lines 54-61) correctly classifies:
- `err == nil` → `"none"` ✓
- `*ParseError` with `ErrSyntax` → `"syntax"` ✓
- `*ParseError` with `ErrRange` → `"range"` ✓

Minor note: If `err` were non-nil but not a `*ParseError`, the result would be `""` (empty string). This cannot happen with the current `ParseHex` implementation, so not a concern.

## 5. No Built-in Hex Parsing Functions

**Status: Correct**

Imports are only `errors` and `math`. No `strconv`, `fmt.Sscanf`, or any other built-in hex parsing is used. ✓

## 6. Edge Cases

**Status: All handled correctly**

| Input | Expected | Result |
|-------|----------|--------|
| `""` (empty) | syntax error | `len(s) < 1` check at line 21 → ErrSyntax ✓ |
| `"peanut"` (invalid start) | syntax error | `'p'` hits default case → ErrSyntax ✓ |
| `"2cg134"` (invalid mid) | syntax error | `'g'` hits default at i=2 → ErrSyntax ✓ |
| `"8000000000000000"` | range error | Guard 1 triggers → ErrRange ✓ |
| `"9223372036854775809"` | range error | Guard 1 triggers → ErrRange ✓ |

## Issues Found

**None.** The implementation is correct and will pass all tests.

## Minor Observations (non-blocking)

1. **Dead code**: The `n1 < n` overflow check (lines 42-44) can never trigger given the Guard 1 check. This is harmless but unnecessary.
2. **HandleErrors robustness**: No fallback for unexpected error types. Safe given the closed implementation but could be fragile if `ParseHex` were modified to return different error types in the future.

Neither observation affects correctness or test passage.
