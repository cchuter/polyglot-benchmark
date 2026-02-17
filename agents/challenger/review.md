# Challenger Review: Bowling Implementation

## Verdict: APPROVE

All 31 test cases pass (21 score + 10 roll). The implementation is correct, clean, and follows the plan.

---

## 1. Correctness

**All tests pass.** I traced through the critical edge cases manually:

- **Perfect game (300):** 12 strikes. Frames 1-9 each complete after 1 roll. Frame 10 gets 3 rolls. Each frame scores 30. Total 300.
- **10th frame spare + bonus:** Frame doesn't complete at 2 rolls when `rawFrameScore == pinsPerFrame`. Waits for 3rd roll. Bonus counted once in Score().
- **10th frame strike + strike + X:** `rawFrameScore` of 20 allowed because `completedFrames == framesPerGame-1` and `isStrike` is true (line 58-59). Third roll validated via `strikeBonus` checks.
- **Consecutive strikes (frames 9-10):** Frame 9 completes after 1 roll (strike). Frame 10 starts at next index. Score correctly walks with `frameStart++` for strikes.
- **Last two strikes + open:** `[...10, 10, 0, 1]` - Frame 9 strike bonus = 10+0 = 10. Frame 10 = 10+0+1 = 11. Total matches expected 31.

### Roll validation edge cases verified:

| Scenario | Expected | Actual |
|----------|----------|--------|
| 10th frame: strike, 6, 10 | Error (6+10>10, second not strike) | Error at line 78-80 |
| 10th frame: strike, 10, 6 | OK (second IS strike, resets pins) | No error, lines 76-80 skipped |
| 10th frame: strike, 5, 6 | Error (5+6>10) | Error at line 78-80 |
| 10th frame: 7, 3 (spare), no bonus | Premature score | Frame not completed, Score() returns error |
| Roll after completed game | Error | Caught at line 43-45 |

## 2. Edge Cases & 10th Frame Logic

The 10th frame logic (lines 49-94) correctly handles all cases:

1. **Strike in 10th (line 49):** NOT caught by the early-return strike handler because `completedFrames() == framesPerGame-1`. Falls through to the `rollsThisFrame` checks.
2. **2 rolls in 10th (lines 55-71):**
   - If first is strike and total > 10: allowed (line 58-59 skips error)
   - If not strike and total > 10: error returned
   - If total < 10: frame completes (open frame, no bonus)
   - If total == 10 (spare): frame NOT completed, awaits 3rd roll
3. **3 rolls in 10th (lines 72-94):**
   - Strike path validates bonus rolls don't exceed pin counts
   - Spare path allows any valid 3rd roll
   - Neither strike nor spare at 3 rolls: `ErrCannotRollAfterGameOver` (unreachable in practice since open frames complete at 2 rolls)

### Subtle correctness note on `isSpare`:
`isSpare(f)` returns true when `rolls[f] + rolls[f+1] == 10`, which would also match a strike followed by 0. However, in `Score()`, `isStrike` is checked first in the switch statement, so a strike is never misidentified as a spare.

## 3. Code Quality

- **Idiomatic Go:** Clean struct with methods, proper error variables, good use of constants.
- **Readable helpers:** `isStrike`, `isSpare`, `rawFrameScore`, `strikeBonus`, `spareBonus` make the core logic self-documenting.
- **Fixed-size array:** `[maxRolls]int` avoids dynamic allocation. `maxRolls = 21` is correct (9 frames * 2 + 3).
- **No unnecessary exports:** Only `Game`, `NewGame`, `Roll`, `Score`, and error variables are exported.

### Minor observations (not blocking):
- The validation order checks `pins > 10` before `pins < 0` (lines 37-42). Functionally equivalent since both return errors, but checking negative first is slightly more conventional. The reference does the same order, so this is consistent.
- The nested conditionals in the 10th frame logic (lines 72-94) are complex but inherent to the problem. The comments adequately explain each branch.
- Line 69: `// For last frame, is it complete now ?` has a space before `?` — trivial style nit matching the reference.

## 4. Adherence to Plan

The implementation follows **Proposal A** (Roll-Array with Frame Tracking) exactly as selected by the judge:

| Plan Element | Implementation |
|---|---|
| Fixed-size `[21]int` array | `[maxRolls]int` where `maxRolls = 21` |
| `nRolls`, `nFrames`, `rFrameStart` tracking | All present in Game struct |
| Strike handling in frames 1-9 | Lines 49-53 |
| 2nd roll validation in frames 1-9 | Lines 55-66 |
| 10th frame with 2-3 rolls | Lines 67-94 |
| Score walking with frameStart | Lines 108-121 |
| Helper methods | Lines 124-131 |

The implementation matches the plan's detailed steps 1-6 precisely.

## 5. Comparison with Reference Solution

The implementation is functionally identical to `.meta/example.go`. The only differences are minor comment variations. Logic, structure, and all method signatures match exactly.

## Summary

No issues found. The implementation is correct, passes all tests, follows the plan, and produces clean idiomatic Go code. Ready to proceed.
