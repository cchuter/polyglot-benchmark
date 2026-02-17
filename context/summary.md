# Context Summary: polyglot-go-bowling

## Status: Complete

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — full implementation of bowling game scorer

## Implementation
- Package: `bowling`
- Struct: `Game` with `rolls [21]int`, `nRolls`, `nFrames`, `rFrameStart`
- Constructor: `NewGame() *Game`
- Methods: `Roll(pins int) error`, `Score() (int, error)`
- Helpers: `rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `isSpare`, `rawFrameScore`, `spareBonus`, `strikeBonus`
- Errors: `ErrNegativeRollIsInvalid`, `ErrPinCountExceedsPinsOnTheLane`, `ErrPrematureScore`, `ErrCannotRollAfterGameOver`

## Test Results
- 31/31 tests pass (10 roll validation + 21 scoring)
- Perfect game (300) works
- All 10th frame edge cases handled

## Branch
- `issue-282` pushed to origin
- Single commit: `f0b30ba Implement bowling game scorer`
