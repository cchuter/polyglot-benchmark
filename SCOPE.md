# Scope: polyglot-go-hexadecimal (Issue #299)

## In Scope

- Implementing `ParseHex(string) (int64, error)` in `hexadecimal.go`
- Implementing `HandleErrors([]string) []string` in `hexadecimal.go`
- Defining error types/values for syntax and range errors
- Handling uppercase and lowercase hex digits (a-f, A-F, 0-9)
- Detecting int64 overflow during conversion
- Passing all existing tests in `hexadecimal_test.go`

## Out of Scope

- Modifying `hexadecimal_test.go` (read-only test file)
- Modifying `go.mod` (already configured)
- Supporting hex prefixes like "0x" or "0X"
- Supporting negative hex values
- Supporting non-hex bases
- Adding external dependencies
- Performance optimization beyond what's needed to pass the benchmark

## Dependencies

- Go standard library: `errors`, `math` (for `math.MaxInt64` constant only — used for overflow detection, not for parsing)
- No external packages
