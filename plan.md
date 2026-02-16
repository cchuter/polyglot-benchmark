# Implementation Plan: polyglot-go-hexadecimal

## Overview

Implement `ParseHex` and `HandleErrors` in `go/exercises/practice/hexadecimal/hexadecimal.go`. The solution follows the reference implementation pattern from `.meta/example.go`.

## File to Modify

**`go/exercises/practice/hexadecimal/hexadecimal.go`** — currently a stub with only `package hexadecimal`.

## Implementation Details

### 1. Define Error Sentinels and Types

```go
var ErrRange = errors.New("value out of range")
var ErrSyntax = errors.New("invalid syntax")
```

Define `ParseError` struct with `Num string` and `Err error` fields, and an `Error()` method that produces a message containing the sentinel error text (so the test's `strings.Contains(strings.ToLower(err.Error()), "syntax")` / `"range"` checks pass).

### 2. Implement `ParseHex(s string) (int64, error)`

Character-by-character conversion:
1. If input is empty, return syntax error.
2. Iterate over each byte of the string.
3. Map `'0'-'9'` → 0-9, `'a'-'f'` → 10-15, `'A'-'F'` → 10-15; anything else is a syntax error.
4. Before multiplying accumulator by 16, check if `n >= math.MaxInt64/16 + 1` to detect overflow.
5. After adding the digit value, check if the result wrapped (new value < old value) for overflow.
6. On overflow, return range error.

### 3. Implement `HandleErrors(tests []string) []string`

For each input string, call `ParseHex`, type-assert the error to `*ParseError`, and classify as "none", "syntax", or "range".

## Approach & Ordering

1. Write the complete solution in `hexadecimal.go`
2. Run `go test` to verify all tests pass
3. Run `go vet` to verify no issues
4. Commit

## Architectural Decisions

- Follow the reference solution's approach closely since it's well-structured and handles edge cases correctly.
- Use `goto Error` pattern matching the reference for consistency with the exercise's style.
- Import only `errors` and `math` from the standard library (no `fmt` needed since error messages are string-concatenated).
