# Scope: polyglot-go-hexadecimal

## In Scope

- Implement `ParseHex(string) (int64, error)` in `go/exercises/practice/hexadecimal/hexadecimal.go`
- Implement `HandleErrors([]string) []string` in the same file
- Handle valid hex conversion (0-9, a-f, A-F)
- Handle syntax errors (empty string, invalid characters)
- Handle range errors (int64 overflow)
- Pass all existing tests in `hexadecimal_test.go`

## Out of Scope

- Modifying test files
- Adding new test cases
- Supporting hex prefixes like "0x"
- Supporting negative hex values
- Any changes outside the `hexadecimal.go` solution file

## Dependencies

- Go standard library only: `errors`, `math` (for `math.MaxInt64`)
- No external packages
- No built-in hex parsing functions
