# Goal: Implement Alphametics Solver in Go

## Problem Statement

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` that solves alphametics puzzles. Alphametics puzzles replace letters in words with digits such that the resulting arithmetic equation is valid.

Example: `SEND + MORE == MONEY` solves to `9567 + 1085 = 10652`.

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

- **Input**: A puzzle string containing uppercase words separated by `+` operators and one `==` operator (e.g., `"SEND + MORE == MONEY"`)
- **Output**: A `map[string]int` mapping each letter to its digit, or an error if no solution exists

## Acceptance Criteria

1. All 10 test cases in `cases_test.go` pass (`go test ./...`)
2. The function correctly solves puzzles with 3 to 10 unique letters
3. Each letter maps to a unique digit (0-9)
4. Leading digits of multi-digit numbers must not be zero
5. Returns an error when no valid solution exists (e.g., `"A == B"`, `"ACA + DD == BD"`)
6. Handles puzzles with many addends (up to 199 addends in test case 10)
7. The solution completes within a reasonable time (benchmark test exists)
8. Code compiles without errors
9. Package name is `alphametics`
10. Only `alphametics.go` is modified; test files are read-only

## Key Constraints

- Solution must be in a single file: `alphametics.go`
- Must use Go 1.18+ (as per go.mod)
- The puzzle format uses `==` (not `=`) as the equality operator
- Words in the puzzle consist entirely of uppercase ASCII letters
