# Scope: Alphametics Solver (Issue #233)

## In Scope

- Implementing the `Solve` function in `go/exercises/practice/alphametics/alphametics.go`
- Parsing the puzzle string format (`WORD + WORD + ... == WORD`)
- Finding a valid digit assignment for all letters
- Enforcing unique digit per letter constraint
- Enforcing no leading zeros constraint
- Returning appropriate errors for unsolvable puzzles
- Passing all provided test cases including the extreme 199-addend puzzle

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new files to the exercise directory
- Supporting non-standard puzzle formats
- Performance benchmarking beyond passing the test suite
- Changes to any other exercise or language directory

## Dependencies

- Go standard library only (no external packages)
- `strings`, `errors`, `unicode` packages (as used in the reference solution)
- Go 1.18 compatibility
