# Goal: polyglot-go-hexadecimal (Issue #299)

## Problem Statement

Implement a hexadecimal-to-decimal converter in Go. The solution must convert hexadecimal strings (e.g. "10af8c") to their `int64` decimal equivalents using first principles — no built-in or external conversion libraries allowed. The program must also handle invalid inputs gracefully, distinguishing between syntax errors (invalid characters, empty input) and range errors (overflow beyond int64).

## Acceptance Criteria

1. **`ParseHex(string) (int64, error)`** correctly converts valid hex strings to int64:
   - `"1"` → 1
   - `"10"` → 16
   - `"2d"` → 45
   - `"012"` → 18
   - `"cfcfcf"` → 13619151
   - `"CFCFCF"` → 13619151 (case insensitive)

2. **`ParseHex`** returns appropriate errors for invalid inputs:
   - Empty string `""` → error containing "syntax"
   - Invalid characters `"peanut"` → error containing "syntax"
   - Partial invalid `"2cg134"` → error containing "syntax"
   - Overflow `"8000000000000000"` → error containing "range"
   - Overflow `"9223372036854775809"` → error containing "range"

3. **`HandleErrors([]string) []string`** calls `ParseHex` on each input and returns a list of `"none"`, `"syntax"`, or `"range"` strings corresponding to the error category of each input.

4. All tests in `hexadecimal_test.go` pass: `TestParseHex`, `TestHandleErrors`, `BenchmarkParseHex`.

5. No use of `strconv`, `fmt.Sscanf`, `math/big`, or any other built-in hex parsing facilities.

## Key Constraints

- Solution must be in `go/exercises/practice/hexadecimal/hexadecimal.go`
- Package name must be `hexadecimal`
- Must use Go 1.18 (per go.mod)
- First principles only — manual character-by-character parsing
- Error type must satisfy the `error` interface
- Error messages must contain the substring "syntax" or "range" (case-insensitive match used in tests)
