# Implementation Plan: Bowling Scoring (Issue #47)

## Overview

Implement a bowling game scorer in `go/exercises/practice/bowling/bowling.go`. The solution tracks rolls, validates input, manages frame state, and calculates the final score following standard 10-pin bowling rules.

## File to Modify

**`go/exercises/practice/bowling/bowling.go`** — Currently contains only `package bowling`. This is the only file to change.

## Architecture

### Data Structure

Use a `Game` struct that stores:
- `rolls [21]int` — Array to store all roll values (max 21 rolls: 9 frames × 2 + 3 for 10th frame)
- `nRolls int` — Count of rolls recorded
- `nFrames int` — Count of completed frames (0-10)
- `rFrameStart int` — Index of the first roll in the current frame

This approach stores raw rolls and tracks frame boundaries, which simplifies both validation and scoring.

### Error Variables

Export four error variables matching the test expectations:
- `ErrNegativeRollIsInvalid`
- `ErrPinCountExceedsPinsOnTheLane`
- `ErrPrematureScore`
- `ErrCannotRollAfterGameOver`

### API Functions

1. **`NewGame() *Game`** — Returns a zero-valued Game pointer
2. **`Roll(pins int) error`** — Records a roll with validation
3. **`Score() (int, error)`** — Calculates total score after game completion

## Approach

### Roll Validation Logic

The `Roll` method must handle these cases in order:

1. **Basic validation**: pins < 0 → `ErrNegativeRollIsInvalid`; pins > 10 → `ErrPinCountExceedsPinsOnTheLane`
2. **Game over check**: if 10 frames completed → `ErrCannotRollAfterGameOver`
3. **Record the roll** in the rolls array and increment nRolls
4. **Frame completion logic**:
   - **Frames 1-9 strike**: If pins == 10 on first roll of frame → complete frame immediately
   - **Frames 1-9 second roll**: If this is the 2nd roll in frame, validate that sum ≤ 10, then complete frame
   - **Frame 10 (special)**:
     - After 2 rolls: if sum < 10, complete frame (no bonus rolls)
     - After 2 rolls: if strike on first roll, validate 2nd+3rd roll pin counts
     - After 3 rolls: validate pin counts for bonus balls, then complete frame

### 10th Frame Pin Validation

The 10th frame has complex validation for bonus rolls:
- After a strike (1st roll = 10):
  - 2nd roll can be 0-10
  - If 2nd roll is NOT a strike, then 2nd + 3rd must ≤ 10
  - If 2nd roll IS a strike, 3rd roll can be 0-10
- After a spare (1st + 2nd = 10, 1st ≠ 10):
  - 3rd roll (bonus) can be 0-10
- If no strike or spare in 10th frame, only 2 rolls allowed

### Score Calculation

Walk through frames using a `frameStart` index:
- **Strike**: score = 10 + next two rolls; advance frameStart by 1
- **Spare**: score = 10 + next one roll; advance frameStart by 2
- **Open frame**: score = sum of two rolls; advance frameStart by 2

### Helper Methods

Small unexported helper methods on `Game`:
- `rollsThisFrame()` — current roll count in frame
- `completeTheFrame()` — increment nFrames, update rFrameStart
- `completedFrames()` — return nFrames
- `isStrike(f)` — check if roll at index f is 10
- `isSpare(f)` — check if rolls[f] + rolls[f+1] == 10
- `rawFrameScore(f)` — rolls[f] + rolls[f+1]
- `strikeBonus(f)` — rolls[f+1] + rolls[f+2]
- `spareBonus(f)` — rolls[f+2]

## Rationale

- The roll-array approach is simpler than a frame-based data structure because scoring needs to look ahead across frame boundaries
- Tracking frame completion explicitly (nFrames, rFrameStart) keeps validation clean
- This architecture matches the reference solution pattern, which is known to pass all tests
- Error variable names match the test explainText strings for clarity

## Ordering

1. Write the complete implementation in `bowling.go`
2. Run `go test -v` to verify all 25 test cases pass
3. Run `go vet` to check for issues
