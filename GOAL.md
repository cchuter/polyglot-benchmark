# Goal: Implement Bowling Scoring in Go

## Problem Statement

Implement a bowling game scorer in Go. The solution must be in `go/exercises/practice/bowling/bowling.go` and pass all tests defined in `bowling_test.go` and `cases_test.go`.

The game consists of 10 frames. Players roll balls to knock down pins. The scoring rules involve strikes (all 10 pins on first roll), spares (all 10 pins across two rolls), and open frames (fewer than 10 pins total). Strikes and spares receive bonus points from subsequent rolls. The 10th frame has special rules allowing up to 3 rolls.

## Required API

The solution must export:

- `NewGame() *Game` — creates a new bowling game
- `(*Game) Roll(pins int) error` — records a roll, returning an error for invalid rolls
- `(*Game) Score() (int, error)` — returns the total score, or an error if the game is incomplete

## Acceptance Criteria

1. **All 20 score test cases pass** — covering all zeros, no strikes/spares, spares, strikes, consecutive strikes/spares, 10th frame bonuses, perfect game (300), incomplete games, and unstarted games
2. **All 12 roll validation test cases pass** — covering negative rolls, exceeding pin counts, frame pin limits, 10th frame bonus validation, and rolling after game over
3. **`go test ./...` exits with code 0** in the bowling exercise directory
4. **Package is named `bowling`** and uses `go 1.18` module
5. **No modifications to test files** — only `bowling.go` is modified

## Key Constraints

- Must implement `Game` struct type
- Must return proper errors for invalid states (negative pins, too many pins, game over, premature scoring)
- 10th frame allows up to 3 rolls if strike or spare is achieved
- Score can only be calculated after the game is complete (all 10 frames finished)
