# Scope: Go Bowling Exercise

## In Scope
- Implement the `Game` struct type in `bowling.go`
- Implement `NewGame() *Game` constructor
- Implement `(g *Game) Roll(pins int) error` method with full input validation
- Implement `(g *Game) Score() (int, error)` method with game-completeness validation
- Handle all bowling scoring rules: open frames, spares, strikes, 10th frame bonus rolls
- Pass all existing tests in `bowling_test.go` and `cases_test.go`

## Out of Scope
- Modifying test files (`bowling_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Adding files beyond `bowling.go`
- Changes to `.docs/` or `.meta/` directories
- Changes to any other exercise in the repository
- Adding external dependencies

## Dependencies
- Go 1.18+ toolchain
- Standard library only (no external packages)
- Existing test infrastructure in `bowling_test.go` and `cases_test.go`
