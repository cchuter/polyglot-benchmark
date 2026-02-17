# Counter Exercise - Build & Test Results

## Build

```
$ go build ./...
(success - no output, exit code 0)
```

## COUNTER_IMPL=4 (All Pass - Expected)

```
$ COUNTER_IMPL=4 go test -v
=== RUN   TestNoAddString
--- PASS: TestNoAddString (0.00s)
=== RUN   TestEmptyString
--- PASS: TestEmptyString (0.00s)
=== RUN   TestSimpleASCIINoNewline
--- PASS: TestSimpleASCIINoNewline (0.00s)
=== RUN   TestASCIIWithNewlineInMiddle
--- PASS: TestASCIIWithNewlineInMiddle (0.00s)
=== RUN   TestStringEndingWithNewline
--- PASS: TestStringEndingWithNewline (0.00s)
=== RUN   TestUnicodeLetters
--- PASS: TestUnicodeLetters (0.00s)
=== RUN   TestMultipleAddStrings
--- PASS: TestMultipleAddStrings (0.00s)
=== RUN   TestOnlyNewlines
--- PASS: TestOnlyNewlines (0.00s)
=== RUN   TestMixedContent
--- PASS: TestMixedContent (0.00s)
PASS
ok  	counter	0.004s
```

**Result: 9/9 PASS**

## COUNTER_IMPL=1 (Failures Expected)

```
$ COUNTER_IMPL=1 go test -v
=== RUN   TestNoAddString
--- PASS: TestNoAddString (0.00s)
=== RUN   TestEmptyString
--- PASS: TestEmptyString (0.00s)
=== RUN   TestSimpleASCIINoNewline
    counter_test.go:32: Lines: got 0, want 1
--- FAIL: TestSimpleASCIINoNewline (0.00s)
=== RUN   TestASCIIWithNewlineInMiddle
    counter_test.go:38: Lines: got 1, want 2
--- FAIL: TestASCIIWithNewlineInMiddle (0.00s)
=== RUN   TestStringEndingWithNewline
--- PASS: TestStringEndingWithNewline (0.00s)
=== RUN   TestUnicodeLetters
--- PASS: TestUnicodeLetters (0.00s)
=== RUN   TestMultipleAddStrings
    counter_test.go:57: Lines: got 1, want 2
--- FAIL: TestMultipleAddStrings (0.00s)
=== RUN   TestOnlyNewlines
--- PASS: TestOnlyNewlines (0.00s)
=== RUN   TestMixedContent
    counter_test.go:69: Lines: got 1, want 2
--- FAIL: TestMixedContent (0.00s)
FAIL
exit status 1
FAIL	counter	0.003s
```

**Result: 5/9 PASS, 4/9 FAIL** (Lines count bug - only counts newlines, not actual lines)

## COUNTER_IMPL=2 (Failures Expected)

```
$ COUNTER_IMPL=2 go test -v
=== RUN   TestNoAddString
--- PASS: TestNoAddString (0.00s)
=== RUN   TestEmptyString
--- PASS: TestEmptyString (0.00s)
=== RUN   TestSimpleASCIINoNewline
--- PASS: TestSimpleASCIINoNewline (0.00s)
=== RUN   TestASCIIWithNewlineInMiddle
--- PASS: TestASCIIWithNewlineInMiddle (0.00s)
=== RUN   TestStringEndingWithNewline
--- PASS: TestStringEndingWithNewline (0.00s)
=== RUN   TestUnicodeLetters
    counter_test.go:50: Letters: got 0, want 13
--- FAIL: TestUnicodeLetters (0.00s)
=== RUN   TestMultipleAddStrings
--- PASS: TestMultipleAddStrings (0.00s)
=== RUN   TestOnlyNewlines
--- PASS: TestOnlyNewlines (0.00s)
=== RUN   TestMixedContent
--- PASS: TestMixedContent (0.00s)
FAIL
exit status 1
FAIL	counter	0.004s
```

**Result: 8/9 PASS, 1/9 FAIL** (Letters count bug - doesn't count unicode letters)

## COUNTER_IMPL=3 (Failures Expected)

```
$ COUNTER_IMPL=3 go test -v
=== RUN   TestNoAddString
--- PASS: TestNoAddString (0.00s)
=== RUN   TestEmptyString
--- PASS: TestEmptyString (0.00s)
=== RUN   TestSimpleASCIINoNewline
--- PASS: TestSimpleASCIINoNewline (0.00s)
=== RUN   TestASCIIWithNewlineInMiddle
--- PASS: TestASCIIWithNewlineInMiddle (0.00s)
=== RUN   TestStringEndingWithNewline
--- PASS: TestStringEndingWithNewline (0.00s)
=== RUN   TestUnicodeLetters
    counter_test.go:50: Characters: got 29, want 16
--- FAIL: TestUnicodeLetters (0.00s)
=== RUN   TestMultipleAddStrings
--- PASS: TestMultipleAddStrings (0.00s)
=== RUN   TestOnlyNewlines
--- PASS: TestOnlyNewlines (0.00s)
=== RUN   TestMixedContent
--- PASS: TestMixedContent (0.00s)
FAIL
exit status 1
FAIL	counter	0.003s
```

**Result: 8/9 PASS, 1/9 FAIL** (Characters count bug - counts bytes instead of runes)

## Summary

| Implementation | Pass | Fail | Status |
|---|---|---|---|
| COUNTER_IMPL=4 (correct) | 9 | 0 | ALL PASS |
| COUNTER_IMPL=1 (lines bug) | 5 | 4 | FAIL (expected) |
| COUNTER_IMPL=2 (letters bug) | 8 | 1 | FAIL (expected) |
| COUNTER_IMPL=3 (chars bug) | 8 | 1 | FAIL (expected) |

All results match expectations. The correct implementation (IMPL=4) passes all tests, and each buggy implementation (IMPL=1,2,3) fails on the tests that target its specific bug.
