# Implementation Plan: polyglot-go-error-handling

## File to Modify

- `go/exercises/practice/error-handling/error_handling.go`

No other files need to be created or modified.

## Implementation

Replace the stub in `error_handling.go` with the `Use` function implementation:

```go
package erratum

func Use(opener ResourceOpener, input string) (err error) {
	// 1. Open the resource, retrying on TransientError
	var r Resource
	for {
		r, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	// 2. Ensure Close is called exactly once after successful open
	defer r.Close()

	// 3. Recover from panics in Frob
	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			err = x.(error)
		}
	}()

	// 4. Call Frob
	r.Frob(input)

	return nil
}
```

## Architectural Decisions

1. **Named return value `err`**: Required so the deferred `recover` function can set the return value.

2. **Infinite loop for TransientError retry**: The spec says to keep retrying on TransientError with no limit. The loop breaks on success or non-transient error.

3. **Defer ordering**: `defer r.Close()` is registered first, `defer func() { recover... }()` second. Since defers execute LIFO, the recover handler runs first (handling Defrob if needed), then Close runs. This ensures Defrob is called before Close.

4. **Type assertion on panic value**: `x.(FrobError)` checks if the panic is a FrobError. `x.(error)` converts the panic value to an error for the return.

## Approach and Ordering

1. Write the `Use` function in `error_handling.go`
2. Run `go test` to verify all 5 tests pass
3. Commit the change
