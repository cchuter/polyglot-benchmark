# Goal: Implement the Connect (Hex) Exercise in Go

## Problem Statement

Implement the `ResultOf` function in `go/exercises/practice/connect/connect.go` that determines the winner of a Hex/Polygon board game. The function receives a board represented as a slice of strings and must determine if player "X" (left-to-right) or player "O" (top-to-bottom) has won by connecting their respective sides with an unbroken chain of their stones, using hexagonal adjacency.

## Acceptance Criteria

1. **Function signature**: `func ResultOf(lines []string) (string, error)` in package `connect`
2. **Returns `"X"`** when player X has a connected path from the left edge (column 0) to the right edge (last column)
3. **Returns `"O"`** when player O has a connected path from the top edge (row 0) to the bottom edge (last row)
4. **Returns `""`** (empty string) when neither player has won
5. **Returns an error** for invalid board input (empty board, empty lines)
6. **All 10 test cases pass** in `cases_test.go`:
   - Empty board → no winner
   - 1x1 board with X → X wins
   - 1x1 board with O → O wins
   - Only edges → no winner
   - Illegal diagonal → no winner
   - Adjacent angles → no winner
   - X wins left to right
   - O wins top to bottom
   - X wins convoluted path
   - X wins spiral path
7. **Hexagonal adjacency**: Each cell has up to 6 neighbors: `{1,0}, {-1,0}, {0,1}, {0,-1}, {-1,1}, {1,-1}`
8. **Board parsing**: Input lines have spaces already stripped by test harness. Characters are `.` (empty), `X`, and `O`.
9. **Benchmark test** `BenchmarkResultOf` runs without error
10. `go test ./...` passes with zero failures in the connect exercise directory

## Key Constraints

- The solution must be in the file `go/exercises/practice/connect/connect.go`
- Package name must be `connect`
- Must not modify test files (`connect_test.go`, `cases_test.go`)
- Must handle boards of varying dimensions (not necessarily square)
- Go version: 1.18 (no external dependencies)
