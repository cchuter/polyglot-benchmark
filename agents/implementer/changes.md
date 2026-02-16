# Changes: markdown exercise implementation

## File changed
- `go/exercises/practice/markdown/markdown.go`

## Summary
Implemented the `Render` function that converts a subset of Markdown to HTML.

### Functions implemented
- **`Render(markdown string) string`** — Main entry point. Splits input by newlines and processes each line as a heading (`#`), list item (`*`), or paragraph. Tracks list state to properly wrap consecutive `<li>` items in `<ul>` tags. Uses `strings.Builder` for efficient string concatenation.
- **`getHeadingWeight(line string) int`** — Counts leading `#` characters (1–6) to determine heading level. Returns -1 for 7+ `#` characters, causing the line to be rendered as a paragraph.
- **`renderHTML(markdownLine string) string`** — Converts inline formatting: `__text__` to `<strong>text</strong>` and `_text_` to `<em>text</em>`. Bold is processed first to avoid conflicts with italic markers. Applied to paragraphs and list items only, not headings.

### Test results
All 17 tests pass.
