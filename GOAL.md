# Goal: polyglot-go-ledger (Issue #303)

## Problem Statement

Implement a ledger formatter in Go that prints a nicely formatted ledger given a locale (American or Dutch) and a currency (US dollar or euro). This is the Exercism "ledger" exercise, which is a refactoring exercise — the goal is to produce clean, well-structured code.

The implementation must provide:
- An `Entry` struct with `Date` (string), `Description` (string), and `Change` (int, in cents) fields
- A `FormatLedger(currency, locale string, entries []Entry) (string, error)` function

## Acceptance Criteria

1. **Empty ledger**: Returns header row only for valid currency/locale with no entries
2. **Single entry**: Formats a single entry correctly with date, description, and change amount
3. **Sorting**: Entries are sorted by date, then by description, then by change amount
4. **Description truncation**: Descriptions longer than 25 characters are truncated to 22 chars + "..."
5. **Currency symbols**: USD → `$`, EUR → `€`
6. **en-US locale formatting**:
   - Header: `Date       | Description               | Change`
   - Date format: `MM/DD/YYYY`
   - Negative amounts in parentheses: `($10.00)`
   - Positive amounts with trailing space: `$10.00 `
   - Thousand separator: comma (`,`)
   - Decimal separator: period (`.`)
7. **nl-NL locale formatting**:
   - Header: `Datum      | Omschrijving              | Verandering`
   - Date format: `DD-MM-YYYY`
   - Negative amounts with trailing minus: `$ 123,45-`
   - Positive amounts with trailing space: `$ 1.234,56 `
   - Thousand separator: period (`.`)
   - Decimal separator: comma (`,`)
   - Space between currency symbol and amount
8. **Error handling**: Return error for:
   - Empty or invalid currency (not "USD" or "EUR")
   - Empty or invalid locale (not "en-US" or "nl-NL")
   - Invalid date format (must be parseable as `YYYY-MM-DD`)
9. **Immutability**: Must NOT modify the input `entries` slice
10. **All tests pass**: `go test` in the ledger directory passes all test cases

## Key Constraints

- Change values are in cents (e.g., 1000 = $10.00)
- Column widths are fixed: Date=10, Description=25, Change=13
- The output format has specific padding and alignment requirements
- Zero amounts display as positive (e.g., `$0.00 ` not `($0.00)`)
