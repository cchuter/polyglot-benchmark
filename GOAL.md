# Goal: polyglot-go-bowling

## Problem Statement

Implement a bowling game scorer in Go. The file `go/exercises/practice/bowling/bowling.go` currently contains only a package declaration stub. It must be completed to pass the existing test suite in `bowling_test.go` and `cases_test.go`.

The implementation must support two operations:
- `Roll(pins int) error` — called each time the player rolls a ball; validates input and tracks game state.
- `Score() (int, error)` — called at the end of the game; returns the total score.

A `NewGame() *Game` constructor must also be provided to create a fresh game instance.

## Acceptance Criteria

1. **All tests pass**: `go test ./...` in the `go/exercises/practice/bowling/` directory must produce zero failures.
2. **Score test cases** (21 cases): Correctly score games with all zeros, no strikes/spares, spares, strikes, consecutive spares/strikes, 10th frame bonuses, perfect game (300), and error cases (incomplete game, premature scoring).
3. **Roll validation test cases** (12 cases): Reject negative rolls, rolls exceeding 10, frame totals exceeding 10, invalid bonus rolls in the 10th frame, and rolls after game completion.
4. **API conformance**: The types and methods must match what the test files expect:
   - `type Game struct` (exported)
   - `func NewGame() *Game`
   - `func (g *Game) Roll(pins int) error`
   - `func (g *Game) Score() (int, error)`

## Key Constraints

- Only `bowling.go` should be modified; test files are read-only.
- Must use Go 1.18+ compatible syntax (per `go.mod`).
- The module name is `bowling` (per `go.mod`).
- A reference solution exists at `.meta/example.go` for validation, but the implementation should be original and clean.
