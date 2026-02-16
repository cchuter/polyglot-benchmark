# Implementation Plan: polyglot-go-counter

## Overview

Write a test suite in `counter_test.go` that validates the `Counter` interface implementations. The test suite must pass on `Impl4` (correct) and detect bugs in `Impl1`, `Impl2`, and `Impl3`.

## Files to Modify

- `go/exercises/practice/counter/counter_test.go` — Replace the stub with a complete test suite

## Architectural Decisions

1. **Helper function**: Use an `assertCounts` helper that checks all three counts (lines, letters, characters) in one call, reducing boilerplate and improving readability.
2. **Test structure**: Each test creates a fresh counter via `makeCounter()`, calls `AddString()` zero or more times, then asserts expected counts.
3. **No table-driven tests**: Since each test case exercises a distinct scenario with different semantics, individual test functions are clearer than a table-driven approach.

## Test Cases (ordered by complexity)

### 1. TestNoAddString — No strings added
- Expected: 0 lines, 0 letters, 0 characters
- Validates: baseline behavior of a fresh counter

### 2. TestEmptyString — Add empty string
- Expected: 0 lines, 0 letters, 0 characters
- Validates: empty input doesn't create a line

### 3. TestSimpleASCIINoNewline — "hello"
- Expected: 1 line, 5 letters, 5 characters
- **Catches Impl1**: Impl1 only counts `\n` as lines, so reports 0 lines for "hello"

### 4. TestASCIIWithNewlineInMiddle — "Hello\nworld!"
- Expected: 2 lines, 10 letters, 12 characters
- **Catches Impl1**: Reports 1 line instead of 2

### 5. TestStringEndingWithNewline — "hello\n"
- Expected: 1 line, 5 letters, 6 characters
- Validates: trailing newline doesn't add an extra line

### 6. TestUnicodeLetters — "здравствуй, мир\n"
- Expected: 1 line, 13 letters, 16 characters
- **Catches Impl2**: Only detects ASCII letters, reports 0 letters
- **Catches Impl3**: Iterates by byte (not rune), gets wrong character count (29 bytes vs 16 runes)

### 7. TestMultipleAddStrings — "hello\n" then "world"
- Expected: 2 lines, 10 letters, 11 characters
- Validates: accumulation across multiple AddString calls

### 8. TestOnlyNewlines — "\n\n\n"
- Expected: 3 lines, 0 letters, 3 characters
- Validates: newline-only strings

### 9. TestMixedContent — "abc 123!@#\ndef"
- Expected: 2 lines, 6 letters, 14 characters
- Validates: mixed alphanumeric and symbols

## Verification Steps

1. `COUNTER_IMPL=4 go test -v` — all tests pass
2. `COUNTER_IMPL=1 go test -v` — fails (line counting bug)
3. `COUNTER_IMPL=2 go test -v` — fails (Unicode letter bug)
4. `COUNTER_IMPL=3 go test -v` — fails (byte vs rune bug)
5. `go vet ./...` — no issues
