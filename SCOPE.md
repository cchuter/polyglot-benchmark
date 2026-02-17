# Scope: Forth Evaluator (Issue #341)

## In Scope

- Implement the `Forth(input []string) ([]int, error)` function in `go/exercises/practice/forth/forth.go`
- Support arithmetic operators: `+`, `-`, `*`, `/`
- Support stack manipulation: `DUP`, `DROP`, `SWAP`, `OVER`
- Support user-defined words with `: word-name definition ;`
- Case-insensitive word handling
- Proper error handling for all edge cases defined in test cases
- Pass all 42 test cases in `cases_test.go` and `forth_test.go`

## Out of Scope

- Modifying test files (`forth_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Supporting Forth features beyond the specified subset (loops, conditionals, strings, etc.)
- Performance optimization beyond what's needed to pass tests
- Any changes to other exercises or non-Go code

## Dependencies

- Go standard library only (`errors`, `strconv`, `strings`, `unicode`)
- No external packages required
