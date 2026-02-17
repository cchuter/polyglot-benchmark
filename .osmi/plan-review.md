# Plan Review: Graph Theory + Backtracking

## Self-Review (no external codex agent available)

### Correctness Assessment

The approach is sound. The domino chain problem is indeed equivalent to finding an Eulerian circuit in a multigraph. The two necessary and sufficient conditions for an Eulerian circuit are:

1. **Even degree**: Every vertex must have even degree. ✓
2. **Connectivity**: All vertices with edges must be in one connected component. ✓

These checks provide O(n) fast rejection for impossible inputs, which handles test cases like "disconnected - simple", "disconnected - double loop", and "disconnected - single isolated".

### Edge Cases Reviewed

1. **Empty input**: Handled explicitly as base case. ✓
2. **Single domino {1,1}**: Sides match → valid. ✓
3. **Single domino {1,2}**: Sides don't match → invalid. ✓
4. **Self-loops {1,1}**: Degree of vertex 1 increases by 2 (both sides), remains even. ✓
5. **Duplicate dominoes**: Backtracking with `used[]` flags handles this correctly. ✓
6. **Need backtrack case**: The DFS backtracking naturally handles this — if a greedy choice leads to a dead end, it backtracks. ✓
7. **Disconnected components with even degrees**: Connectivity check rejects these. ✓

### Potential Issues Identified

1. **Self-loop degree counting**: A domino [a,a] contributes 2 to degree[a] (one per side). This is correct for Euler circuit theory where a self-loop contributes 2 to the vertex degree.

2. **Connectivity check with self-loops**: A domino [a,a] only involves vertex a. The connectivity check must still work — vertex a connects to itself. If we use union-find, a self-loop union(a,a) is fine. If using DFS on an adjacency list, a self-loop adds a to its own adjacency list.

3. **Backtracking efficiency**: With 9 dominoes, worst case is manageable. The Euler checks eliminate impossible inputs before backtracking starts.

### Suggestions

- Use union-find for connectivity check — simpler than building full adjacency lists just for connectivity.
- In backtracking, start with the first domino and try both orientations.
- The chain verification in the test already handles flipped dominoes via canonical form comparison, so we just need to produce a valid chain.

### Verdict

Plan is **approved**. The approach is correct, efficient for the given test sizes, and handles all identified edge cases.
