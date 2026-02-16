# Plan Review (Self-Review — No Codex Agent Available)

## Review of Selected Plan: Coefficient-Based Permutation

### Correctness Analysis

**Coefficient approach is mathematically sound.** The insight that `SEND + MORE == MONEY` reduces to a single linear equation `Σ(coeff[i] * digit[i]) == 0` is correct. Each letter's coefficient is the sum of its place-value contributions across all addend appearances minus its place-value contributions in the result.

**Leading zero detection is correct.** Any letter that is the first character of a multi-digit word must not be assigned 0. This is handled by pre-filtering.

### Performance Analysis

**10-letter, 199-addend puzzle:** The parsing and coefficient computation is O(total_characters) which is fine. The permutation search over 10! = 3,628,800 permutations with a O(10) check each gives ~36M operations. This should complete in well under 10 seconds on modern hardware.

**Key optimization:** The coefficient computation handles the 199 addends at parse time, so the per-permutation cost is independent of the number of addends. This is a significant advantage over column-by-column checking.

### Potential Issues

1. **Permutation generation memory:** Generating all 10! permutations upfront as a slice would use ~140MB. **Must generate permutations lazily or inline.** The reference solution pre-generates all permutations into a slice — this works for smaller counts but is expensive at 10!. Consider using a callback/closure pattern or inline generation.

2. **Leading zero constraint for ALL words, not just the result:** Must check leading letters of ALL words (addends and result), not just the result word. The plan correctly states this.

3. **Single-letter words:** A single-letter word like "A" or "I" — these can be zero since they're single-digit. Actually no — a single-digit number's leading digit IS the number, so it CAN be zero (0 is a valid single-digit number). Wait, let me check the test: `"A == B"` expects an error because A and B must be different but both are single letters... Actually the error is about unique values, not leading zeros. Looking at the test case `"A + A + A + ... + B == BCC"`, A=9 and B=1, C=0. B is leading in BCC so B≠0. A is single-letter so can be any digit. This is correct.

   Actually, re-reading the problem: "the leading digit of a multi-digit number must not be zero". So single-letter words CAN be zero. The constraint only applies to multi-character words.

### Revised Decision

**The plan is sound.** The one adjustment needed is to avoid pre-generating all permutations into memory for the 10-letter case. Use an iterative permutation generator that yields one permutation at a time.

### Verdict: APPROVED with minor adjustment

Generate permutations iteratively rather than collecting them all into a slice.
