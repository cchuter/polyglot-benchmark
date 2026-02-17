# Scope: polyglot-go-ledger (Issue #347)

## In Scope

- Implement the `Entry` struct in `ledger.go`
- Implement `FormatLedger` function in `ledger.go`
- Support `USD` and `EUR` currencies
- Support `en-US` and `nl-NL` locales
- Input validation (currency, locale, date format)
- Entry sorting (date → description → change)
- Locale-specific date formatting
- Locale-specific currency formatting (symbol, separators, negative presentation)
- Description truncation for overlong entries
- Ensure input slice is not mutated

## Out of Scope

- Adding new test cases or modifying `ledger_test.go`
- Supporting additional currencies beyond USD/EUR
- Supporting additional locales beyond en-US/nl-NL
- Modifying `go.mod`
- Performance optimization beyond passing the benchmark
- Any changes outside the `go/exercises/practice/ledger/` directory

## Dependencies

- Go standard library only (`fmt`, `sort`, `strings`, `bytes`, `time`)
- No external packages (module has no dependencies in `go.mod`)
- Test file defines the public API contract (`Entry` struct, `FormatLedger` function signature)
