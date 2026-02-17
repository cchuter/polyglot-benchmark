# Goal: Implement Bowling Game Scoring (Go)

## Problem Statement

Implement a bowling game scorer in Go. The solution must provide a `Game` struct with `Roll(pins int) error` and `Score() (int, error)` methods, plus a `NewGame() *Game` constructor. The implementation must correctly score all 10 frames including strikes, spares, open frames, and the special 10th frame bonus rolls.

## Acceptance Criteria

1. `NewGame()` returns a `*Game` ready to accept rolls
2. `Roll(pins)` records each roll and returns `nil` on valid input
3. `Roll(pins)` returns an error for:
   - Negative pin counts
   - Pin counts exceeding 10
   - Two rolls in a frame exceeding 10 pins total
   - Rolls after the game is complete
   - Invalid bonus rolls in the 10th frame
4. `Score()` returns the total score after all frames are complete
5. `Score()` returns an error if the game is not yet complete
6. Correctly handles spare bonuses (next 1 roll)
7. Correctly handles strike bonuses (next 2 rolls)
8. Correctly handles 10th frame bonus rolls (up to 3 rolls)
9. Perfect game (all strikes) scores 300
10. All tests in `bowling_test.go` and `cases_test.go` pass

## Key Constraints

- Package name must be `bowling`
- Must use Go 1.18 (as specified in go.mod)
- Must match the test API: `NewGame() *Game`, `(*Game).Roll(int) error`, `(*Game).Score() (int, error)`
