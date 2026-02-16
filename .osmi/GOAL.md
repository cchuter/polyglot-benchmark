# Goal: polyglot-go-connect

## Problem Statement

Implement the `ResultOf` function in `go/exercises/practice/connect/connect.go` for the Hex/Polygon/CON-TAC-TIX board game. Given a board representation as a slice of strings, determine which player (if any) has won by connecting their stones to opposite sides.

- Player "O" wins by connecting top to bottom
- Player "X" wins by connecting left to right
- The board is a parallelogram with hexagonal fields
- Boards may not be "fair" (different dimensions, unequal piece counts)

## Acceptance Criteria

1. `ResultOf` accepts `[]string` (board rows with spaces stripped) and returns `(string, error)`
2. Returns `"X"` when X has a connected path from left edge to right edge
3. Returns `"O"` when O has a connected path from top edge to bottom edge
4. Returns `""` when neither player has won
5. All 10 test cases in `cases_test.go` pass, including:
   - Empty board (no winner)
   - 1x1 boards (single stone wins)
   - Edge-only placement (no winner)
   - Illegal diagonal (no winner)
   - Adjacent angles (no winner)
   - Left-to-right X win
   - Top-to-bottom O win
   - Convoluted path X win
   - Spiral path X win
6. `go test ./...` passes with no errors
7. `go vet ./...` passes with no warnings

## Key Constraints

- Solution goes in `connect.go` in package `connect`
- Must export function `ResultOf([]string) (string, error)`
- Hexagonal adjacency: each cell has 6 neighbors: (x±1,y), (x,y±1), (x-1,y+1), (x+1,y-1)
- Input strings have spaces already stripped by the test harness's `prepare` function
