# Implementation Plan: Alphametics Solver

## Proposal A: Column-based Backtracking with Constraint Propagation

**Role: Proponent**

### Approach

Process the puzzle column-by-column (right to left), assigning digits to letters as they are encountered. At each column, determine which letters participate, attempt digit assignments, check column arithmetic consistency (sum of column digits + carry-in ≡ result digit mod 10, with appropriate carry-out), and backtrack on failure.

### Files to Modify

- `go/exercises/practice/alphametics/alphametics.go` — sole implementation file

### Architecture

1. **Parse** the puzzle string: split on `==` to get LHS and RHS, split LHS on `+`, trim whitespace, extract individual words.
2. **Extract unique letters** and identify which letters are leading characters (cannot be zero).
3. **Build column representation**: for each column index (0 = rightmost), collect which letters appear in addend words and the result word at that position, along with their coefficient (how many times each letter appears in the addend column).
4. **Recursive solve**: process columns from right to left. For each column:
   - Identify unassigned letters in this column
   - Try all valid digit assignments for unassigned letters
   - Check: (sum of addend column values + carry_in) mod 10 == result column value, carry_out = (sum + carry_in) / 10
   - If consistent, recurse to next column
5. **Leading zero constraint**: when assigning a digit, skip 0 for letters that lead a multi-digit word.
6. **Return** the mapping on success, or error if all possibilities exhausted.

### Rationale

- Column-based approach prunes the search space early — we only explore assignments consistent with column arithmetic
- Handles the 199-addend puzzle efficiently because column processing reduces the effective branching factor
- Natural handling of carry propagation
- More efficient than brute-force permutation for large puzzles

### Weaknesses Acknowledged

- More complex implementation than simple permutation
- Need careful bookkeeping for column coefficients when letters repeat across many addends

---

## Proposal B: Permutation-based Brute Force with Early Pruning

**Role: Opponent**

### Approach

Extract all unique letters (at most 10), generate permutations of digits 0-9 taken len(letters) at a time, filter by leading-zero constraints, and check if the arithmetic equation holds.

### Files to Modify

- `go/exercises/practice/alphametics/alphametics.go` — sole implementation file

### Architecture

1. **Parse** the puzzle string: split on `==`, split LHS on `+`, trim whitespace.
2. **Extract unique letters** and identify leading letters.
3. **Precompute word structure**: for each word, record the sequence of letter indices.
4. **Recursive permutation generator**: assign digits to letters one at a time, using a `used` bitset to track which digits are taken.
   - Skip digit 0 for leading letters
   - Once all letters assigned, evaluate the equation
   - If it holds, return the solution
5. **Evaluate equation**: convert each word to its numeric value using the current assignment, sum the addends, compare to the result.

### Critique of Proposal A

- Column-based approach is more complex to implement correctly (carry tracking, column coefficient aggregation)
- For puzzles with few unique letters (3-8), the simpler permutation approach is fast enough
- Permutation approach is easier to reason about correctness

### Rationale for Proposal B

- Simpler, more straightforward implementation
- Easier to debug and verify
- With early pruning (assign one letter at a time, check partial constraints), it can be made reasonably fast

### Weaknesses Acknowledged

- For 10 unique letters, worst case is 10! = 3,628,800 permutations — potentially slow
- The 199-addend test case with 10 letters could be very slow if pruning isn't aggressive enough
- Evaluating the full equation on every complete assignment is wasteful

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A (Column-based) | Proposal B (Permutation) |
|-----------|--------------------------|-------------------------|
| Correctness | Both can satisfy all criteria | Both can satisfy all criteria |
| Risk | Medium — more code complexity | High — 10-letter 199-addend test may timeout |
| Simplicity | Moderate | High for small puzzles, but needs optimization for large ones |
| Consistency | Fits Go style (no external deps) | Fits Go style (no external deps) |

### Decision

**Selected: Hybrid approach inspired by Proposal B with key optimizations from Proposal A.**

The permutation-based approach (Proposal B) is simpler and easier to get correct. However, to handle the 199-addend test case efficiently, we'll add two key optimizations:

1. **Precompute letter weights**: Instead of converting words to numbers on every check, precompute the net coefficient of each letter across the equation. For example, in `SEND + MORE == MONEY`, the letter `S` has weight +1000 (from SEND), and `M` has weight -10000+1000 = -9000 (result side negative, addend side positive). A valid solution has the weighted sum equal to zero.

2. **Recursive assignment with partial-sum pruning**: Assign digits to letters one by one. Maintain a running partial sum. After all letters are assigned, check if the total is zero. Additionally, order letters by descending weight magnitude to maximize early pruning.

This gives us the simplicity of the permutation approach with the efficiency needed for large puzzles.

### Detailed Implementation Plan

**File: `go/exercises/practice/alphametics/alphametics.go`**

```go
package alphametics

import (
    "errors"
    "strings"
)
```

#### Step 1: Parse the puzzle

- Split on `==` to get left-hand side and right-hand side
- Split LHS on `+`
- Trim all words
- Validate: exactly one `==`, at least one `+` or single word on LHS

#### Step 2: Extract letters and compute weights

- Iterate over each addend word. For each letter at position `i` from the right, add `10^i` to that letter's weight.
- For the result word, subtract `10^i` for each letter position (since we want addends - result == 0).
- Track which letters are leading characters of multi-digit words (cannot be zero).
- Collect unique letters.

#### Step 3: Sort letters by weight magnitude (descending)

This ensures we assign the most significant letters first, maximizing pruning potential.

#### Step 4: Recursive solver

```
solve(letterIndex, usedDigits, partialSum):
    if letterIndex == len(letters):
        return partialSum == 0
    for digit 0..9:
        if digit is used: skip
        if digit == 0 and letter is leading: skip
        newSum = partialSum + digit * weight[letter]
        recurse with letterIndex+1, mark digit used
```

#### Step 5: Return result

- Build `map[string]int` from the solution assignment
- Return error if no solution found

### Ordering of Changes

1. Write the `Solve` function with parsing logic
2. Write the weight computation
3. Write the recursive solver
4. Test with `go test ./...`
5. Fix any issues
6. Verify with `go vet ./...`
