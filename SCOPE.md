# Scope: Forth Evaluator (Issue #297)

## In Scope

- Implementing the `Forth()` function in `go/exercises/practice/forth/forth.go`
- All supporting types, helper functions, and error variables needed by the implementation
- Integer arithmetic: `+`, `-`, `*`, `/`
- Stack manipulation: `DUP`, `DROP`, `SWAP`, `OVER`
- User-defined word support: `: word-name definition ;`
- Case-insensitive word handling
- Proper error handling for all error conditions tested
- Passing all 45 existing test cases

## Out of Scope

- Modifying test files (`forth_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` files
- Supporting Forth features beyond what the tests require (loops, conditionals, etc.)
- Any changes outside the `go/exercises/practice/forth/` directory
- Documentation files beyond what's required by the workflow

## Dependencies

- Go standard library only (`strconv`, `strings`, `errors`, `unicode`)
- No external packages required
- Go 1.18+ toolchain for building and testing
