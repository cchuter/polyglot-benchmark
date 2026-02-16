# Goal: polyglot-go-counter (Issue #201)

## Problem Statement

Implement the "counter" Exercism exercise in Go. This is a special exercise where, instead of writing code to pass an existing test suite, the task is to **design a test suite** that can detect bugs in multiple provided implementations of a line/letter/character counter tool.

The system under test counts:
- **Lines**: Number of lines in supplied strings
- **Letters**: Number of Unicode letters in supplied strings
- **Characters**: Total number of characters (runes) in supplied strings

Four implementations are provided (`impl1.go` through `impl4.go`):
- **Impl1**: Incorrectly counts lines (counts newlines only, doesn't account for non-newline-terminated lines)
- **Impl2**: Incorrectly determines letters (only counts ASCII letters, misses Unicode)
- **Impl3**: Incorrectly handles Unicode (iterates bytes instead of runes, inflating character counts for multi-byte characters)
- **Impl4**: Correct implementation

## Acceptance Criteria

1. `COUNTER_IMPL=4 go test` passes all tests (correct implementation)
2. `COUNTER_IMPL=1 go test` fails at least one test (detects wrong line counting)
3. `COUNTER_IMPL=2 go test` fails at least one test (detects ASCII-only letter counting)
4. `COUNTER_IMPL=3 go test` fails at least one test (detects byte-level iteration for Unicode)
5. The test suite covers: empty counter, empty string, simple ASCII, newlines, Unicode, multiple AddString calls, edge cases
6. All code compiles without errors

## Key Constraints

- The `interface.go`, `maker.go`, and `impl1-4.go` files are **read-only fixtures** and must not be modified
- The test file `counter_test.go` is the primary deliverable
- The `counter.go` file is a stub (just package declaration) and needs no implementation code
- Tests must use the `Counter` interface and `makeCounter()` factory from the provided fixtures
