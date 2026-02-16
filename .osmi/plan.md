# Implementation Plan: polyglot-go-beer-song

## Branch 1: Switch-Case with Direct String Returns

**Approach**: Use a switch statement in `Verse()` to handle the four special cases (0, 1, 2, default). `Verses()` loops and concatenates. `Song()` delegates to `Verses(99, 0)`.

**Files to modify**: `go/exercises/practice/beer-song/beer_song.go`

**Architecture**:
- `Verse(n)`: switch on n — cases 0, 1, 2 return hardcoded strings; default uses `fmt.Sprintf`
- `Verses(start, stop)`: validate inputs, loop from start down to stop, append each verse + newline
- `Song()`: return `Verses(99, 0)` result, ignore error

**Rationale**: Simplest approach, mirrors the reference solution exactly. Minimal code, easy to read.

**Evaluation**:
- Feasibility: High — straightforward, no dependencies
- Risk: Very low — directly matches reference solution pattern
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~50 lines, 1 file

---

## Branch 2: Template-Based with Helper Functions

**Approach**: Define a verse template and helper functions for pluralization and action text. Build verses by filling in template parameters.

**Files to modify**: `go/exercises/practice/beer-song/beer_song.go`

**Architecture**:
- Helper: `bottleStr(n)` returns "bottle" or "bottles"
- Helper: `countStr(n)` returns "no more" or the number
- Helper: `actionStr(n)` returns the action line
- `Verse(n)`: validates, then assembles from helpers
- `Verses/Song`: same as Branch 1

**Rationale**: More extensible, less hardcoded strings.

**Evaluation**:
- Feasibility: High — standard Go
- Risk: Low — but more code means more surface for bugs
- Alignment: Fully satisfies criteria
- Complexity: ~70 lines, more functions to test mentally

---

## Branch 3: Lookup Table with Pre-computed Verses

**Approach**: Pre-compute all 100 verses into a slice at init time, then lookups are O(1).

**Files to modify**: `go/exercises/practice/beer-song/beer_song.go`

**Architecture**:
- `init()`: populate `verses [100]string` array
- `Verse(n)`: bounds check, return `verses[n]`
- `Verses/Song`: same loop approach

**Rationale**: Optimizes for performance — all string formatting happens once.

**Evaluation**:
- Feasibility: High
- Risk: Medium — init() adds complexity, harder to reason about
- Alignment: Fully satisfies criteria
- Complexity: ~60 lines, slightly over-engineered for this use case

---

## Selected Plan

**Branch 1: Switch-Case with Direct String Returns**

**Rationale**: This is a simple exercise with well-defined special cases. The switch-case approach is the most readable, most maintainable, and has the lowest risk. It matches the idiomatic Go style used in the reference solution. The other branches add complexity without meaningful benefit for this problem.

### Detailed Implementation

**File**: `go/exercises/practice/beer-song/beer_song.go`

**Imports**: `fmt`, `bytes`

**Functions**:

1. **`Verse(n int) (string, error)`**
   - If n < 0 or n > 99, return `("", error)`
   - Switch on n:
     - Case 0: return "No more bottles..." verse
     - Case 1: return "1 bottle..." verse (singular, "Take it down", "no more bottles")
     - Case 2: return "2 bottles..." verse (with "1 bottle" singular on second line)
     - Default: return formatted verse with `fmt.Sprintf` using n and n-1

2. **`Verses(start, stop int) (string, error)`**
   - Validate: start must be 0-99, stop must be 0-99, start >= stop
   - Loop from start down to stop, calling `Verse(i)` each time
   - Join verses with an extra newline between them
   - Use `bytes.Buffer` for efficient concatenation

3. **`Song() string`**
   - Call `Verses(99, 0)`, return result (ignoring error since inputs are known valid)

**Order of implementation**:
1. Add imports
2. Implement `Verse`
3. Implement `Verses`
4. Implement `Song`
5. Run tests
