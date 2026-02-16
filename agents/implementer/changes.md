# Changes: error-handling exercise

## Modified File
- `go/exercises/practice/error-handling/error_handling.go`

## What Was Done
Implemented the `Use` function in the `erratum` package. The function:

1. **Opens a resource** via `ResourceOpener`, retrying in an infinite loop on `TransientError` and returning immediately on any other error.
2. **Defers `r.Close()`** after a successful open to ensure the resource is always closed.
3. **Defers a `recover()` handler** that catches panics from `Frob`. If the panic is a `FrobError`, it calls `r.Defrob(frobErr.defrobTag)` before setting the named return `err` to the panic value cast as `error`.
4. **Calls `r.Frob(input)`** to perform the operation.

Uses a named return value `(err error)` so the deferred recover function can set the return value. Defer ordering ensures recover runs before Close (LIFO).

## Commit
`4484c54` - "Implement Use function for error-handling exercise"
