# Scope: polyglot-go-bowling (Issue #325)

## In Scope

- Implementing the `Game` struct, `NewGame()`, `Roll()`, and `Score()` in `bowling.go`.
- All validation logic required by the test cases (negative rolls, excess pins, game-over detection, incomplete game scoring).
- Passing all existing tests in `bowling_test.go` and `cases_test.go`.

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`).
- Modifying `go.mod`.
- Adding any new test files or documentation.
- Changes to any other exercise or language directory.
- Any external dependencies beyond the Go standard library.

## Dependencies

- Go standard library only (`errors` package).
- No external packages required.
- Tests are pre-written and define the contract.
