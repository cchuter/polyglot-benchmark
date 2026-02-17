# Implementation Plan: polyglot-go-ledger

## Proposal A — Table-Driven Locale/Currency Configuration

**Role: Proponent**

### Approach

Use a table-driven design where locale and currency details are stored in configuration structs looked up by key. The formatting logic is shared across locales using these configuration values.

### Files to Modify

- `go/exercises/practice/ledger/ledger.go` — single file, complete implementation

### Architecture

1. **Entry struct**: Public fields `Date`, `Description`, `Change` as required by tests
2. **Locale config struct**: Contains header strings, date format function, number format function
3. **Currency config map**: Maps "USD" → "$", "EUR" → "€"
4. **Locale config map**: Maps "en-US" and "nl-NL" to their respective formatting behaviors
5. **FormatLedger function**: Validates inputs, copies+sorts entries, formats each row using config

### Detailed Design

```go
type Entry struct {
    Date        string
    Description string
    Change      int
}

type localeConfig struct {
    dateHeader   string
    descHeader   string
    changeHeader string
    formatDate   func(y, m, d string) string
    formatAmount func(symbol string, cents int) string
}

var currencies = map[string]string{"USD": "$", "EUR": "€"}
var locales = map[string]localeConfig{...}
```

### Ordering of Changes

1. Define Entry struct and config types
2. Implement currency/locale lookup with validation
3. Implement date parsing and formatting
4. Implement amount formatting (with thousand separators, negatives)
5. Implement sorting (date → description → change)
6. Implement row formatting with padding/truncation
7. Assemble FormatLedger function

### Rationale

- Table-driven approach minimizes branching and makes it easy to add new locales/currencies
- Config structs with function fields allow locale-specific formatting without if/else chains
- Single file keeps everything co-located as expected by the exercise structure

---

## Proposal B — Direct Procedural Approach with Switch Statements

**Role: Opponent**

### Approach

Use a straightforward procedural style with switch statements for locale/currency differences. No configuration structs — just direct inline logic.

### Files to Modify

- `go/exercises/practice/ledger/ledger.go` — single file

### Architecture

1. **Entry struct**: Same as Proposal A
2. **FormatLedger function**: Validates inputs, copies+sorts entries, then uses switch statements for locale-specific formatting inline

### Detailed Design

```go
func FormatLedger(currency, locale string, entries []Entry) (string, error) {
    // Validate currency
    var symbol string
    switch currency {
    case "USD": symbol = "$"
    case "EUR": symbol = "€"
    default: return "", errors.New("invalid currency")
    }

    // Validate locale and set header
    var header string
    switch locale {
    case "en-US": header = "Date       | Description               | Change\n"
    case "nl-NL": header = "Datum      | Omschrijving              | Verandering\n"
    default: return "", errors.New("invalid locale")
    }

    // Sort, format rows with switch on locale for each formatting detail
    ...
}
```

### Ordering of Changes

1. Write the full FormatLedger function top-to-bottom
2. Use switch statements for currency symbol, header, date format, amount format
3. Helper functions for number formatting only

### Critique of Proposal A

- Over-engineered for just 2 locales and 2 currencies
- Function fields in config structs add indirection without benefit
- The exercise only requires 2 locales — a map of config structs is unnecessary abstraction
- More code, more types to understand

### Rationale

- Simpler, more direct
- Easier to read and debug
- Less code overall
- Good enough for the scope (2 locales, 2 currencies)

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A (Table-Driven) | Proposal B (Procedural) |
|-----------|--------------------------|------------------------|
| Correctness | Both satisfy all criteria | Both satisfy all criteria |
| Risk | Low risk, but more code = more surface area for bugs | Low risk, simple and direct |
| Simplicity | More abstractions than needed for 2 locales | Straightforward, less code |
| Consistency | Codebase exercises vary in style; both acceptable | Slightly more aligned with simpler exercises like food-chain |

### Decision

**Selected: Hybrid approach** — Take the best of both proposals. Use a simple procedural approach (Proposal B's simplicity) but extract number formatting and date formatting into clean helper functions (Proposal A's structure). No config structs or function-valued fields — just well-organized helper functions.

### Final Implementation Plan

**File**: `go/exercises/practice/ledger/ledger.go`

#### 1. Define the Entry struct

```go
type Entry struct {
    Date        string
    Description string
    Change      int
}
```

#### 2. FormatLedger function skeleton

```go
func FormatLedger(currency, locale string, entries []Entry) (string, error) {
    // a) Validate currency → get symbol
    // b) Validate locale → get header
    // c) Validate all entry dates
    // d) Copy and sort entries
    // e) Build output: header + formatted rows
    // f) Return result
}
```

#### 3. Currency validation

Simple switch: "USD" → "$", "EUR" → "€", else error.

#### 4. Locale validation and header

Simple switch with exact header strings (note: header's Change column is NOT padded to 13 chars):
- "en-US": `"Date       | Description               | Change\n"` (47 chars + newline)
- "nl-NL": `"Datum      | Omschrijving              | Verandering\n"` (52 chars + newline)
- else: error

Column widths for data rows: Date=10, Description=25, Change=13 (right-aligned).

#### 5. Date parsing and validation

Helper function `parseDate(date string) (int, int, int, error)` that:
- Validates format is exactly `YYYY-MM-DD` (10 chars, dashes at positions 4 and 7)
- Uses `time.Parse` with layout `"2006-01-02"` for full validation
- Returns year, month, day

#### 6. Date formatting

Helper function `formatDate(locale string, y, m, d int) string`:
- "en-US": `MM/DD/YYYY` with zero-padded month and day
- "nl-NL": `DD-MM-YYYY` with zero-padded day and month

#### 7. Amount formatting

Helper function `formatAmount(locale, symbol string, cents int) string`:
- Determine if negative, work with absolute value
- Split into whole dollars and cents
- Format whole dollars with thousand separator (comma for en-US, period for nl-NL)
- Format cents with decimal separator (period for en-US, comma for nl-NL)
- Apply locale-specific negative/positive wrapping:
  - en-US negative: `(symbol + amount)`
  - en-US positive: `symbol + amount + " "`
  - nl-NL negative: `symbol + " " + amount + "-"`
  - nl-NL positive: `symbol + " " + amount + " "`
- Right-align (left-pad with spaces) to 13 characters total using `fmt.Sprintf("%13s", amount)`

#### 8. Description formatting

Inline: if len > 25, truncate to 22 + "...", else right-pad to 25.

#### 9. Sorting

Copy the entries slice using `make([]Entry, len(entries))` + `copy()` (not slice assignment, which shares the underlying array). Sort by:
1. Date (string comparison works since YYYY-MM-DD)
2. Description (alphabetical)
3. Change (numeric)

Use `sort.Slice` with a multi-level comparator.

#### 10. Row assembly

For each sorted entry: `formatDate + " | " + formatDescription + " | " + formatAmount + "\n"`

### Implementation Order

1. Entry struct
2. `parseDate` helper
3. `formatDate` helper
4. `formatAmount` helper (most complex piece)
5. `FormatLedger` main function (validation, sorting, row assembly)
6. Run tests, iterate
