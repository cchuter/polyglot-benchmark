# Goal: Implement Dominoes Chain Solver (Issue #121)

## Problem Statement

Implement a `MakeChain` function in Go that takes a set of dominoes and determines whether they can be arranged into a valid chain. In a valid chain:
- Adjacent dominoes must match: the right value of one domino equals the left value of the next
- The chain must form a loop: the left value of the first domino equals the right value of the last domino
- Dominoes may be flipped (reversed) to make matches work
- All input dominoes must be used exactly once

## Acceptance Criteria

1. **Function signature matches spec**: `func MakeChain(input []Domino) (chain []Domino, ok bool)` with `type Domino [2]int`
2. **Empty input**: Returns `([], true)` — empty input is a valid chain
3. **Single matching domino**: `{1,1}` returns valid chain
4. **Single non-matching domino**: `{1,2}` returns `ok=false`
5. **Three-element chains**: Correctly chains `{1,2},{3,1},{2,3}` and similar
6. **Domino reversal**: Handles cases where dominoes need to be flipped, e.g. `{1,2},{1,3},{2,3}`
7. **Invalid chains detected**: Returns `ok=false` for unchainable sets like `{1,2},{4,1},{2,3}`
8. **Disconnected graphs detected**: Returns `ok=false` for disconnected sets like `{1,1},{2,2}` or `{1,2},{2,1},{3,4},{4,3}`
9. **Isolated dominoes detected**: Returns `ok=false` for `{1,2},{2,3},{3,1},{4,4}`
10. **Backtracking works**: Handles cases requiring backtracking, e.g. `{1,2},{2,3},{3,1},{2,4},{2,4}`
11. **Separate loops merged**: Handles `{1,2},{2,3},{3,1},{1,1},{2,2},{3,3}`
12. **Large inputs**: Handles 9-element chain correctly
13. **All 12 test cases pass**: `go test` in the dominoes directory passes all tests

## Key Constraints

- Solution must be in package `dominoes` in file `dominoes.go`
- Must use Go 1.18 module
- The `Domino` type is `[2]int`
- Chain validity is verified by the test harness (adjacent matching, end matching, same domino set)
- Multiple valid chains are acceptable; tests only verify validity
