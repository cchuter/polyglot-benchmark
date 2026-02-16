# Goal: polyglot-go-ledger (Issue #133)

## Problem Statement

Implement the `FormatLedger` function in `go/exercises/practice/ledger/ledger.go`. This is the Exercism "ledger" exercise — a refactoring exercise where we implement a ledger printer that formats financial entries for different locales (American en-US and Dutch nl-NL) and currencies (USD and EUR).

The stub file currently contains only the package declaration. We must implement the `Entry` struct and `FormatLedger` function to pass all tests.

## Acceptance Criteria

1. **Entry struct** is defined with fields: `Date` (string), `Description` (string), `Change` (int, in minor units/cents).

2. **FormatLedger(currency, locale string, entries []Entry) (string, error)** returns a formatted ledger string or an error.

3. **Locale support**:
   - `en-US`: Header = `Date       | Description               | Change`, date format = `MM/DD/YYYY`, decimal separator = `.`, thousands separator = `,`, negative = `($X.XX)`, positive = `$X.XX ` (trailing space)
   - `nl-NL`: Header = `Datum      | Omschrijving              | Verandering`, date format = `DD-MM-YYYY`, decimal separator = `,`, thousands separator = `.`, negative = `$ X,XX-`, positive = `$ X,XX ` (trailing space, space between symbol and amount)

4. **Currency support**: `USD` → `$`, `EUR` → `€`

5. **Sorting**: Entries sorted by date, then description, then change amount.

6. **Description truncation**: Descriptions longer than 25 characters are truncated to 22 characters + `...`

7. **Error handling**: Returns error for invalid currency, invalid locale, or invalid date format.

8. **Immutability**: `FormatLedger` must NOT modify the input `entries` slice.

9. **All tests pass**: `go test ./...` in the ledger directory passes all success, failure, and immutability tests.

## Key Constraints

- Package name: `ledger`
- Go module: `ledger` with `go 1.18`
- Only modify `ledger.go` — do not modify test files
- Change amounts are in cents (minor currency units): 1000 = $10.00
