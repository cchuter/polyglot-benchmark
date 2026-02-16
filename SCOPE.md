# Scope: polyglot-go-ledger (Issue #219)

## In Scope

- Implementing the `Entry` struct in `ledger.go`
- Implementing the `FormatLedger` function in `ledger.go`
- All supporting types, functions, and data structures needed
- Ensuring all tests in `ledger_test.go` pass
- Writing clean, idiomatic Go code

## Out of Scope

- Modifying `ledger_test.go` (read-only test file)
- Modifying `go.mod`
- Adding new locales or currencies beyond en-US, nl-NL, USD, EUR
- Adding external dependencies
- Performance optimization beyond reasonable code quality
- Creating additional files beyond `ledger.go`

## Dependencies

- Go standard library only (fmt, sort, strings, time, bytes)
- No external packages
- Must compile with Go 1.18+
