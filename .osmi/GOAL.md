# Goal: polyglot-go-counter

## Problem Statement

Implement the "counter" exercise for the Go polyglot benchmark. This is a special exercise where the task is to design a test suite (not implementation code) for a line/letter/character counter tool.

The system under test counts lines, letters, and total characters in supplied strings. Multiple strings can be added via `AddString()`, and then `Lines()`, `Letters()`, and `Characters()` return the accumulated totals.

Four implementations are provided (impl1-4), where only Impl4 is correct:
- **Impl1**: Incorrect line counting (counts newlines only, never adds 1 for non-newline-ending text)
- **Impl2**: ASCII-only letter detection (misses Unicode letters like Cyrillic)
- **Impl3**: Byte-level iteration instead of rune-level (corrupts multi-byte Unicode)
- **Impl4**: Correct implementation using `range` for runes and `unicode.IsLetter()`

## Acceptance Criteria

1. `COUNTER_IMPL=4 go test ./...` passes (all tests pass with correct implementation)
2. `COUNTER_IMPL=1 go test ./...` fails (tests detect incorrect line counting)
3. `COUNTER_IMPL=2 go test ./...` fails (tests detect ASCII-only letter detection)
4. `COUNTER_IMPL=3 go test ./...` fails (tests detect byte-level iteration bug)
5. `go vet ./...` passes with no warnings
6. The test suite covers: empty input, no-add, simple ASCII, newlines, Unicode, multiple adds, and mixed content

## Key Constraints

- The solution file is `counter.go` (just needs package declaration)
- The test file `counter_test.go` is the actual deliverable
- Must work with the existing `interface.go`, `maker.go`, and `impl*.go` files
- Tests must use the `makeCounter()` factory from `maker.go`
