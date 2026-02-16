# Scope: polyglot-go-alphametics

## In Scope

- Implement `func Solve(puzzle string) (map[string]int, error)` in `go/exercises/practice/alphametics/alphametics.go`
- Parse the puzzle string format: `"WORD + WORD + ... == RESULT"`
- Solve the constraint satisfaction problem: assign unique digits 0-9 to letters
- Enforce the leading-zero constraint
- Return error for unsolvable puzzles
- Pass all existing test cases

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`, `go.mod`)
- Adding new test cases
- Implementing solutions for other exercises
- Changing the package structure or module configuration

## Dependencies

- Go standard library only (no external packages)
- Existing test harness in `alphametics_test.go` and `cases_test.go`
