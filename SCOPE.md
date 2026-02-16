# Scope: polyglot-go-error-handling (Issue #123)

## In Scope

- Implement the `Use` function in `go/exercises/practice/error-handling/error_handling.go`
- Ensure all 5 existing test cases pass
- Follow Go idiomatic patterns for error handling (`defer`, `panic`, `recover`, named returns)

## Out of Scope

- Modifying `common.go` (type definitions are read-only)
- Modifying `error_handling_test.go` (test cases are fixed)
- Modifying `go.mod`
- Adding new test cases
- Changes to any other exercise or language directory

## Dependencies

- `common.go` defines: `TransientError`, `FrobError`, `Resource` interface, `ResourceOpener` type
- `io.Closer` interface (standard library) - embedded in `Resource`
- No external package dependencies (only stdlib `io`)
