# Solo Agent Change Log

## Change 1: Implement Solve function in alphametics.go

**File Modified:** `go/exercises/practice/alphametics/alphametics.go`

**What Changed:**
- Added `Solve(puzzle string) (map[string]int, error)` function
- Parses puzzle string by splitting on whitespace and filtering operators
- Computes algebraic coefficients for each unique letter based on place values
- Uses recursive backtracking to assign digits to letters
- Enforces unique-digit and no-leading-zero constraints
- Returns solution map or error if no solution exists

**Approach:** Coefficient-based reduction — each letter gets a net coefficient from its appearances in addends (+) and the result (-). The puzzle is solved when `Σ(coeff[i] * digit[i]) == 0`.

**Test Results:** All 10 test cases pass in 0.1 seconds. The 199-addend, 10-letter puzzle completes in ~0.02s.
