# Alphametics Exercise - Test Results

## Build (`go build ./...`)

**Result: PASS** - No errors or warnings.

## Vet (`go vet ./...`)

**Result: PASS** - No issues detected.

## Tests (`go test -v -timeout 300s`)

**Result: ALL 10 TESTS PASS** (1.206s total)

| # | Test Case | Result | Duration |
|---|-----------|--------|----------|
| 1 | puzzle_with_three_letters | PASS | 0.00s |
| 2 | solution_must_have_unique_value_for_each_letter | PASS | 0.00s |
| 3 | leading_zero_solution_is_invalid | PASS | 0.00s |
| 4 | puzzle_with_two_digits_final_carry | PASS | 0.00s |
| 5 | puzzle_with_four_letters | PASS | 0.00s |
| 6 | puzzle_with_six_letters | PASS | 0.02s |
| 7 | puzzle_with_seven_letters | PASS | 0.08s |
| 8 | puzzle_with_eight_letters | PASS | 0.19s |
| 9 | puzzle_with_ten_letters | PASS | 0.41s |
| 10 | puzzle_with_ten_letters_and_199_addends | PASS | 0.40s |

## Summary

- **Build**: PASS
- **Vet**: PASS
- **Tests**: 10/10 PASS
- **Total test time**: 1.206s
- **Status**: All checks green
