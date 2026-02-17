# Scope: polyglot-go-bowling

## In Scope

- Implementing the bowling game scorer in `go/exercises/practice/bowling/bowling.go`
- Defining the `Game` struct with necessary fields to track game state
- Implementing `NewGame()`, `Roll()`, and `Score()` methods
- Proper error handling for all invalid inputs and game states
- All 33 test cases (21 score + 12 roll) must pass

## Out of Scope

- Modifying test files (`bowling_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` files
- Adding any additional files beyond `bowling.go`
- Performance optimization beyond correctness
- Any changes to other exercises or languages in the repository

## Dependencies

- Go toolchain (1.18+)
- Standard library only (`errors` package)
- No external dependencies required
