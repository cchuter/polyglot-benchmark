# Goal: polyglot-go-alphametics (Issue #13)

## Problem Statement

Implement a Go function `Solve` that solves alphametics puzzles. Alphametics puzzles replace letters with digits such that the resulting arithmetic equation is valid.

The function signature must be:
```go
func Solve(puzzle string) (map[string]int, error)
```

The puzzle string contains words separated by `+` operators and a `==` operator. Each letter maps to a unique digit (0-9). Leading digits of multi-digit numbers must not be zero.

## Acceptance Criteria

1. `Solve` accepts a puzzle string with zero or more `+` operators and one `==` operator
2. Returns `map[string]int` mapping each letter to its digit value, plus `nil` error on success
3. Returns `nil, error` when no valid solution exists
4. Each letter maps to a unique digit (no two letters share the same digit)
5. Leading digits of multi-digit numbers must not be zero
6. All 10 test cases in `cases_test.go` pass, including:
   - Puzzle with 3 letters: `I + BB == ILL`
   - Unique value constraint: `A == B` (error expected)
   - Leading zero invalid: `ACA + DD == BD` (error expected)
   - Two-digit carry: `A + A + ... + B == BCC`
   - 4, 6, 7, 8, 10 letter puzzles
   - 10-letter puzzle with 199 addends (performance test)
7. `go test` completes within a reasonable timeout (< 5 minutes)
8. Code compiles with `go build` without errors

## Key Constraints

- Package name: `alphametics`
- Go module: `alphametics` with `go 1.18`
- Must handle puzzles with many addends efficiently (199 addends test case)
- Test file and cases file are read-only (auto-generated from Exercism)
