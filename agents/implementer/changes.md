# Changes: Connect (Hex) Exercise

## Summary
Implemented the `ResultOf` function in `go/exercises/practice/connect/connect.go` using a BFS-based approach.

## Details
- **File modified**: `go/exercises/practice/connect/connect.go`
- **Approach**: Breadth-first search (BFS) to determine if a player has a connected path across the hex board
- **Player X** wins by connecting left edge to right edge
- **Player O** wins by connecting top edge to bottom edge
- Hex adjacency uses 6 neighbors: `(0,1), (0,-1), (1,0), (-1,0), (1,-1), (-1,1)`

## Test Results
All 10 test cases pass:
- Empty board (no winner)
- 1x1 board wins (X and O)
- Edge-only configurations (no winner)
- Illegal diagonal (no winner)
- Adjacent angles (no winner)
- Left-to-right crossing (X wins)
- Top-to-bottom crossing (O wins)
- Convoluted path (X wins)
- Spiral path (X wins)

## Commit
`6cb019d` - Implement connect (Hex) exercise using BFS
