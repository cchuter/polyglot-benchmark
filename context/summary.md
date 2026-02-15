# Context Summary: polyglot-go-bowling (Issue #57)

## Status: COMPLETE

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — Full implementation of bowling scorer

## Implementation
- `Game` struct with: `rolls []int`, `frame int`, `rollInFrame int`, `done bool`
- `NewGame()` returns `*Game`
- `Roll(pins int) error` with validation for pin range, frame totals, 10th frame bonuses, game over
- `Score() (int, error)` walks rolls frame-by-frame with strike/spare bonus calculation

## Test Results
- 31/31 tests pass (10 Roll validation + 21 Score)
- Perfect game (300) confirmed
- All error cases handled

## Branch
- `issue-57` pushed to origin
- Single commit: `feat: implement bowling game scorer`
