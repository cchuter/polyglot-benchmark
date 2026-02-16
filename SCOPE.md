# Scope: Alphametics Puzzle Solver (Issue #147)

## In Scope

- Implement `func Solve(puzzle string) (map[string]int, error)` in `go/exercises/practice/alphametics/alphametics.go`
- Parse puzzle input string (words separated by `+` and `==`)
- Solve the alphametics constraint satisfaction problem
- Enforce unique digit assignment per letter
- Enforce no leading zeros on multi-digit numbers
- Return error for unsolvable puzzles
- Pass all existing test cases in `alphametics_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Supporting non-standard puzzle formats
- Supporting lowercase letters
- Adding external dependencies

## Dependencies

- Go standard library only (`errors`, `strings`, `unicode`)
- No external packages
- Existing test harness in `alphametics_test.go` and `cases_test.go`
