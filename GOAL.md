# Goal: polyglot-go-error-handling (Issue #209)

## Problem Statement

Implement the `Use` function in `go/exercises/practice/error-handling/error_handling.go` that demonstrates Go error handling and resource management patterns including error type checking, panic recovery, and guaranteed resource cleanup.

The function signature is:
```go
func Use(opener ResourceOpener, input string) error
```

## Acceptance Criteria

1. **Happy path**: When `opener()` succeeds and `Frob(input)` succeeds, `Use` returns `nil` and `Close()` is called exactly once.

2. **Transient error retry**: When `opener()` returns a `TransientError`, `Use` retries the open operation until it succeeds or a non-transient error occurs.

3. **Non-transient error on open**: When `opener()` returns a non-`TransientError` error, `Use` returns that error immediately (no retry, no Close needed since resource wasn't opened).

4. **FrobError panic handling**: When `Frob(input)` panics with a `FrobError`, `Use` must:
   - Call `Defrob(frobError.defrobTag)` on the resource
   - Call `Close()` exactly once (after Defrob)
   - Return the `FrobError` as an error

5. **Non-FrobError panic handling**: When `Frob(input)` panics with a non-`FrobError` error, `Use` must:
   - NOT call `Defrob`
   - Call `Close()` exactly once
   - Return the panic value as an error

6. **Resource cleanup guarantee**: If the resource was successfully opened, `Close()` must be called exactly once regardless of what happens during `Frob`.

7. **All 5 tests pass**: `TestNoErrors`, `TestKeepTryOpenOnTransient`, `TestFailOpenOnNonTransient`, `TestCallDefrobAndCloseOnFrobError`, `TestCallCloseOnNonFrobError`.

## Key Constraints

- Solution must be in package `erratum` (same package as `common.go` and tests)
- Must use Go's `defer`, `recover`, and named return values pattern
- Must not modify `common.go` or `error_handling_test.go`
- Only `error_handling.go` should be modified
