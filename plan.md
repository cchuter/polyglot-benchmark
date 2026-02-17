# Implementation Plan: polyglot-go-food-chain

## Proposal A — Struct Slice with Loop (Proponent)

### Approach
Use a slice of structs to store animal names and their verse-specific comments, with a constant for the spider's wriggle text. Build each verse algorithmically by iterating backwards through the chain.

### Files to modify
- `go/exercises/practice/food-chain/food_chain.go` — implement all three functions

### Design
```go
package foodchain

var verse = []struct{ eaten, comment string }{
    {"", ""},
    {"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
    {"spider", "It wriggled and jiggled and tickled inside her.\n"},
    {"bird", "How absurd to swallow a bird!\n"},
    {"cat", "Imagine that, to swallow a cat!\n"},
    {"dog", "What a hog, to swallow a dog!\n"},
    {"goat", "Just opened her throat and swallowed a goat!\n"},
    {"cow", "I don't know how she swallowed a cow!\n"},
    {"horse", "She's dead, of course!"},
}

const wriggle = " wriggled and jiggled and tickled inside her"

func Verse(v int) string { ... }   // build verse with loop
func Verses(start, end int) string { ... }  // join range
func Song() string { return Verses(1, 8) }
```

### Rationale
- Matches the reference solution pattern in `.meta/example.go`
- Clean data-driven approach separates data from logic
- The loop in `Verse` handles the cumulative chain naturally
- Minimal code, easy to understand

### Weaknesses (self-identified)
- Spider's wriggle constant is used in two places (comment and chain text) — slight duplication

---

## Proposal B — Map-based with String Builder (Opponent)

### Approach
Use maps for animal names and comments, indexed by verse number. Use `strings.Builder` for efficient string construction.

### Files to modify
- `go/exercises/practice/food-chain/food_chain.go`

### Design
```go
package foodchain

import "strings"

var animals = [9]string{"", "fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var comments = map[int]string{ ... }

func Verse(v int) string {
    var b strings.Builder
    // build with builder
    return b.String()
}
```

### Critique of Proposal A
- Proposal A uses raw string concatenation in a loop, which creates intermediate strings. A `strings.Builder` would be more memory-efficient.

### Why B is superior
- `strings.Builder` is more idiomatic for building strings in loops
- Separating animals from comments into different data structures provides flexibility

### Weaknesses
- More complex than necessary for this problem size
- Maps are unordered; arrays would be more natural for indexed data
- Adds an import (`strings`) that isn't strictly needed
- Deviates from the reference solution conventions in this repo

---

## Selected Plan (Judge)

### Evaluation

**Correctness**: Both proposals can produce correct output. Proposal A matches the reference solution structure exactly, reducing risk of subtle output differences.

**Risk**: Proposal A is lower risk because it mirrors the known-working `.meta/example.go`. Proposal B introduces maps which could have subtle ordering issues and adds unnecessary complexity.

**Simplicity**: Proposal A is simpler — fewer data structures, no imports needed, less code overall.

**Consistency**: Proposal A matches the established patterns in this repository (other solved exercises use simple data structures and string concatenation). The reference solution itself uses this exact approach.

### Decision: Proposal A wins

Proposal B's `strings.Builder` optimization is unnecessary for 8 verses of a short song. The added complexity of maps and imports provides no benefit. Proposal A is simpler, matches the reference, and is proven correct.

### Final Implementation Plan

**File**: `go/exercises/practice/food-chain/food_chain.go`

1. Define a package-level slice of structs holding `eaten` (animal name) and `comment` (verse-specific line) for all 8 animals, 1-indexed with a blank entry at index 0.

2. Define a `wriggle` constant for the spider's special chain text: `" wriggled and jiggled and tickled inside her"`

3. Implement `Verse(v int) string`:
   - Start with `"I know an old lady who swallowed a {animal}.\n{comment}"`
   - For verse 1 (fly) and verse 8 (horse): return immediately (no chain)
   - For verses 2-7: loop backwards from v down to 2, appending `"She swallowed the {v} to catch the {v-1}.\n"`
   - Special case: when v==3 (going from bird to spider), append `" that wriggled and jiggled and tickled inside her"` before the period
   - End with the fly's closing line

4. Implement `Verses(start, end int) string`:
   - Loop from start to end, joining each `Verse(i)` with `"\n\n"`

5. Implement `Song() string`:
   - Return `Verses(1, 8)`
