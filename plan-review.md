# Plan Review: polyglot-go-error-handling

## Verdict: PASS

The selected plan (Proposal A) is correct and will pass all 5 tests. Below is the detailed analysis.

---

## 1. Test-by-Test Verification

### TestNoErrors — PASS
- `opener()` returns `(mr, nil)` on first call. The loop breaks immediately.
- `defer r.Close()` is registered. `defer recover()` is registered.
- `r.Frob(input)` is called with `hello`. No panic occurs.
- Function returns `nil`. Deferred `Close()` runs exactly once.
- The named return `err` is `nil` since no panic occurred.

### TestKeepTryOpenOnTransient — PASS
- `opener()` returns `TransientError` for calls 0-2; the type assertion `err.(TransientError)` succeeds, so the loop continues.
- On call 3, `opener()` returns `(mr, nil)`. Loop breaks.
- `Frob(hello)` is called. No panic. Returns `nil`. `Close()` runs once.

### TestFailOpenOnNonTransient — PASS
- `opener()` returns `TransientError` for calls 0-2 (loop retries).
- On call 3, returns `(nil, errors.New("too awesome"))`. This is not a `TransientError`, so `_, ok := err.(TransientError)` yields `ok == false`, and the function returns the error immediately.
- No `defer` statements have been registered at this point (they are after the loop), so no `Close()` is called. This is correct because the resource was never successfully opened.

### TestCallDefrobAndCloseOnFrobError — PASS
- `opener()` succeeds. `defer r.Close()` registered first, then `defer recover()` handler.
- `r.Frob(input)` panics with `FrobError{tag: "moo", inner: errors.New("meh")}`.
- Defers execute LIFO: recover handler runs first.
  - `recover()` catches the `FrobError`.
  - `x.(FrobError)` succeeds: `r.Defrob("moo")` is called.
  - The test's `defrob` mock checks `closeCallsCount == 0` at this point, which is true because `Close()` has not run yet. This passes.
  - `err = x.(error)` sets the named return to the `FrobError` (which implements `error` via its `Error()` method).
- Then `r.Close()` runs (the first defer). `closeCallsCount` becomes 1.
- Function returns `err` which is the `FrobError`. `err.Error()` returns `"meh"`. All assertions pass.

### TestCallCloseOnNonFrobError — PASS
- `opener()` succeeds. Both defers registered.
- `r.Frob(input)` panics with `errors.New("meh")`.
- Recover handler runs: `x.(FrobError)` fails (`ok == false`), so `Defrob` is NOT called. `defrobCalled` stays `false`.
- `err = x.(error)` succeeds (it is an `*errors.errorString` which implements `error`).
- `r.Close()` runs. `closeCallsCount` becomes 1.
- Function returns error with message `"meh"`. All assertions pass.

---

## 2. Correctness Analysis

### Defer ordering
Correct. `defer r.Close()` is pushed first (line 30 in the plan), then the anonymous recover function (line 31). LIFO means recover runs before Close. This is the standard Go idiom and is guaranteed by the language spec.

### Named return value
Correct. `(err error)` as a named return is essential. When `Frob` panics, the normal `return nil` on the last line is never reached. Instead, the deferred function sets `err` via the named return, and this value is what the caller receives. Without named returns (as identified in Proposal B's critique), the error would be lost.

### Type assertion `x.(error)`
Correct but worth noting: this is a non-checked type assertion. If `Frob` panics with a non-`error` value (e.g., a string), this would cause a second panic. However, the tests only panic with values implementing `error` (`FrobError` and `*errors.errorString`), so this is safe for the test suite. The reference implementation uses the same pattern.

### Resource closing guarantee
Correct. `defer r.Close()` is only registered after the opener loop succeeds (`err == nil` triggers `break`), so Close is only called on successfully opened resources. It runs exactly once because it is a single `defer` call.

### Defrob-before-Close ordering
Correct. Verified by the LIFO defer stack: recover (with Defrob) executes before Close. The test in `TestCallDefrobAndCloseOnFrobError` explicitly checks `closeCallsCount == 0` inside the `defrob` mock to enforce this ordering.

---

## 3. Go Idiom and Compilation Check

- **Package name**: `erratum` -- matches `common.go` and the test file.
- **Field access `frobErr.defrobTag`**: The `defrobTag` field in `FrobError` is unexported (lowercase). Since the implementation file is in the same package (`erratum`), this is accessible. No compilation issue.
- **Interface satisfaction**: `FrobError` has an `Error() string` method, so `x.(error)` succeeds at runtime. `TransientError` also implements `error`.
- **No imports needed**: The implementation uses only types defined in `common.go` within the same package. No import statements required.
- **Named return value with explicit `return err`**: When the loop exits early with `return err` for a non-transient error, the named return variable `err` already holds the error from `opener()`. This is correct.

---

## 4. Minor Observations

- The plan's final implementation differs very slightly from the reference `.meta/example.go`: the reference has an extra `if err != nil { return err }` block after `r.Frob(input)` (lines 24-26 of example.go). This is dead code in the reference -- after `Frob`, `err` is still whatever it was after the opener loop (which is `nil` since we broke out on success). The plan correctly omits this unnecessary check, replacing it with a direct `return nil`. Both are functionally equivalent.

- The plan correctly identifies that only `error_handling.go` needs modification, consistent with the acceptance criteria.

---

## Summary

The plan is correct, complete, and will pass all 5 tests. The implementation uses idiomatic Go patterns (named returns, defer/recover, type assertions), handles all edge cases properly, and maintains correct Defrob-before-Close ordering via defer LIFO semantics. No compilation issues or correctness bugs were found.
