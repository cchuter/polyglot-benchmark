# Solo Agent Change Log

## Change 1: Implement alphametics solver

**File modified**: `go/exercises/practice/alphametics/alphametics.go`

**What was done**:
- Implemented `Solve(puzzle string) (map[string]int, error)` function
- Implemented `search()` recursive backtracking helper
- Parsing: splits puzzle on `==` and `+`, extracts words, identifies unique letters and leading letters
- Algebraic approach: computes positional coefficients for each letter (addends positive, result negative), reducing the puzzle to a single linear equation
- Backtracking search: tries all digit permutations with uniqueness and leading-zero constraints
- Returns `map[string]int` mapping or error if no solution exists

**Test results**: All 10 test cases pass in 0.91s. `go vet` clean.
