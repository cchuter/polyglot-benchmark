# Implementation Plan: polyglot-go-counter

## Proposal A

**Role: Proponent**

### Approach: Accept the existing test suite as-is

The `counter_test.go` already contains a comprehensive, well-designed test suite with 9 tests that correctly:
- Pass for Impl4 (correct implementation)
- Fail for Impl1 (4 test failures detecting line counting bug)
- Fail for Impl2 (1 test failure detecting ASCII-only letter counting)
- Fail for Impl3 (1 test failure detecting byte-level iteration)

**Files to modify:** None. The code is already correct.

**Files to verify:**
- `go/exercises/practice/counter/counter_test.go` - already has complete test suite
- `go/exercises/practice/counter/counter.go` - already has correct `package counter`

**Rationale:**
- The test suite is already comprehensive and correct
- All acceptance criteria are already met (verified by running tests)
- The simplest approach is the best - don't fix what isn't broken
- No risk of introducing regressions by making unnecessary changes

**Steps:**
1. Create feature branch `issue-286`
2. Verify tests pass/fail correctly for all 4 implementations
3. Commit and push (only `.osmi/` artifacts change)
4. Create PR

**Why this is best:** Zero code risk. The solution is already in place. Adding unnecessary modifications would only introduce risk of breaking something.

---

## Proposal B

**Role: Opponent**

### Approach: Enhance the test suite with additional edge cases

While the existing tests work, they could be strengthened:

**Files to modify:**
- `go/exercises/practice/counter/counter_test.go` - add additional test cases

**Proposed additions:**
1. Test with tab characters and other whitespace
2. Test with emoji characters (multi-byte but not letters)
3. Test with very long strings
4. Test with carriage return + line feed (`\r\n`)
5. Add table-driven test structure for better organization

**Critique of Proposal A:**
- The existing suite only has 1 test (`TestUnicodeLetters`) that catches Impl2 and Impl3 bugs separately. If that single test were wrong, both bugs would go undetected.
- No tests for `\r\n` line endings or tab characters
- No tests for emoji (4-byte UTF-8 sequences)

**Rationale:**
- More tests = higher confidence
- Table-driven tests match Go conventions better
- Additional Unicode tests would make Impl3's bug detection more robust

**Why this is better:** More comprehensive coverage provides higher confidence that the test suite catches all edge cases.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | Fully satisfies all acceptance criteria | Also satisfies criteria, adds more coverage |
| Risk | Zero risk - no code changes | Risk of introducing test bugs, modifying working code |
| Simplicity | Maximally simple | More complex, potentially over-engineered |
| Consistency | Matches prior exercise pattern (minimal changes) | Departs from convention |

### Decision: Proposal A wins

**Rationale:**
1. **All acceptance criteria are already met.** The tests pass for Impl4 and fail for Impl1-3. Adding more tests provides marginal value when the core requirement is satisfied.
2. **The prior solution (PR #246) also made zero code changes.** This confirms the convention for this exercise.
3. **Risk vs. reward:** Modifying a working test suite risks introducing errors for no concrete benefit. The existing tests already robustly detect all three implementation bugs.
4. **Opponent's critique is theoretical.** The `TestUnicodeLetters` test correctly catches different bugs in Impl2 (letters=0 vs 13) and Impl3 (characters=29 vs 16). Even if other tests existed, this single test is sufficient.

### Final Plan

1. Create feature branch `issue-286` from current HEAD
2. Verify all acceptance criteria by running tests with each COUNTER_IMPL value
3. Commit `.osmi/` documentation artifacts
4. Push and create PR closing issue #286

**No code modifications needed.** The `counter_test.go` and `counter.go` files are already correct.
