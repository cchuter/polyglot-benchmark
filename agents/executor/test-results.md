# Ledger Exercise - Test Results

## Test Run (`go test -v ./...`)

```
=== RUN   TestFormatLedgerSuccess
=== RUN   TestFormatLedgerSuccess/empty_ledger
=== RUN   TestFormatLedgerSuccess/one_entry
=== RUN   TestFormatLedgerSuccess/credit_and_debit
=== RUN   TestFormatLedgerSuccess/multiple_entries_on_same_date_ordered_by_description
=== RUN   TestFormatLedgerSuccess/final_order_tie_breaker_is_change
=== RUN   TestFormatLedgerSuccess/overlong_descriptions
=== RUN   TestFormatLedgerSuccess/euros
=== RUN   TestFormatLedgerSuccess/Dutch_locale
=== RUN   TestFormatLedgerSuccess/Dutch_negative_number_with_3_digits_before_decimal_point
=== RUN   TestFormatLedgerSuccess/American_negative_number_with_3_digits_before_decimal_point
--- PASS: TestFormatLedgerSuccess (0.00s)
    --- PASS: TestFormatLedgerSuccess/empty_ledger (0.00s)
    --- PASS: TestFormatLedgerSuccess/one_entry (0.00s)
    --- PASS: TestFormatLedgerSuccess/credit_and_debit (0.00s)
    --- PASS: TestFormatLedgerSuccess/multiple_entries_on_same_date_ordered_by_description (0.00s)
    --- PASS: TestFormatLedgerSuccess/final_order_tie_breaker_is_change (0.00s)
    --- PASS: TestFormatLedgerSuccess/overlong_descriptions (0.00s)
    --- PASS: TestFormatLedgerSuccess/euros (0.00s)
    --- PASS: TestFormatLedgerSuccess/Dutch_locale (0.00s)
    --- PASS: TestFormatLedgerSuccess/Dutch_negative_number_with_3_digits_before_decimal_point (0.00s)
    --- PASS: TestFormatLedgerSuccess/American_negative_number_with_3_digits_before_decimal_point (0.00s)
=== RUN   TestFormatLedgerFailure
=== RUN   TestFormatLedgerFailure/empty_currency
=== RUN   TestFormatLedgerFailure/invalid_currency
=== RUN   TestFormatLedgerFailure/empty_locale
=== RUN   TestFormatLedgerFailure/invalid_locale
=== RUN   TestFormatLedgerFailure/invalid_date_(way_too_high_month)
=== RUN   TestFormatLedgerFailure/invalid_date_(wrong_separator)
--- PASS: TestFormatLedgerFailure (0.00s)
    --- PASS: TestFormatLedgerFailure/empty_currency (0.00s)
    --- PASS: TestFormatLedgerFailure/invalid_currency (0.00s)
    --- PASS: TestFormatLedgerFailure/empty_locale (0.00s)
    --- PASS: TestFormatLedgerFailure/invalid_locale (0.00s)
    --- PASS: TestFormatLedgerFailure/invalid_date_(way_too_high_month) (0.00s)
    --- PASS: TestFormatLedgerFailure/invalid_date_(wrong_separator) (0.00s)
=== RUN   TestFormatLedgerNotChangeInput
--- PASS: TestFormatLedgerNotChangeInput (0.00s)
PASS
ok  	ledger	0.004s
```

**Result: ALL 16 TESTS PASS**

## Benchmark Run (`go test -bench=. -benchtime=1s`)

```
goos: linux
goarch: amd64
pkg: ledger
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkFormatLedger-128     	   36340	     32814 ns/op
PASS
ok  	ledger	1.540s
```

## Summary

| Metric | Value |
|--------|-------|
| Tests Passed | 16/16 |
| Tests Failed | 0 |
| Benchmark (ns/op) | 32,814 |
| Benchmark iterations | 36,340 |
| Overall Result | **PASS** |
