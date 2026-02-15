# Challenger Review: Bowling Implementation

## Overall Verdict: PASS

All 31 test cases pass (21 score tests + 10 roll validation tests). The implementation is correct and handles all required edge cases.

## Detailed Analysis

### 1. 10th Frame Bonus Roll Handling - CORRECT

The implementation correctly distinguishes between frames 1-9 and the 10th frame:

- **Strikes in frames 1-9** (line 50-54): Auto-complete the frame with one roll.
- **10th frame strike**: Does NOT auto-complete; requires 2 bonus rolls before the frame is marked done.
- **10th frame spare**: Does NOT auto-complete; requires 1 bonus roll.
- **10th frame normal** (total < 10): Completes after 2 rolls with no bonus.

Verified scenarios:
- Perfect game (12 strikes) scores 300
- 10th frame strike + 7 + 1 = 18 (bonus counted once, not as strike bonus)
- 10th frame spare (7,3) + 10 = 20 (bonus counted once)
- 10th frame strike + strike + 6 = 26 (correct)

### 2. Pin Count Validation - CORRECT

All validation paths verified:

| Scenario | Validation | Line(s) |
|---|---|---|
| Individual roll > 10 | `pins > pinsPerFrame` | 38 |
| Negative roll | `pins < 0` | 41 |
| Frame 1-9: two rolls > 10 | `rawFrameScore > pinsPerFrame` | 57-61 |
| 10th frame: strike + non-strike + X where non-strike+X > 10 | `strikeBonus > pinsPerFrame` (when 2nd is not strike) | 77-80 |
| 10th frame: strike + 6 + 10 (can't strike after non-strike non-spare) | Same check catches this | 77-80 |

### 3. Scoring with Strike/Spare Bonuses - CORRECT

The `Score()` method (lines 101-123) implements standard bowling scoring:

- **Strike**: 10 + next 2 rolls. Frame advances by 1.
- **Spare**: 10 + next 1 roll. Frame advances by 2.
- **Normal**: Sum of 2 rolls. Frame advances by 2.

The frame advancement correctly allows bonus lookups to cross frame boundaries (e.g., strike bonus reads into the next frame's rolls).

Verified with manual trace:
- Consecutive strikes: Frame 1 (10+10+10=30), Frame 2 (10+10+5=25), Frame 3 (10+5+3=18), etc. = 81. Matches test.
- Last two strikes + bonus: 9 frames zeros + strike(20) + strike-frame(11) = 31. Matches test.

### 4. Game Completion Detection - CORRECT

- `completedFrames() == framesPerGame` check prevents rolling after game over (line 44).
- `Score()` returns `ErrPrematureScore` when `completedFrames() != framesPerGame` (line 102).
- Incomplete 10th frame scenarios (strike with 0 or 1 bonus rolls, spare with 0 bonus rolls) correctly report premature score.

### 5. Buffer Safety - CORRECT

- `maxRolls = (2 * 9) + 3 = 21` correctly sizes the array for worst case (all spares + bonus).
- Array accesses in `Score()` are bounded: maximum `frameStart + 2 = 11` for perfect game (all strikes), maximum `frameStart + 2 = 20` for all spares with bonus.

## Minor Observations (Not Bugs)

1. **Partially redundant validation** (lines 83-88): The check for `strikeBonus > pinsPerFrame && < 2*pinsPerFrame` with neither bonus being a strike is logically unreachable when the second roll is not a strike (already caught by lines 77-80). It only adds value when the second roll IS a strike, in which case it correctly allows through valid combinations like (10, 10, 6). Not harmful, just slightly more complex than strictly necessary.

2. **Defensive unreachable check** (lines 89-92): The `!isSpare` branch at line 89 is logically unreachable because if `rawFrameScore >= 10` after 2 rolls and it's not a strike or spare, the frame total exceeds 10, which would have been caught earlier (line 57-61). Good defensive practice but adds dead code.

## Conclusion

The implementation is **correct, complete, and passes all tests**. No bugs or edge case failures found. The code is well-structured with clean helper methods. Ready to merge.
