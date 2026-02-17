# Changes: Connect Exercise Implementation

## File Modified
- `go/exercises/practice/connect/connect.go`

## What Changed
Implemented the `ResultOf` function for the Connect (Hex) board game exercise using a bitmask DFS approach.

### Implementation Details
- **Constants**: Bit flags (`white`, `black`, `connectedWhite`, `connectedBlack`) using `iota`
- **Types**: `colorFlags`, `coord`, and `board` structs
- **Board parsing**: `newBoard` converts input lines to a 2D `int8` grid ('X'→black, 'O'→white)
- **Hex neighbors**: 6-directional adjacency: {1,0}, {-1,0}, {0,1}, {0,-1}, {-1,1}, {1,-1}
- **DFS evaluation**: Recursive search from starting edges, marking visited cells with connected bitmask
- **Win detection**: X wins left→right (start x=0, target x=width-1); O wins top→bottom (start y=0, target y=height-1)

## Test Results
All 10 test cases pass, including the spiral path case.

## Commit
`28145a8` — "Implement ResultOf for Connect (Hex) exercise"
