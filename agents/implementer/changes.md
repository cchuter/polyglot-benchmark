# Implementer Change Log

## Changes Made

### File: `go/exercises/practice/alphametics/alphametics.go`

Implemented the alphametics puzzle solver based on the reference solution (`.meta/example.go`) with critical bug fixes.

### What was implemented

- **`Solve(puzzle string)`** — Entry point that parses a puzzle string and returns a letter-to-digit mapping or an error.
- **`parsePuzzle(puzzle string)`** — Parses the puzzle into a column-based representation. Splits on whitespace, skips `+` and `==` tokens, collects words, builds column-indexed digit arrays, and tracks unique letters.
- **`solvePuzzle()`** — Generates all r-permutations of digits {0..9} (where r = number of unique letters), checks leading zero constraints, then verifies column-by-column arithmetic.
- **`isPuzzleSolution(values []int)`** — Column-by-column sum verification with carry propagation.
- **`puzzleMap()`** — Converts internal letter-value array to `map[string]int`.
- **`permutations(iterable []int, r int)`** — Generates r-length permutations (Python itertools-style).

### Critical fixes applied (vs reference solution)

1. **Added `isLeading [26]bool` to problem struct** — Tracks which letters are leading letters of multi-digit words.
2. **Leading letter detection in `parsePuzzle`** — Marks `field[0]` as a leading letter for any word with `len(field) > 1`. Single-letter words (e.g., "I", "A") are exempt and can be zero.
3. **Leading zero check BEFORE `isPuzzleSolution` in `solvePuzzle`** — The reference only checked the answer word's leading digit after finding a solution. Our implementation checks ALL leading letters before the expensive arithmetic check, both for correctness and performance.
4. **`carry == 0` check at end of `isPuzzleSolution`** — The reference returned `true` without verifying no leftover carry. Added `return carry == 0` for correctness.

### Commit

`feat: implement alphametics puzzle solver` on branch `issue-189`.
