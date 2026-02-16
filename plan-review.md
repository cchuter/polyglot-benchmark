## Review of `.osmi/plan.md` Against `ledger_test.go`

### Findings (ordered by severity)

1. **Medium: Formatting contract is not explicit enough to guarantee test-exact output**
- The plan has widths and separators, but it doesn’t lock down row construction precisely (`.osmi/plan.md:62`, `.osmi/plan.md:66`).
- Tests require exact spacing, including trailing spaces for positive amounts and no padded right-alignment for header labels (`go/exercises/practice/ledger/ledger_test.go:21`, `go/exercises/practice/ledger/ledger_test.go:59`, `go/exercises/practice/ledger/ledger_test.go:157`).
- **Action:** Add an explicit rendering contract:
  - Data rows: date left-padded to 10? Actually left-aligned width 10, description left-aligned width 25, change right-aligned width 13, separator exactly `" | "`, newline at end of each row.
  - Header rows are literal strings per locale (not formatted with `%13s` for change header).

2. **Medium: Truncation behavior needs exact rule**
- Plan says truncate with `...`, but not the exact cutoff (`.osmi/plan.md:48`).
- Test expects 25-char description column with overlong values becoming `22 chars + "..."` (`go/exercises/practice/ledger/ledger_test.go:113`, `go/exercises/practice/ledger/ledger_test.go:125`).
- **Action:** Specify: if description length > 25, take first 22 and append `...`; else left-align/pad to 25.

3. **Medium: Number formatting approach needs stricter implementation guidance**
- Plan is directionally correct (`.osmi/plan.md:56`), but should explicitly forbid float-based formatting.
- **Action:** Specify integer-only math (`units := abs(change)/100`, `cents := abs(change)%100`) and robust thousands grouping. Also note a safe `abs` strategy for extreme negative values (convert to `int64` before abs logic) to avoid overflow edge cases.

4. **Low: Date validation/sort path could be cleaner**
- Plan validates dates with `time.Parse` and then sorts (`.osmi/plan.md:38`, `.osmi/plan.md:39`), which is fine.
- **Action:** Clarify sort key is `Date` string (after validation) then `Description` then `Change`, since `YYYY-MM-DD` is lexicographically sortable. This avoids unnecessary reparsing.

### Answers to your 5 checks

1. **Coverage of tests (success/failure/immutability):**  
Mostly yes. Plan covers invalid currency/locale/date, ordering tie-breakers, localization, and immutability via copy-before-sort (`.osmi/plan.md:35` to `.osmi/plan.md:40`, `.osmi/plan.md:79`).

2. **Missed edge cases:**  
Yes, a few are not called out: exact-25 description boundary, Unicode-safe truncation behavior, and extreme negative `Change` overflow handling.

3. **Architecture soundness:**  
Yes. Single-file + helpers + locale/currency config is clean and idiomatic for this exercise (`.osmi/plan.md:9`, `.osmi/plan.md:44`).

4. **Formatting details correctness:**  
Mostly correct, but too implicit. Needs explicit row/header formatting rules, padding behavior, and guaranteed trailing newline behavior.

5. **Number formatting approach:**  
Conceptually correct for tests (locale separators, sign styles), but should explicitly mandate integer arithmetic and define safe absolute-value handling.