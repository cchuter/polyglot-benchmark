# Goal: Implement Alphametics Puzzle Solver (Issue #147)

## Problem Statement

Implement a Go function `Solve` in the `alphametics` package that solves alphametics puzzles — puzzles where letters in words are replaced with digits to form a valid arithmetic equation.

Given a puzzle string like `"SEND + MORE == MONEY"`, the function must find digit assignments for each letter such that the equation holds, each letter maps to a unique digit, and no multi-digit number has a leading zero.

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

## Acceptance Criteria

1. **Correct solutions**: For valid puzzles, return a `map[string]int` mapping each uppercase letter to its unique digit (0-9).
2. **No solution**: Return an error when no valid assignment exists (e.g., `"A == B"`).
3. **Leading zeros**: Multi-digit numbers must not have a leading zero; such assignments are invalid.
4. **Unique digits**: Each letter must map to a different digit.
5. **All 10 test cases pass**: The implementation must pass all test cases in `cases_test.go`, including:
   - 3-letter puzzles
   - Puzzles with no solution
   - Leading zero rejection
   - Carry propagation (11 addends)
   - 4, 6, 7, 8, 10-letter puzzles
   - Large puzzle with 199 addends and 10 letters
6. **Performance**: The 199-addend test case must complete in reasonable time (not brute-force all 10! permutations naively for 10-letter puzzles).
7. **`go test` passes**: All tests pass with `go test ./...`
8. **`go vet` passes**: No vet warnings.

## Key Constraints

- Package name: `alphametics`
- Module: `alphametics` (Go 1.18)
- Puzzle format: words separated by `+` and `==`, e.g., `"A + B == C"`
- Only uppercase ASCII letters in words
- The last word (after `==`) is the result
