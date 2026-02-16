# Scope: polyglot-go-alphametics

## In Scope

- Implementing the `Solve` function in `go/exercises/practice/alphametics/alphametics.go`
- Parsing the puzzle string format: words joined by `+` and `==`
- Solving the constraint satisfaction problem: unique digit assignment per letter, no leading zeros, arithmetic holds
- Returning the correct `map[string]int` result or error
- Ensuring all 10 existing test cases pass

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Changes to any other exercise directories
- Benchmark optimization beyond passing the test timeout
- Supporting non-uppercase or non-alphabetic puzzle input

## Dependencies

- Go standard library only (no external packages)
- Existing test infrastructure in `alphametics_test.go` and `cases_test.go`
- Go 1.18+ toolchain
