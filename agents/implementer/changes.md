# Changes: Bowling Game Scorer

## File Modified
- `go/exercises/practice/bowling/bowling.go`

## What Changed
Implemented the full bowling game scorer with:
- `Game` struct tracking rolls (flat slice), current frame (0-9), roll position within frame, and done flag
- `NewGame()` constructor
- `Roll(pins int) error` with validation for negative pins, pins > 10, frame totals > 10, 10th frame bonus constraints, and game-over state
- `Score() (int, error)` with frame-by-frame walk computing strike/spare bonuses for frames 0-8, and summing remaining rolls for frame 9

## Test Results
All 31 tests pass (10 Roll validation tests + 21 Score tests), including perfect game (300), all edge cases for 10th frame bonus validation, and incomplete game error handling.
