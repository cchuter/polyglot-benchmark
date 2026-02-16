# Scope: polyglot-go-error-handling (Issue #209)

## In Scope

- Implementing the `Use` function in `go/exercises/practice/error-handling/error_handling.go`
- Handling `TransientError` retries on opener
- Handling `FrobError` panics with `Defrob` calls
- Handling non-`FrobError` panics
- Guaranteed `Close()` via `defer`
- Passing all 5 existing tests

## Out of Scope

- Modifying `common.go` (type definitions are fixed)
- Modifying `error_handling_test.go` (test suite is fixed)
- Adding new test cases
- Modifying any other exercises in the repository
- Adding a `go.mod` file or changing build configuration

## Dependencies

- Go standard library only (`errors`, `io` packages used in common.go/tests)
- Types defined in `common.go`: `TransientError`, `FrobError`, `Resource`, `ResourceOpener`
- No external packages required
