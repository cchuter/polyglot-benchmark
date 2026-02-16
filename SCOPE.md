# Scope: polyglot-go-hexadecimal

## In Scope

- Implementing `ParseHex(string) (int64, error)` in `go/exercises/practice/hexadecimal/hexadecimal.go`
- Implementing `HandleErrors([]string) []string` in the same file
- Error types and sentinel errors as needed
- Overflow detection for int64 range

## Out of Scope

- Modifying `hexadecimal_test.go` — test file is read-only
- Modifying `go.mod` — module definition is fixed
- Other exercises in the repository
- Non-Go language implementations
- Benchmarking optimizations (benchmark test exists but is not an acceptance criterion)
- Supporting bases other than 16

## Dependencies

- Go standard library: `errors`, `math` (for `math.MaxInt64`)
- No external packages
- No built-in hex parsing (`strconv.ParseInt`, `fmt.Sscanf`, etc.)
