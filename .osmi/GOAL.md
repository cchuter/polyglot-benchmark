# Goal: Implement Alphametics Solver (Issue #317)

## Problem Statement

Implement a function `Solve` in the `alphametics` Go package that solves alphametics puzzles. Alphametics puzzles replace letters in words with digits such that the resulting arithmetic equation is valid. Each letter maps to a unique digit (0-9), and leading digits of multi-digit numbers must not be zero.

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

**Input**: A string like `"SEND + MORE == MONEY"` containing words separated by `+` operators and one `==` operator.

**Output**: A `map[string]int` mapping each uppercase letter to its digit, or an error if no valid solution exists.

## Acceptance Criteria

1. `Solve` correctly parses puzzle strings with `+` and `==` operators
2. `Solve` returns a valid mapping where each letter maps to a unique digit (0-9)
3. Leading letters of multi-digit words are never mapped to zero
4. The arithmetic equation holds: sum of all addends equals the result
5. Returns an error when no valid solution exists (e.g., `"A == B"`)
6. All 10 test cases in `cases_test.go` pass, including:
   - 3-letter puzzles
   - Puzzles with no solution
   - Leading zero rejection
   - Carry propagation (11 addends)
   - 4, 6, 7, 8, 10 letter puzzles
   - Large puzzle with 199 addends and 10 letters
7. `go vet` reports no issues
8. Performance: the 199-addend puzzle must complete in reasonable time (not brute-force all 10! permutations naively)

## Key Constraints

- Package name: `alphametics`
- Module: `alphametics` with `go 1.18`
- Only modify `alphametics.go` — test files are auto-generated and read-only
- Must handle up to 10 unique letters (digits 0-9)
