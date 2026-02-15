# Implementation Plan: Go Bowling Exercise

## File to Modify
- `go/exercises/practice/bowling/bowling.go` — the only file that needs changes

## Architecture

### Game Struct Design

```go
type Game struct {
    rolls       []int // all rolls recorded
    currentBall int   // 0 = first ball of frame, 1 = second ball
    currentFrame int  // 0-indexed frame number (0-9)
    done        bool  // game is complete
}
```

The key design decision is to store all rolls in a flat slice and compute the score by walking through the rolls at scoring time, rather than tracking frames explicitly during rolls. This simplifies scoring logic because spare/strike bonuses naturally reference subsequent rolls.

However, we need `currentFrame` and `currentBall` during `Roll()` to validate inputs (e.g., two rolls in a frame can't exceed 10 pins, game-over detection).

### Roll Method Design

`Roll(pins int)` validates and records each roll:

1. **Game over check**: If `done` is true, return error
2. **Basic pin validation**: pins must be 0-10
3. **Frame-aware validation**:
   - **Frames 0-8 (normal frames)**:
     - First ball: pins must be 0-10
     - Second ball: pins + first ball must be <= 10
     - Strike (first ball = 10): advance to next frame immediately
     - Otherwise: advance ball, and if second ball, advance frame
   - **Frame 9 (10th frame)**:
     - First ball: pins must be 0-10
     - Second ball: if first was strike, pins must be 0-10; otherwise first+second <= 10
     - Third ball: depends on first two — if first two are strikes, 0-10; if second was strike after non-strike first, 0-10; if spare (first+second=10), 0-10; otherwise no third ball (but this case means game ended after second ball)
     - After appropriate number of balls, set `done = true`

### 10th Frame Validation Details

The 10th frame is the trickiest part. Tracking:
- `tenthBall` counter (0, 1, 2) within the 10th frame
- After first ball: if strike, pins reset to 10
- After second ball:
  - If first was strike and second is strike: pins reset to 10, allow third
  - If first was strike and second is not strike: remaining = 10 - second, allow third with remaining pins constraint
  - If first + second = 10 (spare): pins reset to 10, allow third
  - If first + second < 10: game over (no bonus)
- After third ball: game always over

### Score Method Design

`Score()` computes the total:

1. If `!done`, return error (game not complete)
2. Walk through rolls by frame:
   ```
   rollIndex = 0
   score = 0
   for frame 0..9:
     if strike: score += 10 + rolls[rollIndex+1] + rolls[rollIndex+2]; rollIndex += 1
     else if spare: score += 10 + rolls[rollIndex+2]; rollIndex += 2
     else: score += rolls[rollIndex] + rolls[rollIndex+1]; rollIndex += 2
   ```
3. Return score, nil

## Implementation Approach

The implementation will be done in a single file edit to `bowling.go`:

1. Define the `Game` struct with fields: `rolls []int`, `currentFrame int`, `currentBall int`, `done bool`
2. Implement `NewGame()` to create and return an initialized `Game`
3. Implement `Roll()` with validation logic for normal and 10th frames
4. Implement `Score()` with frame-walking computation

## Error Messages

Errors should be descriptive but the tests only check `err != nil` vs `err == nil`, so exact messages don't matter. We'll use:
- "negative roll is invalid"
- "pin count exceeds pins on the lane"
- "cannot roll after game is over"
- "score cannot be taken until the end of the game"

## Edge Cases

Key edge cases handled:
- Perfect game (12 strikes = 300)
- All spares with varying bonus rolls
- 10th frame strike followed by non-strike, limiting second bonus roll
- Game completeness checks for scoring
- Preventing rolls after game is over
