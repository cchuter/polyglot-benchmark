# Implementation Plan: polyglot-go-counter

## Branch 1: Verify Existing Test Suite (Minimal Changes)

**Approach**: The test file `counter_test.go` already exists with comprehensive tests. Verify it is correct and sufficient by running all four implementations against it. If all acceptance criteria are met, the exercise is already solved.

**Files to modify**: None (or minimal touch to `counter.go` if needed)

**Steps**:
1. Run `COUNTER_IMPL=4 go test -v` — verify all pass
2. Run `COUNTER_IMPL=1 go test -v` — verify failures for line counting
3. Run `COUNTER_IMPL=2 go test -v` — verify failures for Unicode letters
4. Run `COUNTER_IMPL=3 go test -v` — verify failures for byte-level Unicode
5. If all criteria met, commit as-is

**Evaluation**:
- Feasibility: High — already verified tests pass/fail correctly
- Risk: Low — no code changes needed
- Alignment: Fully satisfies all acceptance criteria (verified above)
- Complexity: Minimal — zero file modifications

## Branch 2: Enhanced Test Suite with Additional Edge Cases

**Approach**: While the existing tests work, add additional edge cases for robustness: tab characters, carriage returns, mixed Unicode scripts, emoji, very long strings, single-character strings.

**Files to modify**: `counter_test.go`

**Steps**:
1. Add tests for tabs, carriage returns, emoji characters
2. Add test for mixed Unicode scripts (CJK + Latin + Cyrillic)
3. Add test for single character strings
4. Verify all still pass with Impl4 and fail with Impl1-3

**Evaluation**:
- Feasibility: High — straightforward test additions
- Risk: Low — adding tests can't break existing behavior
- Alignment: Over-delivers on acceptance criteria
- Complexity: Low-medium — several new test functions

## Branch 3: Table-Driven Test Refactoring

**Approach**: Refactor the entire test suite into a table-driven format (Go convention) with test case structs, making it easier to add cases and maintain. Each case specifies input strings and expected counts.

**Files to modify**: `counter_test.go` (complete rewrite)

**Steps**:
1. Define a test case struct with fields: name, inputs []string, wantLines, wantLetters, wantChars
2. Create a comprehensive test case table
3. Write a single `TestCounter` function that iterates the table
4. Verify all pass with Impl4 and fail with Impl1-3

**Evaluation**:
- Feasibility: High — standard Go pattern
- Risk: Medium — rewriting working tests risks introducing bugs
- Alignment: Satisfies all criteria
- Complexity: Medium — full rewrite of test file

## Selected Plan

**Selected: Branch 1 — Verify Existing Test Suite**

**Rationale**: The existing test suite already:
- Passes all tests with Impl4 (correct implementation)
- Fails for Impl1 (detects wrong line counting in TestSimpleASCIINoNewline, TestASCIIWithNewlineInMiddle, TestMultipleAddStrings, TestMixedContent)
- Fails for Impl2 (detects ASCII-only letters in TestUnicodeLetters)
- Fails for Impl3 (detects byte-level iteration in TestUnicodeLetters — Characters: got 29, want 16)

All acceptance criteria are already met. Adding unnecessary test cases (Branch 2) or refactoring (Branch 3) would be over-engineering — the exercise asks for tests that detect the problems in the implementations, and the current tests do exactly that.

**Implementation Steps**:
1. Create feature branch `issue-201`
2. Verify all four implementations against the test suite
3. Ensure `counter.go` stub is correct (just package declaration)
4. Commit the verified solution
5. Push and create PR

**Files**:
- `go/exercises/practice/counter/counter.go` — already correct (package declaration only)
- `go/exercises/practice/counter/counter_test.go` — already correct (comprehensive test suite)
- No modifications needed to any fixture files
