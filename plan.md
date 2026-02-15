# Implementation Plan: polyglot-go-alphametics

## Current State

The file `go/exercises/practice/alphametics/alphametics.go` already contains a working implementation from commit 761b57d (issue #4). All 10 test cases pass, including the large 199-addend puzzle.

## Analysis of Existing Implementation

The current implementation uses a brute-force permutation approach:

1. **`parsePuzzle`**: Parses the puzzle string into a `problem` struct with:
   - `vDigits`: 2D array of letter indices per word, right-aligned by digit position
   - `letterValues`: array mapping letter index to assigned digit
   - `lettersUsed`: list of unique letters found
   - `maxDigits`: length of the longest word

2. **`solvePuzzle`**: Iterates over all permutations of 10 digits taken `nLetters` at a time, checking each against the puzzle constraints. Checks leading zero only for the answer (last word).

3. **`isPuzzleSolution`**: Validates a candidate assignment by checking column-by-column addition with carry propagation.

4. **`permutations`**: Generates all P(10, r) permutations using the Python `itertools.permutations` algorithm.

## Issues Identified

### Bug: Incomplete leading-zero check
The current code only checks the leading digit of the **answer** (last word) for zero. It does not check leading digits of **addends**. For example, if an addend like "SEND" had S=0, that should be invalid, but the code would not catch it directly.

However, all test cases pass because:
- The test case "ACA + DD == BD" fails for a different reason (answer shorter than longest addend, so column check returns false)
- Other test cases don't have solutions where addend leading digits are zero

### Fix Required
Add comprehensive leading-zero validation for ALL multi-digit words (both addends and the answer) to ensure correctness for all possible inputs.

### Bug: Missing carry overflow check
The `isPuzzleSolution` function does not verify that the final carry is zero after processing all columns. This could cause false positives where a sum overflows beyond the answer's digit count but the overflow carry is silently discarded.

### Fix Required
Change `return true` to `return carry == 0` at the end of `isPuzzleSolution`.

## Plan

### File to Modify
- `go/exercises/practice/alphametics/alphametics.go`

### Changes

1. **Track leading letters**: During `parsePuzzle`, record which letters appear as the leading (first) character of multi-digit words.

2. **Fix `solvePuzzle`**: After finding a candidate solution via `isPuzzleSolution`, check ALL leading letters (not just the answer's leading letter) for zero values. If any leading letter maps to 0, skip the candidate.

3. **Fix carry overflow**: Change `return true` to `return carry == 0` at end of `isPuzzleSolution`.

### Architectural Approach

Add a `leadingLetters` field (a set of rune indices) to the `problem` struct. During parsing, for each word with length > 1, mark its first letter as a leading letter. In `solvePuzzle`, check that no leading letter is assigned the digit 0.

### Ordering

1. Add `leadingLetters` field to `problem` struct
2. Populate it in `parsePuzzle` for all multi-digit words
3. Replace the single leading-zero check in `solvePuzzle` with a loop over all leading letters
4. Fix carry overflow check in `isPuzzleSolution`
5. Run tests to verify all 10 cases still pass
