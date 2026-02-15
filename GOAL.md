# Goal: Implement Bowling Scoring in Go

## Problem Statement

Implement a bowling game scorer in Go that tracks rolls and computes the final score according to standard 10-pin bowling rules. The solution must be placed in `go/exercises/practice/bowling/bowling.go`, which currently contains only a package declaration stub.

## What Needs to Be Built

A `Game` type with the following API:

- `NewGame() *Game` — Creates a new bowling game instance
- `(g *Game) Roll(pins int) error` — Records a roll; returns an error for invalid rolls
- `(g *Game) Score() (int, error)` — Returns the total score; only valid after a complete game

## Acceptance Criteria

1. **All tests pass**: `go test` in `go/exercises/practice/bowling/` passes all test cases in `bowling_test.go` and `cases_test.go`
2. **Score test cases pass** (16 cases):
   - All zeros → 0
   - No strikes or spares → correct sum
   - Spare scoring with bonus from next roll
   - Consecutive spares
   - Spare in last frame with fill ball
   - Strike scoring with bonus from next two rolls
   - Consecutive strikes
   - Strike in last frame with fill balls
   - Perfect game (all strikes) → 300
   - Incomplete game → error (premature score)
   - Bonus roll validation errors
3. **Roll validation test cases pass** (9 cases):
   - Negative roll → error
   - Roll > 10 pins → error
   - Two rolls in a frame exceeding 10 → error
   - Cannot roll after game is over → error
   - Bonus roll pin count validation in 10th frame
4. **Code compiles without errors**
5. **Package name is `bowling`**

## Key Constraints

- Must use Go 1.18 (as specified in go.mod)
- Must export `Game` struct, `NewGame`, `Roll`, and `Score` as the public API
- Error handling must use proper Go error values
- Solution goes in `bowling.go` only — do not modify test files
