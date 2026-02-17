# Changelog - Counter Exercise

## Verification Summary

Verified the counter exercise test suite (`counter_test.go`) against all four
implementations. The test suite contains 9 tests covering:

- Zero-value / no input (`TestNoAddString`, `TestEmptyString`)
- Basic ASCII line counting (`TestSimpleASCIINoNewline`, `TestASCIIWithNewlineInMiddle`)
- Trailing newline handling (`TestStringEndingWithNewline`)
- Unicode letter recognition (`TestUnicodeLetters`)
- Accumulation across multiple calls (`TestMultipleAddStrings`)
- Newline-only strings (`TestOnlyNewlines`)
- Mixed content with digits and punctuation (`TestMixedContent`)

## Test Results

| Implementation | Result | Bug Detected |
|---|---|---|
| Impl1 | 4 FAIL | Line counting: counts `\n` occurrences instead of logical lines |
| Impl2 | 1 FAIL | ASCII-only `IsLetter`: misses Unicode letters (e.g. Cyrillic) |
| Impl3 | 1 FAIL | Byte iteration: counts bytes instead of runes for Characters() |
| Impl4 | 9 PASS | Correct reference implementation |

All buggy implementations are correctly caught by the test suite.
