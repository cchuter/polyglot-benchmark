# Implementation Plan: polyglot-go-ledger

## Branch 1: Direct Implementation (Simplicity-First)

Write a single `ledger.go` file that closely mirrors the reference example solution. Use simple maps for locale/currency config, a straightforward sorting approach with `sort.Slice`, and helper functions for formatting.

### Files to create/modify
- `go/exercises/practice/ledger/ledger.go` (replace stub)

### Approach
1. Define `Entry` struct
2. Define locale config as a struct with date format, translations, currency formatter
3. Store locales and currency symbols in package-level maps
4. Implement `FormatLedger`:
   - Validate currency and locale
   - Copy and sort entries
   - Format header from locale translations
   - Format each entry row (parse date, truncate description, format currency)
5. Helper functions: `moneyToString`, currency formatters per locale

### Evaluation
- **Feasibility**: High. Standard Go, no dependencies.
- **Risk**: Low. Straightforward mapping from example.
- **Alignment**: Fully meets all acceptance criteria.
- **Complexity**: Low. Single file, ~170 lines.

---

## Branch 2: Interface-Based Design (Extensibility)

Define a `Locale` interface with methods for date formatting, currency formatting, and header translations. Each locale becomes a concrete type implementing the interface. Currency symbols stored in a registry.

### Files to create/modify
- `go/exercises/practice/ledger/ledger.go` (replace stub)

### Approach
1. Define `Entry` struct
2. Define `Locale` interface: `DateFormat(time.Time) string`, `CurrencyFormat(symbol string, cents int) string`, `Header() (date, desc, change string)`
3. Implement `americanLocale` and `dutchLocale` types
4. Registry maps for locales and currencies
5. `FormatLedger` uses the interface methods

### Evaluation
- **Feasibility**: High. Standard Go patterns.
- **Risk**: Low, but adds unnecessary abstraction for 2 locales.
- **Alignment**: Fully meets acceptance criteria.
- **Complexity**: Medium. More types and methods, ~200 lines.

---

## Branch 3: Functional Composition (Performance/Unconventional)

Use function composition and closures. Each locale is a collection of functions (date formatter, currency formatter, header provider). Use `strings.Builder` for optimal string building and `sort.Slice` for modern sorting.

### Files to create/modify
- `go/exercises/practice/ledger/ledger.go` (replace stub)

### Approach
1. Define `Entry` struct
2. Define `localeConfig` struct with function fields (no interface)
3. Use `strings.Builder` instead of `bytes.Buffer` for better performance
4. Use `sort.Slice` instead of implementing `sort.Interface`
5. Pre-compute formatted parts using closures

### Evaluation
- **Feasibility**: High. Standard Go.
- **Risk**: Low. `strings.Builder` available since Go 1.10, `sort.Slice` since Go 1.8.
- **Alignment**: Fully meets acceptance criteria.
- **Complexity**: Low-Medium. ~160 lines, slightly more modern Go.

---

## Selected Plan

**Branch 1: Direct Implementation** is selected.

### Rationale
- **Simplicity**: The exercise explicitly says "refactor" — the goal is clean, readable code, not an extensibility framework. Branch 1 achieves this.
- **Risk**: Lowest risk. Closely follows the proven reference solution patterns.
- **Alignment**: Fully satisfies all acceptance criteria with minimum complexity.
- **Branch 2** is over-engineered for 2 locales. Interfaces add abstraction without benefit here.
- **Branch 3** offers marginal performance gains that aren't needed for this exercise, while `strings.Builder` vs `bytes.Buffer` is a minor detail.

### Detailed Implementation Plan

#### File: `go/exercises/practice/ledger/ledger.go`

```go
package ledger
```

**Step 1: Define the Entry struct**
```go
type Entry struct {
    Date        string // "YYYY-MM-DD"
    Description string
    Change      int    // in cents
}
```

**Step 2: Define locale configuration**
```go
type localeInfo struct {
    currency     func(symbol string, cents int, negative bool) string
    dateFormat   string
    translations map[string]string
}
```

With methods:
- `currencyString(symbol string, cents int) string` — handles sign extraction
- `dateString(t time.Time) string` — formats date

**Step 3: Define package-level configuration maps**
- `currencySymbols`: maps "USD" → "$", "EUR" → "€"
- `locales`: maps "en-US" and "nl-NL" to their `localeInfo`

**Step 4: Implement currency formatters**
- `dutchCurrencyFormat(symbol, cents, negative)` — "$ 1.234,56 " or "$ 1.234,56-"
- `americanCurrencyFormat(symbol, cents, negative)` — "$1,234.56 " or "($1,234.56)"

**Step 5: Implement `moneyToString` helper**
- Takes cents (non-negative), thousands separator, decimal separator
- Returns formatted number string like "1,234.56"

**Step 6: Implement sorting**
- Define `entrySlice` type implementing `sort.Interface` (Len, Swap, Less)
- Sort by Date, then Description, then Change

**Step 7: Implement `FormatLedger`**
1. Validate currency → look up symbol
2. Validate locale → look up localeInfo
3. Copy entries slice (to avoid mutating input)
4. Sort the copy
5. Build header line using locale translations
6. For each entry: parse date, truncate description if needed, format currency
7. Return assembled string

**Step 8: Verify**
- Run `go test ./...` to confirm all tests pass
- Run `go vet ./...` for code quality
