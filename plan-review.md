# Plan Review: Go Bowling Exercise

## Review Method
Self-review against test cases and bowling rules (no external codex agent available).

## 1. Scoring Algorithm — Correct

The frame-walking approach for `Score()` is the standard and correct way to score bowling:
- Walk through frames 0-9, using a `rollIndex` to track position in the flat rolls slice
- Strike: 10 + next two rolls, advance rollIndex by 1
- Spare: 10 + next one roll, advance rollIndex by 2
- Open: sum of two rolls, advance rollIndex by 2

This correctly handles all bonus calculations. The 10th frame doesn't need special scoring treatment because the bonus rolls are already captured in the rolls slice — the frame-walk stops at frame 9, and any bonus rolls in the 10th frame are just used for prior frame bonuses.

**Verified against test cases**: Perfect game (12 rolls of 10) correctly yields 300 with this algorithm.

## 2. 10th Frame Validation — Needs Careful Implementation

The plan outlines the 10th frame rules but the implementation details need precision:

### Critical edge case from test: "second bonus roll cannot be strike if first isn't strike"
- `previousRolls: [0,0,...,10, 6]`, `roll: 10` → should error
- After a 10th-frame strike, if second ball is 6 (not a strike), then third ball can be at most 4 (10-6=4), not 10
- The plan mentions this but must be implemented precisely: if first ball is strike and second is NOT strike, third ball is limited to (10 - second ball)

### Two bonus rolls constraint:
- `previousRolls: [0,0,...,10, 5]`, `roll: 6` → should error (5+6=11 > 10)
- After strike + non-strike second, the second and third balls form a mini-frame of 10 pins

### Valid cases that must NOT error:
- `[..., 10, 10, 6]` → valid (26 points) — first two strikes each reset pins
- `[..., 10, 10, 10]` → valid (30 points) — three strikes in 10th frame
- `[..., 10, 7, 3]` → valid (20 points) — strike then spare in bonus
- `[..., 7, 3, 10]` → valid (20 points) — spare then strike bonus

## 3. Game Struct Design — Adequate with Refinement

The proposed struct is reasonable but I recommend tracking the 10th frame state more explicitly:

```go
type Game struct {
    rolls        []int
    currentFrame int   // 0-9
    currentBall  int   // ball within current frame
    done         bool
}
```

This is sufficient. The `currentBall` tracking within frame 9 (values 0, 1, 2) naturally handles the three-ball 10th frame. For frames 0-8, `currentBall` is 0 or 1.

One refinement: consider tracking `pinsUp` (pins remaining) to simplify validation. After a strike in the 10th frame, reset `pinsUp = 10`. After a non-strike first bonus roll, `pinsUp = 10 - pins`.

## 4. Potential Edge Cases

### Covered by tests and plan:
- Negative pins → error
- Pins > 10 → error
- Two rolls in normal frame > 10 → error
- Rolling after game over → error
- Scoring incomplete game → error

### Implementation risks:
1. **10th frame pin reset after strike**: Must reset available pins after each strike in 10th frame
2. **10th frame: non-strike first, spare second, bonus third**: The spare resets pins to 10 for the bonus ball
3. **Game completion detection**: Must correctly identify when game ends:
   - Normal 10th frame (no strike, no spare): after 2 balls
   - Spare in 10th: after 3 balls
   - Strike in 10th: after 3 balls

## Verdict: Plan is sound, proceed to implementation

The core algorithm is correct. The main implementation risk is the 10th frame validation logic, which requires careful tracking of pin resets. Recommend using explicit pin tracking for the 10th frame.
