# Goal: Implement Alphametics Puzzle Solver in Go

## Problem Statement

Implement a `Solve` function in the `alphametics` Go package that solves alphametics puzzles. Alphametics puzzles replace letters in words with digits such that the arithmetic equation holds true.

Given a puzzle string like `"SEND + MORE == MONEY"`, the function must find a mapping of letters to unique digits (0-9) that satisfies the equation, with the constraint that no multi-digit number may have a leading zero.

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

## Acceptance Criteria

1. **AC1**: `Solve` correctly solves simple puzzles (e.g., `"I + BB == ILL"` → `{"B":9, "I":1, "L":0}`)
2. **AC2**: `Solve` returns an error when no valid solution exists (e.g., `"A == B"`)
3. **AC3**: `Solve` returns an error when the only possible solution would require a leading zero on a multi-digit number (e.g., `"ACA + DD == BD"`)
4. **AC4**: `Solve` handles puzzles with many addends and carries (e.g., `"A + A + A + A + A + A + A + A + A + A + A + B == BCC"`)
5. **AC5**: `Solve` handles the classic `"SEND + MORE == MONEY"` puzzle correctly
6. **AC6**: `Solve` handles large puzzles with 10 unique letters and 199 addends (the `FORTRESSES` test case)
7. **AC7**: Each letter maps to a unique digit (0-9); no two letters share the same digit
8. **AC8**: All existing tests in `alphametics_test.go` and `cases_test.go` pass
9. **AC9**: The solution is implemented in `alphametics.go` in the `alphametics` package

## Key Constraints

- The puzzle format uses `+` for addition and `==` for equality
- There may be zero or more `+` operators and exactly one `==` operator
- Letters are uppercase A-Z
- Each letter must map to a unique digit 0-9
- Leading digits of multi-digit numbers must not be zero
- Maximum of 10 unique letters (since only digits 0-9 exist)
- Must handle puzzles with many addends efficiently (the 199-addend test case)
