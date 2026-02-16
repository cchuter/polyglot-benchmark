# Changes: Bowling Game Scorer

## File Modified
- `go/exercises/practice/bowling/bowling.go` — Replaced stub with full implementation

## What Was Done
- Implemented complete bowling game scoring logic including:
  - `NewGame()` constructor returning zero-valued Game struct
  - `Roll(pins int) error` with full validation (negative pins, pin count exceeding lane, game over, 10th frame bonus rules)
  - `Score() (int, error)` with strike/spare bonus calculation across all 10 frames
  - Helper methods: `rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `isSpare`, `rawFrameScore`, `strikeBonus`, `spareBonus`
- All 31 tests pass (10 roll validation tests + 21 score calculation tests)

## Commit
- `d3b34d7` — feat: implement bowling game scorer for issue #77
