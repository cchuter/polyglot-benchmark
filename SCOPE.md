# Scope: Bowling Exercise

## In Scope

- Implementing the `Game` struct, `NewGame`, `Roll`, and `Score` in `bowling.go`
- Handling all standard bowling scoring rules (open frames, spares, strikes)
- Handling the special 10th frame (bonus rolls for strikes and spares)
- Input validation (negative pins, too many pins, game-over detection)
- Passing all existing test cases in `bowling_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding additional test cases
- Any files outside `go/exercises/practice/bowling/bowling.go`
- UI, CLI, or any interface beyond the programmatic API
- Concurrency safety (tests are single-threaded)

## Dependencies

- Go standard library only (`errors` package)
- No external dependencies required
- Tests use only the standard `testing` package
