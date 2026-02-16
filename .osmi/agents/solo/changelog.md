# Solo Agent Change Log

## Implementation: Alphametics Solver

### Changes Made

**File modified**: `go/exercises/practice/alphametics/alphametics.go`

- Implemented `Solve(puzzle string) (map[string]int, error)` function
- **Parsing**: Splits puzzle on `==` to separate addends from result, splits addends on `+`
- **Weight computation**: Computes a net weight for each letter based on its positional value across all words. Addend letters get positive weights, result letters get negative weights. A valid solution has weighted sum = 0.
- **Leading zero tracking**: Letters that appear as the first character of multi-digit words are flagged as non-zero.
- **Optimization**: Letters sorted by descending absolute weight for maximum pruning during backtracking.
- **Backtracking solver**: Recursively assigns digits 0-9 to letters, skipping used digits and leading-zero violations, maintaining a running sum for early termination.

### Test Results

All 10 test cases pass in 0.02 seconds, including the 199-addend/10-letter stress test.
`go vet` passes with no issues.
