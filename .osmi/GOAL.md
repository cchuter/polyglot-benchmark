# Goal: Implement Alphametics Puzzle Solver in Go

## Problem Statement

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` that solves alphametics puzzles. Alphametics puzzles replace letters with digits such that a mathematical equation holds true, with each letter mapping to a unique digit and leading digits being non-zero.

## Acceptance Criteria

1. **Function Signature**: `func Solve(puzzle string) (map[string]int, error)` in package `alphametics`
2. **Parsing**: Correctly parse puzzle strings with `+` operators and `==` equality (e.g., `"SEND + MORE == MONEY"`)
3. **Unique Digits**: Each letter maps to a unique digit (0-9)
4. **No Leading Zeros**: The first letter of any multi-digit word must not map to 0
5. **Correct Solutions**: Return a `map[string]int` mapping each letter to its digit value
6. **Error on No Solution**: Return an error when no valid solution exists
7. **All 10 Test Cases Pass**: The solution must pass all test cases in `cases_test.go`, including:
   - Simple 3-letter puzzles
   - Puzzles with no solution (unique value constraint, leading zero)
   - Puzzles with carry propagation
   - Puzzles with 4, 6, 7, 8, and 10 letters
   - A large puzzle with 199 addends and 10 letters
8. **Performance**: The 199-addend puzzle must complete in reasonable time (not timeout)

## Key Constraints

- Must use Go 1.18+ (per go.mod)
- Must be in package `alphametics`
- Only modify `alphametics.go` - test files are read-only
- Must handle puzzles with multiple addends (not just two)
