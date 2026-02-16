# Goal: Implement Hexadecimal to Decimal Converter (Go)

## Problem Statement

Implement a Go package that converts hexadecimal numbers (represented as strings) to their decimal (int64) equivalents using first principles. No built-in or external conversion libraries may be used. The implementation must also handle invalid input gracefully, distinguishing between syntax errors and range errors.

## Required Functions

1. **`ParseHex(s string) (int64, error)`** — Converts a hex string to int64. Must:
   - Accept both uppercase and lowercase hex digits (0-9, a-f, A-F)
   - Return a syntax error for empty strings or strings with invalid characters
   - Return a range error when the value overflows int64
   - Error messages must contain "syntax" or "range" (lowercase) as appropriate

2. **`HandleErrors(tests []string) []string`** — Takes a list of hex strings, calls ParseHex on each, and returns a list of error classification strings:
   - `"none"` — no error
   - `"syntax"` — syntax error
   - `"range"` — range/overflow error

## Acceptance Criteria

1. All tests in `hexadecimal_test.go` pass (`go test ./...` exits 0)
2. `ParseHex("1")` returns `(1, nil)`
3. `ParseHex("10")` returns `(16, nil)`
4. `ParseHex("2d")` returns `(45, nil)`
5. `ParseHex("012")` returns `(18, nil)`
6. `ParseHex("cfcfcf")` returns `(13619151, nil)`
7. `ParseHex("CFCFCF")` returns `(13619151, nil)` (case insensitive)
8. `ParseHex("")` returns an error containing "syntax"
9. `ParseHex("peanut")` returns an error containing "syntax"
10. `ParseHex("2cg134")` returns an error containing "syntax"
11. `ParseHex("8000000000000000")` returns an error containing "range"
12. `ParseHex("9223372036854775809")` returns an error containing "range"
13. `HandleErrors` correctly classifies all test inputs
14. No use of `strconv`, `fmt.Sscanf`, or any built-in hex parsing functions
15. Solution is in `hexadecimal.go` only (test file is not modified)

## Key Constraints

- First principles only: no built-in or external hex conversion libraries
- Must handle both upper and lowercase hex digits
- Must detect int64 overflow (range errors)
- Error type must satisfy the `error` interface
- Error messages must contain the substring "syntax" or "range" as appropriate
