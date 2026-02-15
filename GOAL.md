# Goal: Implement Bowling Game Scoring (Issue #40)

## Problem Statement

Implement a bowling game scorer in Go. The solution must implement the `Game` struct with `NewGame()`, `Roll(pins int) error`, and `Score() (int, error)` methods in `go/exercises/practice/bowling/bowling.go`.

The current file is a stub containing only `package bowling`. The reference solution exists in `.meta/example.go` and the test suite is already provided in `bowling_test.go` and `cases_test.go`.

## Acceptance Criteria

1. **All score test cases pass** — 20 test cases covering:
   - All zeros game (score 0)
   - No strikes/spares game (score 90)
   - Spare scoring with bonus rolls
   - Consecutive spares
   - Spare in last frame with bonus roll
   - Strike scoring with bonus rolls
   - Consecutive strikes
   - Strike in last frame with bonus rolls
   - Perfect game (score 300)
   - Incomplete/unstarted game errors

2. **All roll validation test cases pass** — 12 test cases covering:
   - Negative pin count rejection
   - Pin count > 10 rejection
   - Two rolls in frame > 10 rejection
   - Bonus roll validation in last frame
   - Cannot roll after game is over

3. **`go test` passes with zero failures** in the `go/exercises/practice/bowling/` directory.

4. **Code is in `bowling.go`** in the `bowling` package — no other files modified.

## Key Constraints

- Must implement `Game` struct, `NewGame() *Game`, `Roll(pins int) error`, `Score() (int, error)`
- Must return appropriate errors for invalid rolls and premature scoring
- Must correctly handle the 10th frame bonus rolls (spare gets 1 bonus, strike gets 2 bonus)
- Must use Go 1.18+ (per go.mod)
