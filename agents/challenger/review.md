# Code Review: ParseOctal Implementation

## Verdict: APPROVED

## Files Reviewed
- `go/exercises/practice/octal/octal.go`

## Review Checklist

### 1. Correctness of octal-to-decimal math
**PASS** — Uses Horner's method (`num = num*8 + int64(ch-'0')`), which correctly accumulates the decimal value left-to-right. Manual verification:
- `"1"` → 1 ✓
- `"10"` → 1×8 + 0 = 8 ✓
- `"1234567"` → 342391 ✓

### 2. Error handling for non-octal characters
**PASS** — The guard `ch < '0' || ch > '7'` correctly rejects any character outside the octal digit range:
- `"carrot"` → 'c' > '7' → returns `(0, error)` ✓
- `"35682"` → '8' > '7' → returns `(0, error)` ✓

On error, the function returns `0` as the numeric value, matching the acceptance criteria.

### 3. No built-in conversion functions
**PASS** — Only `fmt.Errorf` is imported (for error formatting). No `strconv`, `fmt.Sscanf`, or other conversion utilities are used. The conversion is implemented from first principles.

### 4. Adherence to Go conventions
**PASS**
- Exported function with capital letter (`ParseOctal`)
- Idiomatic `(value, error)` return signature
- Clean, minimal code with no unnecessary abstractions
- Proper use of `range` over the string
- `int64` return type matches test expectations

### 5. Edge cases
- **"carrot"**: Rejected on first character ✓
- **"35682"**: Rejected when '8' is encountered ✓
- **Empty string ""**: Returns `(0, nil)` — not tested by the test suite, and not in acceptance criteria, so this is acceptable behavior.

## Summary
The implementation is correct, minimal, idiomatic, and satisfies all acceptance criteria. No issues found.
