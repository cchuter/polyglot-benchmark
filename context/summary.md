# Context Summary: polyglot-go-counter (Issue #115)

## Status: Complete

## Key Decisions

- Wrote 9 test functions covering: no-add, empty string, ASCII without newline, ASCII with newline, trailing newline, Unicode (Cyrillic), multiple AddString calls, only newlines, mixed content
- Used `assertCounts` helper with `t.Helper()` to reduce boilerplate
- All tests use `makeCounter()` factory (env-var driven implementation selection)
- Interesting finding: Impl3's letter count coincidentally matches correct answer for Cyrillic text (lead bytes of 2-byte UTF-8 sequences happen to be Latin letters), but Characters assertion catches it anyway

## Files Modified

- `go/exercises/practice/counter/counter_test.go` — Complete test suite

## Test Results

| Impl | Result | Failures |
|------|--------|----------|
| Impl4 (correct) | PASS | 0/9 |
| Impl1 (line bug) | FAIL | 4/9 |
| Impl2 (letter bug) | FAIL | 1/9 |
| Impl3 (byte bug) | FAIL | 1/9 |

## Branch

- `issue-115` pushed to origin
- Single commit: `4b61af3`
