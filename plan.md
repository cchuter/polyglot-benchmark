# Implementation Plan: polyglot-go-ledger

## File to Modify

- `go/exercises/practice/ledger/ledger.go` — the only file that needs changes

## Architecture

Single file implementation with clean separation of concerns:

### 1. Entry Struct

```go
type Entry struct {
    Date        string // "YYYY-MM-DD"
    Description string
    Change      int    // cents (minor currency units)
}
```

### 2. Locale and Currency Configuration

Use maps to hold locale-specific and currency-specific settings:

- **Currency symbols**: map `"USD"→"$"`, `"EUR"→"€"`
- **Locale headers**: map with literal header strings per locale
- **Locale date formatting**: map with formatting closures or a switch

### 3. FormatLedger Function

```
func FormatLedger(currency string, locale string, entries []Entry) (string, error)
```

Steps:
1. **Validate currency** — lookup in currency symbol map, return error if not found
2. **Validate locale** — check against known locales (`en-US`, `nl-NL`), return error if unknown
3. **Copy entries** — `make([]Entry, len(entries))` + `copy()` to avoid mutating the input
4. **Validate dates** — parse each entry's date with `time.Parse("2006-01-02", ...)`, return error if invalid
5. **Sort entries** — sort by Date string (lexicographic, since YYYY-MM-DD sorts correctly), then Description, then Change (using `sort.Slice` with stable ordering)
6. **Build header** — use literal locale-specific header string + newline
7. **Format each entry** — format date, truncate description, format change amount, build row
8. **Return** — concatenated string via `strings.Builder`

### 4. Helper Functions

- **`formatDate(date string, locale string) string`** — parses "YYYY-MM-DD" and reformats to locale
- **`formatChange(cents int, symbol string, locale string) string`** — formats the monetary amount
- **`truncateDescription(desc string) string`** — truncates/pads to exactly 25 characters

### 5. Explicit Rendering Contract

#### Row Format
Each data row is constructed as:
```
fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, change)
```
- Date: left-aligned, width 10
- Description: left-aligned, width 25
- Change: right-aligned, width 13
- Separator: exactly `" | "`
- Each row ends with `\n`

#### Header Rows (literal strings per locale)
- en-US: `"Date       | Description               | Change\n"`
- nl-NL: `"Datum      | Omschrijving              | Verandering\n"`
These are hardcoded string constants, NOT formatted with Sprintf.

#### Date Formatting
- en-US: `MM/DD/YYYY` (e.g., `01/01/2015`)
- nl-NL: `DD-MM-YYYY` (e.g., `01-01-2015`)

#### Description Truncation
- If `len(description) > 25`: take first 22 characters + `"..."` → always 25 chars
- If `len(description) <= 25`: left-pad/align to 25 chars (Sprintf `%-25s` handles this)

#### Change Formatting — Integer Arithmetic Only (no floats)
1. Compute `negative := cents < 0`
2. Compute absolute value: `abs := cents; if negative { abs = -abs }`
3. Split: `units := abs / 100`, `frac := abs % 100`
4. Format units with thousands separator:
   - Convert units to string
   - Insert separator every 3 digits from the right
   - en-US separator: `,` → e.g., `1,234`
   - nl-NL separator: `.` → e.g., `1.234`
5. Build amount string with decimal:
   - en-US: `units.frac` (decimal = `.`) → e.g., `1,234.56`
   - nl-NL: `units,frac` (decimal = `,`) → e.g., `1.234,56`
6. Apply sign and currency symbol:
   - en-US negative: `"(" + symbol + amount + ")"` → `($1,234.56)`
   - en-US positive/zero: `symbol + amount + " "` → `$1,234.56 ` (trailing space)
   - nl-NL negative: `symbol + " " + amount + "-"` → `$ 1.234,56-`
   - nl-NL positive/zero: `symbol + " " + amount + " "` → `$ 1.234,56 ` (trailing space)

#### Sorting
Sort key order: Date (string comparison) → Description (string comparison) → Change (int comparison). Since dates are validated as YYYY-MM-DD, lexicographic comparison works correctly.

### 6. Implementation Order

1. Define `Entry` struct
2. Define currency symbol map
3. Implement `truncateDescription`
4. Implement `formatDate`
5. Implement thousands-separator helper
6. Implement `formatChange`
7. Implement `FormatLedger` (main function)

## Design Decisions

- **No external dependencies** — only stdlib (`fmt`, `sort`, `strings`, `time`)
- **Integer-only arithmetic** — avoid float precision issues
- **Literal header strings** — exact match with test expectations
- **Copy-then-sort** — ensures input immutability
- **`time.Parse` for validation** — leverages stdlib for robust date validation
- **Lexicographic date sorting** — YYYY-MM-DD format sorts correctly as strings
- **`strings.Builder`** — efficient string concatenation for building output
