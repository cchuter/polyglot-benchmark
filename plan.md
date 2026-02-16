# Implementation Plan: Kindergarten Garden

## Branch 1: Map-Based (Simple, Minimal)

**Approach**: Use `Garden` as a `map[string][]string` type alias. During construction, validate inputs, sort a copy of children, and populate the map with plant names.

**Files to modify**: `kindergarten_garden.go` (single file)

**Architecture**:
- `type Garden map[string][]string` — simple type alias
- `NewGarden` validates, copies+sorts children, iterates rows and columns to fill the map
- `Plants` does a map lookup and returns `(plants, ok)`

**Implementation**:
```go
package kindergarten

import (
    "errors"
    "sort"
    "strings"
)

type Garden map[string][]string

func (g *Garden) Plants(child string) ([]string, bool) {
    p, ok := (*g)[child]
    return p, ok
}

func NewGarden(diagram string, children []string) (*Garden, error) {
    rows := strings.Split(diagram, "\n")
    if len(rows) != 3 || rows[0] != "" {
        return nil, errors.New("diagram must have two rows")
    }
    if len(rows[1]) != len(rows[2]) {
        return nil, errors.New("diagram rows must be same length")
    }
    if len(rows[1])%2 != 0 || len(rows[1]) != 2*len(children) {
        return nil, errors.New("each diagram row must have two cups per child")
    }
    alpha := append([]string{}, children...)
    sort.Strings(alpha)
    g := Garden{}
    for _, n := range alpha {
        g[n] = make([]string, 0, 4)
    }
    if len(g) != len(alpha) {
        return nil, errors.New("no two children can have the same name")
    }
    for _, row := range rows[1:] {
        for nx, n := range alpha {
            for cx := range rows[1:] {
                var p string
                switch row[2*nx+cx] {
                case 'G': p = "grass"
                case 'C': p = "clover"
                case 'R': p = "radishes"
                case 'V': p = "violets"
                default:
                    return nil, errors.New("plant codes must be one of G, C, R, or V")
                }
                g[n] = append(g[n], p)
            }
        }
    }
    return &g, nil
}
```

**Evaluation**:
- Feasibility: High — mirrors the reference solution exactly
- Risk: Very low — proven pattern from `.meta/example.go`
- Alignment: Fully satisfies all acceptance criteria
- Complexity: Minimal — single file, ~55 lines

---

## Branch 2: Struct-Based (Extensible)

**Approach**: Use a `Garden` struct with a `plants` field (map). This allows future extension (e.g., adding metadata, row count, etc.).

**Files to modify**: `kindergarten_garden.go` (single file)

**Architecture**:
- `type Garden struct { plants map[string][]string }`
- `NewGarden` validates inputs and populates `plants`
- `Plants` looks up in the struct's map

**Implementation**: Similar logic to Branch 1 but wrapped in a struct. The `Plants` method would access `g.plants[child]` instead of `(*g)[child]`.

**Evaluation**:
- Feasibility: High — straightforward
- Risk: Low, but the test file expects `*Garden` return type with map-like behavior. The test calls `g.Plants()` on `*Garden` which works with both approaches.
- Alignment: Fully satisfies all criteria
- Complexity: Slightly more code than Branch 1, struct wrapper adds no value for current requirements

---

## Branch 3: Pre-computed Index Array (Performance)

**Approach**: Instead of a map, store plants in a flat `[]string` slice and use index arithmetic to look up children. Store sorted names in a slice and binary search for lookups.

**Files to modify**: `kindergarten_garden.go` (single file)

**Architecture**:
- `type Garden struct { names []string; plants []string }`
- Plants stored as flat array: child i's plants are at indices `[4*i..4*i+3]`
- `Plants` binary-searches `names` for the child, then indexes into `plants`

**Evaluation**:
- Feasibility: High
- Risk: Medium — more complex indexing logic, potential off-by-one errors
- Alignment: Satisfies all criteria
- Complexity: Higher than needed — over-engineered for this problem size (max 12 children)

---

## Selected Plan

**Branch 1: Map-Based (Simple, Minimal)** is the best choice.

**Rationale**:
- It directly mirrors the reference solution in `.meta/example.go`, which is the canonical answer
- Minimal code, minimal risk of bugs
- The test suite expects `*Garden` with a `Plants` method — a map type alias satisfies this cleanly
- No over-engineering; perfectly fits the problem constraints
- Branch 2 adds unnecessary struct wrapping for no current benefit
- Branch 3 is over-engineered for a 12-child problem

**Detailed Implementation Plan**:

1. Edit `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`
2. Add imports: `errors`, `sort`, `strings`
3. Define `type Garden map[string][]string`
4. Implement `Plants` method: map lookup returning `([]string, bool)`
5. Implement `NewGarden`:
   - Split diagram on `\n`, validate 3 parts with empty first part
   - Validate row lengths match
   - Validate row length == 2 * len(children) and even
   - Copy and sort children alphabetically
   - Initialize map entries
   - Check for duplicates via map length
   - Parse each cup code in each row, appending plant names
   - Return error for invalid codes
6. Run `go test` to verify all tests pass
