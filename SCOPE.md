# Scope: polyglot-go-alphametics

## In Scope

- Implementing `func Solve(puzzle string) (map[string]int, error)` in `go/exercises/practice/alphametics/alphametics.go`
- Parsing puzzle strings with `+` and `==` operators
- Solving the constraint satisfaction problem: assigning unique digits to letters
- Enforcing the no-leading-zero constraint for multi-digit numbers
- Returning appropriate errors for unsolvable puzzles
- Passing all 10 test cases in `cases_test.go`

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new dependencies
- Changes to any files outside `go/exercises/practice/alphametics/alphametics.go`
- Optimizing beyond what is needed to pass all tests within reasonable time
- Supporting puzzle formats other than what the tests specify

## Dependencies

- Go standard library only (no external packages)
- Go 1.18+ (as specified in `go.mod`)
- Reference solution available in `.meta/example.go` for guidance
