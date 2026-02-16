# Scope: Alphametics Puzzle Solver

## In Scope

- Implementing the `Solve` function in `go/exercises/practice/alphametics/alphametics.go`
- Parsing the puzzle string (splitting on `+` and `==`)
- Assigning unique digits to letters via backtracking/permutation
- Enforcing the no-leading-zeros constraint
- Returning the correct `map[string]int` or an error
- Ensuring all test cases pass including the large 199-addend puzzle

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding external dependencies
- Changing the function signature
- Supporting puzzle formats beyond the `WORD + WORD == WORD` pattern with `+` and `==` operators
- Benchmark optimization beyond passing all tests in reasonable time

## Dependencies

- Go standard library only (no external packages)
- Existing test infrastructure (`go test`)
- Reference implementation available at `.meta/example.go` for guidance on approach
