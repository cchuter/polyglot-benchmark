# Goal: Implement Dominoes Chain Solver (Issue #164)

## Problem Statement

Implement a `MakeChain` function in Go that determines whether a given set of dominoes can be arranged into a valid chain, and if so, returns one such arrangement.

A valid domino chain has these properties:
- Adjacent dominoes match: the right value of domino `i` equals the left value of domino `i+1`
- The chain is circular: the left value of the first domino equals the right value of the last domino
- Every input domino is used exactly once (dominoes may be flipped)

## Acceptance Criteria

1. **Type definition**: Define `type Domino [2]int` in `dominoes.go`
2. **Function signature**: `func MakeChain(input []Domino) (chain []Domino, ok bool)`
3. **Empty input**: Returns `([]Domino{}, true)` — empty input is a valid chain
4. **Single matching domino**: `{1,1}` returns `([]{1,1}, true)`
5. **Single non-matching domino**: `{1,2}` returns `(nil, false)`
6. **Multi-element chains**: Correctly finds valid chains for 3, 5, 6, and 9 element inputs
7. **Disconnected graphs**: Returns `false` for disconnected domino sets even when all vertices have even degree
8. **Backtracking**: Handles cases requiring backtracking (not just greedy matching)
9. **All tests pass**: `go test ./...` passes in the `dominoes` exercise directory
10. **Code quality**: `go vet ./...` passes with no issues

## Key Constraints

- Dominoes can be flipped (rotated) to match
- Multiple valid chains may exist; the test verifies any valid chain
- The solution must handle duplicate dominoes
- The solution must detect disconnected graphs (even-degree check alone is insufficient)
