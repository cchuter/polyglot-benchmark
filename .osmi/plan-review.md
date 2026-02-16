# Plan Review: Dominoes Chain Solver

**Reviewer**: codex
**Date**: 2026-02-16
**Verdict**: APPROVE with minor suggestions

---

## 1. Correctness

The proposed approach (Union-Find connectivity check + even-degree check + backtracking DFS) is **correct** and will handle all 12 test cases.

### Test-by-test analysis

| Test Case | Handled By |
|---|---|
| Empty input | Early return `([]Domino{}, true)` -- explicitly addressed in Phase 1 |
| Singleton `{1,1}` | Backtracking places it, circularity check passes (`1 == 1`) |
| Singleton `{1,2}` | Even-degree check rejects it (vertex 1 has degree 1, vertex 2 has degree 1) |
| Three elements | Backtracking finds valid ordering with flips |
| Can reverse dominoes | Backtracking tries both orientations per domino |
| Can't be chained `{1,2},{4,1},{2,3}` | Even-degree check rejects (vertex 4 has odd degree) OR connectivity check rejects |
| Disconnected - simple `{1,1},{2,2}` | Connectivity check rejects (two components). Even-degree alone would incorrectly accept. |
| Disconnected - double loop | Connectivity check rejects (two components) |
| Disconnected - single isolated | Connectivity check rejects ({4} is isolated from {1,2,3}) |
| Need backtrack | Backtracking handles this. The plan explicitly acknowledges greedy is insufficient. |
| Separate loops | Backtracking finds valid interleaving of the loops |
| Nine elements | Backtracking with pruning via Phase 1 checks |

**All test cases are covered by the plan.**

## 2. Edge Cases

The plan handles the critical edge cases well:

- **Empty input**: Explicitly addressed.
- **Single domino (both matching and non-matching)**: Handled by the even-degree check (for non-matching) and the backtracking + circularity check (for matching).
- **Disconnected graphs with even degrees**: Explicitly addressed via connectivity check. This is the classic trap mentioned in the test comments, and the plan correctly identifies it.
- **Duplicate dominoes**: Test case "need backtrack" has two `{2,4}` dominoes. The `used []bool` slice indexed by position in the input array handles duplicates correctly since it tracks indices, not domino values.
- **Self-loops like `{1,1}`**: These contribute degree 2 to vertex 1 in the multigraph. The Union-Find union of a vertex with itself is a no-op, but that is fine since the self-loop vertex only needs to be in the same component as other vertices. The plan does not explicitly call this out, but the approach handles it correctly as long as the Union-Find unions both endpoints of every domino (including self-loops, where `union(a, a)` is harmless and the vertex is still registered).

### One subtlety worth noting

The connectivity check must ensure that **all vertices appearing in the input** are in a single connected component. A self-loop `{v, v}` registers vertex `v` but `union(v, v)` does not connect it to anything else. In the "disconnected - single isolated" test case (`{1,2},{2,3},{3,1},{4,4}`), vertex 4 appears only in a self-loop. The Union-Find must detect that vertex 4 is in a separate component from vertices 1, 2, 3. This works naturally as long as:
1. You collect all distinct vertices that appear in any domino.
2. You check that all of those vertices share the same root in the Union-Find.

The plan implicitly handles this, but the implementer should be aware of it.

## 3. Algorithm Choice

### Is backtracking appropriate?

**Yes.** The plan's justification is sound:

- Maximum input size in the test suite is 9 dominoes.
- Backtracking with the pruning from Phase 1 (even-degree + connectivity) means the backtracker only runs on inputs that are known to have a valid Eulerian circuit. For such inputs, the search tree is well-constrained.
- Worst-case complexity is O(n! * 2^n) for the backtracking alone, but with n <= 9 and the guarantee that a solution exists (after Phase 1 passes), this is fast in practice.
- Hierholzer's algorithm is asymptotically better (O(E)) but introduces complexity in mapping the circuit back to oriented dominoes with correct sequencing. The plan correctly identifies this tradeoff.

### Alternative consideration

For production code with large inputs, Hierholzer's would be preferred. For this exercise's constraints, backtracking is the pragmatic choice. No concerns here.

## 4. Potential Issues with the Approach

### 4a. Union-Find implementation detail

The plan says "Use a simple Union-Find (disjoint set)" but does not specify whether path compression or union-by-rank will be used. For n <= 9, this is irrelevant for performance, but path compression is easy to add and is good practice. Not a blocker.

### 4b. Orientation handling in backtracking

The plan says "try both orientations" for each domino. The implementation must be careful to:
- When placing domino `{a, b}` at a position where the required left value is `b`, place it as `{b, a}` in the chain.
- The `used` slice must track that the original domino at index `i` has been used, regardless of orientation.

This is straightforward but is the most common source of bugs in this type of implementation.

### 4c. Starting domino selection

The plan says "Start with the first domino (try both orientations)." This is correct. Since any Eulerian circuit must include all edges, starting from any domino is valid. However, the implementer should note:
- For the first domino, you must try both orientations `{a, b}` and `{b, a}` since there is no preceding domino to constrain the left value.
- For subsequent dominoes, the left value is constrained by the previous domino's right value, so "both orientations" means checking if `a` or `b` matches and placing accordingly.

### 4d. Return value for invalid cases

The GOAL specifies returning `(nil, false)` for invalid cases. The test harness only checks `ok != tc.valid` and only calls `verifyChain` when `ok` is true, so returning `nil` for the chain on failure is fine. However, the plan should ensure it does not accidentally return `([]Domino{}, false)` vs `(nil, false)` -- in Go these are different. Looking at the test, it does not check the chain value when `ok` is false, so either `nil` or empty slice works. No issue here.

### 4e. Phase 1 as an optimization, not just correctness

The Phase 1 checks (even-degree + connectivity) serve dual purposes:
1. **Correctness**: They correctly reject impossible inputs without running the expensive backtracker.
2. **Optimization**: They guarantee that when the backtracker runs, a solution exists, avoiding fruitless exhaustive search on impossible inputs.

This is a good design. Without Phase 1, the backtracker would still return `false` for invalid inputs, but it would be much slower on cases like "disconnected - double loop" where it would exhaust all possibilities before concluding no chain exists.

## 5. Implementation Order

The proposed implementation order is logical:
1. Type definition
2. Edge cases
3. Connectivity (Union-Find)
4. Even-degree check
5. Backtracking DFS
6. Testing

One minor suggestion: implement the even-degree check (step 4) before the connectivity check (step 3), since the even-degree check is simpler and can be tested independently first. But this is a stylistic preference, not a correctness concern.

## Summary

The plan is well-reasoned, correct, and appropriately scoped for the problem. The two-phase approach (fast rejection via graph properties, then backtracking construction) is a clean design. All 12 test cases are accounted for. The only areas requiring care during implementation are:

1. Ensuring the connectivity check correctly handles self-loops (vertices that only appear in self-loop dominoes).
2. Correctly managing domino orientation during backtracking.
3. Properly tracking used dominoes by index (not by value) to handle duplicate dominoes.

None of these are flaws in the plan -- they are implementation details the developer should keep in mind. **The plan is approved for implementation.**
