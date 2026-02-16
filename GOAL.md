# Goal: Implement Alphametics Solver in Go

## Problem Statement

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` that solves alphametics puzzles. An alphametics puzzle replaces letters in words with digits such that the arithmetic equation holds true, each letter maps to a unique digit, and no multi-digit number has a leading zero.

The function signature is:
```go
func Solve(puzzle string) (map[string]int, error)
```

The puzzle string contains words separated by `+` operators and one `==` operator (e.g., `"SEND + MORE == MONEY"`).

## Acceptance Criteria

1. **All 10 test cases pass**: `go test ./...` in the alphametics directory passes with zero failures.
2. **Correct signature**: `func Solve(puzzle string) (map[string]int, error)` in package `alphametics`.
3. **Unique digit mapping**: Each letter maps to a different digit (0-9).
4. **No leading zeros**: The first letter of any multi-digit word must not map to 0.
5. **Error on no solution**: Returns a non-nil error when no valid solution exists.
6. **Correct map return**: Returns `map[string]int` with single uppercase letter keys and digit values.
7. **Handles edge cases**: Works for puzzles with 3 to 10 unique letters, and varying numbers of addends (from 2 to 199+).
8. **Reasonable performance**: The 10-letter / 199-addend test case completes within the Go test timeout (typically 10 minutes, but should target well under 1 minute).

## Key Constraints

- Must be in package `alphametics` in file `alphametics.go`
- Must use Go 1.18+ (per go.mod)
- Must not modify test files (`alphametics_test.go`, `cases_test.go`)
- The stub file currently only has the package declaration; the entire implementation must be added
