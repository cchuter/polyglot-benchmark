# Implementation Plan: Dominoes Chain Solver

## File to Modify

- `go/exercises/practice/dominoes/dominoes.go` — the only file that needs changes

## Algorithm Design

The domino chaining problem is equivalent to finding an Eulerian circuit in a multigraph where:
- Each unique domino value is a node
- Each domino is an edge between its two values

A valid chain exists if and only if:
1. The graph is connected (considering only nodes with edges)
2. Every node has even degree (Eulerian circuit condition)

However, we also need to **produce** the actual chain, not just determine if one exists. We'll use a two-phase approach:

### Phase 1: Feasibility Check (Graph Theory)
1. Build an adjacency representation tracking degree of each node
2. Check all nodes have even degree
3. Check the graph is connected using Union-Find or BFS/DFS

### Phase 2: Chain Construction (Backtracking with Hierholzer's)
If feasible, construct the chain using a backtracking DFS approach:
1. Pick the first domino's left value as the starting node
2. At each step, try placing an unused domino that connects to the current chain end
3. Dominoes can be placed in either orientation (flipped)
4. Backtrack if stuck before all dominoes are placed

### Detailed Algorithm

```
func MakeChain(input []Domino) ([]Domino, bool):
    if len(input) == 0: return [], true
    if len(input) == 1:
        if input[0][0] == input[0][1]: return input, true
        else: return nil, false

    // Check feasibility
    if !isEulerian(input): return nil, false

    // Build chain via backtracking DFS
    chain = backtrack(input)
    if chain != nil: return chain, true
    return nil, false
```

**Feasibility check (`isEulerian`)**:
- Count degree of each node
- All degrees must be even
- Graph must be connected (use DFS/BFS on adjacency list)

**Chain construction (`backtrack`)**:
- Use recursive DFS with a `used` boolean slice
- Start from `input[0][0]`
- At each level, try each unused domino in both orientations
- If the domino's first value matches current end, place it and recurse
- If all dominoes placed and chain end matches chain start, success

### Edge Cases
- Empty input → return `([], true)`
- Single domino with matching sides → return `([domino], true)`
- Single domino with different sides → return `(nil, false)`
- Disconnected graph → detected by connectivity check, return `(nil, false)`

## Approach Order

1. Define `type Domino [2]int`
2. Implement `isConnected` using DFS on adjacency list
3. Implement `hasEvenDegrees` checking all vertex degrees
4. Implement backtracking chain builder
5. Implement `MakeChain` combining feasibility + construction
6. Run tests

## Rationale

- The feasibility check (Euler conditions) quickly rejects impossible inputs without expensive backtracking
- Backtracking DFS is simple and correct for construction; the input sizes are small (max ~9 dominoes in tests) so performance is not a concern
- This avoids the complexity of Hierholzer's algorithm while still being efficient enough for the test cases
