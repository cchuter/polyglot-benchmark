# Scope: polyglot-go-octal (Issue #225)

## In Scope

- Implementing the `ParseOctal(string) (int64, error)` function in `go/exercises/practice/octal/octal.go`
- Handling valid octal strings (digits 0-7 only)
- Returning errors for invalid input (non-octal characters, digits 8-9)
- Ensuring all existing tests pass

## Out of Scope

- Modifying test files (`octal_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Overflow handling (not tested)
- Empty string handling (not tested)
- Supporting octal prefixes like `0o` or `0` (not tested)
- Changes to any other exercises

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- Only Go standard library (`fmt` package for error formatting)
- No external packages
