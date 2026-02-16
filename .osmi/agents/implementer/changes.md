# Changes: Bowling Scorer Implementation

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — Full implementation of bowling game scorer

## Summary
Implemented a complete bowling game scoring engine with:
- Error variables for invalid rolls, pin count exceeded, premature scoring, and game-over conditions
- `Game` struct tracking rolls, roll count, frame count, and frame start index
- `NewGame()` constructor
- `Roll(pins int) error` with full validation for frames 1-9 and the 10th frame (strikes, spares, bonus rolls)
- `Score() (int, error)` walking all frames with strike/spare bonus logic
- Helper methods for frame state tracking and score computation

## Test Results
All 31 tests pass (10 Roll tests + 21 Score tests).
