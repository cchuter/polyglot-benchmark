# Scope: polyglot-go-ledger (Issue #303)

## In Scope

- Implement the `Entry` struct in `go/exercises/practice/ledger/ledger.go`
- Implement the `FormatLedger` function in the same file
- Support two currencies: USD and EUR
- Support two locales: en-US and nl-NL
- Proper sorting, formatting, error handling, and immutability
- All existing tests in `ledger_test.go` must pass

## Out of Scope

- Modifying the test file (`ledger_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Supporting additional currencies or locales beyond USD/EUR and en-US/nl-NL
- Performance optimization beyond what is needed to pass the benchmark test
- Creating additional files beyond `ledger.go`

## Dependencies

- Go standard library only (no external packages)
- `go.mod` specifies Go 1.18
- Must conform to the `Entry` struct and `FormatLedger` function signatures expected by the test file
