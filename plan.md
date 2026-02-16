# Implementation Plan: Alphametics Solver

## File to Modify

- `go/exercises/practice/alphametics/alphametics.go` — the only file that needs changes.

## Algorithm: Column-based backtracking with constraint propagation

### Why not brute-force permutation?

A naive approach generating all permutations of 10 digits for N letters is O(10!/(10-N)!) which for 10 letters is 3,628,800 combinations. The 199-addend test case has 10 unique letters, so brute force is borderline but the real concern is evaluating each candidate — summing 199 words per candidate is expensive.

### Chosen approach: Column-by-column solving with carry propagation

Process the puzzle column by column (right to left, like manual addition). At each column, try to assign digits to any unassigned letters in that column, checking constraints immediately. This prunes the search space dramatically because:

1. We check partial sums column by column, pruning invalid branches early.
2. We enforce digit uniqueness incrementally.
3. We check leading-zero constraints as soon as a leading letter gets assigned.

### Implementation Steps

#### Step 1: Parse the puzzle

1. Split on `" == "` to separate left-hand side (LHS) from right-hand side (RHS).
2. Split LHS on `" + "` to get individual addend words.
3. Collect all unique letters.
4. Identify leading letters (first letter of each multi-digit word) — these cannot be 0.

#### Step 2: Build column structure

For each column index (0 = rightmost):
- Collect which letters appear at that position in each word.
- Build a coefficient map: for each letter, how many times it appears on the LHS minus the RHS at this column position.
- Track which letters first appear in this column (need assignment).

#### Step 3: Recursive column-by-column solver

```
solve(column, carry, assigned_digits, used_digits):
  if column > max_columns:
    return carry == 0  (all columns satisfied, no remaining carry)

  letters_in_this_column = unassigned letters appearing in column

  assign_letters(letters_in_this_column, ...):
    compute column_sum = sum of (letter_coefficients * assigned_digit) + carry_in

    The column constraint is: column_sum mod 10 == result_digit_at_this_column
    and carry_out = column_sum / 10

    recurse to next column with carry_out
```

Actually, a cleaner approach: process columns right-to-left. For each column, determine which letters appear. Some may already be assigned. For unassigned ones, try all valid digit assignments. Once all letters in the column are assigned, check that the column sum (LHS digits + carry_in) mod 10 equals the RHS digit, and compute carry_out.

#### Step 4: Optimization details

- **Early pruning**: If all letters in a column are assigned, immediately check the column constraint before recursing deeper.
- **Leading zero check**: When assigning a digit to a letter, immediately reject 0 if the letter is a leading letter.
- **Digit uniqueness**: Maintain a boolean array `used[10]` tracking which digits are taken.

### Data Structures

```go
type solver struct {
    columns    []column   // indexed by column position (0=rightmost)
    letters    []byte     // all unique letters
    leading    [256]bool  // leading[letter] = true if letter leads a word
    assignment [256]int   // assignment[letter] = digit (-1 if unassigned)
    used       [10]bool   // used[digit] = true if digit is taken
}

type column struct {
    addendLetters []byte  // letters at this column position in LHS words
    resultLetter  byte    // letter at this column position in RHS word
}
```

### Pseudocode for the solve loop

```
func (s *solver) solveColumn(col int, carry int) bool:
    if col == len(s.columns):
        return carry == 0

    // Collect unassigned letters in this column
    unassigned = unique unassigned letters in columns[col]

    return s.assignAndCheck(col, carry, unassigned, 0)

func (s *solver) assignAndCheck(col, carry int, unassigned []byte, idx int) bool:
    if idx == len(unassigned):
        // All letters in this column are assigned; check constraint
        sum = carry
        for each addend letter in column:
            sum += s.assignment[letter]
        result_digit = s.assignment[result_letter]
        if sum % 10 != result_digit:
            return false
        return s.solveColumn(col + 1, sum / 10)

    letter = unassigned[idx]
    startDigit = 0
    if s.leading[letter]: startDigit = 1

    for d = startDigit; d <= 9; d++:
        if !s.used[d]:
            s.assignment[letter] = d
            s.used[d] = true
            if s.assignAndCheck(col, carry, unassigned, idx+1):
                return true
            s.used[d] = false
            s.assignment[letter] = -1

    return false
```

### Edge Cases

- Result word shorter than an addend word: won't happen per puzzle definition, but handle via column structure naturally.
- Puzzles with no `+` (e.g., `"A == B"`): still works; A is the only addend.
- Multiple carries: handled naturally by carry propagation.

## Order of Implementation

1. Write the `Solve` function with parsing logic.
2. Implement the column-based solver.
3. Run tests, iterate on any failures.
4. Verify with `go vet`.
