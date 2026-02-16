# Implementation Plan: Bowling Scorer

## File to Modify

- `go/exercises/practice/bowling/bowling.go` — the only file to change

## Architecture

Use a flat roll-tracking approach with frame counting:

### Data Structure

```go
type Game struct {
    rolls       [21]int  // max 21 rolls in a game (9*2 + 3)
    nRolls      int      // number of rolls recorded
    nFrames     int      // completed frames (0-10)
    rFrameStart int      // index of current frame's first roll
}
```

### Constants and Errors

```go
const pinsPerFrame = 10
const framesPerGame = 10
const maxRollsPerFrame = 2
const maxRollsLastFrame = 3
```

Four error variables:
- `ErrNegativeRollIsInvalid`
- `ErrPinCountExceedsPinsOnTheLane`
- `ErrPrematureScore`
- `ErrCannotRollAfterGameOver`

### Roll Logic

1. Validate pins < 0 → error
2. Validate pins > 10 → error
3. Check game is not over (completedFrames == 10) → error
4. Record the roll
5. For frames 1-9: strike completes frame immediately; 2 rolls complete frame; validate frame total ≤ 10
6. For frame 10: special handling for bonus rolls after strike/spare; validate pin counts considering resets

### Score Logic

1. If game not complete → error
2. Iterate through frames 1-10:
   - Strike: 10 + next 2 rolls
   - Spare: 10 + next 1 roll
   - Open: sum of 2 rolls
3. Frame pointer advances by 1 for strikes, 2 for spares/open

## Approach

The reference solution in `.meta/example.go` provides a well-tested implementation. The implementation will follow the same structural approach — storing rolls in a flat array, tracking frames completed, and computing score by walking the roll array frame-by-frame.

## Order of Changes

1. Write the complete `bowling.go` implementation
2. Run tests to verify
