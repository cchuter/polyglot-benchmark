# Context Summary

## Issue

#317: polyglot-go-alphametics — Implement an alphametics puzzle solver in Go.

## Key Decisions

- **Approach**: Algebraic coefficient method — reduce puzzle to linear equation, then permutation search
- **Why not column-based**: Simpler implementation, equivalent correctness, acceptable performance
- **Leading zeros**: Only constrained for multi-digit words (len > 1), single-letter words can be zero

## Files Modified

- `go/exercises/practice/alphametics/alphametics.go` — Full implementation (Solve + search functions)

## Test Results

All 10 test cases pass. `go vet` clean. Total runtime ~0.91s.

## Branch

`issue-317` — pushed to origin, ready for PR.
