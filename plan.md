# Implementation Plan: Bowling Game Scoring

## Overview

Implement a complete bowling game scorer in `go/exercises/practice/bowling/bowling.go`. The implementation follows the reference solution pattern in `.meta/example.go` with a `Game` struct that tracks rolls, frame count, and frame boundaries.

## File to Modify

- `go/exercises/practice/bowling/bowling.go` — the only file to change (currently a stub with just `package bowling`)

## Architecture

### Data Structures

```go
type Game struct {
    rolls       [21]int  // max 21 rolls in a game (9 frames * 2 + 3 for 10th)
    nRolls      int      // number of rolls recorded
    nFrames     int      // number of completed frames
    rFrameStart int      // index of first roll in current frame
}
```

### Constants and Errors

- `pinsPerFrame = 10`, `framesPerGame = 10`, `maxRollsPerFrame = 2`, `maxRollsLastFrame = 3`
- `maxRolls = (2 * 9) + 3 = 21`
- Error variables: `ErrNegativeRollIsInvalid`, `ErrPinCountExceedsPinsOnTheLane`, `ErrPrematureScore`, `ErrCannotRollAfterGameOver`

### Roll Logic

1. Validate pin count (negative → error, > 10 → error)
2. Check game not over (completed frames == 10 → error)
3. Record the roll
4. For frames 1-9: strike (10 pins on first roll) → complete frame immediately; two rolls → validate sum ≤ 10, complete frame
5. For frame 10 (special): up to 3 rolls allowed if strike/spare; validate pin counts within each "sub-frame" of the 10th frame

### Score Logic

1. If game not complete (nFrames != 10) → return error
2. Walk through frames using a `frameStart` index:
   - Strike: 10 + next two rolls, advance frameStart by 1
   - Spare: 10 + next one roll, advance frameStart by 2
   - Open: sum of two rolls, advance frameStart by 2

### Helper Methods

- `rollsThisFrame()` — rolls in current frame
- `completeTheFrame()` — increment nFrames, update rFrameStart
- `completedFrames()` — return nFrames
- `isStrike(f)`, `isSpare(f)` — check roll patterns at index f
- `rawFrameScore(f)`, `spareBonus(f)`, `strikeBonus(f)` — scoring helpers

## Approach and Ordering

1. Write the complete `bowling.go` file with all types, constants, errors, and methods
2. Run `go test` to verify all 32 test cases pass
3. No incremental steps needed — this is a single-file, self-contained implementation

## Rationale

The reference solution in `.meta/example.go` is well-structured and handles all edge cases correctly. The implementation will follow the same proven approach since the test suite is extensive and validates all bowling scoring rules.
