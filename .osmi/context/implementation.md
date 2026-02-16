# Implementation Context

## Solution Architecture

The alphametics solver in `go/exercises/practice/alphametics/alphametics.go` uses a coefficient-based approach:

1. **Parse** the puzzle string into addend words and the result word
2. **Compute coefficients**: Each letter's coefficient = sum of place-value contributions in addends minus place-value contributions in result
3. **Backtrack**: Recursively assign digits 0-9 to letters, checking uniqueness and leading-zero constraints
4. **Check**: When all letters are assigned, verify `Σ(coeff[i] * digit[i]) == 0`

## Files Modified

- `go/exercises/practice/alphametics/alphametics.go` — sole implementation file

## Test Results

All 10 test cases pass. Total runtime: ~0.1s. The largest test (199 addends, 10 letters) runs in ~0.02s.

## Branch

`issue-92` based on `bench/polyglot-go-alphametics`
