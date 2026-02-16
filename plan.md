# Implementation Plan: Alphametics Solver

## File to Modify

- `go/exercises/practice/alphametics/alphametics.go`

## Approach: Column-based Constraint Solving with Backtracking

### Why Not Brute Force?

A naive approach generating all permutations of digits for N letters would be O(10!/(10-N)!) which is up to 3.6M for 10 letters. This is too slow when combined with 199 addends in the last test case. We need a smarter approach.

### Algorithm: Column-by-Column Backtracking

Process the puzzle column by column (right to left), assigning digits to letters as they are encountered. This allows early pruning: if a partial assignment is inconsistent for a column, we backtrack immediately.

### Detailed Steps

1. **Parse the puzzle**: Split on `==` to get LHS and RHS. Split LHS on `+` to get addend words. Trim whitespace.

2. **Extract columns**: Organize letters by column position (units, tens, hundreds, etc.). For each column, collect which letters appear from each word (addends and result), working from the rightmost character.

3. **Identify leading letters**: Letters that appear as the first character of any multi-digit word cannot be zero.

4. **Collect unique letters** in the order they appear during column processing (right to left). This determines the assignment order for backtracking.

5. **Backtrack column by column**:
   - For each column, determine which letters participate
   - If a letter hasn't been assigned yet, try all available digits (respecting leading-zero constraint)
   - Once all letters in a column are assigned, check if the column sum (including carry from previous column) is consistent with the result letter's digit
   - If consistent, move to the next column with the new carry
   - If all columns are satisfied and final carry is zero, we have a solution

### Implementation Structure

```go
package alphametics

import (
    "errors"
    "strings"
)

func Solve(puzzle string) (map[string]int, error) {
    // 1. Parse puzzle
    // 2. Build column structure
    // 3. Identify unique letters and leading letters
    // 4. Run backtracking solver
    // 5. Return result or error
}
```

### Key Data Structures

- `words []string` — all words (addends + result)
- `addends []string`, `result string` — parsed from puzzle
- `letters []byte` — unique letters in assignment order
- `leading map[byte]bool` — letters that cannot be zero
- `assignment map[byte]int` — current letter→digit mapping
- `used [10]bool` — which digits are currently in use

### Performance Considerations

- Column-by-column processing enables early pruning
- For the 199-addend test case, each column may have many entries from the same letter, but the number of unique letters is still ≤ 10
- The solver processes at most 10 unique letters with at most 10 digit choices each, heavily pruned by column constraints

## Testing

- Run `go test ./...` in the alphametics exercise directory
- All 10 test cases must pass
- The benchmark test should complete reasonably (skipped in short mode)

## Commit Strategy

- Single commit: implement the Solve function
- Commit message: `Closes #80: polyglot-go-alphametics`
