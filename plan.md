# Implementation Plan: Bowling Game Scoring

## Overview

Replace the stub in `go/exercises/practice/bowling/bowling.go` with a complete bowling game scorer implementation. The solution uses an array-based roll storage approach with frame completion tracking for validation and score calculation.

## Architecture

### Data Model

```go
type Game struct {
    rolls       [23]int  // max possible rolls: 9 frames * 2 + 3 (10th frame) = 21, but 23 for safe indexing
    nRolls      int      // number of rolls recorded so far
    nFrames     int      // number of completed frames (0-10)
    rFrameStart int      // index into rolls[] where current frame begins
}
```

**Rationale**: Store raw rolls in a flat array rather than a frame-based structure. This simplifies bonus calculation (just index forward from the frame start) and matches how scoring naturally works — you need to look ahead for strike/spare bonuses.

### API Design

1. **`NewGame() *Game`** — Returns zero-valued Game (all fields default to 0, which is correct initial state)

2. **`Roll(pins int) error`** — Validates and records a roll:
   - Validate: pins >= 0, pins <= 10
   - Validate: game not over (completed frames < 10)
   - Record the roll
   - Determine frame completion based on game state
   - For frames 1-9: strike completes immediately; two rolls complete the frame (validate sum <= 10)
   - For frame 10: special handling for bonus rolls after strike/spare

3. **`Score() (int, error)`** — Calculate total score:
   - Validate: game is complete (10 frames)
   - Iterate through frames 1-9, applying strike/spare bonuses
   - Frame 10 scored at face value (bonuses already counted in pin totals)

### Error Variables

```go
var (
    ErrNegativeRollIsInvalid        = errors.New("Negative roll is invalid")
    ErrPinCountExceedsPinsOnTheLane = errors.New("Pin count exceeds pins on the lane")
    ErrPrematureScore               = errors.New("Score cannot be taken until the end of the game")
    ErrCannotRollAfterGameOver      = errors.New("Cannot roll after game is over")
)
```

### Roll Validation Logic (Frame 10)

The 10th frame is the complex part. After the first roll of frame 10:

- If first roll is a strike: player gets 2 bonus rolls
  - Second roll can be 0-10
  - If second roll is also a strike: third roll can be 0-10
  - If second roll is NOT a strike: third roll must be <= (10 - second roll)
- If first two rolls are a spare: player gets 1 bonus roll (0-10)
- If first two rolls total < 10: frame complete, no bonus rolls

### Score Calculation

```
total = 0
frameStart = 0
for frame 0..8:
    if strike:
        total += 10 + rolls[frameStart+1] + rolls[frameStart+2]
        frameStart += 1
    elif spare:
        total += 10 + rolls[frameStart+2]
        frameStart += 2
    else:
        total += rolls[frameStart] + rolls[frameStart+1]
        frameStart += 2
for frame 9 (10th):
    total += sum of remaining rolls at frameStart
return total
```

Note: The 10th frame is handled the same way by the scoring loop — strikes get +2 bonus, spares get +1 bonus — because the bonus rolls are physically present in the array.

### Helper Methods

- `rollsThisFrame() int` — rolls since frame started
- `completeTheFrame()` — increment frame count, advance frame start
- `completedFrames() int` — return completed frame count
- `isStrike(f int) bool` — roll at f == 10
- `isSpare(f int) bool` — rolls[f] + rolls[f+1] == 10
- `rawFrameScore(f int) int` — rolls[f] + rolls[f+1]
- `strikeBonus(f int) int` — rolls[f+1] + rolls[f+2]
- `spareBonus(f int) int` — rolls[f+2]

## File Changes

| File | Action | Details |
|------|--------|---------|
| `go/exercises/practice/bowling/bowling.go` | Replace | Full implementation (~130 lines) |

## Testing Strategy

1. `go build ./...` in the bowling directory to verify compilation
2. `go test -v ./...` to run all test cases
3. Verify all 36+ tests pass with no failures

## Risks and Mitigations

| Risk | Mitigation |
|------|-----------|
| 10th frame validation edge cases | Reference implementation provides complete logic; test cases cover all edge cases |
| Off-by-one in frame counting | Use 0-based frame counting with clear constants |
| Array bounds on roll storage | Array size 23 covers maximum possible rolls |
