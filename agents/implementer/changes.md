# Implementer Changes

## Bowling Game Scorer Implementation

**File**: `go/exercises/practice/bowling/bowling.go`

### What was implemented

Complete bowling game scoring implementation including:

- **Error variables**: `ErrNegativeRollIsInvalid`, `ErrPinCountExceedsPinsOnTheLane`, `ErrPrematureScore`, `ErrCannotRollAfterGameOver`
- **Constants**: `pinsPerFrame`, `framesPerGame`, `maxRollsPerFrame`, `maxRollsLastFrame`, `maxRolls`
- **Game struct**: Tracks rolls array, roll count, completed frames, and current frame start index
- **NewGame()**: Constructor returning a zero-valued Game
- **Roll(pins int) error**: Records rolls with full validation
- **Score() (int, error)**: Calculates final score with strike/spare bonuses
- **Helper methods**: `rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `isSpare`, `rawFrameScore`, `spareBonus`, `strikeBonus`

### Test results

All tests pass (both Roll and Score test suites).
