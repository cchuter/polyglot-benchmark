# Implementation Plan: error-handling

## Overview

Implement the `Use` function in `error_handling.go`. The function follows a standard Go pattern for error handling with defer/panic/recover.

## File to Modify

- `go/exercises/practice/error-handling/error_handling.go`

## Implementation

The `Use` function signature uses a named return value `err error` to allow the deferred recover function to set the return value.

### Algorithm

```
func Use(opener ResourceOpener, input string) (err error):
  1. Open loop: call opener() in a loop
     - If err is nil, break (success)
     - If err is TransientError, continue looping (retry)
     - If err is any other error, return it immediately
  2. defer r.Close() — ensures Close is called exactly once
  3. defer recover handler:
     - If recover() returns non-nil (a panic occurred):
       - Type-assert to FrobError: if it is, call r.Defrob(frobErr.defrobTag)
       - Type-assert the panic value to error and assign to named return `err`
  4. Call r.Frob(input)
  5. Return nil
```

### Key Design Decisions

1. **Named return value**: Using `(err error)` allows the deferred recover function to modify the return value after a panic
2. **Defer ordering**: `defer r.Close()` is registered first, then the recover handler. Since defers execute LIFO, the recover handler runs first (handling Defrob if needed), then Close runs. This ensures Defrob is called before Close.
3. **Type assertion for TransientError**: Use `_, ok := err.(TransientError)` to check if the error is transient
4. **Type assertion for FrobError in panic**: Use `frobErr, ok := x.(FrobError)` to check the panic value
5. **Infinite retry loop**: The spec says to keep trying on TransientError; the tests show a finite number of transient errors before success

## Verification

- `cd go/exercises/practice/error-handling && go test ./...`
- `cd go/exercises/practice/error-handling && go vet ./...`
