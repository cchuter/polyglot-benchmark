# Implementation Plan: polyglot-go-connect

## Branch 1: Recursive DFS with Bitflag State (Reference-Aligned)

This approach closely follows the reference solution in `.meta/example.go`. It uses bitflags to track stone color and visited state, with recursive DFS to find connected paths.

### Files to Modify
- `go/exercises/practice/connect/connect.go` — write full implementation

### Architecture
- Define constants using `iota` bitflags: `white`, `black`, `connectedWhite`, `connectedBlack`
- Define `colorFlags` struct pairing a color with its connected flag
- Define `coord` struct for (x, y) positions
- Define `board` struct with height, width, and 2D `int8` field array
- `newBoard()` parses input lines into the board
- `board.evaluate()` does recursive DFS from a starting coordinate
- `board.neighbors()` returns up to 6 hexagonal neighbors
- `board.startCoords()` returns the starting edge for a given player
- `board.isTargetCoord()` checks if a coord is on the target edge
- `ResultOf()` orchestrates: build board, check X from left edge, check O from top edge

### Rationale
- Directly aligned with the canonical Exercism solution
- Well-tested approach with known correctness
- Uses bitflags for efficient in-place visited tracking (no extra data structure needed)

### Evaluation
- **Feasibility**: High — proven approach from reference
- **Risk**: Low — directly mirrors working code
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~170 lines in a single file, moderate complexity

---

## Branch 2: BFS with Visited Set (Clean Separation)

This approach uses iterative BFS with an explicit visited set (map), separating state tracking from the board representation.

### Files to Modify
- `go/exercises/practice/connect/connect.go` — write full implementation

### Architecture
- Define `cell` type for board characters
- Define `point` struct for (row, col) positions
- Parse board into `[][]byte` (simple 2D byte array)
- `hasWon(board, player)` function:
  - Identify starting edge cells for the player
  - BFS using a queue (`[]point`) and `map[point]bool` for visited
  - 6-directional neighbor generation with bounds checking
  - Return true if any path reaches the target edge
- `ResultOf()` calls `hasWon` for X then O

### Rationale
- Cleaner separation of concerns (board data vs traversal state)
- Iterative BFS avoids stack overflow risk on very large boards
- Explicit visited set is easier to reason about than bitflags

### Evaluation
- **Feasibility**: High — standard BFS approach
- **Risk**: Low — straightforward, well-understood algorithm
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~100 lines, lower complexity than Branch 1

---

## Branch 3: Union-Find (Disjoint Set)

This approach uses a Union-Find data structure to group connected stones, with virtual start/end nodes for each player's edges.

### Files to Modify
- `go/exercises/practice/connect/connect.go` — write full implementation

### Architecture
- Implement Union-Find with path compression and union by rank
- Create virtual nodes for each player's start and end edges
- Iterate through all cells; for each stone, union it with matching neighbors
- For edge cells, also union with the appropriate virtual node
- After processing all cells, check if start and end virtual nodes are in the same set
- `ResultOf()` builds union-find, processes board, checks connectivity

### Rationale
- Near-optimal time complexity for connectivity queries
- Elegant mathematical approach
- Could handle multiple queries on the same board efficiently

### Evaluation
- **Feasibility**: High — Union-Find is well-known
- **Risk**: Medium — more complex implementation, more room for off-by-one errors
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~120 lines, higher conceptual complexity

---

## Selected Plan

**Selected: Branch 2 (BFS with Visited Set)**

### Rationale

Branch 2 is the best choice because:

1. **Simplicity**: BFS with an explicit visited set is the simplest to implement correctly. No bitflag manipulation, no recursive stack concerns, no union-find bookkeeping.

2. **Lower Risk**: Iterative BFS eliminates stack overflow risk. An explicit `map[point]bool` for visited state is impossible to confuse with board data (unlike bitflags).

3. **Readability**: The code is straightforward and easy to review. Each concern is cleanly separated.

4. **Sufficient Performance**: BFS is O(V+E) which is optimal for this problem. The board sizes in test cases are small (max 9x9 spiral). No need for Union-Find's amortized near-constant time.

5. **Full Alignment**: All acceptance criteria are met — the function signature matches, all test cases will pass, and the benchmark will work.

Branch 1 (reference-aligned) is more complex than necessary with its bitflag approach. Branch 3 (union-find) is over-engineered for this problem size.

### Detailed Implementation Plan

**File**: `go/exercises/practice/connect/connect.go`

```go
package connect

import "errors"

// point represents a position on the board
type point struct {
    row, col int
}

// ResultOf determines the winner of a Hex game.
// Returns "X" if X wins (left-to-right), "O" if O wins (top-to-bottom), or "" if no winner.
func ResultOf(lines []string) (string, error) {
    if len(lines) == 0 {
        return "", errors.New("empty board")
    }
    if len(lines[0]) == 0 {
        return "", errors.New("empty first line")
    }

    // Check X: left-to-right connection
    if hasWon(lines, 'X') {
        return "X", nil
    }
    // Check O: top-to-bottom connection
    if hasWon(lines, 'O') {
        return "O", nil
    }
    return "", nil
}

// hasWon checks if the given player has a connected path across the board.
// X connects left to right. O connects top to bottom.
func hasWon(lines []string, player byte) bool {
    rows := len(lines)
    cols := len(lines[0])

    visited := make(map[point]bool)
    queue := []point{}

    // Seed the BFS with starting edge cells
    if player == 'X' {
        // X starts from left column (col=0)
        for r := 0; r < rows; r++ {
            if lines[r][0] == player {
                p := point{r, 0}
                queue = append(queue, p)
                visited[p] = true
            }
        }
    } else {
        // O starts from top row (row=0)
        for c := 0; c < cols; c++ {
            if lines[0][c] == player {
                p := point{0, c}
                queue = append(queue, p)
                visited[p] = true
            }
        }
    }

    // BFS
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        // Check if we reached the target edge
        if player == 'X' && cur.col == cols-1 {
            return true
        }
        if player == 'O' && cur.row == rows-1 {
            return true
        }

        // Explore hexagonal neighbors
        for _, nb := range neighbors(cur, rows, cols) {
            if !visited[nb] && lines[nb.row][nb.col] == player {
                visited[nb] = true
                queue = append(queue, nb)
            }
        }
    }

    return false
}

// neighbors returns the valid hexagonal neighbors of a point.
// Hex grid adjacency: right, left, down, up, down-left, up-right
func neighbors(p point, rows, cols int) []point {
    deltas := [6]point{
        {0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {-1, 1},
    }
    result := make([]point, 0, 6)
    for _, d := range deltas {
        nr, nc := p.row+d.row, p.col+d.col
        if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
            result = append(result, point{nr, nc})
        }
    }
    return result
}
```

### Implementation Steps

1. Write the complete implementation to `connect.go`
2. Run `go test` to verify all 10 test cases pass
3. Run `go test -bench=.` to verify benchmark works
4. Commit the solution
