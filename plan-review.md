# Plan Review: Beer Song

## Verdict: PASS (with one clarification needed)

The plan is sound and covers all test cases. The approach aligns with the reference solution. A few notes:

### Test Coverage Check

| Test Case | Covered in Plan? |
|---|---|
| Verse(8) - typical verse | Yes (default 3-99) |
| Verse(3) - another typical | Yes (default 3-99) |
| Verse(2) - singular "1 bottle" on second line | Yes (explicit case) |
| Verse(1) - "Take it down", "no more bottles" | Yes (explicit case) |
| Verse(0) - "No more bottles", "Go to the store" | Yes (explicit case) |
| Verse(104) - invalid, expect error | Yes (n < 0 or n > 99) |
| Verses(8,6) - multiple verses | Yes |
| Verses(7,5) - different range | Yes |
| Verses(109,5) - invalid start | Yes |
| Verses(99,-20) - invalid stop | Yes |
| Verses(8,14) - start < stop | Yes |
| Song() - entire song equals Verses(99,0) | Yes |

All test cases are accounted for.

### Trailing Newline / Separator Format - CRITICAL DETAIL

The plan says: "Join verses with an extra `\n` between them (each verse already ends with `\n`, so append `\n` after each)."

This is correct but ambiguous. The key behavior (matching the reference solution and test expectations) is:

- Each verse from `Verse()` ends with `\n`.
- `Verses()` appends an additional `\n` **after every verse including the last one**.
- This means the output of `Verses(8,6)` ends with `...\n\n` (a trailing blank line), NOT just `...\n`.

The test expectations (`verses86`, `verses75`) confirm the output ends with `\n\n`. The reference solution achieves this by unconditionally writing `\n` after each verse in the loop, with no special-casing for the last verse.

**If an implementer interprets "join verses" as only adding separators *between* verses (not after the last one), the tests will fail.** The plan should clarify that `\n` is appended after every verse, not just between them.

### No Issues Found

- The `Verse` cases (0, 1, 2, default) match all expected strings exactly.
- Validation logic for `Verses` covers all three error test cases (invalid start, invalid stop, start < stop).
- `Song()` correctly delegates to `Verses(99, 0)`.
- The `n == 2` case is correctly separated from the default to handle the singular "1 bottle" on the second line.

### Minor Note

The plan does not mention that `Song()` returns `string` (not `(string, error)`). The test calls `Song()` without checking for an error. The reference solution signature is `Song() string`. This is implicit in the plan but worth being explicit about since `Verses` returns `(string, error)`.
