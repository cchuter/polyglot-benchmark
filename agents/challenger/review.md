# Challenger Review: ParseOctal Implementation

## Verdict: PASS

The implementation is correct, clean, idiomatic Go, and faithfully follows the selected plan (Proposal A — Bit-Shift Approach).

---

## 1. Correctness Against Test Cases

All 5 test cases produce the expected results:

| Input | Expected | Trace | Result |
|-------|----------|-------|--------|
| `"1"` | `(1, nil)` | `0<<3 + 1 = 1` | PASS |
| `"10"` | `(8, nil)` | `0<<3+1=1 → 1<<3+0=8` | PASS |
| `"1234567"` | `(342391, nil)` | `1→10→83→668→5349→42798→342391` | PASS |
| `"carrot"` | `(0, error)` | `'c' > '7' → error on first rune` | PASS |
| `"35682"` | `(0, error)` | `'8' > '7' → error at 4th char` | PASS |

### Operator Precedence Verification

The expression `num<<3 + int64(digit-'0')` relies on Go operator precedence. In Go, `<<` (precedence 5) binds tighter than `+` (precedence 4), so this is parsed as `(num<<3) + int64(digit-'0')`. This is **correct** — no parentheses needed.

## 2. Edge Cases

| Edge Case | Behavior | Assessment |
|-----------|----------|------------|
| Empty string `""` | Returns `(0, nil)` | Acceptable — the loop doesn't execute, returning zero. The instructions say "treat invalid input as octal 0". Tests don't cover this case. |
| Leading zeros `"0777"` | Correctly computes `511` | Correct — leading zeros are valid in octal. |
| Very large numbers | Silent int64 overflow | Acceptable — no overflow protection, but tests don't require it and the reference solution doesn't include it either. |
| Multi-byte UTF-8 runes | Returns error | Correct — any non-ASCII rune will fail the `'0'-'7'` range check. |

## 3. Error Handling

- Returns `(0, fmt.Errorf("unexpected rune '%c'", digit))` on invalid digits.
- Error message format matches the plan specification exactly.
- Uses `fmt.Errorf` which is idiomatic Go for formatted errors.
- Early return on first invalid digit is correct — no partial results leaked.

## 4. Go Idioms and Code Quality

- **Bit-shift `<<3`** for multiply-by-8: idiomatic for base-conversion code in Go.
- **`range` over string**: correctly iterates over runes (not bytes), which is the Go-idiomatic way.
- **Minimal imports**: only `fmt`, no unnecessary dependencies.
- **Clean structure**: 16 lines total, single function, no over-engineering.
- **Variable naming**: `num`, `digit`, `octal` are clear and appropriate.

## 5. Adherence to Plan

| Plan Requirement | Implementation | Status |
|-----------------|---------------|--------|
| Bit-shift approach (`<<3`) | `num = num<<3 + int64(digit-'0')` | MATCH |
| First principles (no built-in conversion) | No `strconv`, `fmt.Sscanf`, etc. | MATCH |
| Parameter name `octal` | `func ParseOctal(octal string)` | MATCH |
| Error message `"unexpected rune '%c'"` | `fmt.Errorf("unexpected rune '%c'", digit)` | MATCH |
| Package `octal` | `package octal` | MATCH |
| Import `"fmt"` | `import ("fmt")` | MATCH |

## 6. Comparison to Reference Solution (`.meta/example.go`)

The implementation is functionally identical to the reference solution. The only difference is the absence of inline comments present in `example.go`. This is acceptable — the code is self-explanatory.

## 7. Files Modified

Only `go/exercises/practice/octal/octal.go` was modified, which is within scope. No test files, `go.mod`, or `.meta/` files were touched.

## Summary

The implementation is a clean, correct, minimal solution that:
- Passes all 5 required test cases
- Follows the selected plan exactly
- Matches the reference solution
- Uses idiomatic Go patterns
- Stays within the defined scope

No changes requested.
