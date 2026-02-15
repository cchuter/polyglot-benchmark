# Goal: Implement Go Alphametics Solver

## Problem Statement

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` to solve alphametics puzzles. An alphametics puzzle replaces letters with digits such that a given arithmetic equation holds true (e.g., `SEND + MORE == MONEY`).

The function stub exists but currently panics with "Please implement the Solve function". The test file and test cases are already provided.

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

- **Input**: A puzzle string with words separated by `+` and `==` operators (e.g., `"SEND + MORE == MONEY"`)
- **Output**: A `map[string]int` mapping each uppercase letter to a unique digit (0-9), or an error if no solution exists

## Acceptance Criteria

1. **AC1**: All 10 test cases in `cases_test.go` pass (`go test ./...` in the exercise directory)
2. **AC2**: The function correctly solves puzzles with 3-10 unique letters
3. **AC3**: The function returns an error for puzzles with no valid solution (e.g., `"A == B"`)
4. **AC4**: Leading digits of multi-digit numbers must not be zero (e.g., `"ACA + DD == BD"` has no valid solution because B would need to be 0)
5. **AC5**: Each letter maps to a unique digit (no two letters share the same digit)
6. **AC6**: The solution handles puzzles with many addends (up to 199 addends in test case 10)
7. **AC7**: The solution completes within a reasonable time (the benchmark test should not time out)

## Key Constraints

- Package name must be `alphametics`
- Must use Go module `alphametics` with `go 1.18`
- Must not modify test files (`alphametics_test.go`, `cases_test.go`)
- The function signature must exactly match: `func Solve(puzzle string) (map[string]int, error)`
