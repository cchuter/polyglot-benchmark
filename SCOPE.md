# Scope: polyglot-go-bowling

## In Scope

- Implement the `Game` struct, `NewGame()`, `Roll()`, and `Score()` in `go/exercises/practice/bowling/bowling.go`
- Handle all standard bowling scoring rules (open frames, spares, strikes, 10th frame bonuses)
- Handle all error/validation cases tested by the test suite
- Pass all tests in `bowling_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`, `go.mod`)
- Adding additional files to the bowling exercise
- Modifying any other exercises or code outside the bowling directory
- Adding external dependencies

## Dependencies

- Go toolchain (1.18+)
- No external packages — standard library only (just `errors` or `fmt` for error returns)
