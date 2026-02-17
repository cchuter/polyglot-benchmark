# Goal: polyglot-go-ledger (Issue #347)

## Problem Statement

Implement a ledger printer in Go that formats financial entries into a tabular report. The implementation must support two locales (American `en-US` and Dutch `nl-NL`) and two currencies (US Dollar `USD` and Euro `EUR`), each with locale-specific formatting for dates, currency symbols, thousands separators, decimal separators, and negative number presentation.

The solution file `ledger.go` is currently a stub (package declaration only). The reference implementation exists in `.meta/example.go` and all tests are already written in `ledger_test.go`.

## Acceptance Criteria

1. **Entry type defined**: `Entry` struct with `Date string`, `Description string`, `Change int` (cents) fields.
2. **FormatLedger function**: `func FormatLedger(currency, locale string, entries []Entry) (string, error)` is exported and callable.
3. **All success test cases pass** (10 cases): empty ledger, single entry, credit/debit sorting, same-date ordering by description, tie-breaking by change, overlong description truncation, euros, Dutch locale, Dutch negative numbers, American negative numbers.
4. **All failure test cases pass** (6 cases): empty currency, invalid currency, empty locale, invalid locale, invalid date (bad month), invalid date (wrong separator).
5. **Input immutability**: `FormatLedger` must NOT modify the input `entries` slice.
6. **Benchmark runs**: `BenchmarkFormatLedger` executes without error.
7. **Code quality**: `go vet` passes with no warnings.

## Key Requirements

- **Date formatting**: en-US uses `MM/DD/YYYY`, nl-NL uses `DD-MM-YYYY`
- **Currency symbols**: USD → `$`, EUR → `€`
- **Thousands separators**: en-US uses `,`, nl-NL uses `.`
- **Decimal separators**: en-US uses `.`, nl-NL uses `,`
- **Negative numbers**: en-US wraps in parentheses `($10.00)`, nl-NL appends minus `$ 10,00-`
- **Positive numbers**: trailing space for alignment
- **Description truncation**: max 25 chars, truncate to 22 + `...`
- **Sorting**: by date (ascending), then description (ascending), then change (ascending)
- **Column widths**: Date 10, Description 25, Change 13 (right-aligned)
- **Header row**: locale-specific translations (Date/Datum, Description/Omschrijving, Change/Verandering)
