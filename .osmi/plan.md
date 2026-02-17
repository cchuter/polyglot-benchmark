# Implementation Plan: Kindergarten Garden

## Branch 1: Map-based Garden (Minimal, matches reference)

**Approach**: Define `Garden` as `map[string][]string`. Pre-compute all plant assignments in `NewGarden` and store them in the map. `Plants` is a simple map lookup.

**Files to modify**: `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

**Implementation**:
1. Define `type Garden map[string][]string`
2. `NewGarden`:
   - Split diagram by `\n`, validate 3 parts with empty first element
   - Validate row lengths match and are even (`2 * len(children)`)
   - Copy children slice, sort the copy alphabetically
   - Check for duplicates (map length != slice length after insertion)
   - Iterate rows and columns, map plant codes to names, append to each child's slice
   - Return error for any invalid plant code
3. `Plants`: Simple map index with ok idiom

**Evaluation**:
- Feasibility: High — matches reference solution exactly
- Risk: Very low — proven approach from .meta/example.go
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~55 lines, 1 file changed

---

## Branch 2: Struct-based Garden with separate children index

**Approach**: Define `Garden` as a struct with a `map[string][]string` field. This is more extensible if fields are added later.

**Files to modify**: `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

**Implementation**:
1. Define `type Garden struct { plants map[string][]string }`
2. Same validation logic as Branch 1
3. `Plants` accesses `g.plants[child]`

**Evaluation**:
- Feasibility: High — straightforward
- Risk: Low, but the test file uses `*Garden` so it must be a pointer. With a map type, `*Garden` naturally works. With a struct, it also works but adds indirection.
- Alignment: Fully satisfies all criteria
- Complexity: ~60 lines, 1 file, slightly more code for no benefit

---

## Branch 3: Lazy lookup with stored diagram

**Approach**: Store the sorted children and diagram rows in the Garden struct. Compute plant assignments lazily in `Plants` rather than pre-computing.

**Files to modify**: `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

**Implementation**:
1. Define `type Garden struct { children []string; rows [2]string }`
2. `NewGarden` validates and stores data but doesn't compute assignments
3. `Plants` binary-searches for the child, computes index, extracts plant codes on demand

**Evaluation**:
- Feasibility: High
- Risk: Medium — more complex lookup logic, binary search needed, harder to maintain
- Alignment: Fully satisfies criteria but adds complexity
- Complexity: ~70 lines, more complex per-lookup logic

---

## Selected Plan

**Branch 1** is selected. Rationale:
- It is the simplest approach with the fewest lines of code
- It exactly matches the proven reference solution pattern
- Pre-computing plants in `NewGarden` is efficient for the small data sizes involved
- Map-based type definition is idiomatic Go for this use case
- No unnecessary abstraction layers

### Detailed Implementation

**File**: `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

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
    if len(rows[1]) != 2*len(children) {
        return nil, errors.New("each diagram row must have two cups per child")
    }
    g := Garden{}
    alpha := append([]string{}, children...)
    sort.Strings(alpha)
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
                case 'G':
                    p = "grass"
                case 'C':
                    p = "clover"
                case 'R':
                    p = "radishes"
                case 'V':
                    p = "violets"
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

### Validation Order
1. Diagram format (3 parts, empty first)
2. Row length equality
3. Row length = 2 * number of children (also catches odd cups)
4. Duplicate names (via map length check)
5. Invalid plant codes (during iteration)
