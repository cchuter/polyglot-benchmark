# Scope: polyglot-go-ledger (Issue #133)

## In Scope

- Implement `Entry` struct in `ledger.go`
- Implement `FormatLedger` function in `ledger.go`
- Locale-specific formatting for `en-US` and `nl-NL`
- Currency symbol mapping for `USD` and `EUR`
- Entry sorting (date → description → change)
- Description truncation at 25 characters
- Input validation (currency, locale, date format)
- Ensuring immutability of input entries slice
- All existing tests must pass

## Out of Scope

- Adding new test cases or modifying existing tests
- Supporting additional locales or currencies beyond en-US/nl-NL and USD/EUR
- Modifying `go.mod`
- Performance optimization beyond passing the benchmark
- Changes to any other exercise in the repository

## Dependencies

- Go standard library only (no external packages)
- Test file (`ledger_test.go`) defines the expected API contract
- `go.mod` already configured with `go 1.18`
