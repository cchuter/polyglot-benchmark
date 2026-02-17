# Change Log: polyglot-go-bowling

## Changes Made

### `go/exercises/practice/bowling/bowling.go`
- **Added** error variables: `ErrNegativeRollIsInvalid`, `ErrPinCountExceedsPinsOnTheLane`, `ErrPrematureScore`, `ErrCannotRollAfterGameOver`
- **Added** constants: `pinsPerFrame`, `framesPerGame`, `maxRollsPerFrame`, `maxRollsLastFrame`, `maxRolls`
- **Added** `Game` struct with fixed-size rolls array and frame-tracking counters
- **Added** `NewGame()` constructor
- **Added** `Roll(pins int) error` with full validation logic including 10th frame bonus rules
- **Added** `Score() (int, error)` with frame-by-frame scoring walk
- **Added** helper methods: `rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `rawFrameScore`, `spareBonus`, `strikeBonus`, `isSpare`

## Test Results
- All 31 tests pass (10 Roll + 21 Score)
- `go vet` clean
