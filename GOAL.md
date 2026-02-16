# Goal: polyglot-go-counter

## Problem Statement

Design a comprehensive test suite for a line/letter/character counter tool in Go. This is a special "test design" exercise where instead of writing code to pass existing tests, we write tests that can detect bugs in provided implementations.

Four implementations (`Impl1`–`Impl4`) are provided. The test suite must:
- **Pass** when run against `Impl4` (the correct implementation)
- **Fail** when run against `Impl1`, `Impl2`, and `Impl3` (each has a distinct bug)

## Implementation Analysis

### Interface
```go
type Counter interface {
    AddString(string)
    Lines() int
    Letters() int
    Characters() int
}
```

### Bugs in Each Implementation

| Impl | Bug | Detail |
|------|-----|--------|
| Impl1 | Wrong line counting | Counts `\n` chars only; doesn't add 1 for last line when string doesn't end with `\n`. "Hello\nworld" → Lines()=1 instead of 2 |
| Impl2 | ASCII-only letters | Only counts `A-Z`/`a-z`; misses Unicode letters. "мир" → Letters()=0 instead of 3 |
| Impl3 | Byte iteration | Iterates `s[i]` by byte index instead of `range` by rune; wrong character/letter count for multi-byte UTF-8 |
| Impl4 | None (correct) | Uses `range` for runes, `unicode.IsLetter()`, proper line counting |

## Acceptance Criteria

1. `counter_test.go` contains valid Go test functions
2. `COUNTER_IMPL=4 go test` passes (all tests pass against correct impl)
3. `COUNTER_IMPL=1 go test` fails (detects Impl1's line counting bug)
4. `COUNTER_IMPL=2 go test` fails (detects Impl2's ASCII-only letter bug)
5. `COUNTER_IMPL=3 go test` fails (detects Impl3's byte-iteration bug)
6. Tests use `makeCounter()` factory function from `maker.go`
7. Tests cover: empty state, empty string, ASCII strings, Unicode strings, multiple AddString calls, strings with/without trailing newlines

## Key Constraints

- Tests must be in `counter_test.go` in package `counter`
- Must use `makeCounter()` to get Counter instances (env-var driven impl selection)
- The `counter.go` file is the "solution" stub but may remain minimal since tests are the deliverable
- Module uses `go 1.18`
