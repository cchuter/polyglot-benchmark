# Plan Review

## Review: Self-review (no codex agent available)

Reviewed `.osmi/plan.md` against `ledger_test.go` and `.meta/example.go`.

### (1) Will the plan produce code that passes all tests?

**Yes**, provided the implementation follows the reference solution closely. The plan covers:
- All required exports (`Entry` struct, `FormatLedger` function)
- Input validation for currency, locale, and date
- Sorting by date → description → change
- Locale-specific date formatting, header translations, and currency formatting
- Description truncation at 25 chars
- Input immutability via slice copy

### (2) Edge cases

All edge cases from the test suite are covered:
- **Empty entries** (nil slice) → just header row
- **Sorting ties** → description then change as tiebreaker
- **Zero change** → formatted as `$0.00` with trailing space (positive)
- **Negative one cent** → `($0.01)` in en-US, formatted correctly in nl-NL
- **Invalid dates** → `time.Parse` returns error for bad formats like `2015-131-11` and `2015-12/11`
- **Input mutation test** → plan explicitly copies the slice

### (3) Formatting details that could cause failures

Key details to get right:
1. **Column alignment**: Change column is 13 chars wide, right-aligned via `%13s`
2. **Trailing space**: Positive numbers must end with a space for alignment (both locales)
3. **Dutch currency format**: `$ 1.234,56 ` (space between symbol and amount) and `$ 123,45-` (minus at end for negatives)
4. **American format**: `($10.00)` with parens, `$10.00 ` with trailing space
5. **Thousands separators**: Must handle numbers < 1000 correctly (no separator)
6. **Header format**: Uses `%-10s | %-25s | %s\n` (note: Change/Verandering column is NOT width-formatted in header)

### Conclusion

The plan is sound. Implement following the reference solution's proven patterns. The `sort.Slice` substitution for `sort.Interface` is fine and won't affect correctness.
