# Goal: polyglot-go-bowling (Issue #325)

## Problem Statement

Implement a bowling game scorer in Go. The solution must track rolls and compute the final score according to standard 10-pin bowling rules, including strikes, spares, and the special 10th-frame bonus rolls.

## Acceptance Criteria

1. `NewGame()` returns a `*Game` struct ready to accept rolls.
2. `Game.Roll(pins int) error` records each roll and validates:
   - Negative pin counts return an error.
   - Pin counts > 10 return an error.
   - Two rolls in a frame cannot exceed 10 pins (except 10th frame bonus rules).
   - Rolls after the game is over return an error.
   - 10th frame bonus roll validation (strike/spare fill ball constraints).
3. `Game.Score() (int, error)` returns the total score only when the game is complete:
   - Returns an error if the game is incomplete.
   - Correctly scores open frames, spares (10 + next 1 roll), and strikes (10 + next 2 rolls).
   - Correctly handles the 10th frame with up to 3 rolls.
   - A perfect game (12 strikes) scores 300.
4. All tests in `bowling_test.go` and `cases_test.go` pass (`go test ./...`).
5. Code passes `go vet ./...` with no issues.

## Key Constraints

- The solution must be in the file `go/exercises/practice/bowling/bowling.go`.
- Package name must be `bowling`.
- Must use Go 1.18+ (per `go.mod`).
- The test file and cases file are read-only and define the exact API surface.
