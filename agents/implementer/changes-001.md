# Change Log: Implement alphametics Solve function

## File Modified
- `go/exercises/practice/alphametics/alphametics.go`

## Approach: Weight-based Backtracking

Replaced the buggy column-based stub with a weight-based constraint solver:

1. **Parse** the puzzle: split on `==` and `+`, trim whitespace.
2. **Compute weights**: Each letter gets a weight equal to the sum of its place values across all addend words minus its place values in the result word. A valid solution satisfies `sum(weight[letter] * digit[letter]) == 0`.
3. **Leading-zero constraint**: First character of any multi-character word cannot map to 0.
4. **Backtracking with pruning**: Letters are sorted by absolute weight descending (most impactful first). At each step, loose bounds on the remaining sum are computed using min/max available digits to prune impossible branches early.

## Why Weight-based Over Column-based

The weight-based approach converts the puzzle into a single linear equation, avoiding the complexity of carry tracking between columns. This eliminates bugs in carry undo logic and simplifies the implementation while maintaining excellent performance.

## Test Results

All 10 test cases pass in ~5ms, including the 199-addend puzzle with 10 unique letters.
