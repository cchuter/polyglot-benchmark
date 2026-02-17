# Challenger Review: ledger.go Implementation

## Verdict: PASS - No issues found

The implementation is correct and should pass all test cases.

---

## Test Case Trace-Through

### Success Cases

1. **Empty ledger** - Returns header only. `entries` is nil, loop body never executes. Header string matches expected exactly.

2. **One entry (USD, en-US, -1000)** - Date "2015-01-01" parsed, formatted as "01/01/2015" via `%02d/%02d/%04d`. Description "Buy present" (11 chars < 25) padded to 25. Amount: cents=1000, whole=10, remainder=0, amountStr="10.00", negative so "($10.00)" (8 chars), right-padded to 13 = "     ($10.00)".

3. **Credit and debit (sorting by date)** - Entries are out of date order. String comparison on "YYYY-MM-DD" format sorts correctly: "2015-01-01" < "2015-01-02".

4. **Multiple entries same date (sort by description)** - Same date "2015-01-01", "Buy present" < "Get present" alphabetically.

5. **Tie breaker by change** - Same date, same description "Something". Changes: -1, 0, 1 sorted numerically.
   - Change=-1: "($0.01)" (7 chars) -> "      ($0.01)" (13 chars)
   - Change=0: "$0.00 " (6 chars) -> "       $0.00 " (13 chars)
   - Change=1: "$0.01 " (6 chars) -> "       $0.01 " (13 chars)

6. **Overlong descriptions** - "Freude schoner Gotterfunken" is 27 chars > 25. Truncated: `desc[:22]` = "Freude schoner Gotterf" + "..." = 25 chars total. Amount -123456: whole=1234, len("1234")=4 > 3, thousands grouping produces "1,234", amountStr="1,234.56", formatted="($1,234.56)" (11 chars) -> "  ($1,234.56)" (13 chars).

7. **Euros** - Currency "EUR" maps to symbol "€". Multi-byte UTF-8 char, but Go's `fmt.Sprintf("%13s", ...)` counts runes not bytes, so padding is correct: "(€10.00)" = 8 runes -> "     (€10.00)" = 13 runes.

8. **Dutch locale** - Header "Datum      | Omschrijving              | Verandering\n" matches. Date formatted as DD-MM-YYYY: "12-03-2015". Amount 123456 positive in nl-NL: thousandSep=".", decimalSep=",", wholeStr="1.234", amountStr="1.234,56", formatted="$ 1.234,56 " (11 chars) -> "  $ 1.234,56 " (13 chars).

9. **Dutch negative (-12345)** - whole=123, len("123")=3, NOT > 3 so no thousands separator. amountStr="123,45". Negative nl-NL: "$ 123,45-" (9 chars) -> "    $ 123,45-" (13 chars).

10. **American negative (-12345)** - whole=123, amountStr="123.45". Negative en-US: "($123.45)" (9 chars) -> "    ($123.45)" (13 chars).

### Failure Cases

11. **Empty currency ""** - Falls to `default` in currency switch -> returns error.
12. **Invalid currency "ABC"** - Same `default` branch -> error.
13. **Empty locale ""** - Falls to `default` in locale switch -> error.
14. **Invalid locale "nl-US"** - Same `default` branch -> error.
15. **Invalid date "2015-131-11"** - `time.Parse("2006-01-02", ...)` fails (131 not valid month) -> returns error.
16. **Invalid date "2015-12/11"** - `time.Parse("2006-01-02", ...)` fails (wrong separator) -> returns error.

### Immutability Test

- `sorted := make([]Entry, len(entries)); copy(sorted, entries)` creates a separate copy.
- `sort.Slice(sorted, ...)` sorts the copy, not the original.
- `Entry` is a struct with string and int fields (value types), so `copy()` is a true deep copy.
- Original `entries` slice remains unmodified.

## Design Quality

- Clean separation: `formatAmount` helper isolates amount formatting logic.
- Thousand-separator logic correctly handles numbers with >3 digits using a loop that groups from the right.
- Date validation happens before sorting, so invalid dates are caught early.
- `strings.Builder` used for efficient string concatenation.

## Minor Observations (Not Bugs)

- `len(desc)` uses byte count rather than rune count for description truncation. This works for all test cases since the only multi-byte description ("Freude schoner Gotterfunken" in success tests) is ASCII. The immutability test has Unicode chars but doesn't check formatting output.
- The `time` package is used for both validation and formatting. An alternative would be manual string parsing, but `time.Parse` is idiomatic Go.

## Conclusion

Implementation is correct, clean, and handles all edge cases required by the test suite. No changes needed.
