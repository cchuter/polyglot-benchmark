# Changes: Alphametics Solver Implementation

## File Modified
- `go/exercises/practice/alphametics/alphametics.go`

## What Changed
Implemented `Solve(puzzle string) (map[string]int, error)` using coefficient-based backtracking with bounds pruning.

### Algorithm
1. **Parse**: Split puzzle on whitespace, filter out `+` and `==` tokens. All words except the last are addends; the last word is the result.
2. **Coefficient computation**: Each letter gets a coefficient based on its positional value. Addend letters get positive coefficients (+1, +10, +100...), result letters get negative (-1, -10, -100...). A valid solution has total weighted sum = 0.
3. **Leading-letter constraint**: First character of any multi-character word cannot be assigned digit 0.
4. **Letter ordering**: Letters sorted by |coefficient| descending for early pruning of high-impact assignments.
5. **Backtracking with bounds pruning**: For each unassigned letter, try digits 0-9 (skipping used digits and 0 for leading letters). After each tentative assignment, compute relaxed min/max bounds for remaining letters to check if sum=0 is still reachable. Prune branches where it's impossible.

### Helper Function
- `availMinMax(used int, isLeading bool) (int, int)`: Returns smallest and largest available digits given a used-digit bitmask and leading constraint.

## Test Results
All 10 test cases pass in ~5ms, including the 199-addend puzzle with 10 unique letters.
