# Challenger Review: markdown/markdown.go

## Summary

**PASS** — The implementation is correct, clean, and idiomatic. All 17 test cases verified by manual trace.

## Test Case Verification

| # | Description | Result |
|---|-------------|--------|
| 1 | parses normal text as a paragraph | PASS |
| 2 | parsing italics | PASS |
| 3 | parsing bold text | PASS |
| 4 | mixed normal, italics and bold text | PASS |
| 5 | with h1 header level | PASS |
| 6 | with h2 header level | PASS |
| 7 | with h3 header level | PASS |
| 8 | with h4 header level | PASS |
| 9 | with h5 header level | PASS |
| 10 | with h6 header level | PASS |
| 11 | h7 header level is a paragraph | PASS |
| 12 | unordered lists | PASS |
| 13 | With a little bit of everything | PASS |
| 14 | markdown symbols in header not reinterpreted | PASS |
| 15 | markdown symbols in list item not reinterpreted | PASS |
| 16 | markdown symbols in paragraph not reinterpreted | PASS |
| 17 | unordered lists close properly with preceding/following lines | PASS |

## Edge Case Analysis

- **h7 fallback**: `headerLevel` returns 7, which is > 6, so the line is correctly rendered as a paragraph with the raw `#######` text preserved.
- **Bold/italic ordering**: `applyInline` processes `__` (bold) before `_` (italic). This is essential — if `_` were processed first, `__bold__` would incorrectly become `<em></em>bold<em></em>`. Processing `__` first avoids ambiguity.
- **List close semantics**: When a non-list line follows list items, `</ul>` is written before the new line's content. When input ends mid-list, the post-loop `if inList` guard closes it.
- **Markdown symbols not reinterpreted**: `#` and `*` in content are safe because header/list detection only checks line prefixes, and `applyInline` only processes `_`/`__` markers.

## Code Quality

- Clean, readable, idiomatic Go.
- Efficient use of `strings.Builder` throughout.
- Minimal imports (only `"strings"`).
- Well-named functions with clear single responsibilities (`headerLevel`, `applyInline`, `replaceMarker`).
- `replaceMarker` is a clean generic abstraction for paired delimiter replacement.
- No unnecessary complexity or dead code.

## Verdict

No issues found. The implementation should pass all 17 test cases.
