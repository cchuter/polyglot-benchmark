# Implementation Plan: Bottle Song

## Proposal A: Lookup Table with Format Strings

**Role: Proponent**

### Approach

Use a simple lookup table (slice or map) mapping integers 0-10 to their English word equivalents, then build verses using `fmt.Sprintf`.

### Files to Modify

- `go/exercises/practice/bottle-song/bottle_song.go` — add the `Recite` function and helper utilities

### Implementation Details

1. Define a `numberWords` slice: `[]string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}`
2. Create a helper `bottlePhrase(n int) string` that returns `"{Word} green bottle(s)"` with correct capitalization and pluralization
3. Implement `Recite(startBottles, takeDown int) []string`:
   - Loop from `startBottles` down for `takeDown` iterations
   - For each verse, append 4 lines to the result slice
   - Between verses, append an empty string `""`
4. Use `strings.Title` (or manual capitalization) for the first line of each verse where the number word needs title case

### Rationale

- Straightforward and minimal code
- Easy to verify correctness by inspecting the lookup table
- No complex logic; the number range is fixed (0-10)
- Follows the pattern seen in beer-song where each verse is constructed independently

---

## Proposal B: Recursive Verse Builder

**Role: Opponent**

### Approach

Build a single `verse(n int) []string` function that returns one verse, then recursively chain verses together.

### Files to Modify

- `go/exercises/practice/bottle-song/bottle_song.go` — add `Recite` and recursive helpers

### Implementation Details

1. Define a map `var words = map[int]string{0: "no", 1: "one", ...}`
2. Create `verse(n int) []string` returning the 4 lines for a single verse
3. Implement `Recite` recursively:
   ```go
   func Recite(start, take int) []string {
       if take == 0 { return nil }
       v := verse(start)
       if take == 1 { return v }
       return append(append(v, ""), Recite(start-1, take-1)...)
   }
   ```

### Critique of Proposal A

- Proposal A's iterative approach requires managing separator logic (empty string between verses but not after the last verse), which is error-prone
- Using `strings.Title` is deprecated and adds complexity

### Rationale for Proposal B

- Recursive structure mirrors the song's structure naturally
- The base case handles the separator issue cleanly (no trailing separator)
- Using a map instead of a slice is more robust if extended beyond 10

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | Both can satisfy all test cases | Both can satisfy all test cases |
| Risk | Low — simple loop | Medium — recursive append with variadic spread creates intermediate allocations |
| Simplicity | Slightly more manual separator logic | Elegant but recursion is overkill for max 10 iterations |
| Consistency | Matches the iterative style of beer-song | Diverges from codebase conventions |

### Decision

**Proposal A wins** with refinements from Proposal B.

The iterative approach is simpler, more idiomatic Go, and consistent with the existing beer-song exercise. The recursion in Proposal B adds no real benefit for a max of 10 iterations and introduces unnecessary allocations. However, Proposal B's point about `strings.Title` deprecation is valid — we'll use manual capitalization instead.

### Final Plan

**File to modify:** `go/exercises/practice/bottle-song/bottle_song.go`

**Implementation:**

1. Add a package-level `numberWord` slice mapping ints 0-10 to lowercase English words
2. Add a helper `capitalize(s string) string` that uppercases the first rune
3. Implement `Recite(startBottles, takeDown int) []string`:
   - Pre-allocate a result slice
   - Loop `i` from `startBottles` down to `startBottles - takeDown + 1`
   - For each iteration, determine the current number word (capitalized) and next number word (lowercase)
   - Determine singular/plural: "bottle" if count == 1, "bottles" otherwise
   - Append 4 lines per verse:
     - `"{N} green bottle(s) hanging on the wall,"`
     - `"{N} green bottle(s) hanging on the wall,"`
     - `"And if one green bottle should accidentally fall,"`
     - `"There'll be {N-1} green bottle(s) hanging on the wall."`
   - If not the last verse, append `""` as separator
   - Return the result slice

**Key details:**
- `numberWord` = `[]string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}`
- Singular "bottle" used when count is exactly 1; "bottles" for all other counts (including 0)
- First two lines use capitalized number word; last line uses lowercase
- Third line always says "one green bottle" (singular, lowercase)
- No external dependencies
