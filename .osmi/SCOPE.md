# Scope: polyglot-go-alphametics

## In Scope

- Implementing the `Solve` function in `go/exercises/practice/alphametics/alphametics.go`
- Parsing the puzzle string format: `WORD + WORD + ... == WORD`
- Solving the constraint satisfaction problem (letter → digit mapping)
- Enforcing unique digit assignment per letter
- Enforcing no leading zeros on multi-digit words
- Returning appropriate errors for unsolvable puzzles
- Passing all existing test cases in `alphametics_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new dependencies
- Implementing solutions for other exercises
- Benchmarking or performance optimization beyond passing the 199-addend test

## Dependencies

- Go standard library only (no external packages)
- Existing test harness in `alphametics_test.go` and `cases_test.go`
- Go 1.18+ toolchain
