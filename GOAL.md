# Goal: polyglot-go-ledger (Issue #219)

## Problem Statement

Implement a ledger printer in Go that formats financial entries into a nicely formatted ledger table. The solution must support two locales (American `en-US` and Dutch `nl-NL`) and two currencies (US dollar `USD` and euro `EUR`).

The current `ledger.go` file is a stub containing only `package ledger`. The solution must implement the `Entry` struct and `FormatLedger` function to pass all existing tests.

## Acceptance Criteria

1. **Entry struct** is defined with fields:
   - `Date string` (format: "YYYY-MM-DD")
   - `Description string`
   - `Change int` (amount in cents)

2. **FormatLedger function** with signature `FormatLedger(currency, locale string, entries []Entry) (string, error)`:
   - Returns a formatted ledger string with header and entry rows
   - Returns an error for invalid currency, locale, or date formats

3. **Locale support**:
   - `en-US`: Header "Date | Description | Change", date format MM/DD/YYYY, American currency formatting (negative in parentheses)
   - `nl-NL`: Header "Datum | Omschrijving | Verandering", date format DD-MM-YYYY, Dutch currency formatting (negative with trailing minus)

4. **Currency support**:
   - `USD`: Symbol `$`
   - `EUR`: Symbol `€`

5. **Formatting rules**:
   - Amounts use thousands separators (`,` for en-US, `.` for nl-NL)
   - Decimal separators (`.` for en-US, `,` for nl-NL)
   - Negative amounts: `($10.00)` for en-US, `$ 10,00-` for nl-NL
   - Positive amounts: `$10.00 ` for en-US (trailing space), `$ 10,00 ` for nl-NL (trailing space)
   - Description truncated to 22 chars + "..." if longer than 25 chars
   - Column widths: Date=10, Description=25, Change=13

6. **Sorting**: Entries sorted by date, then description, then change amount

7. **Input safety**: FormatLedger must NOT modify the input entries slice

8. **Error handling**: Return error for:
   - Empty or invalid currency
   - Empty or invalid locale
   - Invalid date formats

9. **All tests pass**: `go test ./...` in the ledger directory must pass

## Key Constraints

- Solution file: `go/exercises/practice/ledger/ledger.go`
- Must use `package ledger`
- Go module version: 1.18
- Test file is read-only (do not modify)
- Code should be clean and well-structured (this is a refactoring exercise)
