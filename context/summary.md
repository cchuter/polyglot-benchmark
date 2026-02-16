# Context: Book Store Exercise (Issue #193)

## Key Decisions

1. **Algorithm choice**: Greedy + Adjustment over recursive or DP approaches. The greedy layer-peeling approach with 5+3→4+4 adjustment is O(n), simplest, and provably correct for this discount structure.
2. **Pre-computed cost table**: Used a `[6]int` array for group costs instead of computing dynamically, avoiding repeated arithmetic.
3. **No built-in min**: Implemented inline min comparison since `go.mod` targets Go 1.18 (built-in `min` requires Go 1.21).

## Files Modified

- `go/exercises/practice/book-store/book_store.go` — Full implementation (46 lines)

## Test Results

- All 18 test cases pass
- `go vet` clean
- Verifier verdict: PASS

## Branch

- Feature branch: `issue-193`
- Pushed to origin
