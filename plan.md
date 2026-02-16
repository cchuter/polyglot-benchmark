# Implementation Plan: Alphametics Solver

## Overview

Implement the `Solve` function in `alphametics.go` using a coefficient-based backtracking approach with partial-sum bounding. The reference implementation in `.meta/example.go` uses brute-force permutation generation which is correct but slow for 10-letter puzzles. Our approach reduces the puzzle to a single linear equation and uses optimized backtracking with pruning.

## File to Modify

- `go/exercises/practice/alphametics/alphametics.go` — the only file to change

## Architectural Approach

### Parsing

Use `strings.Fields()` to tokenize the puzzle string, filtering out `+` and `==` tokens. This robustly handles whitespace. All tokens before `==` (except `+`) are addend words; the token after `==` is the result word. Identify leading letters (first character of each multi-digit word) that cannot be zero.

### Coefficient Computation

For each letter, compute a net coefficient based on positional value:
- Each occurrence as the i-th digit from the right in an addend word contributes `+10^i`
- Each occurrence in the result word contributes `-10^i`
- Accumulate across all words (handles repeated addends like the 199-addend test case naturally)

Example: `SEND + MORE == MONEY` gives: `S*1000 + E*100 + N*10 + D + M*1000 + O*100 + R*10 + E - M*10000 - O*1000 - N*100 - E*10 - Y = 0`

### Solving Strategy: Backtracking with Partial-Sum Bounding

1. **Sort letters by descending absolute coefficient** — assigns high-impact letters first, producing tighter bounds earlier and pruning more aggressively
2. **Recursive backtracking** — assign digits 0-9 to letters one at a time:
   - Skip digit 0 for leading letters
   - Skip already-used digits (tracked via `[10]bool`)
3. **Partial-sum bounding** — after each assignment, compute whether zero is still reachable given remaining unassigned coefficients and available digits:
   - Compute min possible sum by pairing most-negative remaining coefficients with largest available digits and most-positive with smallest
   - Compute max possible sum by the opposite pairing
   - If `[partial + min_remaining, partial + max_remaining]` doesn't contain 0, prune
4. **Terminal check** — when all letters assigned, verify weighted sum equals 0 exactly

### Data Structures

- `[]byte` for unique letters (since all are uppercase ASCII)
- `[26]int` for letter-to-coefficient mapping (indexed by `letter - 'A'`, avoids map overhead in hot path)
- `[10]bool` for tracking used digits
- `[26]bool` for leading letter flags
- `[]int` for sorted coefficient values (for bounding computation)

### Output Conversion

Convert internal `[26]int` values to `map[string]int` with single-character string keys (e.g., `"S"`, `"M"`) matching test expectations.

### Error Handling

Return `errors.New("no solution found")` when backtracking exhausts all possibilities. Any non-nil error satisfies the test harness.

## Ordering of Implementation

1. Write the `Solve` function as the entry point
2. Write parsing logic: `strings.Fields()` tokenization, coefficient computation, leading letter identification
3. Sort unique letters by descending absolute coefficient
4. Write the backtracking solver with partial-sum bounding
5. Write the bounding function (min/max reachable sum calculation)
6. Return the solution map or error

## Performance Expectations

- Without bounding: 10! = 3,628,800 nodes worst case (too slow)
- With coefficient-magnitude ordering + partial-sum bounding: typically resolves in thousands to tens of thousands of nodes
- The 199-addend puzzle has the same 10 unique letters as the simpler 10-letter test; the coefficient precomputation absorbs the addend count

## Rationale

- Coefficient-based reduction converts any puzzle (regardless of addend count) into a single weighted-sum-equals-zero problem
- Backtracking avoids O(n!) memory for permutation storage
- Partial-sum bounding is the critical optimization for 10-letter performance
- Array-based data structures (`[26]int`, `[10]bool`) avoid map overhead in the recursive hot path
- `strings.Fields()` for parsing is robust and matches the reference implementation's approach
- Standard library only: `strings`, `errors`, `sort`
