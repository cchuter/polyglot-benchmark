# Changes: Bowling Game Scorer

## Summary
Implemented the bowling game scoring solution in `go/exercises/practice/bowling/bowling.go`.

## What was implemented
- **Error variables**: `ErrNegativeRollIsInvalid`, `ErrPinCountExceedsPinsOnTheLane`, `ErrPrematureScore`, `ErrCannotRollAfterGameOver`
- **Constants**: `pinsPerFrame`, `framesPerGame`, `maxRollsPerFrame`, `maxRollsLastFrame`, `maxRolls`
- **Game struct**: with `rolls`, `nRolls`, `nFrames`, `rFrameStart` fields
- **NewGame()**: constructor returning a zero-valued Game
- **Roll(pins int) error**: records rolls with full validation (negative pins, pin count exceeding lane, game over detection, 10th frame special handling)
- **Score() (int, error)**: computes final score with strike and spare bonus logic
- **Helper methods**: `rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `rawFrameScore`, `spareBonus`, `strikeBonus`, `isSpare`

## Test results
All tests pass (`go test ./...` — ok).
