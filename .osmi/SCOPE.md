# Scope: polyglot-go-octal

## In Scope

- Implementing `ParseOctal` function in `go/exercises/practice/octal/octal.go`
- Handling valid octal strings (digits 0-7) by converting to decimal int64
- Handling invalid input (non-digit characters, digits 8-9) by returning 0 and an error
- Passing all existing tests in `octal_test.go`

## Out of Scope

- Modifying test files (`octal_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` files
- Implementing other exercises
- Overflow handling beyond what tests require
- Negative number support (not tested)

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- Standard library `fmt` or `errors` package for error creation
- No external packages
