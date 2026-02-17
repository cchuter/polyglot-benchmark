# Scope: Forth Evaluator (Issue #256)

## In Scope

- Implementing the `Forth` function in `go/exercises/practice/forth/forth.go`
- Integer arithmetic operators: `+`, `-`, `*`, `/`
- Stack manipulation words: `DUP`, `DROP`, `SWAP`, `OVER`
- User-defined word definitions (`: word-name definition ;`)
- Proper error handling for all edge cases defined in tests
- Case-insensitive word handling
- Passing all 42 test cases in `cases_test.go`

## Out of Scope

- Modifying test files (`forth_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Supporting Forth features beyond those tested (loops, conditionals, etc.)
- External dependencies (solution uses only stdlib)
- Other languages or exercises in the repository

## Dependencies

- Go standard library: `errors`, `strconv`, `strings`, `unicode`
- No external packages required
