# Changes: ledger.go implementation

## Files Modified
- `go/exercises/practice/ledger/ledger.go` — Complete implementation

## What Changed
- Defined `Entry` struct with `Date`, `Description`, and `Change` fields
- Implemented `FormatLedger(currency, locale string, entries []Entry) (string, error)`:
  - Currency validation: "USD" → "$", "EUR" → "€", else error
  - Locale validation with exact headers for "en-US" and "nl-NL"
  - Date validation via `time.Parse("2006-01-02", ...)`
  - Copies input slice before sorting (preserves caller's slice)
  - Sorts by date, then description, then change
  - Formats dates: en-US as MM/DD/YYYY, nl-NL as DD-MM-YYYY
  - Truncates descriptions >25 chars to 22+"..."
  - Right-aligns amounts to 13 characters
- Implemented `formatAmount` helper:
  - Thousand separators (comma for en-US, period for nl-NL)
  - Decimal separators (period for en-US, comma for nl-NL)
  - Negative wrapping: en-US uses parentheses, nl-NL uses trailing dash

## Test Results
All 17 tests pass (10 success cases, 6 failure cases, 1 mutation guard).
