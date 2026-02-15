# Plan Review: bottle-song Implementation (Issue #45)

## Review Method

Self-review comparing plan against test cases in `cases_test.go` and reference solution in `.meta/example.go`. No external codex agent was available for review.

## Verdict: PASS

### 1. Correctness — PASS

The plan correctly identifies the key components:
- Number-to-word mapping (0-10) including "no" for 0
- `capitalize` helper for lines 1-2
- `bottleWord` helper for singular/plural
- Verse generation using fmt.Sprintf
- Multi-verse assembly with empty string separators

The approach cleanly handles all edge cases through the bottleWord helper rather than special-casing.

### 2. Test Coverage — PASS

All 7 test cases are addressed:
1. First generic verse (start=10, take=1) — "Ten"/"nine", plural throughout
2. Last generic verse (start=3, take=1) — "Three"/"two", plural throughout
3. Verse with 2 bottles (start=2, take=1) — "Two" plural, "one" singular in line 4
4. Verse with 1 bottle (start=1, take=1) — "One" singular, "no" plural in line 4
5. First two verses (start=10, take=2) — multi-verse with empty string separator
6. Last three verses (start=3, take=3) — includes transitions through 2 and 1
7. All verses (start=10, take=10) — full song, 10 verses

### 3. Edge Cases — PASS

- "no green bottles" for count 0: handled by mapping 0 -> "no" and bottleWord(0) -> "bottles"
- Singular "bottle" when count is 1: handled by bottleWord(1) -> "bottle"
- Capitalization in lines 1-2 vs lowercase in line 4: capitalize helper used only for lines 1-2
- Empty string separator between verses (not after last): loop logic handles this

### 4. Capitalize Helper — PASS

Using `strings.ToUpper(s[:1]) + s[1:]` is safe for ASCII number words. This avoids deprecated `strings.Title` and is self-contained.

### 5. Import Requirements — PASS

Only `fmt` and `strings` needed. No external dependencies.

### 6. Potential Issues — NONE FOUND

- The plan's generalized approach with bottleWord() is actually cleaner than the reference solution's switch-case approach
- No risk of off-by-one in the separator logic since the plan specifies "between verses, not after the last"

### Summary

The plan is correct, complete, and well-scoped. It covers all test cases and edge cases. Approved for implementation.
