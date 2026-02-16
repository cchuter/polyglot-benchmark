# Implementation Plan: polyglot-go-connect

## File to Modify

- `go/exercises/practice/connect/connect.go`

## Approach

Use a flood-fill (DFS) algorithm with visited tracking via bit flags on a board representation.

### Data Structures

1. **Constants** using `iota` bit flags for stone colors and connected state:
   - `white`, `black`, `connectedWhite`, `connectedBlack`

2. **`board` struct**: holds height, width, and a 2D `[][]int8` field array

3. **`coord` struct**: simple x,y pair

### Algorithm

1. Parse the input strings into a `board`, mapping 'X' to `black` and 'O' to `white`
2. For black (X), start from all cells in the leftmost column (x=0); target is rightmost column (x=width-1)
3. For white (O), start from all cells in the top row (y=0); target is bottom row (y=height-1)
4. DFS from each start cell: if the cell has the correct color and hasn't been visited (connected flag not set), mark it connected and recurse into all 6 hex neighbors
5. If any DFS reaches a target cell, that player wins
6. Check black first, then white. Return "" if neither wins.

### Hex Adjacency

Six neighbors of (x,y): `(x+1,y), (x-1,y), (x,y+1), (x,y-1), (x-1,y+1), (x+1,y-1)`

### Functions

- `newBoard([]string) (board, error)` — parse input
- `(board).at(coord, colorFlags) (bool, bool)` — check color and connected status
- `(board).markConnected(coord, colorFlags)` — set connected flag
- `(board).validCoord(coord) bool` — bounds check
- `(board).neighbors(coord) []coord` — return valid hex neighbors
- `(board).startCoords(colorFlags) []coord` — starting edge cells
- `(board).isTargetCoord(coord, colorFlags) bool` — check if at winning edge
- `(board).evaluate(coord, colorFlags) bool` — recursive DFS
- `ResultOf([]string) (string, error)` — main entry point

## Ordering

Single file change. Write the complete implementation in `connect.go`.
