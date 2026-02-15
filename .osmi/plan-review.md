# Plan Review: bottle-song Implementation (Issue #38)

## Review Method

Self-review comparing plan against test cases in `cases_test.go` and reference solution in `.meta/example.go`.

## Verdict: PASS

### 1. Correctness — PASS

The plan correctly identifies the three key components:
- Number-to-word mapping (0-10)
- Verse generation with singular/plural handling
- Multi-verse assembly with separators

The reference solution uses explicit `switch` cases for n==1, n==2, and a default case. The plan proposes the same approach. This correctly handles:
- n==1: singular "bottle" in lines 1-2, "no green bottles" in line 4
- n==2: plural "bottles" in lines 1-2, singular "one green bottle" in line 4
- n>=3: plural "bottles" throughout

### 2. Test Coverage — PASS

All 7 test cases are addressed:
1. First generic verse (start=10, take=1) — default case
2. Last generic verse (start=3, take=1) — default case
3. Verse with 2 bottles (start=2, take=1) — n==2 case
4. Verse with 1 bottle (start=1, take=1) — n==1 case
5. First two verses (start=10, take=2) — multi-verse with separator
6. Last three verses (start=3, take=3) — multi-verse including n==2 and n==1
7. All verses (start=10, take=10) — full song

### 3. Edge Cases — PASS

- The plan handles "no green bottles" for count 0 (in n==1 case, line 4)
- Singular "bottle" when count is 1 is handled in dedicated cases
- Empty string separator between verses is correctly specified

### 4. Capitalize Helper — PASS

The plan proposes a `capitalize` function using `strings.ToUpper(s[:1]) + s[1:]`. This is safe because all number words are single-byte ASCII. This avoids the deprecated `strings.Title`.

Note: The reference solution uses `Title()` from the test file. The plan's `capitalize` approach is equivalent and more self-contained.

### 5. Import Requirements — PASS

Only `fmt` and `strings` are needed. No external dependencies.

### Summary

The plan is correct, complete, and well-scoped. It covers all test cases and edge cases. Approved for implementation.
