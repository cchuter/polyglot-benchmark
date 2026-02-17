# Goal: polyglot-go-hexadecimal

## Problem Statement

Implement a hexadecimal-to-decimal converter in Go. The solution must convert a hexadecimal number represented as a string (e.g. "10af8c") to its decimal equivalent using first principles — no built-in or external conversion libraries allowed.

The solution must also handle invalid hexadecimal strings, distinguishing between syntax errors (invalid characters, empty input) and range errors (values that overflow int64).

## Acceptance Criteria

1. **`ParseHex(string) (int64, error)`** correctly converts valid hex strings to int64 values:
   - Single digit: `"1"` -> `1`
   - Multi-digit: `"10"` -> `16`
   - Mixed digits/letters: `"2d"` -> `45`
   - Leading zeros: `"012"` -> `18`
   - Long strings: `"cfcfcf"` -> `0xcfcfcf`
   - Case insensitive: `"CFCFCF"` -> same as `"cfcfcf"`

2. **Error handling for syntax errors**:
   - Empty string returns error containing "syntax"
   - Invalid characters (e.g. "peanut") returns error containing "syntax"
   - Partially invalid (e.g. "2cg134") returns error containing "syntax"

3. **Error handling for range errors**:
   - Values exceeding int64 max (e.g. "8000000000000000") return error containing "range"
   - Very large values (e.g. "9223372036854775809") return error containing "range"

4. **`HandleErrors([]string) []string`** processes a list of hex strings and returns a matching list of error classifications: "none", "syntax", or "range".

5. All tests pass: `go test ./go/exercises/practice/hexadecimal/`

## Key Constraints

- Must use first principles — no `strconv.ParseInt`, `fmt.Sscanf`, or similar built-in hex parsing
- Error messages must contain the substrings "syntax" or "range" (case-insensitive match in tests)
- Must handle both uppercase and lowercase hex digits
- Must detect int64 overflow correctly
