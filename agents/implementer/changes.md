# Changes: Implement FormatLedger for ledger exercise

## File Modified
- `go/exercises/practice/ledger/ledger.go`

## What Was Implemented
- `Entry` struct with Date, Description, and Change fields
- `FormatLedger(currency, locale string, entries []Entry) (string, error)` main function
- `formatDate` helper for locale-specific date formatting (en-US: MM/DD/YYYY, nl-NL: DD-MM-YYYY)
- `truncateDescription` helper to cap descriptions at 25 chars (22 + "...")
- `formatChange` helper for locale-specific currency formatting with integer arithmetic
- `formatWithThousands` helper for inserting thousands separators

## Key Design Decisions
- Integer-only arithmetic to avoid float precision issues
- Copy-then-sort to preserve input immutability
- Literal header strings for exact test matching
- stdlib only: fmt, sort, strings, time, errors

## Test Results
All 17 tests pass (10 success, 6 failure, 1 input mutation check).
