# Challenger Review: counter exercise

## Date
2026-02-16

## Summary

**PASS** — The test suite in `counter_test.go` is correct and complete. It successfully detects all three known bugs in the incorrect implementations while passing all tests on the correct implementation (Impl4). All fixture files are unmodified.

## Verified Test Results (actual execution)

| Implementation | Result | Failing Tests |
|---|---|---|
| Impl4 (correct) | 9/9 PASS | — |
| Impl1 (wrong line counting) | 5/9 PASS | TestSimpleASCIINoNewline, TestASCIIWithNewlineInMiddle, TestMultipleAddStrings, TestMixedContent |
| Impl2 (ASCII-only letters) | 8/9 PASS | TestUnicodeLetters |
| Impl3 (byte-level iteration) | 8/9 PASS | TestUnicodeLetters |

## Fixture File Integrity

All fixture files verified unchanged from main via `git diff`:
- `interface.go` — unchanged
- `maker.go` — unchanged
- `impl1.go` — unchanged
- `impl2.go` — unchanged
- `impl3.go` — unchanged
- `impl4.go` — unchanged

## Bug Detection Analysis

### Impl1: Wrong Line Counting
Impl1 counts lines by incrementing on each `\n` character only. It fails to account for content that doesn't end with a newline (e.g., "hello" should be 1 line, but Impl1 returns 0).

**Detected by:**
- TestSimpleASCIINoNewline — Lines: got 0, want 1
- TestASCIIWithNewlineInMiddle — Lines: got 1, want 2
- TestMultipleAddStrings — Lines: got 1, want 2
- TestMixedContent — Lines: got 1, want 2

### Impl2: ASCII-only Letter Counting
Impl2 checks `(char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')` which misses all non-ASCII Unicode letters (Cyrillic, etc.).

**Detected by:**
- TestUnicodeLetters — Letters: got 0, want 13

### Impl3: Byte-level Iteration
Impl3 iterates with `for i := 0; i < len(s); i++` and casts individual bytes to rune via `rune(s[i])`. For multi-byte UTF-8 characters (like Cyrillic, 2 bytes each), this produces incorrect character counts (counts bytes instead of runes).

**Detected by:**
- TestUnicodeLetters — Characters: got 29, want 16

## Test Coverage Assessment

The 9 test cases cover:
1. **Zero state** — TestNoAddString
2. **Empty string** — TestEmptyString
3. **Simple ASCII** — TestSimpleASCIINoNewline
4. **Newline in middle** — TestASCIIWithNewlineInMiddle
5. **Trailing newline** — TestStringEndingWithNewline
6. **Unicode content** — TestUnicodeLetters (critical for detecting Impl2 and Impl3 bugs)
7. **Multiple AddString calls** — TestMultipleAddStrings (tests accumulation)
8. **Only newlines** — TestOnlyNewlines
9. **Mixed content** — TestMixedContent (letters, digits, symbols)

## Correctness of Expected Values

All test expectations verified arithmetically against Impl4's correct logic:
- Lines = newlines + 1 (if trailing non-newline content exists), else newlines, else 0
- Letters = count of `unicode.IsLetter(r)` runes
- Characters = total rune count

## Code Quality
- Clean, idiomatic Go test code
- `assertCounts` helper with `t.Helper()` for clean error reporting
- Descriptive test names following Go conventions
- Correct package declaration (`package counter`)
