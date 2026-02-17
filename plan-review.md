# Plan Review: polyglot-go-ledger

## Overall Assessment

The plan is well-structured and covers most of the key concerns. The hybrid approach (clean helper functions without over-abstracted config structs) is a sound design decision. However, there are several specific issues ranging from a critical terminology error to minor gaps in specificity.

---

## 1. Column Widths and Padding

**Verdict: Mostly correct, but header detail is underspecified.**

The plan correctly identifies the three columns and mentions padding/truncation in step 8 (description formatting). However, it does not explicitly state the column widths anywhere in the final implementation plan. The widths must be:

- Date column: exactly 10 characters
- Description column: exactly 25 characters (left-aligned, right-padded with spaces)
- Change column: exactly 13 characters in data rows

Proposal B's inline header string is correct:
```
"Date       | Description               | Change\n"
```
This gives Date=10, Description=25, Change=6 (just the word, not padded to 13). The header's Change column is NOT padded to 13 characters -- it is just the bare word "Change" (6 chars) or "Verandering" (11 chars). This means header lines are shorter than data lines (47 vs 54 chars for en-US, 52 vs 54 chars for nl-NL). The plan does not call this out, though the inline header strings in Proposal B happen to be correct. The final selected plan should explicitly state these header strings to avoid ambiguity.

**Recommendation**: Add explicit column width constants (10, 25, 13) and note that the header's final column is not padded to 13.

---

## 2. Sorting Order

**Verdict: Correct.**

The plan (step 9) correctly specifies sorting by:
1. Date (string comparison, which works for YYYY-MM-DD format)
2. Description (alphabetical)
3. Change (numeric)

This matches both the "multiple entries on same date ordered by description" test case (same date, different descriptions, sorted alphabetically: "Buy present" before "Get present") and the "final order tie breaker is change" test case (same date, same description, sorted by change: -1, 0, 1).

No issues here.

---

## 3. Amount Formatting for Both Locales

**Verdict: Correct in substance, but see issue in point 4.**

The plan correctly identifies:
- en-US: comma as thousands separator, period as decimal separator
- nl-NL: period as thousands separator, comma as decimal separator

The plan's step 7 lists the formatting patterns:
- en-US positive: `symbol + amount + " "`
- en-US negative: `(symbol + amount)`
- nl-NL positive: `symbol + " " + amount + " "`
- nl-NL negative: `symbol + " " + amount + "-"`

These patterns are verified correct against the test cases:
- `$10.00 ` (en-US positive, with trailing space)
- `($10.00)` (en-US negative, parentheses)
- `$ 1.234,56 ` (nl-NL positive, space between symbol and amount, trailing space)
- `$ 123,45-` (nl-NL negative, space between symbol and amount, trailing minus)

The trailing space on positive amounts serves as a placeholder for the negative indicator character (closing paren or minus sign), which ensures consistent right-alignment.

---

## 4. Negative Number Formatting and Right-Alignment

**Verdict: CRITICAL ERROR in the plan -- says "Right-pad" but must be "Right-align" (left-pad).**

Step 7 of the plan states:

> Right-pad to 13 characters total

This is **wrong**. The test data clearly shows the Change column is **right-aligned** (left-padded with spaces) to 13 characters. For example:

| Formatted amount | Content | Content len | Leading spaces |
|---|---|---|---|
| `     ($10.00)` | `($10.00)` | 8 | 5 |
| `      $10.00 ` | `$10.00 ` | 7 | 6 |
| `  ($1,234.56)` | `($1,234.56)` | 11 | 2 |
| `  $ 1.234,56 ` | `$ 1.234,56 ` | 11 | 2 |
| `    $ 123,45-` | `$ 123,45-` | 9 | 4 |

Every amount is left-padded with spaces so the total is 13 characters, making it right-aligned. If you right-padded instead, the amounts would be left-aligned and all tests would fail.

**Recommendation**: Change "Right-pad to 13 characters total" to "Right-align (left-pad with spaces) to 13 characters total". In Go, this is `fmt.Sprintf("%13s", formattedAmount)`.

---

## 5. Zero Amount Handling

**Verdict: Correct.**

The "final order tie breaker is change" test case includes `Change: 0`, which is displayed as `$0.00 ` (positive format with trailing space). Since zero is not negative (0 >= 0), it correctly falls into the positive formatting path. The plan's logic of "Determine if negative, work with absolute value" handles this naturally -- zero is not negative, so it gets positive formatting.

