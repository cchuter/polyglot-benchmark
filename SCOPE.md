# Scope: Issue #77 - polyglot-go-bowling

## In Scope

- Implementing the bowling game logic in `go/exercises/practice/bowling/bowling.go`
- Exported types and functions: `Game` struct, `NewGame()`, `Roll()`, `Score()`
- Error handling with descriptive error messages for:
  - Negative pin values
  - Pin count exceeding available pins
  - Premature scoring (game not complete)
  - Rolling after game is over
- Frame tracking including 10th frame special rules
- Strike and spare bonus calculation
- All 36+ test cases passing

## Out of Scope

- Modifying any test files (`bowling_test.go`, `cases_test.go`)
- Modifying exercise metadata (`.meta/` directory)
- Modifying `go.mod`
- Adding external dependencies
- Creating additional Go files beyond `bowling.go`
- Changes to any other exercise directories
- Changes to non-Go language directories

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- Standard library only (`errors` package)
- No external packages required

## Files to Modify

| File | Action |
|------|--------|
| `go/exercises/practice/bowling/bowling.go` | Replace stub with full implementation |

## Files Read-Only (reference)

| File | Purpose |
|------|---------|
| `bowling_test.go` | Test harness |
| `cases_test.go` | Auto-generated test cases |
| `.meta/example.go` | Reference implementation |
