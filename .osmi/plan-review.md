# Plan Review: polyglot-go-markdown

## Review Method
Manual review of the selected plan (Branch 1: Clean Imperative) against all 17 test cases in `cases_test.go`.

## Test Case Coverage Analysis

### 1. "parses normal text as a paragraph" — `"This will be a paragraph"` → `<p>This will be a paragraph</p>`
- **Plan handles**: Yes. Default case wraps in `<p>` tags.

### 2. "parsing italics" — `"_This will be italic_"` → `<p><em>This will be italic</em></p>`
- **Plan handles**: Yes. `renderInline` replaces `_` pairs with `<em>`.

### 3. "parsing bold text" — `"__This will be bold__"` → `<p><strong>This will be bold</strong></p>`
- **Plan handles**: Yes. `renderInline` replaces `__` pairs with `<strong>` before `_`.

### 4. "mixed normal, italics and bold text" — `"This will _be_ __mixed__"` → `<p>This will <em>be</em> <strong>mixed</strong></p>`
- **Plan handles**: Yes. Bold before italic ordering is correct.

### 5-10. Headings h1-h6
- **Plan handles**: Yes. `headingLevel` counts `#` chars, returns 1-6.

### 11. "h7 header level is a paragraph" — `"####### This will not be an h7"` → `<p>####### This will not be an h7</p>`
- **CRITICAL ISSUE**: The plan says `headingLevel` returns 0 for 7+ hashes, then the line falls through to paragraph. But the paragraph case calls `renderInline(content)` — the `#` and `*` symbols in the text should NOT be interpreted. Since `#` and `*` aren't processed by `renderInline` (which only handles `_` and `__`), this is actually fine. The raw line including the `#######` prefix will be wrapped in `<p>` tags. ✅

### 12. "unordered lists" — `"* Item 1\n* Item 2"` → `<ul><li>Item 1</li><li>Item 2</li></ul>`
- **Plan handles**: Yes. `strings.HasPrefix(line, "* ")` detects list items, `inList` flag manages `<ul>` open/close.

### 13. "With a little bit of everything" — header + bold list item + italic list item
- **Plan handles**: Yes. Header rendered first, then list items with inline markup.

### 14. "with markdown symbols in the header text that should not be interpreted"
- **Plan handles**: Yes. Header content after `# ` is NOT passed through `renderInline` in the reference, and the plan's heading branch doesn't call `renderInline` either. The `#` and `*` in text are just literal characters. ✅

### 15. "with markdown symbols in the list item text that should not be interpreted"
- **Plan handles**: Yes. `renderInline` only processes `_` and `__`, not `#` or `*`. ✅

### 16. "with markdown symbols in the paragraph text that should not be interpreted"
- **Plan handles**: Yes. Same reasoning as above. ✅

### 17. "unordered lists close properly with preceding and following lines"
- Input: `"# Start a list\n* Item 1\n* Item 2\nEnd a list"`
- Expected: `<h1>Start a list</h1><ul><li>Item 1</li><li>Item 2</li></ul><p>End a list</p>`
- **Plan handles**: Yes. The `inList` flag ensures `</ul>` is written when a non-list line follows list items.

## Potential Issues Identified

### Issue 1: Heading content extraction
The plan says to write `<hN>content</hN>` after counting `#` chars. Need to ensure we skip the space after the `#` markers. For `"# This will be an h1"`, the heading level is 1, so content starts at index `headerLevel + 1` (skipping `# `). This matches the reference: `line[headerWeight+1:]`.

### Issue 2: List item content extraction
The plan uses `strings.HasPrefix(line, "* ")` and extracts content as `line[2:]`. This is correct for all test cases.

### Issue 3: No edge case for empty input
No test case for empty input, so this is not a concern.

## Conclusion

The selected plan (Branch 1) correctly handles all 17 test cases. No revisions needed. The implementation should follow the reference solution's logic closely with improved naming and structure.

**Recommendation**: Proceed with implementation as planned.
