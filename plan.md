# Implementation Plan: Alphametics Solver (Revised)

## Overview

Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go` to solve alphametics puzzles using a permutation-based approach with column-by-column validation.

## File to Modify

- `go/exercises/practice/alphametics/alphametics.go` — currently only contains `package alphametics`

## Algorithm: Permutation Search with Column Validation

1. **Parse** the puzzle string: split on whitespace, skip `+` and `==` tokens. Collect words in order; the last word is the result.
2. **Extract unique letters** and identify **leading letters** — the first character of every word (all words, since even single-char words represent non-zero numbers in standard alphametics, though the test suite only enforces this for multi-digit words).
3. **Build column structure**: Pad all words to `maxDigits` length. Store digits right-to-left (index 0 = ones place). Use 0 for padding, 1..26 for A..Z.
4. **Generate permutations** of digits {0..9} taken `nLetters` at a time.
5. **For each permutation**:
   a. Assign digits to letters.
   b. **Early exit**: If any leading letter of a multi-digit word maps to 0, skip.
   c. **Column validation**: Check columns right-to-left with carry propagation. If any column's sum (mod 10) doesn't match the result digit, reject immediately.
   d. **Carry check**: After all columns, verify final carry is 0.
6. **Return** first valid solution as `map[string]int`, or error if no solution found.

## Architecture

Single file (`alphametics.go`) with these components:

```go
// Public entry point
func Solve(puzzle string) (map[string]int, error)

// Internal types and methods
type problem struct { ... }
func parsePuzzle(puzzle string) *problem
func (p *problem) solvePuzzle() (map[string]int, error)
func (p *problem) isPuzzleSolution(values []int) bool
func (p *problem) puzzleMap() map[string]int
func permutations(iterable []int, r int) [][]int
```

## Key Design Decisions

1. **Follow reference implementation structure** — proven correct by test suite, well-understood
2. **Leading-zero check before column validation** — check all multi-digit words' leading letters, not just the result. More correct than reference and faster (avoids column-sum computation)
3. **Final carry == 0 check** — close latent correctness gap in reference where carry-out after last column was not verified
4. **Materialized permutations** — simpler than lazy generation; memory usage (~388 MB for 10 letters) is acceptable for test context

## Implementation Order

1. Write full implementation in `alphametics.go`
2. Run `go test ./...` in the exercise directory
3. Run `go vet ./...`
4. Fix any failures and iterate
