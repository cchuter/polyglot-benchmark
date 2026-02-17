# Goal: polyglot-go-counter

## Problem Statement

Implement the Go "counter" exercise from the Exercism polyglot benchmark. This is a test-design exercise where the task is to have a comprehensive test suite (`counter_test.go`) that can detect bugs in multiple provided implementations of a line/letter/character counter tool.

The system under test counts lines, letters, and total characters in supplied strings via an `AddString` method, with `Lines()`, `Letters()`, and `Characters()` methods to retrieve totals.

Four implementations are provided:
- **Impl1**: Bug - incorrectly counts lines (only counts `\n` characters, missing the +1 for non-newline-ending strings)
- **Impl2**: Bug - only counts ASCII letters, not Unicode letters
- **Impl3**: Bug - iterates by byte index (`len(s)`) instead of runes, breaking multi-byte Unicode character handling
- **Impl4**: Correct implementation

## Acceptance Criteria

1. `COUNTER_IMPL=4 go test` passes (all tests pass for the correct implementation)
2. `COUNTER_IMPL=1 go test` fails (tests detect Impl1's line counting bug)
3. `COUNTER_IMPL=2 go test` fails (tests detect Impl2's ASCII-only letter counting)
4. `COUNTER_IMPL=3 go test` fails (tests detect Impl3's byte-level iteration bug)
5. The test suite covers edge cases: empty string, no AddString, Unicode, multiple AddString calls, newline-only strings, mixed content
6. `counter.go` exists with `package counter` declaration

## Key Constraints

- The `counter_test.go` file is the primary deliverable (it IS the solution for this test-design exercise)
- The `interface.go`, `maker.go`, and `impl*.go` files must NOT be modified
- Tests use `makeCounter()` from `maker.go` which requires `COUNTER_IMPL` env var
- The solution must work with Go 1.18+ (as specified in `go.mod`)
