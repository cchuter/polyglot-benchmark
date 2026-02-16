# Scope: polyglot-go-error-handling

## In Scope

- Implementing the `Use` function in `go/exercises/practice/error-handling/error_handling.go`
- Handling TransientError retries on opener
- Handling FrobError panics (with Defrob call)
- Handling non-FrobError panics
- Ensuring resource Close is called exactly once when opened
- Passing all existing tests

## Out of Scope

- Modifying `common.go` (read-only type definitions)
- Modifying `error_handling_test.go` (read-only test definitions)
- Modifying `go.mod`
- Adding new test cases
- Adding new type definitions

## Dependencies

- `common.go`: Provides `Resource`, `ResourceOpener`, `FrobError`, `TransientError` types
- `error_handling_test.go`: Provides the test suite
- Go standard library only (no external dependencies)
