# Implementer Changes

## File Modified
- `go/exercises/practice/alphametics/alphametics.go`

## Changes Made

### 1. Added `leadingLetters` field to `problem` struct (line 17)
Added `leadingLetters []rune` to track the letter indices of the first character of every multi-digit word in the puzzle.

### 2. Populated `leadingLetters` in `parsePuzzle` (lines 50-60)
After collecting all words, iterate over `valueStrings` and for each word with length > 1, record its first letter's index (`rune(field[0]) - 'A'`) into `p.leadingLetters`. A `seen` map prevents duplicates.

### 3. Replaced single leading-zero check in `solvePuzzle` (lines 91-101)
The old code only checked the answer word's leading digit for zero. The new code loops over ALL entries in `p.leadingLetters` (covering both addends and the answer) and rejects any solution where any leading letter maps to 0.

### 4. Fixed carry overflow check in `isPuzzleSolution` (line 142)
Changed `return true` to `return carry == 0` so that solutions where the column addition produces a leftover carry (overflow beyond the answer's digit count) are correctly rejected.

## Verification
- `go build ./...` - compiles successfully
- `go test -v -timeout 300s` - all 10 test cases pass (1.26s total)
- Commit: `fbebfed` on branch `issue-13`
