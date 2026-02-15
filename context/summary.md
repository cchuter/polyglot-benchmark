# Context: polyglot-go-alphametics (Issue #13)

## Key Decisions

1. The alphametics solver already existed from issue #4 with all tests passing
2. Identified two latent correctness bugs not exposed by current test cases:
   - Leading-zero check only validated the answer word, not addend words
   - `isPuzzleSolution` did not verify final carry was zero (carry overflow)
3. Fixed both bugs with minimal, targeted changes to `alphametics.go`

## Files Modified

- `go/exercises/practice/alphametics/alphametics.go` - Added `leadingLetters` field, comprehensive leading-zero check, carry overflow fix

## Test Results

- All 10 test cases pass
- `go build` clean
- `go vet` clean
- Total test time: ~1.2s

## Architecture

- Package: `alphametics` (Go 1.18)
- Approach: Brute-force permutation solver using P(10, r) digit permutations
- Column-by-column arithmetic validation with carry propagation
- Standard library only (errors, strings, unicode)

## Branch

- `issue-13` pushed to origin
- Single commit: `fbebfed`
