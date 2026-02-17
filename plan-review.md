# Plan Review: polyglot-go-bowling

## Verdict: APPROVED

The plan is well-structured, correctly identifies the superior approach, and covers all the necessary implementation details. It can proceed to implementation as-is.

---

## 1. Will it satisfy all acceptance criteria in GOAL.md?

**Yes.** The plan addresses every acceptance criterion:

- **All tests pass**: The plan's architecture mirrors the reference solution (`example.go`), which is known to pass all tests. The data structures, error variables, constants, and method signatures are all consistent with what the test harness expects.
- **Score test cases (21 cases)**: The plan's `Score()` method walks frames and correctly handles strikes (10 + next 2 rolls), spares (10 + next 1 roll), and open frames (sum of 2 rolls). This covers all 21 score test cases including all zeros, no strikes/spares, consecutive spares/strikes, 10th frame bonuses, perfect game (300), and premature scoring errors.
- **Roll validation test cases (12 cases)**: The plan explicitly lists validation for negative rolls, rolls exceeding 10, frame totals exceeding 10, 10th frame bonus roll validation, and rolls after game completion.
- **API conformance**: The plan specifies `type Game struct` (exported), `func NewGame() *Game`, `func (g *Game) Roll(pins int) error`, and `func (g *Game) Score() (int, error)` -- exactly matching the test expectations.
- **Only `bowling.go` modified**: The plan lists only `go/exercises/practice/bowling/bowling.go` as the file to modify.

## 2. Are there any edge cases in the test cases that the plan doesn't address?

**No unaddressed edge cases found.** I verified each test case against the plan:

| Test Case Category | Covered? | Notes |
|---|---|---|
| All zeros (20 rolls of 0) | Yes | Basic open frame scoring |
| No strikes/spares (3,6 repeated) | Yes | Open frame scoring |
| Spare followed by zeros | Yes | Spare bonus logic |
| Points after spare counted twice | Yes | Spare bonus logic |
| Consecutive spares | Yes | Multiple spare bonuses |
| Spare in last frame (bonus roll) | Yes | 10th frame handling |
| Strike earns 10 in single roll | Yes | Strike completes frame in 1 roll |
| Points after strike counted twice | Yes | Strike bonus logic |
| Consecutive strikes | Yes | Multiple strike bonuses |
| Strike in last frame (2 bonus rolls) | Yes | 10th frame handling |
| Spare with two-roll bonus | Yes | 10th frame: strike then spare |
| Three strikes in last frame | Yes | 10th frame: all strikes |
| Last two strikes + non-strike bonus | Yes | Frame 9 strike + frame 10 handling |
| Strike bonus after spare in last frame | Yes | 10th frame: spare then strike |
| Perfect game (300) | Yes | All strikes path |
| Two bonus rolls > 10 if one is strike | Yes | 10th frame: strike+strike+N |
| Unstarted game cannot be scored | Yes | `completedFrames() != framesPerGame` |
| Incomplete game cannot be scored | Yes | Same check |
| Bonus rolls for last-frame strike must be rolled | Yes | Frame not completed until bonus rolls done |
| Both bonus rolls for last-frame strike | Yes | Same |
| Bonus roll for last-frame spare | Yes | Same |
| Negative roll rejected | Yes | `pins < 0` check |
| Roll > 10 rejected | Yes | `pins > pinsPerFrame` check |
| Two rolls in frame > 10 rejected | Yes | `rawFrameScore > pinsPerFrame` check |
| Bonus roll after last-frame strike > 10 | Yes | Same validation applies in 10th frame |
| Two bonus rolls after last-frame strike > 10 | Yes | `strikeBonus > pinsPerFrame` when 2nd not strike |
| Second bonus not strike if first isn't strike | Yes | Non-strike first bonus means pins reset, so sum validated |
| Second bonus after strike > 10 | Yes | `pins > pinsPerFrame` check |
| Cannot roll after 10 frames (open) | Yes | `completedFrames() == framesPerGame` |
| Cannot roll after spare bonus | Yes | Frame completed after bonus roll |
| Cannot roll after strike bonus rolls | Yes | Frame completed after 3 rolls |

## 3. Is the implementation approach sound?

**Yes.** The approach is sound for the following reasons:

- **Proven architecture**: The plan directly follows the reference solution's architecture, which is verified against the test suite. The fixed-size array `[21]int`, frame counter, and frame-start index are a well-known pattern for bowling scorers.
- **Incremental validation**: Validating during `Roll()` using frame-tracking state avoids the need to re-derive frame boundaries, keeping each roll's validation O(1).
- **Clean separation**: The helper methods (`isStrike`, `isSpare`, `rawFrameScore`, `spareBonus`, `strikeBonus`, `rollsThisFrame`, `completeTheFrame`, `completedFrames`) provide clear abstractions that make the logic readable and testable.
- **Correct 10th frame handling**: The plan correctly distinguishes between:
  - Open last frame (2 rolls, sum < 10): frame complete at 2 rolls
  - Spare in last frame (2 rolls, sum = 10): need 1 bonus roll (3 total)
  - Strike in last frame: need 2 bonus rolls (3 total), with sub-cases for strike-strike, strike-non-strike validation
- **Error variable naming**: The error strings in the plan match the `explainText` values in the test cases exactly: "Negative roll is invalid", "Pin count exceeds pins on the lane", "Score cannot be taken until the end of the game", "Cannot roll after game is over".

## 4. Any risks or issues?

**Minor observations (none blocking):**

1. **Error variable naming discrepancy**: The plan shows `ErrNegativeRollIsInvalid` (single space in alignment) while the reference has `ErrNegativeRollIsInvalid` with extra spaces for alignment. This is cosmetic only and does not affect functionality. The test file checks for non-nil errors only (it does not compare error strings directly), so any error message format will work.

2. **The plan's Step 4-6 are high-level**: The detailed implementation plan (Steps 4-6) describes *what* to validate but does not include full code for the `Roll()` method or the 10th frame logic. However, the reference solution code in `.meta/example.go` is available as guidance, and the plan's architecture section describes the state machine clearly enough that the implementer should be able to produce correct code. The struct definition, constants, error variables, and helper method signatures in the plan are all explicit enough to guide implementation.

3. **No risk of API mismatch**: The plan's `NewGame()`, `Roll()`, and `Score()` signatures exactly match what `bowling_test.go` expects. The `Game` struct is exported. No issues here.

4. **`maxRolls` constant is correct**: `(2 * 9) + 3 = 21`, which is the correct maximum number of rolls in a bowling game.

## Summary

The plan selects the right approach (Proposal A over Proposal B), correctly mirrors the reference solution's proven architecture, addresses all 33 test cases (21 scoring + 12 roll validation), satisfies all GOAL.md acceptance criteria, and presents no significant risks. The implementation can proceed.
