# Implementation Plan: polyglot-go-counter

## Branch 1: Minimal Direct Tests

**Approach**: Write focused, individual test functions that each test a specific behavior. Each test creates a counter, calls AddString, and asserts Lines/Letters/Characters counts.

**Files to modify**:
- `go/exercises/practice/counter/counter_test.go` — replace stub with test suite
- `go/exercises/practice/counter/counter.go` — clean up stub (remove TODO comment if present)

**Architecture**:
- One helper function `assertCounts(t, c, lines, letters, chars)` to reduce repetition
- ~8-9 individual `Test*` functions, each testing one scenario
- Tests cover: no input, empty string, simple ASCII, newlines, Unicode, multiple AddStrings, only newlines, mixed content

**Evaluation**:
- Feasibility: High — straightforward, no dependencies
- Risk: Low — simple test functions
- Alignment: Fully satisfies all acceptance criteria
- Complexity: Low — 1 file changed, ~70 lines of test code

---

## Branch 2: Table-Driven Test Suite

**Approach**: Use Go's table-driven test pattern with `t.Run` subtests. Define a slice of test cases with input strings and expected counts, then iterate through them.

**Files to modify**:
- `go/exercises/practice/counter/counter_test.go` — replace stub with table-driven tests
- `go/exercises/practice/counter/counter.go` — clean up stub

**Architecture**:
- Define a `testCase` struct with fields: `name`, `inputs []string`, `wantLines`, `wantLetters`, `wantChars`
- Single `TestCounter` function that iterates over cases with `t.Run`
- All test data in a single slice literal

**Evaluation**:
- Feasibility: High — standard Go pattern
- Risk: Low — but slightly harder to see which specific impl bug each test targets
- Alignment: Satisfies criteria, but the `makeCounter()` factory calls `log.Fatalf` on missing env var, so each subtest still needs its own counter
- Complexity: Low-Medium — 1 file changed, ~80 lines, slightly more complex structure

---

## Branch 3: Exhaustive Multi-Impl Validation

**Approach**: Write tests that explicitly run against all four implementations in a single test run, validating that Impl4 passes and Impl1-3 fail in specific ways.

**Files to modify**:
- `go/exercises/practice/counter/counter_test.go` — comprehensive validation suite
- `go/exercises/practice/counter/counter.go` — clean up stub

**Architecture**:
- Bypass the `makeCounter()` factory and directly instantiate `&Impl1{}` through `&Impl4{}`
- Test each implementation against every test case
- Assert that Impl4 passes all tests, and that specific impls fail specific tests

**Evaluation**:
- Feasibility: Medium — works but bypasses the intended factory pattern
- Risk: Medium — couples tests to specific impl details, fragile
- Alignment: Over-satisfies criteria but changes the testing model from the exercise design
- Complexity: High — more code, more complex assertions

---

## Selected Plan

**Selected: Branch 1 — Minimal Direct Tests**

**Rationale**: Branch 1 is the clear winner because:
1. It matches the exercise's intended design — tests use `makeCounter()` and the `Counter` interface
2. It's the simplest approach with the lowest risk
3. It directly maps each test to a specific behavior, making it easy to verify acceptance criteria
4. It follows the same pattern used in the previous successful solution for this exercise
5. Branch 2 adds unnecessary structure for just 8-9 test cases
6. Branch 3 violates the exercise design by bypassing the factory

### Detailed Implementation Plan

**Step 1**: Create feature branch `issue-245` from current HEAD.

**Step 2**: Write `counter_test.go` with the following test functions:

```
assertCounts(t, c, lines, letters, chars)  — helper
TestNoAddString                             — zero value, expects 0/0/0
TestEmptyString                             — AddString(""), expects 0/0/0
TestSimpleASCIINoNewline                    — "hello", expects 1/5/5 (catches Impl1)
TestASCIIWithNewlineInMiddle                — "Hello\nworld!", expects 2/10/12 (catches Impl1)
TestStringEndingWithNewline                 — "hello\n", expects 1/5/6
TestUnicodeLetters                          — "здравствуй, мир\n", expects 1/13/16 (catches Impl2, Impl3)
TestMultipleAddStrings                      — two AddString calls, expects 2/10/11 (catches Impl1)
TestOnlyNewlines                            — "\n\n\n", expects 3/0/3
TestMixedContent                            — "abc 123!@#\ndef", expects 2/6/14 (catches Impl1)
```

**Step 3**: Ensure `counter.go` is just `package counter`.

**Step 4**: Verify:
- `COUNTER_IMPL=4 go test -v` → all pass
- `COUNTER_IMPL=1 go test -v` → failures (line counting)
- `COUNTER_IMPL=2 go test -v` → failures (unicode letters)
- `COUNTER_IMPL=3 go test -v` → failures (unicode characters)
- `go vet ./...` → clean

**Step 5**: Commit changes with descriptive message.
