# Implementation Plan: polyglot-go-dominoes

## Proposal A: Graph-Theoretic Approach (Euler Circuit)

**Role: Proponent**

### Approach

Model the dominoes as edges in a multigraph where each domino value is a vertex. A valid domino chain exists if and only if:
1. All vertices have even degree (Eulerian circuit condition)
2. The graph is connected (considering only vertices with non-zero degree)

If these conditions hold, construct the chain using Hierholzer's algorithm to find an Eulerian circuit.

### Files to Modify

- `go/exercises/practice/dominoes/dominoes.go` — sole file to implement

### Implementation Details

1. Define `type Domino [2]int`
2. Implement `MakeChain`:
   - Empty input: return `[]Domino{}, true`
   - Single domino: return based on whether sides match
   - Build adjacency list (multigraph: vertex → list of (neighbor, domino-index) edges)
   - Check all vertices have even degree; if not, return `nil, false`
   - Check connectivity via DFS/BFS on vertices with edges; if disconnected, return `nil, false`
   - Run Hierholzer's algorithm to find Eulerian circuit
   - Convert the vertex path back to a sequence of oriented dominoes

### Rationale

- **Correctness**: Euler's theorem guarantees that a multigraph has an Eulerian circuit iff all vertices have even degree and the graph is connected. This maps exactly to the domino chain problem.
- **Efficiency**: O(V + E) time where E = number of dominoes. No factorial blowup.
- **Elegance**: Clean separation between validation (degree + connectivity) and construction (Hierholzer).

### Weaknesses (acknowledged)

- More complex to implement correctly (edge tracking, path splicing)
- Hierholzer's requires careful bookkeeping of used edges
- Converting vertex path to oriented domino chain needs care

---

## Proposal B: Backtracking Search with Pruning

**Role: Opponent**

### Approach

Use depth-first backtracking to try placing dominoes one at a time, flipping each domino as needed. This is essentially a brute-force search with early termination.

### Files to Modify

- `go/exercises/practice/dominoes/dominoes.go` — sole file to implement

### Implementation Details

1. Define `type Domino [2]int`
2. Implement `MakeChain`:
   - Empty input: return `[]Domino{}, true`
   - Single domino: return based on whether sides match
   - Pick the first domino (try both orientations), then recursively try each remaining unused domino that matches the current chain end
   - Track used dominoes with a boolean slice
   - At the end, verify the first and last values match
3. Optimization: Pre-check that all vertices have even degree before searching (early rejection of impossible cases)

### Rationale

- **Simplicity**: Straightforward recursive DFS; easy to understand and debug
- **Proven**: The `.meta/example.go` uses a similar approach (permutation-based), showing this style is accepted
- **Correctness**: Exhaustive search guarantees finding a solution if one exists

### Critique of Proposal A

- Hierholzer's algorithm is tricky to implement correctly, especially with multigraph edges and the need to track which specific edge (domino) has been used
- The conversion from vertex circuit to oriented domino sequence adds complexity
- For the test cases (max 9 dominoes), the efficiency advantage is irrelevant
- More bug-prone: off-by-one errors in edge tracking, incorrect path splicing

### Advantages over A

- Far simpler to implement and verify
- Backtracking with constraint propagation (matching next edge to current end) prunes the search space aggressively — much better than brute-force permutations
- The `.meta/example.go` approach generates ALL permutations; this approach is already much better since it only explores matching branches

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A (Euler/Hierholzer) | Proposal B (Backtracking) |
|-----------|-------------------------------|---------------------------|
| **Correctness** | Mathematically proven | Exhaustive search guarantees correctness |
| **Risk** | Higher — complex algorithm, edge tracking bugs | Lower — simple recursive structure |
| **Simplicity** | Moderate — needs adjacency list, connectivity check, Hierholzer, path conversion | High — single recursive function with used-tracking |
| **Consistency** | Different style from example.go | Similar style to example.go |
| **Performance** | O(V+E) | Exponential worst case, but irrelevant for test sizes |

### Decision: Proposal B (Backtracking) — with key improvements

Proposal B wins on simplicity, lower risk, and consistency with the existing codebase. The test cases have at most 9 dominoes, making performance irrelevant. The backtracking approach with matching-based pruning is far more efficient than the example's full-permutation approach while being much simpler to implement correctly than Hierholzer's algorithm.

**Key improvement over naive backtracking**: Add the even-degree and connectivity pre-check from Proposal A as a fast rejection filter. This catches impossible cases (like disconnected graphs) in O(V+E) without needing to run the full search.

### Detailed Implementation Plan

**File**: `go/exercises/practice/dominoes/dominoes.go`

```go
package dominoes

type Domino [2]int

func MakeChain(input []Domino) (chain []Domino, ok bool) {
    // 1. Handle edge cases
    //    - Empty: return []Domino{}, true
    //    - Single: return input[0] if sides match, else nil, false

    // 2. Pre-check: all vertices must have even degree
    //    - Count degree of each value across all dominoes
    //    - If any value has odd degree, return nil, false

    // 3. Pre-check: graph must be connected
    //    - Build adjacency info and use DFS/union-find to verify connectivity
    //    - If disconnected, return nil, false

    // 4. Backtracking search
    //    - Start with first domino (try both orientations sets the "start" value)
    //    - Recursively try each unused domino that has a side matching the current chain end
    //    - On finding complete chain, verify first[0] == last[1]
    //    - Return first valid chain found
}
```

#### Step-by-step:

1. **Type definition**: `type Domino [2]int`

2. **MakeChain function**:
   - Handle len 0 → return `[]Domino{}, true`
   - Handle len 1 → check `input[0][0] == input[0][1]`
   - Call `canChain(input)` for pre-validation (degree + connectivity)
   - If pre-validation fails, return `nil, false`
   - Call `solve(input)` for backtracking search
   - Return result

3. **canChain helper** (pre-validation):
   - Build degree map: for each domino `[a, b]`, increment degree[a] and degree[b]
   - Check all degrees are even
   - Check connectivity: use union-find or simple DFS on the values present
   - For connectivity, use union-find: union(a, b) for each domino, then check all values share the same root

4. **solve helper** (backtracking):
   - Allocate `used []bool` of length `len(input)`
   - Allocate `chain []Domino` of length `len(input)`
   - For each domino index `i`, try placing it first (both orientations)
   - Mark used[i] = true, set chain[0] = domino (possibly flipped)
   - Call recursive `search(input, used, chain, 1)` where 1 is the next position to fill
   - In `search`: iterate over unused dominoes, if either orientation matches `chain[pos-1][1]`, place it, recurse
   - At depth == len(input), check `chain[0][0] == chain[len-1][1]`

5. **Union-Find** (for connectivity check):
   - Simple array-based union-find with path compression
   - Map domino values to indices
