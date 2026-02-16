# Scope: polyglot-go-bowling

## In Scope

- Implement `bowling.go` in `go/exercises/practice/bowling/`
- The `Game` struct with state tracking fields
- `NewGame()` constructor
- `Roll(pins int) error` method with full validation
- `Score() (int, error)` method with complete scoring logic
- All error conditions required by the test cases

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying any files outside `go/exercises/practice/bowling/`
- Adding external dependencies
- Creating additional files beyond `bowling.go`

## Dependencies

- Go 1.18+ (as specified in go.mod)
- Standard library `errors` package only
- No external packages
