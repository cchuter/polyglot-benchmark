# Plan Review

## Reviewer: Self-review (no codex agent available)

### Coverage of Test Cases

The plan covers all test scenarios in `cases_test.go`:

**Score tests (20 cases):**
- All zeros game ✓
- No strikes/spares game ✓
- Spare followed by zeros ✓
- Points after spare counted twice ✓
- Consecutive spares ✓
- Spare in last frame with bonus ✓
- Strike scoring ✓
- Points after strike counted twice ✓
- Consecutive strikes ✓
- Strike in last frame with bonus ✓
- Spare with two-roll bonus in last frame ✓
- Three strikes in last frame ✓
- Last two strikes with non-strike bonus ✓
- Strike with one-roll bonus after spare in last frame ✓
- Perfect game (300) ✓
- Two bonus rolls after last-frame strike scoring >10 ✓
- Unstarted game error ✓
- Incomplete game error ✓
- Bonus rolls for last-frame strike must be rolled ✓
- Both bonus rolls must be rolled ✓
- Bonus roll for spare must be rolled ✓

**Roll tests (12 cases):**
- Negative roll error ✓
- Roll > 10 error ✓
- Two rolls in frame > 10 error ✓
- Bonus roll after last-frame strike > 10 error ✓
- Two bonus rolls after last-frame strike > 10 error ✓
- Second bonus after strike can't be strike if first isn't ✓
- Second bonus after strike > 10 error ✓
- Cannot roll after game over (open frame) ✓
- Cannot roll after spare bonus ✓
- Cannot roll after strike bonus ✓

### 10th Frame Edge Cases

The plan correctly identifies the complex 10th frame scenarios:
1. Strike + non-strike + fill: second and third rolls can't exceed 10
2. Strike + strike + fill: any value 0-10 valid for third roll
3. Spare + fill: single fill ball, any value 0-10
4. Open frame (<10 on two rolls): no bonus, frame complete

### Assessment

The plan is sound. Using the reference solution pattern is the lowest-risk approach. No issues identified.

### Recommendation

Proceed with implementation as planned.
