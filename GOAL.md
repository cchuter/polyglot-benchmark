# Goal: Implement Alphametics Puzzle Solver in Go

## Problem Statement

Implement a `Solve` function in the `alphametics` Go package that solves alphametics puzzles. An alphametics puzzle replaces letters in words with digits such that the resulting arithmetic equation is valid. Each letter maps to a unique digit (0-9), and leading digits of multi-digit numbers must not be zero.

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

- **Input**: A string like `"SEND + MORE == MONEY"` containing words joined by `+` operators with a `==` separator before the result word.
- **Output**: A `map[string]int` mapping each single letter to its digit, or an error if no solution exists.

## Acceptance Criteria

1. All 10 test cases in `cases_test.go` pass (`go test ./...`).
2. `go vet ./...` reports no issues.
3. The function correctly handles:
   - Simple puzzles (3 letters, e.g., `"I + BB == ILL"`)
   - Puzzles with no solution (returns error)
   - Puzzles where a leading-zero solution would be the only option (returns error)
   - Puzzles with carry propagation (e.g., `"A + A + A + ... + B == BCC"`)
   - Large puzzles with 10 unique letters and many addends (199 addends test case)
4. Each letter maps to a unique digit (0-9).
5. Leading letters of multi-digit numbers are not assigned 0.
6. Performance: the solver must complete the 199-addend test case in reasonable time (the benchmark test exists).

## Key Constraints

- Must use package name `alphametics`
- Must use Go 1.18 (as specified in go.mod)
- Must not modify test files (`alphametics_test.go`, `cases_test.go`)
- Only `alphametics.go` should be modified
