# Challenger Review: Alphametics Solver

## Summary

**Verdict: PASS — No bugs found that would cause test failures.**

The implementation is a correct, clean coefficient-based backtracking solver with bounds pruning. All key algorithmic components are sound.

---

## Detailed Review

### 1. Parsing (lines 14-23) — CORRECT

- `strings.Fields` splits on whitespace, correctly handling multiple spaces.
- Filters `+` and `==` tokens; remaining tokens are words.
- `len(words) < 2` guards against degenerate input.
- Last word treated as result, all others as addends. Correct.

### 2. Coefficient Computation (lines 29-44) — CORRECT

Traced through multiple test cases:

**"I + BB == ILL"**: I = +1 - 100 = -99, B = +11, L = -11
- Check: -99(1) + 11(9) + -11(0) = -99 + 99 = 0 ✓

**"SEND + MORE == MONEY"**: S=+1000, E=+91, N=-90, D=+1, M=-9000, O=-900, R=+10, Y=-1
- Check: 1000(9)+91(5)+(-90)(6)+1(7)+(-9000)(1)+(-900)(0)+10(8)+(-1)(2) = 9542-9542 = 0 ✓

**"A + A + ... + B == BCC"** (11 A's): A=+11, B=+1-100=-99, C=-11
- Check: 11(9)+(-99)(1)+(-11)(0) = 99-99 = 0 ✓

The sign/place logic at lines 33-39 is correct: addends get sign=+1, result gets sign=-1. Place values accumulate correctly right-to-left.

### 3. Leading Zero Constraint (lines 41-43) — CORRECT

- First character of words with `len(word) > 1` marked as leading.
- Single-letter words (like "A", "I") correctly NOT marked as leading. These CAN be 0.
- Enforced at line 83-85 during backtracking: `if d == 0 && isLeading[idx] { continue }`.

**Test case "ACA + DD == BD"**: A, D, B all leading. The equation 101A + 10C + 10D - 10B = 0 requires 101A ≡ 0 (mod 10), which is impossible for A ∈ {1..9}. Correctly returns error. ✓

**Test case "A == B"**: Neither A nor B is leading (single-char). Need A=B but digits must be unique. Correctly returns error. ✓

### 4. Sorting by |coefficient| (lines 47-61) — CORRECT

Standard `sort.Slice` with absolute value comparison. Letters with highest impact assigned first, maximizing pruning effectiveness.

### 5. Backtracking Solver (lines 74-119) — CORRECT

- Iterates digits 0-9 in order for each letter.
- Used-digit bitmask `used & (1<<d)` correctly tracks uniqueness.
- `newSum` accumulates running total.
- Base case at `idx == n` checks `sum == 0`. Correct.

### 6. Bounds Pruning (lines 89-111) — CORRECT (sound, optimistic)

The pruning computes lo/hi bounds on what remaining letters can contribute:

- **Positive coefficient**: lo = c * minAvailDigit, hi = c * maxAvailDigit ✓
- **Negative coefficient**: lo = c * maxAvailDigit (most negative), hi = c * minAvailDigit (least negative) ✓
- **Prune condition**: `newSum+lo > 0 || newSum+hi < 0` means 0 ∉ [newSum+lo, newSum+hi] ✓

**Important subtlety**: The bounds are optimistic because `availMinMax(newUsed, ...)` uses the same `newUsed` bitmask for ALL remaining letters — it doesn't account for digits that will be consumed by intermediate letters. This means:
- `lo` may be lower than the actual minimum (multiple letters can't all use the same minimum digit)
- `hi` may be higher than the actual maximum

**This is safe for pruning**: the bounds are wider than reality, so we never prune a valid branch. We might fail to prune some dead branches, but correctness is preserved.

### 7. `availMinMax` function (lines 135-155) — CORRECT

- For leading letters, starts search from digit 1 (skips 0). ✓
- Correctly returns (-1, -1) when no digits available. ✓
- Caller handles this via the `feasible` flag at lines 95-98. ✓

### 8. Result Construction (lines 125-129) — CORRECT

Maps `letters[i]` (byte) to `assignment[i]` (int). Uses `string(ch)` to convert byte to single-char string for the map key. Correct.

---

## Edge Cases Verified

| Test Case | Status |
|---|---|
| "I + BB == ILL" (3 letters) | ✓ Coefficients and solution correct |
| "A == B" (no solution, single letters) | ✓ Returns error — unique digit constraint prevents A=B |
| "ACA + DD == BD" (leading zero makes it impossible) | ✓ Returns error — 101A mod 10 ≠ 0 for A>=1 |
| "A + A + ... + A + B == BCC" (repeated addend) | ✓ Coefficients accumulate correctly across 11 copies of "A" |
| "SEND + MORE == MONEY" (classic 8-letter) | ✓ All coefficients verified |
| 10-letter / 199-addend puzzle | ✓ Coefficient approach handles this efficiently — just accumulates per-letter weights across all 199 addends |

## Potential Concerns (non-blocking)

1. **No input validation for non-uppercase characters**: Not needed for exercism tests.
2. **Optimistic pruning bounds**: As discussed above, safe but could be tighter with per-digit accounting. Not needed for performance on these test cases.
3. **Deterministic solution order**: The solver finds the first valid assignment in digit-order iteration. Since all test puzzles have unique solutions, the specific iteration order doesn't matter.

## Conclusion

The implementation is clean, correct, and efficient. All test cases should pass.
