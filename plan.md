# Implementation Plan: Alphametics Solver

## Branch 1: Permutation-Based Brute Force (Simplicity)

### Approach
Directly follows the reference solution pattern. Parse the puzzle, extract unique letters, generate all permutations of digits for those letters, and test each permutation.

### Files to Modify
- `go/exercises/practice/alphametics/alphametics.go` — full implementation in single file

### Architecture
1. Parse puzzle string by splitting on whitespace, filtering `+` and `==`
2. Extract unique letters and identify leading letters (cannot be 0)
3. Store words in a right-aligned column format for easy column-wise checking
4. Generate all P(10, n) permutations where n = number of unique letters
5. For each permutation, check leading-zero constraint, then check arithmetic column by column
6. Return first valid solution or error

### Evaluation
- **Feasibility**: High — straightforward to implement, single file
- **Risk**: Low for correctness, but HIGH for performance on the 10-letter test case. P(10,10) = 3,628,800 permutations, each requiring arithmetic check over many columns. The 199-addend case multiplies work per check. May time out.
- **Alignment**: Meets all criteria if fast enough
- **Complexity**: Low — ~150 lines, simple algorithm

---

## Branch 2: Column-by-Column Backtracking with Pruning (Extensibility)

### Approach
Use a constraint-propagation backtracking approach. Process the puzzle column by column (right to left), assigning digits one letter at a time and pruning invalid branches early. This is far more efficient than brute force for large puzzles.

### Files to Modify
- `go/exercises/practice/alphametics/alphametics.go` — full implementation in single file

### Architecture
1. **Parse**: Split puzzle on whitespace, filter operators. Last word is the result; all others are addends.
2. **Extract columns**: For each column position (right to left), collect the letters that appear in addends and the expected result letter. Compute a "coefficient" for each letter: +1 for each appearance in an addend column, -1 for the result column position (scaled by powers of 10 isn't needed since we work column by column with carry).
3. **Letter ordering**: Order letters by the column in which they first appear (rightmost first). This ensures that when we backtrack, we assign letters that are needed soonest.
4. **Backtracking solver**:
   - For each letter in order, try each unused digit (0-9)
   - Skip digit 0 for leading letters
   - After assigning all letters in a column, check: `sum of addend digits + carry_in ≡ result digit (mod 10)` and compute carry_out
   - If column check fails, prune (backtrack)
   - If all letters assigned and all columns check out with final carry = 0, solution found
5. **Return** the letter→digit mapping or error

### Evaluation
- **Feasibility**: High — well-known algorithm, fits naturally
- **Risk**: Medium — more complex logic but well-understood technique. The column-based pruning dramatically reduces search space.
- **Alignment**: Fully meets all acceptance criteria, including performance
- **Complexity**: Medium — ~200 lines, more data structures but clean separation

---

## Branch 3: Coefficient-Based Backtracking (Performance)

### Approach
Reduce the puzzle to a single linear equation: assign each letter a coefficient based on its positional value in each word (+1×place for addends, -1×place for result). Then solve `Σ(coefficient_i × digit_i) = 0` via backtracking with pruning.

### Files to Modify
- `go/exercises/practice/alphametics/alphametics.go` — full implementation in single file

### Architecture
1. **Parse**: Extract words, identify addends vs result
2. **Compute coefficients**: For each letter, sum up its positional weight. E.g., in SEND+MORE==MONEY: S has coefficient +1000, E has +100+1-10=+91, N has +10-1000+100=-890, etc. Addend letters get positive weights, result letters get negative weights.
3. **Sort letters**: by |coefficient| descending — assign highest-impact letters first for better pruning
4. **Backtracking**: Try digits 0-9 for each letter (skip 0 for leading), track used digits, maintain running sum. Prune when the remaining unassigned letters cannot possibly make the sum reach 0 (bounds checking).
5. **Bounds pruning**: For remaining unassigned letters, compute min/max possible contribution using available digits. If 0 is not in [current_sum + min, current_sum + max], prune.

### Evaluation
- **Feasibility**: High — straightforward math, single equation
- **Risk**: Low — the coefficient approach is mathematically clean. Bounds pruning makes it very fast.
- **Alignment**: Fully meets all criteria, excellent performance
- **Complexity**: Medium — ~180 lines, clean and efficient

---

## Selected Plan

**Branch 3: Coefficient-Based Backtracking** is selected.

### Rationale
- **vs Branch 1 (Brute Force)**: Branch 1 risks timeout on the 10-letter/199-addend test. Generating all 3.6M permutations in memory before checking is wasteful. Branch 3 uses lazy backtracking with pruning, handling 10 letters easily.
- **vs Branch 2 (Column-by-Column)**: Branch 2 is solid but more complex to implement correctly (managing carry state across columns, partial column checks). Branch 3 reduces the entire problem to a single equation `Σ(coeff_i × digit_i) = 0`, which is simpler to reason about and debug. Bounds pruning on a single sum is simpler than column-based carry tracking.
- **Performance**: Branch 3's bounds pruning cuts the search space dramatically. For the 10-letter case, most branches are pruned within 2-3 assignments.
- **Correctness**: The coefficient approach is mathematically exact — if `Σ(coeff_i × digit_i) = 0`, the arithmetic is correct by construction.

### Detailed Implementation Plan

#### File: `go/exercises/practice/alphametics/alphametics.go`

```go
package alphametics
```

**Step 1: Parse the puzzle**
- Split puzzle string on whitespace
- Filter out `+` and `==` tokens
- All tokens except the last are addends; the last is the result
- Validate: all characters are uppercase letters

**Step 2: Compute letter coefficients**
- For each addend word, iterate right to left, giving each letter a coefficient of +1, +10, +100, etc.
- For the result word, iterate right to left, giving each letter a coefficient of -1, -10, -100, etc.
- Accumulate coefficients per letter (a letter appearing in multiple positions gets the sum)

**Step 3: Identify leading letters**
- The first character of every multi-digit word (addends and result) cannot be 0
- Store in a `leadingLetters` set

**Step 4: Prepare for backtracking**
- Create a sorted list of unique letters (sorted by |coefficient| descending for better pruning)
- Create an array of their coefficients in the same order
- Track `usedDigits` as a bitmask (10 bits for digits 0-9)

**Step 5: Backtracking solver**
```
func solve(index, currentSum, usedDigits):
    if index == len(letters):
        return currentSum == 0
    for digit 0..9:
        if digit is used: skip
        if digit == 0 and letter is leading: skip
        newSum = currentSum + coefficient[index] * digit
        // Bounds pruning: check if remaining letters can make newSum reach 0
        if canReachZero(newSum, remaining coefficients, available digits):
            mark digit used
            if solve(index+1, newSum, usedDigits): return true
            unmark digit
    return false
```

**Step 6: Bounds pruning function**
- For each remaining coefficient, compute min/max contribution using available digits
- For positive coefficients: min uses smallest available digit, max uses largest
- For negative coefficients: min uses largest available digit (most negative), max uses smallest
- Check if `newSum + minRemaining <= 0 <= newSum + maxRemaining`

**Step 7: Return result**
- If solution found, build `map[string]int` from letter assignments
- If no solution, return error

### Implementation Order
1. Write the parsing logic and data structures
2. Write the coefficient computation
3. Write the backtracking solver with bounds pruning
4. Write the main `Solve` function connecting everything
5. Run tests
