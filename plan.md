# Implementation Plan: Go Alphametics Solver

## Overview

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` to solve alphametics puzzles using a constraint-based backtracking approach with pruning for performance.

## File to Modify

- `go/exercises/practice/alphametics/alphametics.go` — Replace the stub with the full implementation

## Architectural Approach

### Strategy: Column-based constraint backtracking

Rather than brute-force permutation of all possible digit assignments (which is O(10!/(10-n)!) and too slow for 10 letters), use a smarter approach that leverages the structure of the addition problem:

1. **Parse** the puzzle string into addend words and a result word
2. **Extract** unique letters and identify which letters are leading digits (cannot be zero)
3. **Solve** using recursive backtracking with digit assignment, checking constraints as early as possible

### Algorithm Details

The approach mirrors the reference implementation in `.meta/example.go`:

1. **Parse the puzzle**: Split on whitespace, skip `+` and `==` tokens, collect words. The last word is the result; all others are addends.
2. **Build a problem struct** that stores:
   - The words as digit columns (reversed for right-to-left processing)
   - Letter-to-digit mapping array
   - List of unique letters used
   - Maximum word length
3. **Generate permutations** of digits 0-9 taken `n` at a time (where `n` is the number of unique letters)
4. **For each permutation**, check if it satisfies the puzzle equation column-by-column with carry propagation
5. **Validate leading zeros**: reject solutions where any word's leading letter maps to 0

### Key Design Decisions

- **Use the proven approach from example.go**: The reference implementation uses permutation-based search which is correct and handles all test cases. The permutation generator avoids generating all permutations in memory at once — it builds them incrementally.
- **All leading zeros must be checked**, not just the result word. The test case `"ACA + DD == BD"` expects an error because there's no valid solution (B would need to be a leading digit with value 0).
- **Keep it simple**: The implementation follows the same structure as the reference. No external dependencies.

## Implementation Steps

1. Define the `problem` struct with fields for digit columns, letter values, letters used, and max digits
2. Implement `parsePuzzle()` to parse the input string into the problem struct
3. Implement `solvePuzzle()` to iterate over permutations and check each
4. Implement `isPuzzleSolution()` to verify a digit assignment column-by-column with carry
5. Implement `puzzleMap()` to convert the internal representation to the output map
6. Implement `permutations()` helper to generate r-length permutations from digits 0-9
7. Wire everything together in the `Solve()` function

## Testing

- Run `go test -v ./...` from the exercise directory
- All 10 test cases must pass
- The 199-addend test case (test 10) must complete without timeout
