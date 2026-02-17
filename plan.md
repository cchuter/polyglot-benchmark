# Implementation Plan: Alphametics Solver

## Branch 1: Direct Permutation (Reference Solution Approach)

Follow the reference solution closely. Parse the puzzle into a columnar representation, generate all permutations of digits for the number of unique letters, and test each permutation.

### Files to modify
- `go/exercises/practice/alphametics/alphametics.go`

### Approach
1. Parse puzzle string by splitting on whitespace, skipping `+` and `==` tokens
2. Extract unique letters and track which letters appear in leading positions
3. Store words in a columnar (right-aligned) format for efficient column-by-column checking
4. Generate all permutations of 10 digits taken `n` at a time (where `n` = number of unique letters)
5. For each permutation, check column sums with carry propagation
6. Reject solutions where leading letters map to zero

### Evaluation
- **Feasibility**: High — this is the proven reference approach
- **Risk**: Low — directly based on the working example solution
- **Alignment**: Satisfies all acceptance criteria
- **Complexity**: Medium — ~200 lines, single file, permutation generator needed

### Concern
For the 10-letter puzzle, this generates P(10,10) = 3,628,800 permutations. Each must be checked. This is the brute-force upper bound but the reference solution uses this approach, so it should be acceptable.

---

## Branch 2: Backtracking with Column Pruning

Use a depth-first backtracking search that assigns digits to letters one at a time, processing columns from right to left. Prune invalid assignments early by checking partial column sums.

### Files to modify
- `go/exercises/practice/alphametics/alphametics.go`

### Approach
1. Parse puzzle into words and result, extract unique letters
2. Order letters by the column they first appear in (rightmost first)
3. Use recursive backtracking: assign a digit to the next unassigned letter
4. After each assignment, check if any column is now fully assigned; if so, verify the column sum
5. Prune early: if a fully-assigned column doesn't match, backtrack
6. Enforce leading-zero constraint during assignment (not just at the end)

### Evaluation
- **Feasibility**: High — backtracking is a standard approach for constraint satisfaction
- **Risk**: Medium — more complex to implement correctly, especially carry tracking across partially-assigned columns
- **Alignment**: Satisfies all acceptance criteria, potentially faster for the 199-addend puzzle
- **Complexity**: High — more sophisticated logic, harder to debug

---

## Branch 3: Optimized Permutation with Coefficient Reduction

Reduce the puzzle to a single linear equation over letter coefficients, then use permutation search with early termination. Instead of checking column-by-column, compute a single weighted sum.

### Files to modify
- `go/exercises/practice/alphametics/alphametics.go`

### Approach
1. Parse the puzzle and compute a coefficient for each letter based on its positional value
   - For `SEND + MORE == MONEY`: S has coefficient 1000, E has coefficient (100 + 1 - 10) = 91, etc.
   - Addend letters get positive coefficients; result letters get negative coefficients
2. The puzzle is solved when `sum(coefficient[letter] * digit[letter]) == 0`
3. Order letters by descending absolute coefficient value (most constrained first)
4. Use recursive digit assignment with pruning:
   - Track running partial sum
   - If remaining digits can't possibly bring sum to zero, prune
5. Enforce leading-zero and unique-digit constraints during search

### Evaluation
- **Feasibility**: High — well-known optimization technique for alphametics
- **Risk**: Medium — coefficient computation and pruning bounds require care
- **Alignment**: Satisfies all criteria; likely fastest approach for the 199-addend, 10-letter puzzle
- **Complexity**: Medium — cleaner than Branch 2, about same as Branch 1

---

## Selected Plan

**Branch 3: Optimized Permutation with Coefficient Reduction**

### Rationale

Branch 3 is selected because:
1. It reduces the puzzle to a single equation, making the check trivial (sum == 0)
2. The coefficient approach handles 199 addends naturally — just add coefficients for repeated letters
3. Pruning by partial sum bounds dramatically reduces the search space vs. Branch 1's exhaustive permutations
4. It's cleaner than Branch 2's column-by-column backtracking with partial carries
5. The implementation complexity is manageable — roughly on par with Branch 1

### Detailed Implementation

#### File: `go/exercises/practice/alphametics/alphametics.go`

```go
package alphametics

import (
    "errors"
    "strings"
)
```

**Step 1: Parse the puzzle**
- Split on `==` to get left-hand side and right-hand side
- Split left-hand side on `+` to get addend words
- Trim whitespace from each word
- Collect unique letters
- Identify leading letters (first character of each word with length > 1)

**Step 2: Compute coefficients**
- For each addend word, each letter gets `+10^position` added to its coefficient (position 0 = ones, 1 = tens, etc.)
- For the result word, each letter gets `-10^position` added to its coefficient
- The puzzle is solved when `sum(coeff[i] * digit[i]) == 0` for all letters `i`

**Step 3: Backtracking solver**
- Sort letters by descending absolute coefficient (most impactful first)
- Maintain a `used` bitmask of digits 0-9 already assigned
- Maintain a `leadingLetters` set — digits assigned to these cannot be 0
- Recurse: for each unassigned letter, try each unused digit (skipping 0 for leading letters)
  - Add `coeff * digit` to running sum
  - Prune: compute max/min possible contribution from remaining letters
    - For each remaining letter, find the best/worst unused digit * coefficient
    - If 0 is not reachable from the current partial sum + possible range, prune
  - If all letters assigned and sum == 0, return solution

**Step 4: Build and return result map**
- Convert internal letter indices to `map[string]int` with letter strings as keys

### Pruning Detail

For remaining unassigned letters with coefficients c_i, compute:
- `maxRemaining`: sum of (max positive contribution) for each remaining letter
- `minRemaining`: sum of (min/most negative contribution) for each remaining letter
- If `partialSum + minRemaining > 0` or `partialSum + maxRemaining < 0`, prune

This is an O(remaining) check per node, but dramatically reduces the tree.

### Implementation Order
1. Write the parsing logic
2. Write the coefficient computation
3. Write the backtracking solver with pruning
4. Write the result map builder
5. Test with `go test ./...`
6. Optimize if the 199-addend test is too slow
