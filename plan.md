# Implementation Plan: Dominoes Chain Solver

## Branch 1: Graph Theory + Backtracking (Simple & Correct)

**Approach**: Use Euler circuit theory for fast rejection, then backtracking DFS to find the actual chain.

**Rationale**: The domino chain problem is equivalent to finding an Eulerian circuit in a multigraph where each domino value is a vertex and each domino is an edge. An Eulerian circuit exists if and only if (1) all vertices have even degree and (2) the graph is connected (considering only vertices with edges).

**Steps**:
1. Define `type Domino [2]int` in `dominoes.go`
2. Handle base cases (empty input, single domino)
3. Build adjacency degree counts and check even-degree condition
4. Check graph connectivity using DFS/BFS on vertices
5. If both conditions met, use recursive backtracking to construct the actual chain
6. Return the chain or nil/false

**Files**: Only `go/exercises/practice/dominoes/dominoes.go`

**Evaluation**:
- Feasibility: High — straightforward graph theory approach
- Risk: Low — well-understood algorithm, two-phase approach (fast reject then construct)
- Alignment: Fully satisfies all criteria; handles disconnected cases, backtracking cases
- Complexity: ~80-100 lines of Go, single file

## Branch 2: Hierholzer's Algorithm (Extensible & Elegant)

**Approach**: Use Hierholzer's algorithm to directly find an Eulerian circuit, which naturally produces the domino chain.

**Rationale**: Hierholzer's algorithm is the standard O(E) algorithm for finding Eulerian circuits. It builds the circuit by following edges until stuck, then splicing in sub-circuits. More extensible if future tests add larger inputs.

**Steps**:
1. Define `type Domino [2]int`
2. Handle base cases
3. Check even-degree and connectivity (required preconditions for Hierholzer's)
4. Build adjacency lists (multimap of edges indexed by vertex)
5. Run Hierholzer's to produce a vertex sequence
6. Convert vertex sequence back to domino chain, flipping as needed
7. Return result

**Files**: Only `go/exercises/practice/dominoes/dominoes.go`

**Evaluation**:
- Feasibility: High — well-known algorithm
- Risk: Medium — more complex implementation; edge tracking with duplicates requires care; converting vertex path back to oriented dominoes adds complexity
- Alignment: Fully satisfies all criteria
- Complexity: ~100-130 lines; more data structures needed (adjacency lists, edge tracking)

## Branch 3: Permutation-Based (Match Reference Solution)

**Approach**: Generate all permutations of dominoes and try to arrange each into a valid chain, closely following the `.meta/example.go` reference.

**Rationale**: This is the approach used by the reference solution. It's conceptually simple — try all orderings and for each ordering, try to match adjacent dominoes by flipping.

**Steps**:
1. Copy the approach from `example.go` — define `Domino`, permutation generator, and chain arranger
2. Implement `dominoPermutations` function
3. Implement `arrangeChain` that tries to build chain from a given ordering
4. Iterate permutations until a valid chain is found

**Files**: Only `go/exercises/practice/dominoes/dominoes.go`

**Evaluation**:
- Feasibility: High — proven by reference solution
- Risk: Low — known to pass all tests
- Alignment: Fully satisfies all criteria
- Complexity: ~130-150 lines; the permutation generator is verbose
- **Major downside**: O(n! * n) complexity. For 9 elements, that's 362,880 permutations. Works for the test suite but is inherently inefficient. The reference solution exists as a correctness baseline, not as the ideal approach.

---

## Selected Plan

**Branch 1: Graph Theory + Backtracking** is the best choice.

**Rationale**:
- **Simpler than Branch 2**: Hierholzer's algorithm requires maintaining adjacency lists with edge deletion and circuit splicing, plus converting a vertex path back to oriented dominoes. The backtracking approach is more straightforward.
- **Far more efficient than Branch 3**: Permutation-based is O(n!), while graph checks + backtracking prunes the search space dramatically. The even-degree and connectivity checks reject impossible inputs in O(n) time.
- **Lower risk than Branch 2**: The backtracking DFS is simple to implement correctly. Hierholzer's has more subtle edge cases with duplicate edges and self-loops.
- **Fully meets acceptance criteria**: Handles all test cases including disconnected graphs, backtracking requirements, and duplicate dominoes.

### Detailed Implementation Plan

**File**: `go/exercises/practice/dominoes/dominoes.go`

```go
package dominoes

type Domino [2]int

func MakeChain(input []Domino) (chain []Domino, ok bool) {
    // Base cases
    // - empty: return [], true
    // - single: return input, true if sides match; nil, false otherwise

    // Graph validation (Euler circuit necessary conditions):
    // 1. All vertices must have even degree
    // 2. All edges must be in a single connected component
    // If either fails, return nil, false

    // Backtracking DFS to build the chain:
    // - Start with any domino
    // - At each step, try to extend the chain with an unused domino that matches
    // - If chain uses all dominoes and first/last ends match, return it
    // - Otherwise backtrack
}
```

**Helper functions**:
1. `canChain(input []Domino) bool` — checks even-degree + connectivity
2. `solve(chain []Domino, used []bool, input []Domino, count int) bool` — recursive backtracking

**Algorithm details**:
- Build degree map: for each domino [a,b], increment degree[a] and degree[b]
- Check all degrees are even
- Check connectivity via union-find or DFS on vertices that appear in dominoes
- For backtracking: try each unused domino, flip if needed to match the chain's current end
- When all dominoes used, check if chain[0][0] == chain[last][1]

**Estimated size**: ~80-100 lines of Go code.
