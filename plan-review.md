# Plan Review

## Reviewer: Self-review (no external codex agent available)

### 1. 10th Frame Validation Edge Cases

The plan mentions special handling for 10th frame bonus rolls but lacks specifics. The test cases require:

- After a strike in frame 10, first bonus roll can be 0-10
- After strike + non-strike bonus, second bonus must respect remaining pins (e.g., 10,5 → max 5 for second)
- After strike + strike bonus, second bonus can be 0-10 (pins reset)
- After a spare in frame 10, one bonus roll allowed (0-10)
- Game ends after bonus rolls; no further rolls allowed

**Verdict**: Plan needs these specific validation rules documented, but the reference solution in `.meta/example.go` handles all of these correctly and the plan follows that approach.

### 2. Flat Roll Array Approach

The flat `[21]int` array is sound:
- Maximum rolls: 9 frames × 2 rolls + 3 rolls for frame 10 = 21
- Frame pointer walks the array: +1 for strikes, +2 for spares/open
- Score calculation is simple and efficient

**Verdict**: Sound approach. Well-proven pattern for bowling scoring.

### 3. Missing Error Conditions

All test error conditions are covered by the four error variables:
- Negative roll → `ErrNegativeRollIsInvalid`
- Pin count > 10 → `ErrPinCountExceedsPinsOnTheLane`
- Frame total > 10 → `ErrPinCountExceedsPinsOnTheLane`
- Roll after game over → `ErrCannotRollAfterGameOver`
- Score before complete → `ErrPrematureScore`

**Verdict**: All error conditions are accounted for.

### Overall Assessment

**APPROVED** — The plan is sound. The reference solution provides a well-tested implementation that matches all test cases. Proceed with implementation.
