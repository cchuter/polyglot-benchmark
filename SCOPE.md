# Scope: polyglot-go-alphametics

## In Scope

- Implementing the `Solve` function in `go/exercises/practice/alphametics/alphametics.go`
- Parsing the puzzle string to extract operand words and the result word
- Implementing a solving algorithm (permutation-based or constraint-based)
- Handling all edge cases: unique digit constraint, leading zero constraint, no-solution case
- Ensuring all 10 test cases pass
- Performance: the solution must handle the 199-addend puzzle in reasonable time

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.docs/` or `.meta/` content
- Adding new files beyond `alphametics.go`
- Implementing solutions for other exercises
- Changes to non-Go code in the repository

## Dependencies

- Go standard library only (no external packages)
- Packages likely needed: `errors`, `strings`, `unicode`
- No dependencies on other exercises or external services
