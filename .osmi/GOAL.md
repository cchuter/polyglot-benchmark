# Goal: Implement Dominoes Chain Solver

## Problem Statement

Implement the `MakeChain` function in `go/exercises/practice/dominoes/dominoes.go` that determines whether a given set of dominoes can be arranged into a valid chain, and if so, returns one such arrangement.

A valid domino chain has these properties:
1. Adjacent dominoes must match: the right value of one domino equals the left value of the next.
2. The chain must be circular: the left value of the first domino equals the right value of the last domino.
3. Dominoes may be flipped (reversed) to make connections work.
4. All input dominoes must be used exactly once.

## Acceptance Criteria

1. `MakeChain` returns `([]Domino{}, true)` for empty input.
2. `MakeChain` returns `(input, true)` for a single domino with matching sides (e.g., `{1,1}`).
3. `MakeChain` returns `(nil, false)` for a single domino with non-matching sides (e.g., `{1,2}`).
4. `MakeChain` correctly chains 3+ dominoes, flipping as needed.
5. `MakeChain` returns `(nil, false)` for unchainable sets (disconnected graphs, odd-degree vertices).
6. `MakeChain` handles duplicate dominoes correctly.
7. `MakeChain` handles backtracking cases where greedy approaches fail.
8. All 11 test cases in `cases_test.go` pass.
9. `go vet ./...` reports no issues.

## Key Constraints

- Must define `type Domino [2]int` in the solution file.
- Function signature: `func MakeChain(input []Domino) (chain []Domino, ok bool)`.
- The solution must handle the 9-element test case efficiently enough for tests to pass.
