# Implementation Plan: polyglot-go-error-handling

## Branch 1: Minimal defer/recover with named return (Simplicity)

**Approach**: Use a named return value `err`, a `for` loop for transient retries, `defer Close()`, and a separate `defer func()` with `recover()` that handles both FrobError and other panics.

**File to modify**: `go/exercises/practice/error-handling/error_handling.go`

**Implementation**:
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

**Rationale**: This is the canonical Go solution. Named return `err` allows the deferred recovery function to set the return value. Two defers ensure Close runs after recovery (LIFO order: recovery runs first, then Close). The for loop handles transient retries.

**Evaluation**:
- Feasibility: High - straightforward idiomatic Go
- Risk: Low - well-established pattern
- Alignment: Fully satisfies all 7 acceptance criteria
- Complexity: Minimal - single file, ~20 lines

## Branch 2: Explicit close management without defer for Close (Extensibility)

**Approach**: Use defer only for recover, manage Close explicitly in each code path to give maximum control over ordering.

**Implementation**:
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
    defer func() {
        if x := recover(); x != nil {
            if frobErr, ok := x.(FrobError); ok {
                r.Defrob(frobErr.defrobTag)
            }
            err = x.(error)
            r.Close()
        }
    }()
    r.Frob(input)
    r.Close()
    return nil
}
```

**Rationale**: More explicit about when Close happens. Could be extended to handle Close errors differently in panic vs normal paths.

**Evaluation**:
- Feasibility: High - valid Go
- Risk: Medium - Close is called in two places (DRY violation), easy to forget a path. Also, test `TestCallDefrobAndCloseOnFrobError` checks that Close is called AFTER Defrob but within the same flow - this works but is fragile.
- Alignment: Satisfies criteria but the double-Close-call pattern is error-prone
- Complexity: Slightly higher maintenance burden

## Branch 3: Channel-based or helper-function abstraction (Unconventional)

**Approach**: Extract retry logic into a helper and use a wrapping function for the Frob+recover logic.

**Implementation**:
```go
package erratum

func openWithRetry(opener ResourceOpener) (Resource, error) {
    for {
        r, err := opener()
        if err == nil {
            return r, nil
        }
        if _, ok := err.(TransientError); !ok {
            return nil, err
        }
    }
}

func Use(opener ResourceOpener, input string) (err error) {
    r, err := openWithRetry(opener)
    if err != nil {
        return err
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

**Rationale**: Separates concerns - retry logic is isolated. More testable individually if needed.

**Evaluation**:
- Feasibility: High - valid Go
- Risk: Low - same core pattern as Branch 1
- Alignment: Fully satisfies all criteria
- Complexity: Higher - adds a helper function that's unnecessary for this small scope. Over-engineered for the exercise.

---

## Selected Plan

**Branch 1** is selected.

**Rationale**: Branch 1 is the canonical, idiomatic Go solution. It's the simplest approach with the least code, uses well-established Go patterns (named returns, defer, recover), and fully satisfies all acceptance criteria. Branch 2 introduces a DRY violation and fragility. Branch 3 over-engineers with an unnecessary helper function. Since this is an Exercism exercise, the idiomatic approach is the correct one.

### Detailed Implementation Steps

1. **Edit** `go/exercises/practice/error-handling/error_handling.go`:
   - Replace the stub `package erratum` with the full `Use` function
   - Use named return value `(err error)` to allow deferred recovery to set the return
   - Implement retry loop for `TransientError` on `opener()`
   - `defer r.Close()` to guarantee cleanup
   - `defer func() { recover... }()` to handle panics from `Frob`
   - In recovery: type-assert for `FrobError` to call `Defrob`, then set `err = x.(error)`
   - Call `r.Frob(input)` and return nil on success

2. **Verify** by running `go test` in the exercise directory

3. **Commit** with message: `Closes #209: polyglot-go-error-handling`

### Key Design Decisions

- **Defer ordering**: `defer r.Close()` is declared before `defer func() { recover... }()`. Since defers execute LIFO, the recovery function runs first (handling Defrob if needed), then Close runs. This satisfies the test requirement that Defrob is called before Close.
- **Named return**: Essential for the deferred function to modify the return value
- **Infinite loop for retry**: Standard pattern for transient error retry without backoff (matches test expectations)
- **Type assertion in recover**: `x.(error)` works because all test panics use `error` types
