# Scope: polyglot-go-hexadecimal

## In Scope

- Implement `ParseHex(string) (int64, error)` in `go/exercises/practice/hexadecimal/hexadecimal.go`
- Implement `HandleErrors([]string) []string` in the same file
- Define supporting error types (`ErrRange`, `ErrSyntax`, `ParseError`) as needed by the test suite
- Ensure all tests in `hexadecimal_test.go` pass
- Ensure `go vet` passes

## Out of Scope

- Modifying the test file `hexadecimal_test.go`
- Modifying `go.mod`
- Changes to any other exercises or languages
- Adding new test cases
- Adding benchmarks beyond what already exists

## Dependencies

- Go 1.18+ (specified in go.mod)
- Standard library only: `errors`, `math`, `strings` (no external packages)
- The reference solution in `.meta/example.go` serves as a guide but does not need to be copied verbatim
