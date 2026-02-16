# Verification Report: polyglot-go-error-handling

## Verdict: PASS

## Build Verification

- `go build ./...` completes with zero errors.
- `go test -v ./...` completes with zero errors.

## Test Results (5/5 PASS)

| Test | Result |
|------|--------|
| TestNoErrors | PASS |
| TestKeepTryOpenOnTransient | PASS |
| TestFailOpenOnNonTransient | PASS |
| TestCallDefrobAndCloseOnFrobError | PASS |
| TestCallCloseOnNonFrobError | PASS |

## Acceptance Criteria Verification

### 1. Happy path
**PASS** - `Use` opens the resource via `opener()`, calls `Frob(input)`, defers `Close()` exactly once, and returns `nil` when no errors occur. Verified by `TestNoErrors`.

### 2. Transient errors on open
**PASS** - The `for` loop (lines 6-14) retries `opener()` whenever the error is a `TransientError`. It continues looping until success or a non-transient error. Verified by `TestKeepTryOpenOnTransient` (retries 3 times before success).

### 3. Non-transient errors on open
**PASS** - When `opener()` returns a non-`TransientError`, the type assertion `err.(TransientError)` fails and the function returns the error immediately (line 12). Verified by `TestFailOpenOnNonTransient`.

### 4. FrobError panic
**PASS** - The deferred `recover()` function (lines 20-27) catches panics from `Frob`. If the panic value is a `FrobError`, it calls `r.Defrob(frobErr.defrobTag)` before `Close()` runs. Due to LIFO defer ordering, the recover defer executes before the Close defer, ensuring Defrob is called first. The error is returned via the named return value. Verified by `TestCallDefrobAndCloseOnFrobError` (which also asserts Close was not called before Defrob).

### 5. Non-FrobError panic
**PASS** - When `Frob` panics with a non-`FrobError` value, the recover function catches it, skips the Defrob call (the type assertion fails), and sets `err = x.(error)`. Close is still called via defer. Verified by `TestCallCloseOnNonFrobError` (which asserts Defrob was NOT called).

### 6. Close exactly once
**PASS** - `Close()` is called via a single `defer r.Close()` (line 17), placed after the resource is successfully opened and before any code that could panic. This guarantees exactly one call. Both `TestNoErrors` and `TestCallDefrobAndCloseOnFrobError` and `TestCallCloseOnNonFrobError` explicitly verify `closeCallsCount == 1`.

### 7. Implementation quality
- Package is `erratum` as required.
- Uses `defer` for Close and recovery.
- Uses named return value `err` to allow the recover handler to set the error.
- Defrob is called before Close due to correct LIFO defer ordering.
- Code is clean, minimal, and idiomatic Go.
