# Scope: Alphametics Solver (Issue #317)

## In Scope

- Implementing `func Solve(puzzle string) (map[string]int, error)` in `alphametics.go`
- Parsing puzzle input format: `"WORD + WORD + ... == WORD"`
- Constraint satisfaction: unique digit assignment, no leading zeros
- Passing all existing test cases in `cases_test.go` and `alphametics_test.go`
- Performance optimization sufficient for the 199-addend, 10-letter test case

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding external dependencies
- Supporting puzzle formats beyond `+` and `==` operators
- Subtraction, multiplication, or other arithmetic operators
- Any changes to other exercises or languages in the repository

## Dependencies

- Go standard library only (no external packages)
- `go 1.18` compatibility (no generics beyond what 1.18 supports)
- Existing test infrastructure (`cases_test.go` defines `testCases`)
