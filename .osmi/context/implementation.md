# Context: Dominoes Implementation

## Key Decisions
- **Algorithm**: Backtracking DFS with even-degree + connectivity pre-checks
- **Pre-checks use union-find**: Map-based union-find with path compression for connectivity
- **Starting domino fixed to input[0]**: Only try both orientations of the first domino, not all dominoes

## Files Modified
- `go/exercises/practice/dominoes/dominoes.go` — full implementation

## Test Results
- 11/11 tests pass
- `go vet` clean
- Branch: `issue-335`, commit: `cd9b37f`

## Architecture
- `MakeChain()` — public entry point, handles edge cases, calls canChain + search
- `canChain()` — even-degree check + union-find connectivity check
- `search()` — recursive backtracking, tries each unused domino in both orientations
