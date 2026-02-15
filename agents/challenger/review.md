# Adversarial Review: polyglot-go-alphametics

## Verdict: PASS

All changes are correct, match the plan, and introduce no regressions.

---

## 1. Correctness: Leading-Zero Fix

### What changed
- Added `leadingLetters []rune` field to `problem` struct (line 17).
- In `parsePuzzle` (lines 50-60), iterates ALL `valueStrings` and for each word with `len(field) > 1`, records its first letter index (`rune(field[0]) - 'A'`) into `leadingLetters`, deduplicated via a `seen` map.
- In `solvePuzzle` (lines 91-101), replaced old single-answer check with a loop over ALL `leadingLetters`.

### Analysis
**Correct.** The old code only checked the answer (last word):
```go
r := p.vDigits[len(p.vDigits)-1][p.maxDigits-1]
if p.letterValues[r-1] == 0 { continue }
```
The new code checks every multi-digit word's leading letter. This covers addends AND the answer. The `seen` map correctly prevents duplicate entries. The `valueStrings` slice is populated from ALL fields (excluding "+" and "=="), so no words are missed.

### Edge case: `rune(field[0])` byte-to-rune cast
`field[0]` returns a `byte`, not a `rune`. The cast `rune(field[0])` is safe because the parser validates all characters via `unicode.IsUpper(r)` which ensures ASCII A-Z only. No issue.

### Edge case: could old code panic?
The old code accessed `p.vDigits[last][p.maxDigits-1]` which would return 0 if the answer is shorter than the longest addend. Then `p.letterValues[0-1]` would be an out-of-bounds panic. This was only safe because `isPuzzleSolution` always returns false first in such cases (the `r == 0` check at line 138). The new code eliminates this latent risk entirely since `leadingLetters` is computed from the words themselves, not from vDigits positions. This is a subtle improvement.

---

## 2. Correctness: `return carry == 0`

### What changed
Line 142: `return true` changed to `return carry == 0`.

### Analysis
**Correct.** After the column-by-column loop, if `carry != 0` the sum exceeds the answer's digit count. The old `return true` would accept such overflow silently. The fix ensures arithmetic correctness.

Verified against test case 4 ("A + A + ... + B == BCC"):
- Column 0: 11*9 + 1 = 100, carry=10, digit=0=C ✓
- Column 1: carry=10, carry=1, digit=0=C ✓
- Column 2: carry=1, carry=0, digit=1=B ✓
- Final carry=0 → returns true ✓

No test case has a valid expected solution with leftover carry, so no regression.

---

## 3. Edge Cases

| Case | Handled? | Reasoning |
|------|----------|-----------|
| Single-letter words (e.g., "A", "I") | ✓ | `len(field) > 1` correctly excludes them from `leadingLetters`. Per alphametics rules, single-letter words CAN be zero. |
| Puzzle "A == B" (no multi-digit words) | ✓ | `leadingLetters` is empty. Permutations ensure unique values, so A≠B always → error. |
| "ACA + DD == BD" (answer shorter than addend) | ✓ | `isPuzzleSolution` rejects all candidates at column 2 (answer has no digit there, `r==0` triggers false). |
| Same letter leads multiple words (e.g., "SEND + SOME == MONEY", S leads both) | ✓ | The `seen` map deduplicates, but even without dedup the check would still work (just redundant iterations). |
| Puzzle with all single-letter words (e.g., "A + B == C") | ✓ | `leadingLetters` is empty, no leading-zero check applies. Single-letter values can be 0. |

---

## 4. Adherence to Plan

| Plan Step | Implementation | Match? |
|-----------|---------------|--------|
| Add `leadingLetters` field to struct | Line 17: `leadingLetters []rune` | ✓ |
| Populate in `parsePuzzle` for multi-digit words | Lines 50-60 | ✓ |
| Replace single check with loop over all leading letters | Lines 91-101 in `solvePuzzle` | ✓ |
| Fix carry overflow: `return carry == 0` | Line 142 | ✓ |
| Run tests | All 10 pass per changes.md | ✓ |

**Full adherence.** No deviations from plan.

---

## 5. Regression Analysis (All 10 Test Cases)

| # | Puzzle | Expected | Leading Letters Checked | carry==0? | Pass? |
|---|--------|----------|------------------------|-----------|-------|
| 1 | I + BB == ILL | B=9,I=1,L=0 | B(9),I(1) | ✓ | ✓ |
| 2 | A == B | error | (none) | N/A | ✓ |
| 3 | ACA + DD == BD | error | A,D,B (never reached) | N/A | ✓ |
| 4 | 11*A + B == BCC | A=9,B=1,C=0 | B(1) | carry=0 ✓ | ✓ |
| 5 | AS + A == MOM | A=9,M=1,O=0,S=2 | A(9),M(1) | ✓ | ✓ |
| 6 | NO + NO + TOO == LATE | N=7,O=4,T=9,L=1,A=0,E=2 | N(7),T(9),L(1) | ✓ | ✓ |
| 7 | HE + SEES + THE == LIGHT | expected values | H(5),S(9),T(7),L(1) | ✓ | ✓ |
| 8 | SEND + MORE == MONEY | expected values | S(9),M(1) | ✓ | ✓ |
| 9 | AND + ... == DEFENSE | expected values | A(5),S(6),O(2),G(8),D(3) | ✓ | ✓ |
| 10 | 199-addend puzzle | expected values | multiple, all non-zero | ✓ | ✓ |

**No regressions.** All expected solutions remain valid under the new checks.

---

## 6. Code Quality

- **Clean and minimal**: Only the necessary changes were made, no unnecessary refactoring.
- **Good comments**: The "Record leading letters of multi-digit words" comment is clear.
- **Consistent style**: Matches the existing codebase style (imperative comments, similar loop patterns).
- **No new dependencies**: No imports added or removed.

### Minor optimization opportunity (not a bug)
The leading-zero check currently runs AFTER `isPuzzleSolution` (the expensive column arithmetic). It could be moved before the call since `isPuzzleSolution` sets `letterValues` as its first action. However:
1. This would require extracting the letter-value assignment from `isPuzzleSolution`, adding complexity.
2. Performance is acceptable (1.26s for all 10 tests including the 199-addend case).
3. Not a correctness issue.

**Recommendation:** Not worth changing. Current approach is correct and readable.

---

## Final Assessment

**PASS.** The implementation is correct, complete, matches the plan exactly, and introduces no regressions. Both bugs identified in the plan (incomplete leading-zero check and missing carry overflow validation) are properly fixed. Edge cases are handled correctly.
