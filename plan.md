# Implementation Plan: polyglot-go-counter

## Proposal A

**Role: Proponent**

### Approach: Keep existing test suite as-is

The `counter_test.go` file already contains a comprehensive test suite with 8 test functions and a helper that correctly:
- Passes with Impl4 (correct implementation)
- Fails with Impl1 (detects incorrect line counting on 4 tests)
- Fails with Impl2 (detects ASCII-only letter detection via Unicode test)
- Fails with Impl3 (detects byte-level iteration via Unicode test)

The `counter.go` stub only needs `package counter` since there's no implementation code to write.

**Files to modify:** None — the solution is already complete.

**Rationale:** The existing test suite already satisfies all acceptance criteria. The tests are well-structured with a helper function `assertCounts()`, clear test names, and good coverage including edge cases (empty strings, only newlines, Unicode, multiple adds). Modifying working tests adds unnecessary risk.

**Changes:**
1. Verify `counter.go` has correct package declaration (already does)
2. Verify all tests pass with COUNTER_IMPL=4 (confirmed)
3. Verify all tests fail with COUNTER_IMPL=1,2,3 (confirmed)
4. Commit and push

---

## Proposal B

**Role: Opponent**

### Approach: Rewrite tests as table-driven with subtests

Replace the existing test suite with table-driven tests using `t.Run()` for subtests, following the pattern used in other exercises (e.g., bowling, book-store). This would consolidate the 8 separate test functions into 1-2 table-driven test functions.

**Weaknesses of Proposal A:**
- The test names use individual functions rather than the table-driven pattern common in Go
- No `t.Run()` subtests means less granular output
- The tests don't follow the exact same convention as other exercises in the repo

**Files to modify:** `counter_test.go` — rewrite with table-driven pattern.

**Changes:**
1. Define a `testCases` slice with description, inputs, expected lines/letters/chars
2. Use `t.Run(tc.description, ...)` for each case
3. Consolidate the `assertCounts()` helper into the loop body
4. Test `makeCounter()` separately for the no-add case

**Rationale:** Table-driven tests are more idiomatic Go and easier to extend. They produce cleaner output with `go test -v` and match the convention in bowling and other exercises.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion     | Proposal A (Keep as-is) | Proposal B (Rewrite as table-driven) |
|---------------|------------------------|--------------------------------------|
| Correctness   | Already verified passing all acceptance criteria | Would need re-verification; risk of introducing bugs |
| Risk          | Zero risk — no changes | Medium risk — rewriting working tests could introduce regressions |
| Simplicity    | Minimal effort, already done | Additional work with no functional benefit |
| Consistency   | Test structure is slightly different from table-driven exercises, but this exercise IS different (it's about writing tests, not code) | More consistent with other exercises' test patterns |

### Decision: Proposal A wins

**Rationale:** The existing test suite is already correct, comprehensive, and verified. The counter exercise is fundamentally different from other exercises — the tests ARE the solution, not boilerplate. Rewriting them in table-driven format adds risk for no functional benefit. The existing tests clearly detect all bugs in Impl1-3 while passing Impl4. The individual test function style is actually more appropriate here since each test targets a specific scenario that a correct counter must handle.

### Final Plan

1. **Verify `counter.go`** — Confirm it contains just `package counter` (already verified)
2. **Verify `counter_test.go`** — Confirm all tests pass with COUNTER_IMPL=4 and fail with COUNTER_IMPL=1,2,3 (already verified)
3. **Run `go vet`** — Confirm no issues (already verified)
4. **Create feature branch** `issue-329`
5. **Commit** the solution with descriptive message
6. **Push** to origin

No code changes are needed. The solution files are already correct.
