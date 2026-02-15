# Context: Bowling Exercise Implementation

## Status
Complete — all tests pass, verifier issued PASS verdict.

## Branch
`issue-11` — pushed to origin

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — full implementation of bowling scorer

## Implementation Summary
- `Game` struct: `rolls []int`, `currentFrame int`, `currentBall int`, `done bool`
- `NewGame()`: returns initialized `*Game`
- `Roll(pins int) error`: validates pins, tracks frame/ball state, handles 10th frame bonus
- `Score() (int, error)`: walks frames 0-9 using rollIndex, applies strike/spare bonuses

## Key Design Decisions
1. **Flat rolls slice**: All rolls stored in `[]int`, scored by walking frames at `Score()` time
2. **10th frame tracking**: `currentBall` goes 0→1→2 in frame 9, with pin-reset logic after strikes
3. **Validation**: Negative pins, pins > 10, frame overflow, game-over all return errors
4. **Game completion**: `done` flag set when all required rolls (including bonus) are recorded

## Test Results
- 31/31 tests pass (10 roll validation + 21 scoring)
- Perfect game (300), consecutive strikes (81), all spare/strike combos verified
- All error cases validated (negative, overflow, incomplete game, post-game rolls)

## Commit
`0928bb7` — feat: implement bowling game scorer with Game struct, Roll, and Score methods
