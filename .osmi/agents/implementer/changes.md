# Changes

## Implement `Use` function for error-handling exercise

- Added `Use(opener ResourceOpener, input string) (err error)` in `error_handling.go`
- Retries `opener()` in a loop when `TransientError` is returned; returns immediately on other errors
- Defers `r.Close()` after successful open
- Deferred recover handler checks for `FrobError` (calls `r.Defrob`) and propagates panic as error via named return
- Calls `r.Frob(input)` and returns nil on success
- All 5 tests pass
