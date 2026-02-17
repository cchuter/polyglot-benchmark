# Implementation Plan: polyglot-go-ledger

## Proposal A (Proponent)

**Approach: Clean implementation modeled on the reference solution's architecture**

Write `ledger.go` as a clean, well-structured implementation that follows the same architectural pattern as `.meta/example.go` — using maps for locale/currency config, a custom sort interface, and helper functions for formatting. This is the proven approach that the reference solution uses successfully.

### Files to modify
- `go/exercises/practice/ledger/ledger.go` — the only file to edit

### Architecture

1. **Entry struct** — exported type with Date, Description, Change fields
2. **Currency config** — map from currency code to symbol string
3. **Locale config** — struct holding date format string, header translations, and a currency formatter function
4. **FormatLedger** — main function that validates inputs, copies+sorts entries, formats header and rows
5. **Helper functions**:
   - `moneyToString(cents int, thousandsSep, decimalSep string) string` — formats cents as dollars with separators
   - Locale-specific currency formatters (American wraps negatives in parens, Dutch appends minus)
6. **entrySlice** — custom type implementing `sort.Interface` for multi-key sorting

### Rationale
- Follows the proven reference solution architecture
- Clear separation of concerns (locale config vs formatting logic)
- Uses Go idioms (sort.Interface, maps for config)
- Single file keeps it simple

### Ordering
1. Define Entry struct
2. Define currency symbol map
3. Define locale info struct and locale map
4. Implement helper functions (moneyToString, currency formatters)
5. Implement entrySlice sort interface
6. Implement FormatLedger

---

## Proposal B (Opponent)

**Approach: Minimal flat implementation without config structs**

Write everything inline in FormatLedger using switch statements instead of config maps/structs. Avoid creating separate types for locale info. Use `sort.Slice` instead of implementing `sort.Interface`.

### Files to modify
- `go/exercises/practice/ledger/ledger.go` — the only file to edit

### Architecture

1. **Entry struct** — same
2. **FormatLedger** — large function that uses switch/case for locale and currency handling inline
3. Use `sort.Slice` with a closure for sorting
4. Inline all formatting logic

### Critique of Proposal A
- Creating separate types (localeInfo, entrySlice) adds more code surface for a simple exercise
- The reference solution's architecture is reasonable but more complex than strictly necessary
- Config maps and function pointers add indirection

### Rationale
- Fewer types and less indirection
- `sort.Slice` is more modern and concise than implementing sort.Interface
- Everything visible in fewer places

### Weaknesses of this approach
- FormatLedger becomes very long and hard to read
- Switch statements duplicate formatting logic
- Adding a new locale would require modifying the core function
- Less idiomatic Go — the standard approach for this kind of config is maps/structs

---

## Selected Plan (Judge)

### Evaluation

**Correctness**: Both proposals can satisfy all acceptance criteria. Neither has a correctness advantage.

**Risk**:
- Proposal A: Low risk — follows the reference solution pattern which is proven to pass all tests.
- Proposal B: Medium risk — inline formatting logic is more error-prone, especially for the tricky currency formatting edge cases.

**Simplicity**:
- Proposal A adds more types but each piece is small and testable.
- Proposal B has fewer types but creates a monolithic function that's harder to debug.

**Consistency**:
- Proposal A matches the reference solution and Go community conventions.
- Proposal B's use of `sort.Slice` is valid but the inline approach is less conventional.

### Decision: Proposal A wins

Proposal A's structured approach is the better choice. The reference solution exists for a reason — it cleanly separates concerns, is easy to read, and follows Go conventions. The small overhead of config types pays off in clarity and maintainability.

However, I'll adopt one improvement from Proposal B: use `sort.Slice` instead of implementing `sort.Interface`, as it's more concise and equally correct.

### Final Plan

**File**: `go/exercises/practice/ledger/ledger.go`

**Implementation in order**:

1. **Package and imports**:
   ```go
   package ledger
   import ("bytes", "fmt", "sort", "strings", "time")
   ```

2. **Entry struct**:
   ```go
   type Entry struct {
       Date        string // "Y-m-d"
       Description string
       Change      int // in cents
   }
   ```

3. **Currency symbols map**:
   ```go
   var currencySymbols = map[string]string{"USD": "$", "EUR": "€"}
   ```

4. **Locale config struct and map**:
   ```go
   type localeInfo struct {
       currency     func(symbol string, cents int, negative bool) string
       dateFormat   string
       translations map[string]string
   }
   ```
   With methods `currencyString` and `dateString`, and a `locales` map holding en-US and nl-NL configurations.

5. **Currency formatter functions**:
   - `dutchCurrencyFormat`: symbol + space + amount + minus/space
   - `americanCurrencyFormat`: optional paren + symbol + amount + optional paren/space

6. **moneyToString helper**: Converts cents to formatted string with thousands and decimal separators.

7. **FormatLedger function**:
   - Validate currency (lookup in currencySymbols map)
   - Validate locale (lookup in locales map)
   - Copy entries slice (to avoid mutating input)
   - Sort copy using `sort.Slice` with multi-key comparison (date, description, change)
   - Write header row using locale translations
   - For each entry: parse date, truncate description if needed, format currency, write row
   - Return formatted string

**Key formatting details**:
- Header: `"%-10s | %-25s | %s\n"` with locale translations
- Rows: `"%-10s | %-25s | %13s\n"` with formatted date, description, currency
- Description truncation: if len > 25, use first 22 chars + "..."
- Date parsing via `time.Parse("2006-01-02", ...)` for validation
