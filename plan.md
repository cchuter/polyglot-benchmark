# Implementation Plan: Kindergarten Garden

## Proposal A (Proponent)

**Approach: Map-based Garden (match reference solution pattern)**

Define `Garden` as `map[string][]string` and pre-compute all plant assignments in `NewGarden`.

### Files to modify
- `go/exercises/practice/kindergarten-garden/kindergarten_garden.go` — implement full solution

### Architecture
- `Garden` is a type alias for `map[string][]string`, mapping child name to their 4 plants
- `NewGarden` validates the diagram, copies and sorts children, then iterates rows and columns to build the map
- `Plants` is a simple map lookup returning `([]string, bool)`

### Implementation detail
```go
type Garden map[string][]string

func NewGarden(diagram string, children []string) (*Garden, error) {
    // 1. Split diagram by newline, validate 3 parts (empty + row1 + row2)
    // 2. Validate row lengths match
    // 3. Copy children slice, sort copy alphabetically
    // 4. Validate row length == 2 * len(children) (catches odd cups too)
    // 5. Check for duplicate names via map size vs slice length
    // 6. Iterate rows, for each child pick 2 chars per row, decode plant codes
    // 7. Return error on invalid plant code
    // Return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
    p, ok := (*g)[child]
    return p, ok
}
```

### Rationale
- Directly follows the reference solution pattern in `.meta/example.go`
- Simple, minimal code — map provides O(1) lookup
- Pre-computing in the constructor means `Plants` is trivially fast
- Matches test expectations exactly (pointer to map type)

---

## Proposal B (Opponent)

**Approach: Struct-based Garden with computed lookups**

Define `Garden` as a struct holding the raw diagram rows and a sorted name-to-index mapping. Compute plants on the fly in `Plants`.

### Files to modify
- `go/exercises/practice/kindergarten-garden/kindergarten_garden.go` — implement full solution

### Architecture
```go
type Garden struct {
    rows     [2]string
    nameIdx  map[string]int
}
```
- `NewGarden` validates and stores the diagram rows, builds name→index map
- `Plants` computes the 4 plants from stored rows on each call

### Rationale for preferring over Proposal A
- Separates storage from computation
- Uses less memory (stores diagram once instead of expanded strings)
- More extensible if requirements change

### Critique of Proposal A
- Using a type alias for a map is unconventional Go
- Pre-computing all plants up front when only some may be queried wastes work

---

## Selected Plan (Judge)

### Evaluation

| Criterion     | Proposal A (Map)                    | Proposal B (Struct)                     |
|---------------|------------------------------------|-----------------------------------------|
| Correctness   | Matches reference solution exactly | Correct but more code needed            |
| Risk          | Very low — mirrors proven solution | Medium — more places for off-by-one     |
| Simplicity    | ~50 lines, straightforward         | ~60 lines, more complex Plants method   |
| Consistency   | Matches `.meta/example.go` exactly | Deviates from established pattern       |

**Proposal B's criticisms are weak:**
- "Type alias for map is unconventional" — it's the exact pattern used in the reference solution and works perfectly with Go's type system
- "Pre-computing wastes work" — the benchmark tests call Plants on all children, and the test suite queries most children, so pre-computation is optimal
- "Separates storage from computation" — unnecessary complexity for this problem

**Winner: Proposal A**

The map-based approach is simpler, matches the reference solution, has fewer places for bugs, and is optimal for the benchmark test cases.

### Final Implementation Plan

**File:** `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

**Step-by-step:**

1. Define `Garden` as `map[string][]string`
2. Implement `Plants` method on `*Garden`: simple map dereference and lookup
3. Implement `NewGarden`:
   a. Split diagram on `"\n"`, verify exactly 3 parts with empty first part
   b. Verify both row strings have equal length
   c. Copy the `children` slice (to avoid mutation), sort the copy
   d. Verify row length equals `2 * len(sortedChildren)` — this also catches odd-cup scenarios
   e. Initialize Garden map, insert each child with an empty slice
   f. Check `len(g) == len(sortedChildren)` to detect duplicates
   g. Iterate over each row (rows[1], rows[2]), for each child at index `nx`:
      - For each cup offset (0, 1): decode `row[2*nx + offset]` to plant name
      - Valid codes: G→"grass", C→"clover", R→"radishes", V→"violets"
      - Return error on any other character
      - Append plant name to child's slice
   h. Return `&g, nil`

**Imports:** `errors`, `sort`, `strings`

**Validation error cases (from tests):**
- `"RC\nGG"` — no leading newline → error (rows[0] != "")
- `"\nRCCC\nGG"` — mismatched row lengths → error
- `"\nRCC\nGGC"` — 3 cups per row with 1 child (3 != 2*1) → error (odd cups)
- `"\nRCCC\nGGCC"` with `["Alice","Alice"]` — duplicate → error (map size < slice size)
- `"\nrc\ngg"` — lowercase codes → error (invalid plant code)
