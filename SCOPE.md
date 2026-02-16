# Scope: Forth Evaluator (Issue #213)

## In Scope

- Implementing the `Forth` function in `go/exercises/practice/forth/forth.go`
- Integer arithmetic operators: `+`, `-`, `*`, `/`
- Stack manipulation words: `DUP`, `DROP`, `SWAP`, `OVER`
- User-defined word support with `: word-name definition ;` syntax
- Case-insensitive word handling
- Error handling for all edge cases defined in test cases
- Passing all existing tests in `forth_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`forth_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Supporting Forth features beyond what the tests require (loops, conditionals, strings, etc.)
- Supporting floating-point numbers
- Adding new test cases
- Changes to any files outside `go/exercises/practice/forth/forth.go`

## Dependencies

- Go standard library only (`errors`, `strconv`, `strings`, `unicode`)
- No external packages required
- Module: `forth` with `go 1.18`
