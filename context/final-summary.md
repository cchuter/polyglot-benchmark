# Context: polyglot-go-error-handling (Issue #293)

## Status: Complete

## Branch: issue-293 (pushed to origin)

## What Was Done
Implemented the `Use(opener ResourceOpener, input string) error` function in `go/exercises/practice/error-handling/error_handling.go` for the Exercism error-handling exercise.

## Implementation Details
The function uses three key Go patterns:
1. **Retry loop**: Calls `opener()` in an infinite loop, retrying on `TransientError`, returning immediately on other errors
2. **Two defers**: `defer r.Close()` registered first, then a deferred anonymous function with `recover()` — LIFO ordering ensures Defrob runs before Close
3. **Named return value**: `(err error)` allows the deferred recover handler to set the return value after a panic

## Test Results
All 5 tests pass:
- TestNoErrors
- TestKeepTryOpenOnTransient
- TestFailOpenOnNonTransient
- TestCallDefrobAndCloseOnFrobError
- TestCallCloseOnNonFrobError

## Files Modified
- `go/exercises/practice/error-handling/error_handling.go` (only file)

## Commit
`bcfc1ee Implement Use function for error-handling exercise`
