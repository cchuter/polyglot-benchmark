# Test Results: Go Alphametics Solver

**Date:** 2026-02-14
**Command:** `cd /tmp/osmi-bench-polyglot-go-alphametics-282051489/go/exercises/practice/alphametics && go test -v ./...`
**Result:** ALL TESTS PASSED

## Full Output

```
=== RUN   TestSolve
=== RUN   TestSolve/puzzle_with_three_letters
=== RUN   TestSolve/solution_must_have_unique_value_for_each_letter
=== RUN   TestSolve/leading_zero_solution_is_invalid
=== RUN   TestSolve/puzzle_with_two_digits_final_carry
=== RUN   TestSolve/puzzle_with_four_letters
=== RUN   TestSolve/puzzle_with_six_letters
=== RUN   TestSolve/puzzle_with_seven_letters
=== RUN   TestSolve/puzzle_with_eight_letters
=== RUN   TestSolve/puzzle_with_ten_letters
=== RUN   TestSolve/puzzle_with_ten_letters_and_199_addends
--- PASS: TestSolve (1.13s)
    --- PASS: TestSolve/puzzle_with_three_letters (0.00s)
    --- PASS: TestSolve/solution_must_have_unique_value_for_each_letter (0.00s)
    --- PASS: TestSolve/leading_zero_solution_is_invalid (0.00s)
    --- PASS: TestSolve/puzzle_with_two_digits_final_carry (0.00s)
    --- PASS: TestSolve/puzzle_with_four_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_six_letters (0.02s)
    --- PASS: TestSolve/puzzle_with_seven_letters (0.08s)
    --- PASS: TestSolve/puzzle_with_eight_letters (0.23s)
    --- PASS: TestSolve/puzzle_with_ten_letters (0.43s)
    --- PASS: TestSolve/puzzle_with_ten_letters_and_199_addends (0.37s)
PASS
ok  	alphametics	1.216s
```

## Summary

| Test Case | Result | Duration |
|-----------|--------|----------|
| puzzle_with_three_letters | PASS | 0.00s |
| solution_must_have_unique_value_for_each_letter | PASS | 0.00s |
| leading_zero_solution_is_invalid | PASS | 0.00s |
| puzzle_with_two_digits_final_carry | PASS | 0.00s |
| puzzle_with_four_letters | PASS | 0.00s |
| puzzle_with_six_letters | PASS | 0.02s |
| puzzle_with_seven_letters | PASS | 0.08s |
| puzzle_with_eight_letters | PASS | 0.23s |
| puzzle_with_ten_letters | PASS | 0.43s |
| puzzle_with_ten_letters_and_199_addends | PASS | 0.37s |

**Total: 10/10 tests passed in 1.216s**

- No build errors
- No warnings
- No test failures
