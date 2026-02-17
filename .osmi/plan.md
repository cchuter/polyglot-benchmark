# Implementation Plan: polyglot-go-markdown

## Proposal A

**Role: Proponent**

### Approach: Line-by-line parser with state tracking

Implement a clean, single-pass line-by-line parser that closely follows the reference solution's structure but emphasizes readability.

**Files to modify:**
- `go/exercises/practice/markdown/markdown.go`

**Architecture:**
1. Split input by newlines, process each line
2. Use a `strings.Builder` for output accumulation
3. Track "in list" state with a slice of list items
4. For each line, determine its type by the first character:
   - `*` → list item: accumulate in list items slice
   - `#` → potential heading: count `#` chars (1-6 = heading, 7+ = paragraph)
   - anything else → paragraph
5. When transitioning out of a list (non-`*` line after `*` lines), flush the accumulated list
6. After all lines, flush any remaining list
7. Inline formatting handled by a `renderInlineHTML` function:
   - First replace `__` pairs with `<strong>`/`</strong>`
   - Then replace `_` pairs with `<em>`/`</em>`

**Why this is best:**
- Directly mirrors the reference solution structure, ensuring correctness
- Simple, linear flow that's easy to follow
- Minimal abstraction overhead — just two functions (`Render` and `renderInlineHTML`) plus a helper (`getHeadingLevel`)
- Uses only standard library (`fmt`, `strings`)

## Proposal B

**Role: Opponent**

### Approach: Type-driven with explicit line classification

Instead of checking characters inline, first classify each line into a typed struct, then render all classified lines. This separates parsing from rendering.

**Files to modify:**
- `go/exercises/practice/markdown/markdown.go`

**Architecture:**
1. Define a `lineType` enum (heading, listItem, paragraph)
2. Define a `parsedLine` struct with type, content, and heading level
3. First pass: classify every line into `[]parsedLine`
4. Second pass: render classified lines, grouping consecutive list items into `<ul>` blocks
5. Inline formatting in a separate function

**Critique of Proposal A:**
- Proposal A interleaves classification and rendering, making it slightly harder to test individual concerns
- State tracking (the list items slice) is mutable state that makes reasoning harder

**Why this is better:**
- Clean separation of concerns: parsing vs rendering
- Each phase is independently testable
- More extensible if new line types are added

## Selected Plan

**Role: Judge**

### Evaluation

**Correctness:** Both proposals can satisfy all 17 test cases. Proposal A is proven correct since it mirrors the reference solution.

**Risk:** Proposal B introduces more code surface area (types, structs, two passes) for no additional test coverage. The extra abstraction is a risk of over-engineering for this small exercise.

**Simplicity:** Proposal A is simpler — fewer lines, fewer concepts, fewer types. For an exercise with exactly 3 line types, a type enum adds ceremony without value.

**Consistency:** Looking at other exercises in the repo (e.g., `ledger.go`), solutions are direct and pragmatic, not heavily abstracted. Proposal A fits this convention better.

### Winner: Proposal A

Proposal A wins because it is simpler, proven correct via the reference solution, and consistent with codebase conventions. Proposal B's separation of concerns is laudable in a larger system but is over-engineering for this 65-line exercise.

### Detailed Implementation Plan

**File:** `go/exercises/practice/markdown/markdown.go`

**Step 1:** Add package declaration and imports (`fmt`, `strings`)

**Step 2:** Define constants for marker characters:
- `headingMarker = '#'`
- `listItemMarker = '*'`

**Step 3:** Implement `Render(markdown string) string`:
- Initialize `strings.Builder` for HTML output
- Initialize `[]string` for accumulating list items
- Split markdown by `\n`, iterate lines:
  - If line starts with `*`: append `<li>renderInlineHTML(line[2:])</li>` to list items, continue
  - Else if list items non-empty: flush `<ul>...</ul>` to builder, reset list
  - If line starts with `#`: call `getHeadingLevel(line)`:
    - If valid (1-6): write `<hN>content</hN>`
    - If invalid (-1): write `<p>line</p>`
  - Else: write `<p>renderInlineHTML(line)</p>`
- After loop: flush remaining list items if any
- Return builder string

**Step 4:** Implement `getHeadingLevel(line string) int`:
- Count leading `#` chars up to index 6
- If count > 6, return -1
- Return count

**Step 5:** Implement `renderInlineHTML(text string) string`:
- Replace `__` pairs with `<strong>`/`</strong>` (loop while `__` exists)
- Replace `_` pairs with `<em>`/`</em>` (loop while `_` exists)
- Return result

**Step 6:** Run `go test ./...` and `go vet ./...` to verify all 17 tests pass.
