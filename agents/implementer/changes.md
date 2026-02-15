# Implementer Changes

## bowling.go — Full Implementation

Replaced stub code with complete bowling game scorer:

- **Game struct**: Tracks `rolls []int`, `currentFrame`, `currentBall`, and `done` state
- **NewGame()**: Returns initialized `*Game`
- **Roll(pins int) error**: Validates and records each roll with full rule enforcement:
  - Rejects negative pins, pins > 10, and rolls after game over
  - Normal frames (0-8): enforces two-roll pin limit, auto-advances on strikes
  - 10th frame: handles strike/spare bonus rolls with correct pin reset logic
- **Score() (int, error)**: Walks flat rolls slice by frame, applying strike and spare bonuses; returns error if game incomplete

All 31 tests pass (21 score tests + 10 roll validation tests).
