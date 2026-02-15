# Scope: polyglot-go-alphametics (Issue #13)

## In Scope

- Implementation of `func Solve(puzzle string) (map[string]int, error)` in `go/exercises/practice/alphametics/alphametics.go`
- Parsing puzzle strings with `+` and `==` operators
- Solving the alphametics constraint satisfaction problem
- Validating uniqueness of digit assignments
- Enforcing no-leading-zero constraint on all multi-digit words
- Passing all test cases in `alphametics_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` directories
- Adding new test cases
- Optimizing beyond what is needed to pass the 199-addend test case within timeout
- Changes to any other exercise directories

## Dependencies

- Go 1.18+ toolchain
- Standard library only (no external packages)
- Test framework: `testing` package (standard library)
