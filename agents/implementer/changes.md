# Changes: Implement bowling game scorer

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — replaced stub with full implementation

## Summary
Implemented a complete bowling game scorer following the plan's Proposal A (roll-array with frame tracking). The implementation uses a fixed-size array of 21 rolls and tracks game progress via frame count and frame-start index.

### Key components:
- **Error variables**: `ErrNegativeRollIsInvalid`, `ErrPinCountExceedsPinsOnTheLane`, `ErrPrematureScore`, `ErrCannotRollAfterGameOver`
- **Game struct**: Tracks rolls in a `[21]int` array with `nRolls`, `nFrames`, and `rFrameStart` counters
- **NewGame()**: Returns a zero-valued Game pointer
- **Roll()**: Records rolls with full validation — negative pins, pin count > 10, game-over detection, frame total validation, and 10th frame special handling (bonus rolls for strikes/spares)
- **Score()**: Walks frames computing strike bonuses (next 2 rolls), spare bonuses (next 1 roll), and open frame sums
- **Helper methods**: `rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `isSpare`, `rawFrameScore`, `spareBonus`, `strikeBonus`

## Test Results
All 31 tests pass (10 Roll tests + 21 Score tests).
