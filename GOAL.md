# Goal: polyglot-go-error-handling (Issue #123)

## Problem Statement

Implement the `Use` function in `go/exercises/practice/error-handling/error_handling.go` that demonstrates Go error handling and resource management patterns including `defer`, `panic`, and `recover`.

The function signature is:
```go
func Use(opener ResourceOpener, input string) error
```

It must:
1. Open a resource using the provided `ResourceOpener`
2. Call `Frob(input)` on the opened resource
3. Close the resource in all cases (exactly once)
4. Properly handle errors and panics

## Acceptance Criteria

1. **Happy path**: When no errors occur, `Use` opens the resource, calls `Frob(input)`, closes the resource exactly once, and returns `nil`.

2. **Transient errors on open**: If `opener()` returns a `TransientError`, retry opening. Keep retrying until it succeeds or a non-transient error occurs.

3. **Non-transient errors on open**: If `opener()` returns a non-`TransientError` error, return that error immediately (do not retry).

4. **FrobError panic**: If `Frob` panics with a `FrobError`, call `Defrob(frobError.defrobTag)` on the resource, then close the resource, and return the error.

5. **Non-FrobError panic**: If `Frob` panics with a non-`FrobError` error, close the resource (do NOT call `Defrob`), and return the error.

6. **Close exactly once**: If the resource was successfully opened, `Close()` must be called exactly once regardless of errors or panics.

7. **All 5 test cases pass**: `TestNoErrors`, `TestKeepTryOpenOnTransient`, `TestFailOpenOnNonTransient`, `TestCallDefrobAndCloseOnFrobError`, `TestCallCloseOnNonFrobError`.

## Key Constraints

- Solution must be in package `erratum` in file `error_handling.go`
- Must use Go's `defer`, `recover` for panic handling
- Must use named return values to allow `recover` handler to set the error
- `Defrob` must be called before `Close` when handling `FrobError`
- Types `Resource`, `ResourceOpener`, `FrobError`, `TransientError` are defined in `common.go`
