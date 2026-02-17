# Implementation Plan: Connect Exercise

## Proposal A — Bitmask DFS (Reference-Aligned)

**Role: Proponent**

### Approach

Follow the same architecture as the reference implementation in `.meta/example.go`: use bitmask flags to represent stone color and connected state, with a recursive DFS to find winning paths.

### Files to Modify

- `go/exercises/practice/connect/connect.go` — replace stub with full implementation

### Architecture

1. **Constants**: Use `iota` bit flags: `white`, `black`, `connectedWhite`, `connectedBlack`
2. **Types**:
   - `colorFlags` struct — pairs a color flag with its connected flag
   - `coord` struct — (x, y) position
   - `board` struct — height, width, 2D slice of `int8` fields
3. **Board Parsing** (`newBoard`): Parse each line character by character, setting `black` for 'X' and `white` for 'O'
4. **Neighbor Calculation**: 6-directional hexagonal adjacency using direction vectors `{1,0}, {-1,0}, {0,1}, {0,-1}, {-1,1}, {1,-1}`
5. **DFS Evaluation**: From each starting coordinate (left column for X, top row for O), recursively explore neighbors of the same color, marking visited cells with the `connected` flag. Return true if we reach the target edge.
6. **ResultOf**: Parse board, evaluate X first (left→right), then O (top→bottom). Return winner string or empty string.

### Rationale

- Directly mirrors the proven reference implementation
- Bitmask approach is memory-efficient and avoids separate visited sets
- Recursive DFS is simple and handles all test cases including the spiral path
- Minimal code — no extra data structures beyond what's needed
- Consistent with Exercism Go conventions

---

## Proposal B — BFS with Visited Set

**Role: Opponent**

### Approach

Use a BFS (breadth-first search) with an explicit `map[coord]bool` visited set instead of bitmask mutation. Separate the concerns more cleanly.

### Files to Modify

- `go/exercises/practice/connect/connect.go` — replace stub with full implementation

### Architecture

1. **Types**:
   - `coord` struct — (x, y) position
   - Simple `[][]byte` board representation (store raw characters)
2. **Board Parsing**: Store the board as `[][]byte` directly from input strings
3. **BFS Function**: `hasPath(board [][]byte, starts []coord, isTarget func(coord) bool, stone byte) bool`
   - Uses a queue and `map[coord]bool` visited set
   - Starts from all starting positions simultaneously (multi-source BFS)
   - Returns true if any path reaches a target cell
4. **ResultOf**: Parse board, run `hasPath` for 'X' (left→right) then 'O' (top→bottom)

### Critique of Proposal A

- The bitmask approach mutates the board state during evaluation, which means evaluating X modifies the board before evaluating O. While the reference handles this because the connected flags are separate bits, it's fragile and harder to reason about.
- Recursive DFS can stack overflow on very large boards (Go's default goroutine stack is 1MB, grows to ~1GB but deep recursion is still a concern).
- The reference implementation has more types and indirection than needed.

### Rationale for Proposal B

- BFS is iterative — no stack overflow risk
- Explicit visited set is cleaner than mutating board state
- Simpler types — just `[][]byte` and `coord`
- Multi-source BFS is elegant for checking connectivity from an entire edge
- Easier to understand and debug

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A (Bitmask DFS) | Proposal B (BFS Visited Set) |
|-----------|-------------------------|------------------------------|
| Correctness | Proven correct (mirrors reference) | Correct in principle, but needs careful implementation |
| Risk | Low — follows known-working approach | Low — BFS is well-understood |
| Simplicity | Moderate — bitmask logic adds complexity | Moderate — map allocation, but simpler types |
| Consistency | High — matches reference exactly | Moderate — diverges from reference patterns |

### Decision

**Selected: Proposal A (Bitmask DFS)** with minor simplifications.

Rationale:
- The reference implementation is proven to pass all tests. Following its architecture minimizes risk.
- Stack overflow is not a realistic concern for the test boards (max 9x9 spiral).
- The bitmask approach avoids map allocations, which matters for the benchmark test.
- Consistency with the reference solution is valuable — future maintainers can compare.

However, I'll incorporate one element from Proposal B: clearer function documentation.

### Final Implementation Plan

**File**: `go/exercises/practice/connect/connect.go`

**Step 1: Define constants and types**
```go
const (
    white = 1 << iota
    black
    connectedWhite
    connectedBlack
)

type colorFlags struct {
    color     int8
    connected int8
}

type coord struct {
    x, y int
}

type board struct {
    height, width int
    fields        [][]int8
}
```

**Step 2: Implement board parsing**
- `newBoard(lines []string) (board, error)` — validates input, creates 2D int8 grid
- 'X' → `black`, 'O' → `white`, everything else → 0

**Step 3: Implement board methods**
- `at(c coord, cf colorFlags) (bool, bool)` — check color and connected state
- `markConnected(c coord, cf colorFlags)` — set connected bit
- `validCoord(c coord) bool` — bounds check
- `neighbors(c coord) []coord` — 6 hexagonal directions
- `startCoords(cf colorFlags) []coord` — starting edge cells
- `isTargetCoord(c coord, cf colorFlags) bool` — target edge check
- `evaluate(c coord, cf colorFlags) bool` — recursive DFS

**Step 4: Implement ResultOf**
- Parse board
- Evaluate X (black) from left column → right column
- Evaluate O (white) from top row → bottom row
- Return winner or empty string

**Step 5: Verify**
- Run `go test ./...` in the connect directory
- Confirm all 10 test cases pass
- Confirm benchmark runs without error
