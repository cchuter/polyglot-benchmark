# Goal: Implement Alphametics Puzzle Solver (Go)

## Problem Statement

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` that solves alphametics puzzles. Alphametics puzzles replace letters in words with digits such that the resulting arithmetic equation is valid.

The function signature is:
```go
func Solve(puzzle string) (map[string]int, error)
```

The puzzle string contains words separated by `+` operators and one `==` operator. For example: `"SEND + MORE == MONEY"`.

## Acceptance Criteria

1. **All tests pass**: `go test ./...` in the alphametics exercise directory passes all test cases defined in `cases_test.go` and `alphametics_test.go`.

2. **Correct function signature**: `func Solve(puzzle string) (map[string]int, error)` accepting a puzzle string and returning a letter-to-digit mapping or an error.

3. **Unique digit assignment**: Each letter maps to a unique digit (0-9). No two letters share the same digit.

4. **No leading zeros**: The leading letter of any multi-digit word must not map to zero.

5. **Error on no solution**: Returns a non-nil error when no valid solution exists (e.g., `"A == B"` or `"ACA + DD == BD"`).

6. **Handles all test cases**: Including:
   - Simple 3-letter puzzles (`"I + BB == ILL"`)
   - Puzzles with many addends (`"A + A + A + ... + B == BCC"`)
   - Classic SEND + MORE == MONEY
   - Large 10-letter puzzle with 199 addends
   - Invalid puzzles that should return errors

7. **Performance**: Must complete all test cases within a reasonable time, including the 199-addend puzzle with 10 unique letters.

## Key Constraints

- Package name must be `alphametics`
- Must use Go 1.18+ (as specified in go.mod)
- Only modify `alphametics.go`; test files are read-only
- Solution must handle puzzles with 1-10 unique letters
- Digits range from 0-9 (at most 10 unique letters)
