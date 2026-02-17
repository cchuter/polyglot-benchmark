# Plan Review: polyglot-go-dominoes

## Reviewer: Self-review (no external agents available)

### Algorithm Correctness

**Pre-checks (Even Degree + Connectivity)**: Correct. A set of dominoes can form a valid chain (Eulerian circuit in the multigraph) if and only if:
1. Every vertex (domino value) has even degree
2. The graph is connected (among vertices with edges)

These are necessary AND sufficient conditions for Eulerian circuits. The pre-check will correctly reject all invalid inputs without needing to run the full backtracking search.

**Backtracking Search**: Correct for finding the actual chain arrangement. Since the pre-checks already guarantee a solution exists, the backtracking is guaranteed to find one. The search space is greatly pruned by only considering dominoes whose sides match the current chain end.

### Edge Cases

- **Empty input**: Handled — returns `[]Domino{}, true` ✓
- **Single matching domino `{1,1}`**: Handled — returns the domino ✓
- **Single non-matching `{1,2}`**: Handled — returns `nil, false` ✓
- **Disconnected graphs**: Caught by connectivity pre-check ✓
- **Self-loops `{4,4}`**: Degree contribution is +2 (even), connectivity handled correctly if domino `[a,a]` only connects `a` to itself ✓
- **Duplicate dominoes**: Backtracking uses index-based tracking, so duplicates are handled correctly ✓

### Potential Issues

1. **Union-Find for self-loops**: A domino `[a,a]` unions `a` with itself — this is a no-op. If `{4,4}` is the ONLY domino involving value 4, and other dominoes don't touch 4, the connectivity check correctly identifies it as disconnected. ✓

2. **Starting domino choice**: The plan says "try each domino as the first domino." For efficiency, we only need to try the first domino in both orientations since if a solution exists, it can start with any domino. Actually, since we're guaranteed a solution exists (pre-checks passed), and any Eulerian circuit visits all edges, we can fix the starting domino to `input[0]` and just try both orientations. This reduces the outer loop from O(n) to O(1). **Recommendation: Only try input[0] as the starting domino.**

3. **Chain end matching**: At depth == len(input), must verify `chain[0][0] == chain[len-1][1]`. This is essential since the pre-check guarantees a circuit exists but the particular arrangement found by backtracking might not close the loop at every attempt.

### Simplifications

- Since pre-checks guarantee a solution, the backtracking will always succeed. No need to return false from the search phase.
- Fix starting domino to input[0]; only try both orientations of it.

### Verdict

**Plan is sound.** Proceed with implementation. Apply the optimization of fixing the starting domino to `input[0]`.
