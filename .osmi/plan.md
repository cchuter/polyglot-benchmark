# Implementation Plan: Bowling Scoring

## File to Modify

- `go/exercises/practice/bowling/bowling.go` — replace the stub with the full implementation.

## Architecture

### Data Model

Use a `Game` struct that stores:
- `rolls [21]int` — array of pin counts (max 21 rolls: 9 frames × 2 + last frame × 3).
- `nRolls int` — number of rolls recorded so far.
- `nFrames int` — number of completed frames (0–10).
- `rFrameStart int` — index into `rolls` where the current frame began.

### Constants

- `pinsPerFrame = 10`
- `framesPerGame = 10`
- `maxRollsPerFrame = 2`
- `maxRollsLastFrame = 3`
- `maxRolls = 21` (derived from above)

### Error Variables

Four sentinel errors:
- `ErrNegativeRollIsInvalid`
- `ErrPinCountExceedsPinsOnTheLane`
- `ErrPrematureScore`
- `ErrCannotRollAfterGameOver`

### Roll Logic

1. Validate: pins < 0 → error; pins > 10 → error; game over → error.
2. Record the roll.
3. For frames 1–9: a strike completes the frame immediately; otherwise complete after 2 rolls, validating the sum ≤ 10.
4. For frame 10: allow up to 3 rolls with special validation:
   - After 2 rolls: if sum < 10, frame is complete (no bonus roll).
   - After 3 rolls: validate that non-strike bonus rolls don't exceed 10 pins; complete the frame.

### Score Logic

Walk frames 1–10 using a `frameStart` cursor:
- Strike: 10 + next 2 rolls; advance cursor by 1.
- Spare: 10 + next 1 roll; advance cursor by 2.
- Open: sum of 2 rolls; advance cursor by 2.

Return error if `completedFrames != 10`.

### Helper Methods

Small helper methods on `*Game` for readability:
- `rollsThisFrame()`, `completeTheFrame()`, `completedFrames()`
- `isStrike(f)`, `isSpare(f)`, `rawFrameScore(f)`
- `strikeBonus(f)`, `spareBonus(f)`

## Ordering

1. Write the full implementation in `bowling.go`.
2. Run `go test` and `go vet` to verify.
3. Commit.
