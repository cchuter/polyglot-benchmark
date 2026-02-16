# Test Results - Counter Exercise

## IMPL=4 (Correct Implementation) - ALL PASS

```
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

## IMPL=1 (Buggy Line Counter) - FAIL (4 failures)

```
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
FAIL	counter	0.004s
```

## IMPL=2 (Buggy Unicode Letter Counter) - FAIL (1 failure)

```
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
FAIL	counter	0.005s
```

## IMPL=3 (Buggy Unicode Character Counter) - FAIL (1 failure)

```
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

## Summary

| Implementation | Result | Failures | Details |
|---|---|---|---|
| IMPL=4 (correct) | PASS | 0 | All 9 tests pass |
| IMPL=1 (buggy lines) | FAIL | 4 | Line counting fails: counts newlines not lines |
| IMPL=2 (buggy letters) | FAIL | 1 | Unicode letter counting fails: got 0, want 13 |
| IMPL=3 (buggy characters) | FAIL | 1 | Unicode character counting fails: got 29 (bytes), want 16 (runes) |
