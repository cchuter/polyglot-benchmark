# Implementation Plan: Alphametics Puzzle Solver

## Branch 1: Direct Permutation (Simplest Approach)

Closely follows the reference solution in `.meta/example.go`. Uses brute-force permutation generation to try all possible digit assignments.

### Approach
1. Parse the puzzle string into words (addends + result)
2. Extract unique letters and identify leading letters (cannot be zero)
3. Generate all permutations of digits for the number of unique letters
4. For each permutation, check if the equation holds with carry propagation
5. Skip permutations where leading letters are assigned zero

### Files to Modify
- `go/exercises/practice/alphametics/alphametics.go` — implement all logic

### Architecture
- `Solve()` function as entry point
- `parsePuzzle()` helper to extract words and letter metadata
- `permutations()` generator function (itertools-style)
- Column-by-column arithmetic verification

### Evaluation
- **Feasibility**: High — proven approach from the reference solution
- **Risk**: Low — straightforward implementation. Main risk is performance on the 199-addend, 10-letter puzzle (10! = 3.6M permutations)
- **Alignment**: Meets all acceptance criteria
- **Complexity**: Medium — ~200 lines, single file

---

## Branch 2: Constraint-Based Backtracking with Pruning

Uses recursive backtracking with constraint propagation. Assigns digits to letters one at a time, pruning impossible branches early using column-based constraints.

### Approach
1. Parse puzzle into column-based representation with coefficients
2. Compute net coefficient for each letter (sum of place-value contributions across all addend occurrences minus result occurrences)
3. Sort letters by most-constrained-first heuristic
4. Recursively assign digits, checking partial sums at each step
5. Prune when partial assignment violates constraints (leading zeros, used digits)

### Files to Modify
- `go/exercises/practice/alphametics/alphametics.go` — implement all logic

### Architecture
- `Solve()` entry point with coefficient computation
- `backtrack()` recursive function with used-digit tracking
- Column constraint checking for early termination

### Evaluation
- **Feasibility**: High — well-known technique for CSP problems
- **Risk**: Medium — more complex logic, but much better worst-case performance
- **Alignment**: Meets all acceptance criteria
- **Complexity**: Higher — ~150-200 lines, more intricate logic

---

## Branch 3: Coefficient-Based Permutation with Leading-Zero Pre-filtering

A hybrid approach: compute algebraic coefficients for each letter, then use permutation search but with aggressive pre-filtering. Instead of checking column-by-column arithmetic, compute a single weighted sum.

### Approach
1. Parse puzzle to extract words
2. Compute a coefficient for each unique letter based on its position and whether it appears in an addend (+) or result (-)
3. Identify letters that cannot be zero (leading letters of multi-digit words)
4. Generate permutations but skip any where a leading letter gets zero
5. For each valid permutation, check if `sum(coefficient[letter] * digit[letter]) == 0`

### Files to Modify
- `go/exercises/practice/alphametics/alphametics.go` — implement all logic

### Architecture
- `Solve()` entry point
- Simple parsing with `strings.Fields`
- Coefficient calculation: letter at position p in addend gets `+10^p`, in result gets `-10^p`
- Single equation check: `sum(coeff[i] * assigned_digit[i]) == 0`
- Permutation generator (same as Branch 1)

### Key Insight
Instead of simulating column-by-column addition, the entire puzzle reduces to one linear equation:
`SEND + MORE == MONEY` becomes:
`1000S + 100E + 10N + D + 1000M + 100O + 10R + E - 10000M - 1000O - 100N - 10E - Y = 0`

This makes the check O(num_letters) per permutation instead of O(num_letters * max_digits).

### Evaluation
- **Feasibility**: High — mathematically elegant and simple to implement
- **Risk**: Low — simple correctness check, same permutation approach but faster per-check
- **Alignment**: Meets all acceptance criteria. The coefficient approach handles any number of addends naturally
- **Complexity**: Low — ~120-150 lines, cleanest code of all three branches

---

## Selected Plan

**Branch 3: Coefficient-Based Permutation with Leading-Zero Pre-filtering**

### Rationale

Branch 3 is superior because:

1. **Simplicity**: The coefficient approach reduces the puzzle to a single linear equation, eliminating the complex column-by-column checking logic. This means fewer bugs and cleaner code.

2. **Performance**: Each permutation check is a simple dot product rather than iterating through columns with carry propagation. This matters for the 199-addend puzzle.

3. **Correctness**: The mathematical formulation naturally handles any number of addends without special cases. Adding coefficients for each word occurrence is straightforward.

4. **Handles the hard test case**: The 199-addend puzzle has 10 unique letters, so we need 10! = 3,628,800 permutations. With a fast O(10) check per permutation and leading-zero pre-filtering, this is feasible.

Compared to Branch 1 (direct permutation), Branch 3 has a faster per-permutation check. Compared to Branch 2 (backtracking with pruning), Branch 3 is simpler to implement and debug while still being fast enough for all test cases.

### Detailed Implementation

#### Step 1: Parse the puzzle
- Split on whitespace using `strings.Fields`
- Separate into addend words and result word (word after `==`)
- Extract unique letters into a sorted slice
- Identify "leading" letters (first letter of any multi-digit word)

#### Step 2: Compute coefficients
- For each addend word, iterate right-to-left, assigning `+10^position` to each letter's coefficient
- For the result word, iterate right-to-left, assigning `-10^position` to each letter's coefficient
- Store as `map[byte]int` or use letter-indexed array

#### Step 3: Generate permutations and check
- Generate permutations of 10 digits taken `n` at a time (where n = number of unique letters)
- For each permutation:
  - Skip if any leading letter is assigned 0
  - Compute `sum = Σ coefficient[i] * digit[i]`
  - If sum == 0, we found the solution
- If no permutation works, return error

#### Step 4: Build result map
- Convert the winning assignment to `map[string]int`

### Code Structure

```go
package alphametics

import (
    "errors"
    "strings"
)

func Solve(puzzle string) (map[string]int, error) {
    // 1. Parse
    // 2. Compute coefficients
    // 3. Try permutations
    // 4. Return result or error
}

func permutations(n, r int) // generates permutations of 0..9 taken r at a time
```

Estimated: ~120-150 lines total in `alphametics.go`.
