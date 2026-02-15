# Changes: Alphametics Solver Implementation

## Summary
Replaced the stub in `go/exercises/practice/alphametics/alphametics.go` with a complete alphametics puzzle solver.

## What was implemented

- **`problem` struct**: Stores parsed puzzle state including digit columns (`vDigits`), letter-to-value mappings (`letterValues`), and the set of unique letters used (`lettersUsed`).
- **`Solve(puzzle string) (map[string]int, error)`**: Entry point that parses the puzzle and delegates to the solver.
- **`parsePuzzle(puzzle string) *problem`**: Parses the input string (e.g. `"SEND + MORE == MONEY"`) into the `problem` struct, extracting words, validating uppercase letters, and building right-aligned digit columns.
- **`solvePuzzle() (map[string]int, error)`**: Iterates through all permutations of digit assignments and checks each for validity, rejecting solutions with leading zeros.
- **`isPuzzleSolution(values []int) bool`**: Validates a candidate digit assignment by summing columns with carry propagation and comparing against the result word.
- **`puzzleMap() map[string]int`**: Converts the internal letter-value array into the output map format.
- **`permutations(iterable []int, r int) [][]int`**: Generates r-length permutations of the input slice, modeled after Python's `itertools.permutations()`.

## Commit
`feat: implement alphametics solver for Go exercise` on branch `issue-3`
