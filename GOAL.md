# Goal: Implement Bowling Scoring in Go

## Problem Statement

Implement a bowling game scorer in Go. The solution must track rolls and compute
the final score according to standard bowling rules, including strikes, spares,
and the special 10th frame logic.

## Required API

The solution must implement these types/functions in `go/exercises/practice/bowling/bowling.go`:

- `Game` struct — tracks the state of a bowling game
- `NewGame() *Game` — creates a new game
- `(g *Game) Roll(pins int) error` — records a roll; returns error for invalid rolls
- `(g *Game) Score() (int, error)` — returns the final score; errors if game is incomplete

## Acceptance Criteria

1. All tests in `bowling_test.go` and `cases_test.go` pass (`go test ./...`)
2. **Score tests** (16 cases): all zeros, no strikes/spares, spares, strikes, consecutive strikes/spares, 10th frame bonus rolls, perfect game (300)
3. **Roll validation tests** (11 cases): negative pins, pins > 10, frame total > 10, 10th frame bonus validation, game-over detection
4. Error cases return non-nil error; valid cases return nil error
5. `Score()` returns error when called before game is complete

## Key Constraints

- Package name: `bowling`
- Module: `bowling` (go 1.18)
- Must handle 10th frame fill balls correctly
- Must validate pin counts against remaining pins in the current frame
- Must detect game-over state and reject further rolls
