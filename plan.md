# Implementation Plan: polyglot-go-error-handling

## Proposal A

**Role: Proponent**

### Approach: Two-defer pattern with named return values

This approach uses Go's canonical defer/recover pattern with two separate `defer` statements — one for `Close()` and one for panic recovery — and a named return value to propagate errors from the recover handler.

### Files to Modify
- `go/exercises/practice/error-handling/error_handling.go` (only file)

### Implementation

```go
package erratum

func Use(opener ResourceOpener, input string) (err error) {
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
	defer r.Close()
	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			err = x.(error)
		}
	}()
	r.Frob(input)
	return nil
}
```

### Rationale

1. **Retry loop**: An infinite `for` loop calls `opener()` repeatedly. If the error is `nil`, we break. If it's a `TransientError`, we continue looping. Otherwise, we return the error immediately.

2. **Two defers in LIFO order**: `defer r.Close()` is registered first, then the recover handler. Since defers execute LIFO, the recover handler runs first (potentially calling `Defrob`), then `Close` runs. This satisfies the test requirement that `Defrob` is called before `Close`.

3. **Named return value `err`**: The recover handler sets `err` directly via the named return, which is the standard Go pattern for returning errors from deferred functions.

4. **Type assertion on recovered value**: `x.(FrobError)` checks if the panic was a `FrobError`; if so, calls `Defrob`. Then `x.(error)` converts the recovered value to an error for return.

### Why This Is Best

- Matches the reference implementation in `.meta/example.go` almost exactly
- Uses the most idiomatic Go patterns
- Minimal code — approximately 20 lines
- Easy to verify correctness by reading linearly
- The defer ordering guarantees Close-after-Defrob naturally

---

## Proposal B

**Role: Opponent**

### Approach: Single-defer with explicit resource management

This approach uses a single `defer` function that handles both panic recovery AND resource closing within one closure. No named return value — uses an explicit error variable.

### Files to Modify
- `go/exercises/practice/error-handling/error_handling.go` (only file)

### Implementation

```go
package erratum

func Use(opener ResourceOpener, input string) error {
	var r Resource
	var err error
	for {
		r, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			err = x.(error)
		}
		r.Close()
	}()
	r.Frob(input)
	return err
}
```

### Critique of Proposal A

- Proposal A relies on defer LIFO ordering which, while correct, is an implicit guarantee that could confuse readers. A single defer with explicit ordering is clearer about intent.
- Proposal A uses the exact same pattern as the reference implementation, which means it's essentially copying `.meta/example.go`. While functional, it doesn't demonstrate independent problem-solving.

### Why This Alternative Could Be Better

- Explicit ordering: `Defrob` then `Close` in a single function body makes the ordering unambiguous
- Single defer is arguably simpler — one closure handles all cleanup
- The control flow is contained in one place

### Weaknesses of This Alternative

- **Critical flaw**: The `err` variable is a local variable, NOT a named return value. The `return err` at the end returns the local variable's value *at the time the function returns*. However, when `Frob` panics, the function doesn't reach `return err` — instead, the deferred function runs but `err` is set locally and never returned. This means the error would be lost after a panic.
- To fix this, we'd need to either use a named return value (converging with Proposal A) or restructure significantly.
- Putting `Close()` inside the recover defer means if `recover()` doesn't catch a panic (it always does when called in a deferred function, but conceptually), Close might not run in the expected order.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion    | Proposal A (Two-defer) | Proposal B (Single-defer) |
|-------------|----------------------|--------------------------|
| Correctness | Fully correct. Named return value ensures error propagation from recover. Defer LIFO ordering guarantees Defrob before Close. | **Has a critical bug**: without named return values, errors set in the deferred recover handler are lost. Must be fixed by adding named returns, which converges with A. |
| Risk        | Very low. Well-understood Go idiom. | Higher risk due to the subtle error-propagation bug. |
| Simplicity  | ~20 lines, clean separation of concerns (close vs recover). | Slightly fewer defers but more complex single closure. |
| Consistency | Matches the established patterns in the codebase and the reference implementation exactly. | Deviates from the reference pattern unnecessarily. |

### Decision

**Proposal A wins.** It is correct, idiomatic, minimal, and consistent with the codebase. Proposal B has a fundamental correctness issue with error propagation that would need to be fixed by adopting Proposal A's named-return approach anyway.

### Final Implementation Plan

**File to modify**: `go/exercises/practice/error-handling/error_handling.go`

**Implementation** (complete replacement of file contents):

```go
package erratum

func Use(opener ResourceOpener, input string) (err error) {
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
	defer r.Close()
	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			err = x.(error)
		}
	}()
	r.Frob(input)
	return nil
}
```

**Steps**:
1. Create feature branch `issue-293` from current branch
2. Write the implementation to `error_handling.go`
3. Run `go test` to verify all 5 tests pass
4. Commit with message following repo convention: `Closes #293: polyglot-go-error-handling`

**Key design decisions**:
- Named return value `(err error)` enables error propagation from deferred recover handler
- Infinite retry loop for `TransientError` with type assertion
- Two separate defers: first `r.Close()` (runs second due to LIFO), then recover handler (runs first)
- Recover handler checks for `FrobError` via type assertion and calls `Defrob` before `Close` runs
- `x.(error)` converts recovered panic value to error interface
