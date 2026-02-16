# Goal: polyglot-go-hexadecimal

## Problem Statement

Implement the `hexadecimal` Go exercise: convert a hexadecimal string (e.g. "10af8c") to its decimal `int64` equivalent using first principles — no built-in or external conversion libraries.

The solution must also handle invalid input and overflow, returning typed errors.

## Acceptance Criteria

1. **`ParseHex(s string) (int64, error)`** correctly converts valid hex strings (uppercase and lowercase) to their `int64` decimal value.
2. **Empty strings** return a `*ParseError` whose `.Err` is `ErrSyntax` (error message contains "syntax").
3. **Invalid characters** (e.g. "peanut", "2cg134") return a `*ParseError` whose `.Err` is `ErrSyntax`.
4. **Overflow values** (e.g. "8000000000000000", "9223372036854775809") return a `*ParseError` whose `.Err` is `ErrRange` (error message contains "range").
5. **`HandleErrors(tests []string) []string`** calls `ParseHex` on each input and returns a list of "none", "syntax", or "range" strings corresponding to the error category.
6. All tests pass: `go test ./...` in the hexadecimal exercise directory.
7. `go vet ./...` passes with no issues.

## Key Constraints

- No use of `strconv.ParseInt`, `fmt.Sscanf`, or any built-in/external hex parsing.
- Must use first-principles character-by-character conversion.
- Must handle both upper and lowercase hex digits (a-f, A-F).
- Must detect int64 overflow and return range errors.