The sorting also handles zero correctly: with Change values of -1, 0, and 1, the expected order is -1, 0, 1 (ascending numeric sort), which the plan's "Change (numeric)" sort produces.

No issues here.

---

## 6. Immutability (Not Modifying Input)

**Verdict: Correct.**

The plan explicitly states in step 9:

> Copy the entries slice (to avoid modifying input).

And step 2d of the FormatLedger skeleton says:

> d) Copy and sort entries

This directly addresses the `TestFormatLedgerNotChangeInput` test, which verifies via `reflect.DeepEqual` that the original entries slice is unchanged after calling `FormatLedger`.

**Minor note**: The implementation must copy the slice contents (not just the slice header). A correct copy would be:
```go
sorted := make([]Entry, len(entries))
copy(sorted, entries)
```
Simply assigning `sorted := entries` or `sorted := entries[:]` would share the underlying array and fail the test. The plan does not spell this out, but the intent is clear.

---

## 7. Date Validation and Failure Test Cases

**Verdict: Correct, all failure cases accounted for.**

The plan covers:
- **Empty currency** and **invalid currency**: Step 3 uses a switch with a default error case, which handles both `""` and `"ABC"`.
- **Empty locale** and **invalid locale**: Step 4 uses a switch with a default error case, which handles both `""` and `"nl-US"`.
- **Invalid date (way too high month)**: `"2015-131-11"` -- Step 5 uses `time.Parse` with layout `"2006-01-02"`, which will reject month=131.
- **Invalid date (wrong separator)**: `"2015-12/11"` -- Step 5 checks that dashes are at positions 4 and 7, which catches the slash at position 7.

The plan's two-layer validation (structural format check + `time.Parse`) is robust. The structural check catches the wrong separator, and `time.Parse` catches semantically invalid dates.

**Minor note**: `time.Parse("2006-01-02", "2015-131-11")` -- this string is 11 characters, not 10, so the "exactly 10 chars" check in step 5 would catch it before even reaching `time.Parse`. This is fine; belt-and-suspenders validation.

---

## 8. nl-NL Amount Formatting Detail (Space Between Symbol and Amount)

**Verdict: Correct.**

The plan explicitly states in step 7:
- nl-NL positive: `symbol + " " + amount + " "`
- nl-NL negative: `symbol + " " + amount + "-"`

This matches the test expectations:
- `$ 1.234,56 ` (positive: dollar, space, amount, trailing space)
- `$ 123,45-` (negative: dollar, space, amount, trailing minus)

No issues here.

---

## 9. Right-Alignment of Change Column

**Verdict: See critical error in point 4.**

This is the same issue identified in point 4. The plan says "Right-pad to 13 characters total" but the test data requires right-alignment (left-padding). This is the single most important correction needed in the plan.

Additionally, step 9 (Section title "Right-alignment of Change column" in the review checklist) relates to how the final formatted string for the Change column is padded. The plan should clarify that the amount string (including its trailing space or negative indicator) is first fully constructed, then left-padded with spaces to exactly 13 characters.

---

## Summary of Findings

| Check | Status | Notes |
|-------|--------|-------|
| Column widths and padding | Minor gap | Widths not explicitly stated as constants; header not-padded behavior not called out |
| Sorting order | Pass | Correctly specified |
| Amount formatting (both locales) | Pass | Patterns correctly identified |
| Negative number formatting | **CRITICAL** | Says "right-pad" but must be "right-align" (left-pad) |
| Zero amount handling | Pass | Naturally handled by positive path |
| Immutability | Pass | Copy before sort is stated; implementation must use `copy()` not slice assignment |
| Date validation / failure cases | Pass | All 6 failure cases covered |
| nl-NL space between symbol and amount | Pass | Explicitly stated |
| Right-alignment of Change column | **CRITICAL** | Same as negative number formatting issue |

### Verdict

The plan is solid in its overall structure and covers nearly all edge cases correctly. There is one critical error: **step 7 says "Right-pad to 13 characters" when it must be "Right-align (left-pad with spaces) to 13 characters."** If implemented as written (right-pad = left-align), every single data row would fail. This must be corrected before implementation.

The remaining findings are minor: explicit column width constants and a note about the header not padding its last column would improve clarity but are not blockers.
