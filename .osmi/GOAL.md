# Goal: polyglot-go-dominoes (Issue #335)

## Problem Statement

Implement the `MakeChain` function in Go that determines whether a given set of dominoes can be arranged into a valid chain, and if so, returns one such arrangement.

A valid domino chain satisfies:
1. Adjacent dominoes match: the right value of domino `i` equals the left value of domino `i+1`
2. The chain forms a loop: the left value of the first domino equals the right value of the last domino
3. Dominoes may be flipped (reversed) to achieve a valid chain
4. All input dominoes must be used exactly once

## Acceptance Criteria

1. **Type definition**: `type Domino [2]int` is defined in the `dominoes` package
2. **Function signature**: `func MakeChain(input []Domino) (chain []Domino, ok bool)` is implemented
3. **Empty input**: Returns `([]Domino{}, true)` for empty input
4. **Single matching domino**: Returns the domino and `true` if both sides are equal (e.g., `{1,1}`)
5. **Single non-matching domino**: Returns `(nil, false)` if sides differ (e.g., `{1,2}`)
6. **Valid chains**: For chainable inputs, returns a valid chain where adjacent dominoes match and ends match
7. **Invalid chains**: Returns `(nil, false)` for unchainable inputs
8. **Disconnected graphs**: Correctly identifies disconnected domino sets as unchainable
9. **Backtracking**: Handles cases requiring backtracking (not just greedy matching)
10. **All 11 test cases pass**: `go test ./...` passes in the dominoes exercise directory
11. **Code vets cleanly**: `go vet ./...` passes

## Key Constraints

- Solution must be in `go/exercises/practice/dominoes/dominoes.go`
- Must use package name `dominoes`
- The `Domino` type must be `[2]int`
- Dominoes can be reversed/rotated to form the chain
- Multiple valid chains may exist; only one needs to be returned
- Duplicate dominoes are possible in the input
