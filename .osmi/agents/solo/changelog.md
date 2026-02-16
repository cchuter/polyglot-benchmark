# Change Log: polyglot-go-connect

## Changes Made

### `go/exercises/practice/connect/connect.go`
- Implemented `ResultOf([]string) (string, error)` function
- Added board parsing (`newBoard`) that converts string input to a 2D int8 grid
- Used bit flags (`white`, `black`, `connectedWhite`, `connectedBlack`) for efficient color and visited tracking
- Implemented DFS flood-fill (`evaluate`) with hex adjacency (6 neighbors per cell)
- Player X: starts from left edge (x=0), wins at right edge (x=width-1)
- Player O: starts from top edge (y=0), wins at bottom edge (y=height-1)

## Test Results
- All 10 test cases pass
- `go vet` clean
