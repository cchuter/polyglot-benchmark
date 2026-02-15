# Scope: Bowling Game (Issue #40)

## In Scope

- Implement bowling game logic in `go/exercises/practice/bowling/bowling.go`
- `Game` struct with state tracking for rolls, frames, and game completion
- `NewGame()` constructor returning a `*Game`
- `Roll(pins int) error` method with full input validation
- `Score() (int, error)` method with game-completion checking and correct scoring
- All 32 test cases (20 score + 12 roll) must pass

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding external dependencies
- Modifying any files outside `go/exercises/practice/bowling/bowling.go`
- UI, CLI, or any interface beyond the programmatic API
- Performance optimization beyond correctness

## Dependencies

- Go 1.18+ toolchain
- No external packages — standard library only (`errors` package)
- Existing test harness in `bowling_test.go` and `cases_test.go`
