# Change Log: error-handling

## Changes Made

### `go/exercises/practice/error-handling/error_handling.go`
- Implemented `Use(opener ResourceOpener, input string) error` function
- Retry loop for `TransientError` on resource open
- `defer r.Close()` ensures resource cleanup in all cases
- `defer func()` with `recover()` handles panics from `Frob`:
  - `FrobError`: calls `Defrob(defrobTag)` then returns error
  - Other errors: returns error directly
- LIFO defer ordering ensures Defrob runs before Close

## Test Results
- All 5 tests pass
- `go vet` clean
