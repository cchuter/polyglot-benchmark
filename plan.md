# Implementation Plan: Alphametics Solver

## Overview

Implement a permutation-based alphametics solver in `go/exercises/practice/alphametics/alphametics.go`. The approach follows the reference example in `.meta/example.go` — parsing the puzzle, generating digit permutations for unique letters, and checking each permutation against the arithmetic equation with leading-zero constraints.

## File to Modify

- `go/exercises/practice/alphametics/alphametics.go` — currently an empty stub with only `package alphametics`

## Architecture

### Data Structures

```go
type problem struct {
    vDigits      [][]rune   // per-word digit slots, reversed (LSB first), using 1-26 for A-Z, 0 for empty
    maxDigits    int        // max word length across all operands/result
    letterValues [26]int    // current digit assignment per letter (index 0 = 'A')
    lettersUsed  []rune     // list of unique letters (as 0-25 indices)
    nLetters     int        // count of unique letters
}
```

### Algorithm

1. **Parse**: Split puzzle on whitespace, skip `+` and `==` tokens. Collect words, track unique letters, determine max word length. Validate all characters are uppercase.

2. **Build digit grid**: For each word, create a reversed array of letter indices (1-based, 0 = no letter in this position), padded to `maxDigits`.

3. **Solve**: Generate all r-permutations of digits 0-9 (where r = number of unique letters). For each permutation:
   - Assign digits to letters
   - Check column-by-column addition with carry
   - If arithmetic checks out, verify no leading zeros on any multi-digit word
   - Return first valid solution

4. **Leading zero check**: After finding an arithmetic match, verify that the first letter of every multi-digit word is not assigned digit 0.

5. **Error handling**: Return `errors.New("no solution")` if no permutation works; return `errors.New("invalid puzzle")` if parsing fails.

### Key Functions

| Function | Purpose |
|----------|---------|
| `Solve(puzzle string) (map[string]int, error)` | Public entry point |
| `parsePuzzle(puzzle string) *problem` | Parse and validate puzzle input |
| `(p *problem) solvePuzzle() (map[string]int, error)` | Iterate permutations, find solution |
| `(p *problem) isPuzzleSolution(values []int) bool` | Check if assignment satisfies the equation |
| `(p *problem) puzzleMap() map[string]int` | Convert internal state to result map |
| `permutations(iterable []int, r int) [][]int` | Generate r-length permutations of digits |

### Performance Consideration

The permutation-based approach generates P(10, r) permutations where r is the number of unique letters. For the worst case (10 letters), this is 10! = 3,628,800 permutations. The column-by-column check provides early termination. The reference solution handles the 199-addend test case this way.

A more optimized approach (backtracking with constraint propagation) could be faster, but the permutation approach matches the reference and passes all tests including the large one.

## Approach and Ordering

1. Write the complete implementation in `alphametics.go` based on the reference pattern
2. The implementation adapts the reference `.meta/example.go` but improves the leading-zero check to cover ALL words (not just the answer), ensuring test case 3 (`"ACA + DD == BD"`) correctly returns an error
3. Run `go test` to verify all 10 test cases pass
4. Run `go vet` to check for code issues

## Imports Required

```go
import (
    "errors"
    "strings"
    "unicode"
)
```
