# Implementation Plan: Dominoes Chain Solver

## File to Modify

- `go/exercises/practice/dominoes/dominoes.go` — the only file that needs changes

## Algorithm: Eulerian Circuit via Backtracking DFS

The domino chaining problem is equivalent to finding an Eulerian circuit in a multigraph where:
- Each unique number is a vertex
- Each domino `[a, b]` is an edge between vertices `a` and `b`

### Approach

Use a two-phase approach:

**Phase 1: Validity checks (fast rejection)**
1. Empty input → return `([]Domino{}, true)` immediately
2. Check all vertices have even degree (necessary condition for Eulerian circuit)
3. Check graph connectivity using Union-Find or BFS/DFS (necessary because disconnected graphs with even degrees can't form a single chain)

**Phase 2: Construct the chain via backtracking DFS**
- Use recursive backtracking with a `used` boolean slice to track which dominoes have been placed
- Start with the first domino (try both orientations)
- At each step, try to extend the chain by finding an unused domino whose left value matches the current chain's right end
- When all dominoes are placed, verify the chain is circular (first left == last right)

### Detailed Design

```go
type Domino [2]int

func MakeChain(input []Domino) (chain []Domino, ok bool) {
    // Handle empty case
    // Check necessary conditions (even degree, connectivity)
    // Backtrack to find a valid chain
}
```

#### Connectivity Check
Use a simple Union-Find (disjoint set) to verify all domino values are in the same connected component. This is simpler than BFS for this use case.

#### Backtracking
- `used []bool` tracks which input dominoes are already placed
- `chain []Domino` accumulates the current chain
- Try each unused domino that can connect (matching left side to current right end), in both orientations
- Base case: all dominoes placed, check circularity

### Why Backtracking Over Hierholzer's

Hierholzer's algorithm finds Eulerian circuits in O(E) time, but it works on the multigraph of edges, not on the original domino sequence. Converting back to a domino sequence with correct orientations adds complexity. For the small input sizes in the test cases (max 9 dominoes), backtracking is simple, correct, and fast enough.

## Implementation Order

1. Define the `Domino` type
2. Implement `MakeChain` with edge cases (empty, single)
3. Implement connectivity check via Union-Find
4. Implement even-degree check
5. Implement backtracking DFS to construct the chain
6. Run tests and verify
