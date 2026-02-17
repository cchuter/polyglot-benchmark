# Implementation Context

## Algorithm: Coefficient Reduction + Backtracking

The alphametics solver works by:
1. Reducing the puzzle to a linear equation: `sum(coeff[letter] * digit[letter]) == 0`
2. Addend letters get positive coefficients (+10^position), result letters get negative
3. Letters are sorted by descending |coefficient| for optimal pruning
4. Backtracking assigns digits to letters one at a time
5. At each node, min/max bounds are computed for remaining unassigned letters
6. If the current partial sum + bounds cannot include 0, the branch is pruned

## Files Modified
- `go/exercises/practice/alphametics/alphametics.go` — sole implementation file

## Test Results
- 10/10 tests pass
- Runtime: ~5ms total
- `go vet`: clean

## Branch: `issue-233`
- Single commit: `1198df4` — "Closes #233: polyglot-go-alphametics"
