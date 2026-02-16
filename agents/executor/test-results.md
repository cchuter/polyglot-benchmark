# Ledger Exercise - Build & Test Results

## Build

```
$ go build ./...
```

**Result: PASS** (exit code 0, no errors)

## Tests

```
$ go test -v ./...
```

### TestFormatLedgerSuccess (10/10 PASS)

| # | Subtest | Status |
|---|---------|--------|
| 1 | empty_ledger | PASS |
| 2 | one_entry | PASS |
| 3 | credit_and_debit | PASS |
| 4 | multiple_entries_on_same_date_ordered_by_description | PASS |
| 5 | final_order_tie_breaker_is_change | PASS |
| 6 | overlong_descriptions | PASS |
| 7 | euros | PASS |
| 8 | Dutch_locale | PASS |
| 9 | Dutch_negative_number_with_3_digits_before_decimal_point | PASS |
| 10 | American_negative_number_with_3_digits_before_decimal_point | PASS |

### TestFormatLedgerFailure (6/6 PASS)

| # | Subtest | Status |
|---|---------|--------|
| 1 | empty_currency | PASS |
| 2 | invalid_currency | PASS |
| 3 | empty_locale | PASS |
| 4 | invalid_locale | PASS |
| 5 | invalid_date_(way_too_high_month) | PASS |
| 6 | invalid_date_(wrong_separator) | PASS |

### TestFormatLedgerNotChangeInput (1/1 PASS)

| # | Subtest | Status |
|---|---------|--------|
| 1 | (main test) | PASS |

**Result: ALL 17 TESTS PASS** (exit code 0)

## Vet

```
$ go vet ./...
```

**Result: PASS** (exit code 0, no issues found)

## Summary

- **Build**: PASS
- **Tests**: 17/17 PASS (0 failures)
- **Vet**: PASS (no issues)
