# Goal: Implement the Connect (Hex) Exercise in Go

## Problem Statement

Implement a function `ResultOf` in `go/exercises/practice/connect/connect.go` that computes the winner of a game of Hex / Polygon / CON-TAC-TIX.

The game is played on a parallelogram board with hexagonal fields. Two players (`O` and `X`) place stones on the board. Player `O` wins by connecting the top edge to the bottom edge. Player `X` wins by connecting the left edge to the right edge. The function must determine if either player has won, or if there is no winner yet.

The board is represented as a slice of strings where:
- `O` represents Player O's stones
- `X` represents Player X's stones
- `.` represents empty hexagonal fields
- Spaces are stripped before input (handled by test helper)

The hexagonal adjacency means each cell has up to 6 neighbors (not 4 as in a square grid). The six neighbor directions are: right, left, down, up, down-left, up-right (in terms of coordinate offsets: `{1,0}, {-1,0}, {0,1}, {0,-1}, {-1,1}, {1,-1}`).

## Acceptance Criteria

1. The function `ResultOf(lines []string) (string, error)` must be exported from package `connect`
2. Returns `"X"` when Player X has a connected path from left to right
3. Returns `"O"` when Player O has a connected path from top to bottom
4. Returns `""` (empty string) when no player has won
5. Returns an error for invalid board input (empty board, empty lines)
6. All 10 test cases in `cases_test.go` must pass:
   - Empty board → no winner
   - 1x1 board with X → X wins
   - 1x1 board with O → O wins
   - Only edges → no winner
   - Illegal diagonal → no winner
   - Adjacent angles → no winner
   - X crossing left to right → X wins
   - O crossing top to bottom → O wins
   - X convoluted path → X wins
   - X spiral path → X wins
7. `go test` must pass with zero failures
8. The benchmark `BenchmarkResultOf` must run without errors

## Key Constraints

- Must be in package `connect`
- Must use Go 1.18 (as specified in go.mod)
- Must implement `ResultOf` as the sole exported function (matching test expectations)
- Board input has spaces already stripped (the test helper `prepare()` does this)
- Must handle non-square boards (different width and height)
- Must correctly model hexagonal adjacency (6 neighbors, not 4)
