# Plan Review

## Review of counter_test.go against Impl1, Impl2, Impl3 bugs

### (1) Bug Detection Analysis

**Impl1 — Incorrect line counting** (counts `\n` only, no +1 for non-newline-ending text):
- `TestSimpleASCIINoNewline("hello")`: expects Lines=1, Impl1 returns 0 — **DETECTED**
- `TestASCIIWithNewlineInMiddle("Hello\nworld!")`: expects Lines=2, Impl1 returns 1 — **DETECTED**
- `TestMultipleAddStrings("hello\n" + "world")`: expects Lines=2, Impl1 returns 1 — **DETECTED**
- `TestMixedContent("abc 123!@#\ndef")`: expects Lines=2, Impl1 returns 1 — **DETECTED**

**Impl2 — ASCII-only letter detection** (uses `>='A'&&<='Z'` instead of `unicode.IsLetter()`):
- `TestUnicodeLetters("здравствуй, мир\n")`: expects Letters=13, Impl2 returns 0 — **DETECTED**

**Impl3 — Byte-level iteration** (iterates bytes not runes, corrupting multi-byte chars):
- `TestUnicodeLetters("здравствуй, мир\n")`: expects Characters=16, Impl3 returns 29 — **DETECTED**

All three buggy implementations are correctly detected by the test suite.

### (2) Risk Assessment

**No risk in keeping tests as-is.** The tests:
- Already pass with COUNTER_IMPL=4 (verified)
- Already fail with COUNTER_IMPL=1,2,3 (verified)
- Use standard Go testing patterns
- Have clear, descriptive test names
- Include a well-factored `assertCounts()` helper

### (3) Edge Case Analysis

The test suite covers:
- No add (zero state) ✓
- Empty string ✓
- Simple ASCII without newline ✓
- ASCII with newline in middle ✓
- String ending with newline ✓
- Unicode (Cyrillic) text ✓
- Multiple AddString calls ✓
- Only newlines ✓
- Mixed content (letters, digits, special chars) ✓

**Potential gaps (non-blocking):**
- No test for very long strings, but this is not needed for correctness
- No test for tabs/special whitespace, but tests already cover non-letter char counting
- No test for emoji/surrogate pairs, but Cyrillic test already validates rune iteration

**Conclusion:** The existing test suite is comprehensive and sufficient. No changes needed.

## Verdict

Plan is approved as-is. Proceed with implementation (which requires no code changes).
