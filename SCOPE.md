# Scope: polyglot-go-connect

## In Scope

- Implementing the `ResultOf` function in `go/exercises/practice/connect/connect.go`
- Any necessary helper types, functions, and constants in the same file
- Ensuring all tests in `connect_test.go` and `cases_test.go` pass
- Ensuring the benchmark runs without errors

## Out of Scope

- Modifying test files (`connect_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` files
- Implementing solutions for any other exercises
- Adding external dependencies
- Optimizing beyond what is needed to pass tests and benchmarks
- Adding CLI or main package functionality

## Dependencies

- Go standard library only (`errors`, `strings`, `fmt`)
- No external packages required
- The solution file must be in package `connect`
- Tests use `testing` and `strings` from standard library
