# Goal: Implement Bowling Scoring in Go

## Problem Statement

Implement a bowling game scorer in Go. The solution must track rolls and compute the final score according to standard 10-pin bowling rules, including strikes, spares, and the special 10th frame.

## Acceptance Criteria

1. `NewGame()` returns a `*Game` struct ready to track rolls.
2. `Roll(pins int) error` records each ball thrown:
   - Returns error for negative pin counts.
   - Returns error if pin count exceeds pins remaining on the lane.
   - Returns error if game is already over.
   - Correctly handles the 10th frame's bonus rolls.
3. `Score() (int, error)` returns the total game score:
   - Returns error if the game is not yet complete.
   - Correctly scores open frames, spares, and strikes (including bonuses).
   - Correctly scores the 10th frame (up to 3 rolls, no cascading bonuses).
   - A perfect game (12 strikes) scores 300.
4. All tests in `bowling_test.go` and `cases_test.go` pass (`go test ./...`).
5. Code passes `go vet ./...` with no issues.

## Key Constraints

- Package name must be `bowling`.
- Must use Go 1.18+ (per `go.mod`).
- Must export `Game` type, `NewGame`, `Roll`, and `Score` as defined by the test harness.
