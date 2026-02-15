# Goal: Implement Bowling Game Scorer in Go

## Problem Statement

Implement a bowling game scorer in Go at `go/exercises/practice/bowling/bowling.go`. The solution must track rolls and compute the final score according to standard 10-pin bowling rules, including strikes, spares, and the special 10th frame with bonus rolls.

## Required API

The test suite expects the following public API in package `bowling`:

- `NewGame() *Game` — constructor returning a new Game instance
- `(g *Game) Roll(pins int) error` — records a roll; returns an error for invalid rolls
- `(g *Game) Score() (int, error)` — returns the total score; returns an error if the game is incomplete

## Acceptance Criteria

1. **All zeros**: 20 rolls of 0 → score 0
2. **No strikes/spares**: e.g. 10 frames of (3,6) → score 90
3. **Spare scoring**: spare value = 10 + next roll
4. **Strike scoring**: strike value = 10 + next two rolls
5. **Consecutive spares/strikes**: each gets its own bonus
6. **10th frame special cases**: strike gets 2 bonus rolls, spare gets 1 bonus roll; bonus rolls are counted only once
7. **Perfect game**: 12 strikes → score 300
8. **Error: negative pins**: Roll(-1) returns error
9. **Error: pins > 10**: Roll(11) returns error
10. **Error: frame total > 10**: e.g. Roll(5) then Roll(6) returns error
11. **Error: game over**: rolling after game is complete returns error
12. **Error: incomplete game**: Score() before all frames complete returns error
13. **10th frame bonus validation**: bonus rolls must respect pin count constraints
14. All 35 test cases in `cases_test.go` and `bowling_test.go` pass

## Key Constraints

- Package name must be `bowling`
- Must use Go 1.18+ (per go.mod)
- Only modify `bowling.go` — test files are read-only
- The `Game` type must be a struct (tests use `*Game`)
