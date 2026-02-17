# Plan Review: polyglot-go-counter

## Review (Self-Review — no Codex agent available)

### Strengths
1. **Complete bug coverage**: Each buggy implementation is caught by at least one test:
   - Impl1 (wrong line counting): caught by TestSimpleASCIINoNewline, TestASCIIWithNewlineInMiddle, TestMultipleAddStrings, TestMixedContent
   - Impl2 (ASCII-only letters): caught by TestUnicodeLetters
   - Impl3 (byte-level iteration): caught by TestUnicodeLetters (wrong character count)
2. **Clean design**: Uses a simple helper function and individual test functions — idiomatic Go
3. **Correct expected values**: All expected values verified against Impl4 (correct implementation)
4. **Edge cases covered**: Empty counter, empty string, only newlines, trailing newlines, multiple AddString calls

### Potential Gaps
1. **Impl2 vs Impl3 differentiation**: Both fail on TestUnicodeLetters but for different reasons (Impl2: wrong letter count, Impl3: wrong character count). The test detects both, but through different assertion failures — this is fine.
2. **No test for strings with only non-letter characters** (e.g., "123!@#"): Not strictly needed since TestMixedContent covers non-letter characters, and the impls don't have bugs in this area.
3. **No test for multi-byte characters without newlines**: Could add but TestUnicodeLetters already covers this adequately.

### Verdict
The plan is **sound and sufficient**. The test suite:
- Passes all tests for Impl4 ✓
- Detects Impl1's line counting bug ✓
- Detects Impl2's ASCII-only letter detection ✓
- Detects Impl3's byte-level iteration bug ✓
- Covers all reasonable edge cases ✓

**Recommendation**: Proceed with implementation as planned.
