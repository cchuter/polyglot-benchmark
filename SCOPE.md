# Scope: polyglot-go-octal (Issue #309)

## In Scope

- Implement `ParseOctal` function in `go/exercises/practice/octal/octal.go`
- Function must convert octal string to int64 decimal value
- Proper error handling for invalid input
- All existing tests in `octal_test.go` must pass

## Out of Scope

- Modifying test files (`octal_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` files
- Adding new test cases
- Implementing other exercises
- Changes to any files outside `go/exercises/practice/octal/octal.go`

## Dependencies

- Go 1.18+ (as specified in go.mod)
- Only Go standard library packages (e.g., `fmt` for error formatting)
- No external dependencies
