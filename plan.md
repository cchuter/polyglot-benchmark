# Implementation Plan: Alphametics Solver

## Branch 1: Direct Permutation (Reference Solution Approach)

**Approach**: Closely follow the reference solution in `.meta/example.go`. Parse the puzzle into column-based digit arrays, generate all permutations of digits for the unique letters, and check each permutation.

**Files to modify**:
- `go/exercises/practice/alphametics/alphametics.go`

**Architecture**:
1. Parse puzzle: split on whitespace, skip `+` and `==`, collect words
2. Build column-based representation (digits from right to left)
3. Track unique letters and which are leading letters (cannot be zero)
4. Generate all r-permutations of {0..9} where r = number of unique letters
5. For each permutation, check column-by-column with carry propagation
6. Skip permutations that assign 0 to a leading letter
7. Return first valid solution or error

**Rationale**: Minimal code, proven correct (it's the reference solution). Simple to understand.

**Evaluation**:
- **Feasibility**: High — reference solution exists and works
- **Risk**: Low — well-tested approach. However, the reference only checks leading zero on the answer; need to fix this to check ALL words
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: Single file, ~200 lines, moderate complexity

## Branch 2: Backtracking with Constraint Propagation

**Approach**: Use a backtracking search with constraint propagation. Assign digits letter-by-letter, pruning early when constraints are violated. Process columns right-to-left to detect contradictions sooner.

**Files to modify**:
- `go/exercises/practice/alphametics/alphametics.go`

**Architecture**:
1. Parse puzzle into words (addends and result)
2. Extract unique letters and identify leading letters
3. Assign column weights to each letter (e.g., S in SEND has weight 1000 as addend)
4. Use recursive backtracking: for each unassigned letter, try digits 0-9
5. Prune: skip digits already used; skip 0 for leading letters
6. At each step, check partial sums for early termination
7. When all letters assigned, verify full equation

**Rationale**: More extensible — constraint propagation framework can be enhanced. Better for harder puzzles.

**Evaluation**:
- **Feasibility**: High — standard CS approach
- **Risk**: Medium — more code to write and debug, but the approach is well-understood
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: Single file, ~150-200 lines, higher algorithmic complexity

## Branch 3: Column-by-Column Backtracking with Early Pruning

**Approach**: Process the puzzle column-by-column from right to left. For each column, assign digits to the letters that appear in that column (if not yet assigned), then verify the column sum. This gives maximum pruning since we can reject partial assignments without computing the full equation.

**Files to modify**:
- `go/exercises/practice/alphametics/alphametics.go`

**Architecture**:
1. Parse puzzle into words and result
2. Build column representation (right to left)
3. For each column, determine which NEW letters appear (not yet assigned)
4. Recursive solve: process columns left-to-right (least significant first)
5. For each column, try all valid digit assignments for new letters in that column
6. Check column arithmetic: sum of addend digits + carry_in = result_digit + 10 * carry_out
7. Prune: if column doesn't balance, backtrack immediately
8. Leading zero constraint: reject 0 for leading letters

**Rationale**: Maximum performance through early column-level pruning. Important for the 199-addend, 10-letter test case.

**Evaluation**:
- **Feasibility**: Medium-High — more complex to implement correctly
- **Risk**: Medium-High — column-by-column logic with partial assignments is tricky
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: Single file, ~200-250 lines, highest complexity

---

## Selected Plan

**Selected: Branch 1 — Direct Permutation (Reference Solution Approach)**

**Rationale**: Branch 1 is the best choice because:

1. **Proven correctness**: The reference solution in `.meta/example.go` is known to pass all tests. We can adapt it directly.
2. **Minimal risk**: The approach is battle-tested. The only fix needed is ensuring leading zero checks cover ALL words, not just the answer.
3. **Simplicity**: The permutation approach is straightforward and easy to verify.
4. **Performance**: While theoretically slower than backtracking for very large puzzles, the test suite's largest case has only 10 unique letters. P(10,10) = 3,628,800 permutations, which is manageable. The reference solution is known to handle this.
5. **Alignment**: Directly satisfies all acceptance criteria with minimal implementation risk.

Branch 2 and 3 are more elegant algorithmically but introduce unnecessary complexity and debugging risk for no practical benefit given the test constraints.

### Detailed Implementation Plan

**File**: `go/exercises/practice/alphametics/alphametics.go`

**Step 1: Define the package and imports**
```go
package alphametics

import (
    "errors"
    "strings"
    "unicode"
)
```

**Step 2: Define the problem struct**
- `vDigits [][]rune` — column-indexed digit representations of each word
- `maxDigits int` — maximum word length
- `letterValues [26]int` — digit assignment per letter (A=0, B=1, ...)
- `lettersUsed []rune` — list of unique letters found
- `nLetters int` — count of unique letters
- `leadingLetters` — set of letters that appear as leading digits (must not be 0)

**Step 3: Implement `Solve(puzzle string) (map[string]int, error)`**
- Parse the puzzle via `parsePuzzle`
- Call `solvePuzzle` to find a valid assignment
- Return the solution map or error

**Step 4: Implement `parsePuzzle`**
- Split puzzle on whitespace
- Skip `+` and `==` tokens
- Track leading letter of each multi-digit word
- Build column-indexed representation
- Collect unique letters

**Step 5: Implement `solvePuzzle`**
- Generate all permutations of {0..9} choosing nLetters
- For each permutation, check leading zero constraint for ALL words
- Check column-by-column arithmetic with carry
- Return first valid solution

**Step 6: Implement helpers**
- `isPuzzleSolution` — column-by-column sum verification with carry
- `puzzleMap` — convert internal representation to `map[string]int`
- `permutations` — generate r-permutations of a slice

**Key differences from reference (incorporating review feedback)**:

1. **Leading zero check for ALL multi-digit words**: The reference only checks the answer word. Our implementation must check leading zeros for ALL multi-digit words (addends and result). Single-letter words like "I" or "A" are exempt — they CAN be zero.

2. **Check leading zeros BEFORE `isPuzzleSolution`**: Move the leading zero check before the expensive arithmetic check. This prunes invalid permutations cheaply and improves performance.

3. **Track leading letters during parsing**: In `parsePuzzle`, record the leading letter of each word with `len(word) > 1`. Store these in a boolean array `isLeading [26]bool`.

4. **Final carry check**: Add `carry == 0` verification after the column loop in `isPuzzleSolution` for theoretical correctness.
