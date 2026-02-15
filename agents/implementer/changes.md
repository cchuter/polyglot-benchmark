# Bowling Implementation Changes

## Commit: 256f8aa - feat: implement bowling game scoring (issue #40)

### File Changed
- `go/exercises/practice/bowling/bowling.go`

### Summary
Implemented the complete bowling game scoring logic based on the reference solution. The implementation includes:

- **Game struct**: Tracks rolls, roll count, completed frames, and frame start index using a fixed-size array for up to 21 rolls.
- **NewGame()**: Constructor returning a zero-valued Game.
- **Roll(pins int) error**: Records each roll with full validation:
  - Rejects negative pin counts
  - Rejects pin counts exceeding 10
  - Rejects rolls after game is over
  - Handles strikes (auto-completes frame for frames 1-9)
  - Validates pin totals per frame (cannot exceed 10 for non-last frames)
  - Special handling for the 10th frame: allows up to 3 rolls for strikes/spares with proper validation
- **Score() (int, error)**: Computes the final score after all frames are complete:
  - Strike bonus: next 2 rolls added
  - Spare bonus: next 1 roll added
  - Returns error if game is incomplete
- Helper methods for frame state tracking
