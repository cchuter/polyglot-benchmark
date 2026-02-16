# Scope: Alphametics Puzzle Solver

## In Scope

- Implementing the `Solve` function in `go/exercises/practice/alphametics/alphametics.go`
- Parsing puzzle strings (handling `+`, `==`, and whitespace)
- Finding digit assignments that satisfy the equation
- Enforcing unique digit constraint and no-leading-zero constraint
- Returning appropriate errors for unsolvable puzzles
- Passing all existing test cases including the large 199-addend puzzle

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding external dependencies
- Changing the function signature
- Supporting puzzle formats other than what tests expect

## Dependencies

- Go standard library only (no external packages)
- Existing test infrastructure in the repository
