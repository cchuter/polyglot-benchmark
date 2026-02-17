# Goal: Implement Alphametics Solver (Issue #233)

## Problem Statement

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` that solves alphametics puzzles. An alphametics puzzle substitutes letters for digits in an arithmetic equation. The solver must find a mapping of letters to digits (0-9) such that:

1. Each letter maps to a unique digit
2. The leading digit of any multi-digit number is not zero
3. The arithmetic equation holds true

## Function Signature

```go
func Solve(puzzle string) (map[string]int, error)
```

- Input: A string like `"SEND + MORE == MONEY"` with `+` operators and one `==` operator
- Output: A map from single-letter strings to digit values, or an error if unsolvable

## Acceptance Criteria

1. **AC1**: `Solve` correctly solves simple puzzles (e.g., `"I + BB == ILL"` returns `{"B":9, "I":1, "L":0}`)
2. **AC2**: `Solve` returns an error for puzzles with no valid solution (e.g., `"A == B"`)
3. **AC3**: `Solve` rejects solutions where a leading digit is zero (e.g., `"ACA + DD == BD"` has no valid solution)
4. **AC4**: `Solve` handles puzzles with many addends (e.g., `"A + A + A + ... + B == BCC"`)
5. **AC5**: `Solve` handles the classic `"SEND + MORE == MONEY"` puzzle correctly
6. **AC6**: `Solve` handles the extreme 10-letter, 199-addend puzzle within a reasonable time
7. **AC7**: All tests pass via `go test ./...` in the exercise directory
8. **AC8**: Code passes `go vet ./...` with no issues

## Key Constraints

- Package name must be `alphametics`
- Only the `alphametics.go` file should be modified
- Must work with Go 1.18 (as specified in go.mod)
- No external dependencies allowed (standard library only)
