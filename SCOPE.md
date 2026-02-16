# Scope: Forth Evaluator (Issue #170)

## In Scope

- Implement the `Forth` function in `go/exercises/practice/forth/forth.go`
- Support arithmetic: `+`, `-`, `*`, `/`
- Support stack manipulation: `DUP`, `DROP`, `SWAP`, `OVER`
- Support user-defined words with `: name definition ;` syntax
- Proper error handling for all edge cases defined in test cases
- Case-insensitive word matching

## Out of Scope

- Modifying test files (`forth_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Supporting any Forth features beyond what the tests require
- Floating-point arithmetic
- String handling
- Any language other than Go

## Dependencies

- Go standard library only (`errors`, `strconv`, `strings`, `unicode`)
- No external packages needed
- Module: `forth` with `go 1.18`
