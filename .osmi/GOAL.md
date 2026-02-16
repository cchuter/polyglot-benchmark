# Goal: polyglot-go-counter

## Problem Statement

Design a test suite (`counter_test.go`) for a line/letter/character counter tool in Go. This is a special exercise where the solution is the test code itself, not production code. Four implementations (`impl1.go` through `impl4.go`) are provided; the test suite must pass on the correct implementation (`impl4`) and detect bugs in the incorrect ones (`impl1`, `impl2`, `impl3`).

## Bugs in Each Implementation

- **Impl1**: Counts lines incorrectly — only counts `\n` characters, missing that a non-empty string without a trailing newline still constitutes a line.
- **Impl2**: Only counts ASCII letters (`A-Z`, `a-z`), missing Unicode letters. Uses `unicode.IsLetter` incorrectly — actually doesn't use it at all.
- **Impl3**: Iterates by byte index (`for i := 0; i < len(s); i++`) instead of by rune (`for _, char := range s`), causing incorrect character/letter counts for multi-byte Unicode characters.
- **Impl4**: Correct implementation — uses `range` for proper Unicode handling, counts lines correctly (newlines + 1 for non-empty non-newline-terminated content), and uses `unicode.IsLetter` for letter detection.

## Acceptance Criteria

1. `counter_test.go` contains a comprehensive test suite covering:
   - No strings added (zero counts)
   - Empty string (zero counts)
   - Simple ASCII string without newline
   - ASCII string with newline in the middle
   - String ending with newline
   - Unicode (non-ASCII) letters
   - Multiple `AddString` calls (accumulation)
   - Only newlines
   - Mixed content (letters, digits, symbols, newlines)
2. Tests pass with `COUNTER_IMPL=4`: `COUNTER_IMPL=4 go test ./...` succeeds
3. Tests fail with `COUNTER_IMPL=1`: detects incorrect line counting
4. Tests fail with `COUNTER_IMPL=2`: detects failure to count Unicode letters
5. Tests fail with `COUNTER_IMPL=3`: detects incorrect byte-vs-rune counting
6. `go vet ./...` passes with no issues
7. Code compiles cleanly

## Key Constraints

- The solution file is `counter_test.go` (not `counter.go`)
- Must use the `Counter` interface defined in `interface.go`
- Must use the `makeCounter()` factory from `maker.go`
- Tests run via environment variable `COUNTER_IMPL` to select implementation
