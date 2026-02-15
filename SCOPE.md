# Scope: polyglot-go-bowling

## In Scope

- Implementing the bowling game scorer in `go/exercises/practice/bowling/bowling.go`
- The `Game` struct with internal state tracking
- `NewGame()` constructor function
- `Roll(pins int) error` method with full input validation
- `Score() (int, error)` method with complete scoring logic
- Error variables for all error conditions tested
- Passing all existing tests in `bowling_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying any `.meta/` or `.docs/` files
- Modifying any other exercise's files
- Adding external dependencies
- Adding benchmark tests
- Creating additional test files
- Changes to any non-Go files in the repository

## Dependencies

- Go 1.18+ toolchain (for building and testing)
- Standard library `errors` package (for error values)
- No external packages required
