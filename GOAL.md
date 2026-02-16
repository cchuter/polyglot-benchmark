# Goal: Implement Alphametics Solver in Go

## Problem Statement

Implement a `Solve` function in the `alphametics` Go package that solves alphametics puzzles. Alphametics puzzles replace letters in arithmetic expressions with digits, where each letter maps to a unique digit and leading letters of multi-digit numbers cannot be zero.

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

## Input Format

- A string containing an arithmetic expression with uppercase letters
- Zero or more `+` operators on the left-hand side
- One `==` operator separating the left-hand side from the right-hand side result
- Example: `"SEND + MORE == MONEY"`

## Output

- On success: a `map[string]int` mapping each unique letter to its digit (0-9), and `nil` error
- On failure (no valid solution): `nil` map and a non-nil error

## Acceptance Criteria

1. All 10 test cases in `cases_test.go` pass (`go test` exits with 0)
2. The function correctly parses puzzle strings with `+` and `==` operators
3. Each letter maps to a unique digit (0-9)
4. Leading digits of multi-digit numbers are non-zero
5. The arithmetic equation holds when letters are replaced with their assigned digits
6. Returns an error when no valid solution exists (e.g., `"A == B"`)
7. Handles puzzles with varying numbers of letters (3 to 10 unique letters)
8. Handles puzzles with many addends (up to 199 addends in the test suite)
9. Solution completes within reasonable time (test suite should not timeout)

## Key Constraints

- Must be in package `alphametics`
- Must use Go 1.18 (as specified in go.mod)
- Must implement exactly `func Solve(puzzle string) (map[string]int, error)`
- The stub file `alphametics.go` already declares the package
- Test files are read-only and auto-generated
