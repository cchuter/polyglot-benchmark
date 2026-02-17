# Plan Review

## Review Method
No external codex agent was available in the tmux environment. Self-review conducted against the reference solution, test cases, and edge cases.

## Review of Selected Plan (Branch 3: Coefficient Reduction)

### Correctness of Coefficient Approach
- **CORRECT**: The linear equation `sum(coeff[i] * digit[i]) == 0` correctly encodes the alphametics constraint. For addend words, letters at position `p` (from right) contribute `+10^p`; for the result word, `-10^p`. This is mathematically equivalent to checking the addition.
- **VERIFIED**: For `SEND + MORE == MONEY`:
  - S: +1000, E: +100+1-10 = +91, N: +10-100 = -90, D: +1
  - M: +1000-10000 = -9000, O: +100-1000 = -900, R: +10, Y: -1
  - With solution S=9,E=5,N=6,D=7,M=1,O=0,R=8,Y=2:
  - 9000 + 455 - 540 + 7 - 9000 - 0 + 80 - 2 = 0 ✓

### Pruning Strategy
- **SOUND**: Bounding the remaining partial sum with min/max possible contributions from unassigned letters is correct. Using sorted absolute coefficients (largest first) maximizes early pruning.
- **IMPROVEMENT NEEDED**: The pruning bounds should account for the leading-zero constraint — when computing min/max for a leading letter, digit 0 should not be considered.

### Edge Cases Identified
1. **Single-letter words**: `"I"` in `"I + BB == ILL"` — leading zero constraint applies (I ≠ 0). ✓ Handled by tracking leading letters.
2. **Equality-only puzzles**: `"A == B"` — two different letters must map to the same value, which is impossible with unique digits. ✓ Handled: coeff A=+1, B=-1, and with only 2 letters this has no solution where A≠B and A-B=0.
   - Wait, actually A=+1, B=-1. We need A-B=0, meaning A=B. But digits must be unique. So error returned. ✓
3. **Leading zeros in result**: `"ACA + DD == BD"` — B is a leading letter, so B≠0. Also D is leading, D≠0. Need to check all multi-digit words, not just the result.
   - **FIX NEEDED**: Must track leading letters for ALL words (addends AND result), not just result.
4. **Many addends**: The 199-addend puzzle accumulates large coefficients for repeated letters. The coefficient approach handles this naturally.
5. **Single-letter equality**: `"A == A"` — coefficient would be 0 for A. Any digit works. But this isn't in the test cases, and likely not a concern.

### Revised Notes
- Leading zero tracking must apply to ALL words with length > 1 (addends + result)
- Actually, re-reading: leading constraint is "the leading digit of a multi-digit number must not be zero" — so single-character words CAN be zero
- The pruning bound computation for leading letters should exclude 0 from available digits

### Verdict
The plan is sound with the minor corrections noted above. Ready for implementation.
