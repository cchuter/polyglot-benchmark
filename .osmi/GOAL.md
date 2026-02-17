# Goal: polyglot-go-counter

## Problem Statement

Implement a comprehensive test suite for a line/letter/character counter tool in Go. This is an inverted exercise: instead of writing code to pass existing tests, the task is to write tests that validate counter implementations.

The counter system has:
- A `Counter` interface (`interface.go`) with methods: `AddString(string)`, `Lines() int`, `Letters() int`, `Characters() int`
- Four implementations (`impl1.go` through `impl4.go`), where `Impl4` is correct and `Impl1-3` contain specific bugs
- A factory function `makeCounter()` (`maker.go`) that selects an implementation via the `COUNTER_IMPL` environment variable

### Bug Analysis

- **Impl1**: Counts lines by counting `\n` only — misses that a non-empty string without a trailing newline still constitutes a line
- **Impl2**: Uses ASCII-only letter detection (`A-Z`, `a-z`) — misses Unicode letters (e.g., Cyrillic)
- **Impl3**: Iterates by byte index (`s[i]`) instead of using `range` — breaks multi-byte UTF-8 characters, producing incorrect character and letter counts for Unicode input
- **Impl4**: Correct implementation using `range` iteration and `unicode.IsLetter`

## Acceptance Criteria

1. `counter_test.go` contains a test suite that **passes all tests** when run with `COUNTER_IMPL=4`
2. The test suite **detects bugs in Impl1** (at least one test fails for `COUNTER_IMPL=1`)
3. The test suite **detects bugs in Impl2** (at least one test fails for `COUNTER_IMPL=2`)
4. The test suite **detects bugs in Impl3** (at least one test fails for `COUNTER_IMPL=3`)
5. `go vet ./...` passes with no issues
6. `go test` compiles and runs without errors (when COUNTER_IMPL is set)

## Key Constraints

- The solution is the **test file** (`counter_test.go`), not the production code
- Tests must use the `Counter` interface and `makeCounter()` factory
- The `counter.go` file should remain as a minimal package declaration
- Tests must cover: empty input, single strings, multiple `AddString` calls, newline handling, Unicode characters, and mixed content
