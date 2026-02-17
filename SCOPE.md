# Scope: polyglot-go-error-handling (Issue #293)

## In Scope

- Implement the `Use` function in `go/exercises/practice/error-handling/error_handling.go`
- The implementation must handle:
  - Retry loop for `TransientError` from `ResourceOpener`
  - Immediate return of non-transient errors from `ResourceOpener`
  - Panic recovery from `Frob` calls
  - `FrobError`-specific handling (call `Defrob` with `defrobTag`)
  - Guaranteed `Resource.Close()` via `defer`
- All 5 existing tests must pass

## Out of Scope

- Modifying `common.go` (editor/support file)
- Modifying `error_handling_test.go` (test file)
- Modifying `.meta/` directory contents
- Modifying `go.mod`
- Adding new test cases
- Adding new files to the exercise directory
- Changes to any other exercise or directory

## Dependencies

- `common.go` — provides `Resource` interface, `ResourceOpener` type, `TransientError` struct, `FrobError` struct
- `io.Closer` — standard library interface embedded in `Resource`
- Go 1.18+ toolchain for building and testing
