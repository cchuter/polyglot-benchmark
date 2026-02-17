# Goal: polyglot-go-error-handling (Issue #293)

## Problem Statement

Implement the `Use` function in `go/exercises/practice/error-handling/error_handling.go` for the Exercism "error-handling" exercise. The function must demonstrate proper Go error handling, panic recovery, and resource management patterns.

The function signature is:
```go
func Use(opener ResourceOpener, input string) error
```

## Acceptance Criteria

1. **All 5 tests pass** (`go test` in the exercise directory returns PASS):
   - `TestNoErrors` — happy path: opens resource, calls `Frob(input)`, closes resource exactly once, returns nil
   - `TestKeepTryOpenOnTransient` — retries `opener()` on `TransientError` until it succeeds
   - `TestFailOpenOnNonTransient` — returns non-transient errors from `opener()` immediately
   - `TestCallDefrobAndCloseOnFrobError` — on `FrobError` panic from `Frob`: calls `Defrob(defrobTag)` before `Close`, closes exactly once, returns the error
   - `TestCallCloseOnNonFrobError` — on non-`FrobError` panic from `Frob`: calls `Close` exactly once (no `Defrob`), returns the error

2. **Resource is always closed** exactly once if it was successfully opened, regardless of errors or panics.

3. **`Defrob` is called before `Close`** when a `FrobError` panic occurs (the test verifies ordering).

4. **Only `error_handling.go` is modified** — `common.go`, test files, and `.meta/` are read-only.

## Key Constraints

- Package must be `erratum`
- Must use Go idioms: named return values, `defer`, `recover()`, type assertions
- No external dependencies (standalone module with `go 1.18`)
- `common.go` provides `Resource`, `ResourceOpener`, `TransientError`, and `FrobError` types — do not modify
