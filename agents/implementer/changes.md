# Bowling Exercise Implementation

## Changes Made

- **File**: `go/exercises/practice/bowling/bowling.go`
- **Commit**: `feat: implement bowling game scorer`

## Summary

Implemented the complete bowling game scorer based on the reference solution. The implementation includes:

### Error Variables
- `ErrNegativeRollIsInvalid` - returned when a negative pin count is rolled
- `ErrPinCountExceedsPinsOnTheLane` - returned when pin count exceeds available pins
- `ErrPrematureScore` - returned when score is requested before game completion
- `ErrCannotRollAfterGameOver` - returned when rolling after all frames are complete

### Constants
- `pinsPerFrame` (10), `framesPerGame` (10), `maxRollsPerFrame` (2), `maxRollsLastFrame` (3), `maxRolls` (21)

### Types and Functions
- `Game` struct with rolls array, nRolls, nFrames, rFrameStart fields
- `NewGame()` constructor returning a zero-valued Game
- `Roll(pins int) error` method with full validation for all frames including 10th frame edge cases
- `Score() (int, error)` method computing final score with strike/spare bonuses
- Helper methods: `rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `isSpare`, `rawFrameScore`, `spareBonus`, `strikeBonus`
