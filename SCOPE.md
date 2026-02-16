# Scope: polyglot-go-bowling

## In Scope

- Implementing the bowling game scorer in `go/exercises/practice/bowling/bowling.go`
- The `Game` struct with `Roll` and `Score` methods
- `NewGame` constructor function
- Error handling for all invalid states
- Passing all existing tests in `bowling_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` files
- Adding new test cases
- Implementing exercises other than bowling
- Any UI, CLI, or network components

## Dependencies

- Go 1.18+ (as specified in go.mod)
- Standard library only (specifically `errors` package)
- No external dependencies required
